package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ClickHouse/clickhouse-go"
	_ "github.com/go-sql-driver/mysql"
)

var clickDb *sql.DB
var db *sql.DB

const (
	mysqlDsn      = "root:root123@/hack?charset=utf8mb4&parseTime=True&loc=Local"
	clickhouseDsn = "tcp://192.168.3.8:9000?debug=false&username=&password=123456&compress=true"
)

func init() {
	var err error
	db, err = sql.Open("mysql", mysqlDsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(100)
	clickDb, err = sql.Open("clickhouse", clickhouseDsn)
	if err != nil {
		panic(err)
	}
	clickDb.SetMaxOpenConns(90)
	if err := clickDb.Ping(); err != nil {
		if ex, ok := err.(*clickhouse.Exception); ok {
			log.Printf("[%d] %s \n%s\n", ex.Code, ex.Message, ex.StackTrace)
		}
		panic(err)
	}
}
func qq() {
	source := "D:\\迅雷下载\\buffhack\\6.9更新总库.txt"
	f, _ := os.Open(source)
	scanner := bufio.NewScanner(f)
	count := 1
	go func() {
		t := time.NewTicker(time.Second * 10)
		for range t.C {
			log.Printf("当前已新增%d条数据\n", count-1)
		}
	}()
	dataChan := make(chan string, 100)
	sqlWg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go func() {
			sqlWg.Add(1)
			defer sqlWg.Done()
			for str := range dataChan {
				if _, err := db.Exec(str); err != nil {
					log.Println("exec sql: ", err, len(str))
				}
			}
		}()
	}
	lineChan := make(chan string, 1000)
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			var arr []string
			sqlSb := &strings.Builder{}
			sqlSb.WriteString("INSERT INTO `qq2` (`qq`, `phone`) VALUES ")
			for line := range lineChan {
				line = strings.TrimSpace(line)
				arr = strings.Split(line, "----")
				if len(arr) == 3 {
					arr[1] = arr[2]
				}
				if len(arr) != 2 && len(arr) != 3 {
					log.Printf("%s 不是规范的行\n", line)
					continue
				}
				_, _ = sqlSb.WriteString("('" + arr[0] + "','" + arr[1] + "'" + "),")
				if count%1001 == 0 {
					sqlStr := sqlSb.String()
					dataChan <- sqlStr[:len(sqlStr)-1]
					sqlSb.Reset()
					sqlSb.WriteString("INSERT INTO `qq2` (`qq`, `phone`) VALUES ")
				}
			}
			sqlSb.WriteString("('buff','buff'),")
			sqlStr := sqlSb.String()
			dataChan <- sqlStr[:len(sqlStr)-1]
		}()
	}
	begin := time.Now()
	for scanner.Scan() {
		lineChan <- scanner.Text()
		count++
	}
	close(lineChan)
	wg.Wait()
	close(dataChan)
	sqlWg.Wait()
	log.Printf("共插入%d条数据,每秒%.2f条", count, float64(count)/time.Since(begin).Seconds())
}
func jd() {
	source := "D:\\迅雷下载\\buffhack\\1\\京东快递解压密码pncldyerk4gqofhp.onion\\www_jd_com_12g.txt"
	f, _ := os.Open(source)
	scanner := bufio.NewScanner(f)
	count := 1
	go func() {
		t := time.NewTicker(time.Second * 10)
		for range t.C {
			log.Printf("当前已新增%d条数据\n", count-1)
		}
	}()
	dataChan := make(chan string, 100)
	sqlWg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go func() {
			sqlWg.Add(1)
			defer sqlWg.Done()
			for str := range dataChan {
				if _, err := db.Exec(str); err != nil {
					log.Println("exec sql: ", err, len(str))
				}
			}
		}()
	}
	lineChan := make(chan string, 1000)
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			var arr []string
			sqlSb := &strings.Builder{}
			sqlSb.WriteString(
				"INSERT INTO `jd` (`name`, `nick`, `email`,`idcard`, `phone1`, `phone2`) VALUES ")
			for line := range lineChan {
				line = strings.TrimSpace(line)
				arr = strings.Split(line, "---")
				if len(arr) != 7 {
					// log.Printf("%s 不是规范的行\n", line)
					continue
				}
				sqlSb.WriteString("('")
				sqlSb.WriteString(arr[0])
				sqlSb.WriteString("','")
				sqlSb.WriteString(arr[1])
				sqlSb.WriteString("','")
				sqlSb.WriteString(arr[3])
				sqlSb.WriteString("','")
				sqlSb.WriteString(arr[4])
				sqlSb.WriteString("','")
				sqlSb.WriteString(arr[5])
				sqlSb.WriteString("','")
				sqlSb.WriteString(arr[6])
				sqlSb.WriteString("'),")
				if count%1001 == 0 {
					sqlStr := sqlSb.String()
					dataChan <- sqlStr[:len(sqlStr)-1]
					sqlSb.Reset()
					sqlSb.WriteString(
						"INSERT INTO `jd` (`name`, `nick`, `email`,`idcard`, `phone1`, " +
							"`phone2`) VALUES ")
				}
			}
			sqlSb.WriteString("('buff','buff','buff','buff','buff','buff'),")
			sqlStr := sqlSb.String()
			dataChan <- sqlStr[:len(sqlStr)-1]
		}()
	}
	begin := time.Now()
	for scanner.Scan() {
		lineChan <- scanner.Text()
		count++
	}
	close(lineChan)
	wg.Wait()
	close(dataChan)
	sqlWg.Wait()
	// 141639667
	log.Printf("共插入%d条数据,每秒%.2f条", count, float64(count)/time.Since(begin).Seconds())
}
func weibo() {
	source := "D:\\迅雷下载\\buffhack\\1\\微博五亿2019.txt"
	f, _ := os.Open(source)
	scanner := bufio.NewScanner(f)
	count := 1
	go func() {
		t := time.NewTicker(time.Second * 10)
		for range t.C {
			log.Printf("当前已新增%d条数据\n", count-1)
		}
	}()
	dataChan := make(chan string, 100)
	sqlWg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go func() {
			sqlWg.Add(1)
			defer sqlWg.Done()
			for str := range dataChan {
				if _, err := db.Exec(str); err != nil {
					log.Println("exec sql: ", err, len(str))
				}
			}
		}()
	}
	lineChan := make(chan string, 10)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			var arr []string
			sqlSb := &strings.Builder{}
			sqlSb.WriteString("INSERT INTO `weibo` (`phone`, `uid`) VALUES ")
			for line := range lineChan {
				line = strings.TrimSpace(line)
				arr = strings.Split(line, "\t")
				if len(arr) != 2 {
					log.Printf("%q 不是规范的行\n", line)
					continue
				}
				_, _ = sqlSb.WriteString("('" + arr[0] + "','" + arr[1] + "'" + "),")
				if count%1001 == 0 {
					sqlStr := sqlSb.String()
					dataChan <- sqlStr[:len(sqlStr)-1]
					sqlSb.Reset()
					sqlSb.WriteString("INSERT INTO `weibo` (`phone`, `uid`) VALUES ")
				}
			}
			sqlSb.WriteString("('buff','buff'),")
			sqlStr := sqlSb.String()
			dataChan <- sqlStr[:len(sqlStr)-1]
		}()
	}
	begin := time.Now()
	for scanner.Scan() {
		lineChan <- scanner.Text()
		count++
	}
	close(lineChan)
	wg.Wait()
	close(dataChan)
	sqlWg.Wait()
	log.Printf("共插入%d条数据,每秒%.2f条", count, float64(count)/time.Since(begin).Seconds())
}
func shunfeng() {
	source := "D:\\迅雷下载\\buffhack\\1\\shunfeng\\script.sql"
	f, _ := os.Open(source)
	// f2 := transform.NewReader(f, simplifiedchinese.GBK.NewDecoder())
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	reader := bufio.NewReader(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
		reader.ReadLine()
	}
}
func main() {
	shunfeng()
}
