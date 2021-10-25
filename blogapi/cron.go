package main

import (
	"blogapi/models"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("开始执行清除已被删除的标签")
		models.CleanDeletedTags()
		log.Println("执行清除已被删除的标签完毕!")
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("开始执行清除已被删除的文章")
		models.CleanDeletedArticles()
		log.Println("执行清除已被删除的文章完毕!")
	})
	c.Start()

	// 阻塞主线程
	t := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t.C:
			t.Reset(time.Second * 10)
		}
	}
}
