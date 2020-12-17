/*
 *  @Author : huangzj
 *  @Time : 2020/12/16 17:10
 *  @Description：
 */

package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"testing"
)

//测试单个定时任务
func TestCron(t *testing.T) {
	i := 0
	c := cron.New()
	spec := "0 */1 * * * ?" //一分钟运行一次
	_ = c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
	})
	c.Start()

	select {}
}

//多个定时任务
func TestMoreCron(t *testing.T) {
	i := 0
	c := cron.New()
	//AddFunc
	spec := "*/5 * * * * ?"
	_ = c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})

	//AddJob方法
	_ = c.AddJob(spec, TestJob{})
	_ = c.AddJob(spec, Test2Job{})

	//启动计划任务
	c.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select {}
}

type TestJob struct {
}

func (t TestJob) Run() {
	fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (t Test2Job) Run() {
	fmt.Println("testJob2...")
}
