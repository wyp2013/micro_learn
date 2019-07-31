package utils

import (
	"fmt"
	"os"
	"time"
)

type XormLogger struct {
	F      *os.File
	File   string
	Suffix string
}

//获取当前小时数
func getNowHour() string {
	//获取时间戳
	timestamp := time.Now().Unix()

	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006010215")
}

func NewXormLogger(file string) *XormLogger {
	suffix := getNowHour()
	fileName := file + "." + suffix

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	x := XormLogger{F: f, File: file, Suffix: suffix}

	//启动断续器,定期变更日志文件句柄
	ticker := time.NewTicker(time.Second * 1)
	go func(l *XormLogger) {
		for range ticker.C {
			if l.Suffix != getNowHour() {
				fmt.Println("Change log file")
				su := getNowHour()
				fileName := l.File + "." + su
				tf, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
				if err != nil {
					fmt.Println("Reopen failed!", err.Error())
				} else {
					go func(file *os.File) {
						time.Sleep(30 * time.Second)
						fmt.Println("Old file close")
						file.Close()
					}(l.F)
					l.F = tf
					l.Suffix = su
				}
			}
			//fmt.Println(t)
		}
	}(&x)

	return &x
}

//实现Writer接口
func (x *XormLogger) Write(p []byte) (n int, err error) {
	return x.F.Write(p)
}
