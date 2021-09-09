package taskManger

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"go.uber.org/atomic"
)

var (
	ErrFinished = errors.New("task manger is finished")
)

type (
	taskItem struct {
		Title string
		Data  interface{}
	}
	taskManger struct {
		i        *atomic.Uint32
		finished bool
		tasks    map[string]taskItem
	}
)

func New() *taskManger {
	return &taskManger{
		i:     atomic.NewUint32(0),
		tasks: map[string]taskItem{},
	}
}
func (t taskManger) Go(fn func(), title string, dataParam ...interface{}) error {
	if t.finished {
		return ErrFinished
	}
	go func() {
		t.i.Inc()
		id := uuid.New().String()
		var data interface{}
		if len(dataParam) > 0 {
			data = dataParam[0]
		}
		t.tasks[id] = taskItem{
			Title: title,
			Data:  data,
		}
		fn()
		delete(t.tasks, id)
		t.i.Dec()
	}()
	return nil
}
func (t taskManger) Wait(ctx context.Context) error {
	for {
		if t.i.Load() == 0 {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}
func (t taskManger) Tasks() map[string]taskItem {
	return t.tasks
}
func (t *taskManger) SetFinish() {
	t.finished = true
}
func Example() {
	tm := New()
	_ = tm.Go(f1, "f1", "buffge qqqqqqq")
	_ = tm.Go(f2, "f2")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tm.SetFinish()
	if err := tm.Wait(ctx); err != nil {
		log.Println("等待所有任务完成失败,未完成任务如下: ")
		for _, item := range tm.Tasks() {
			log.Println("title:", item.Title, ",data:", item.Title)
		}
	}
}
func f1() {
	time.Sleep(11 * time.Second)
}
func f2() {
	time.Sleep(3 * time.Second)
}
