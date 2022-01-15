package cron

import (
	"time"
)

func Crontab(f func(), h int , m int) {
	for true {
		t := time.Now()
		loc, _ := time.LoadLocation("Asia/Taipei")
		n := t.In(loc)
		hour := n.Hour()
		min := n.Minute()
		if hour == h && min == m {
			go f()
			time.Sleep(23 * time.Hour)
		}
	}
}