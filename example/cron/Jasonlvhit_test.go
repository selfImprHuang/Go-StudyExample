/*
 *  @Author : huangzj
 *  @Time : 2020/12/16 17:23
 *  @Description：
 */

package cron

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"testing"
	"time"
)

func TestGoCron(t *testing.T) {
	s := gocron.NewScheduler()
	_ = s.Every(1).Seconds().Do(task)
	_ = s.Every(4).Seconds().Do(superWang)

	sc := s.Start() // keep the channel
	go test(s, sc)  // wait
	<-sc            // it will happens if the channel is closed
}

func task() {
	fmt.Println("I am runnning task.", time.Now())
}
func superWang() {
	fmt.Println("I am runnning superWang.", time.Now())
}

func test(s *gocron.Scheduler, sc chan bool) {
	time.Sleep(8 * time.Second)
	s.Remove(task) //remove task
	time.Sleep(8 * time.Second)
	s.Clear()
	fmt.Println("All task removed")
	close(sc) // close the channel
}

func TestFunc(t *testing.T) {
	// Do jobs without params
	_ = gocron.Every(1).Second().Do(task)
	_ = gocron.Every(2).Seconds().Do(task)
	_ = gocron.Every(1).Minute().Do(task)
	_ = gocron.Every(2).Minutes().Do(task)
	_ = gocron.Every(1).Hour().Do(task)
	_ = gocron.Every(2).Hours().Do(task)
	_ = gocron.Every(1).Day().Do(task)
	_ = gocron.Every(2).Days().Do(task)
	_ = gocron.Every(1).Week().Do(task)
	_ = gocron.Every(2).Weeks().Do(task)

	// Do jobs with params
	_ = gocron.Every(1).Second().Do(taskWithParams, 1, "hello")

	// Do jobs on specific weekday
	_ = gocron.Every(1).Monday().Do(task)
	_ = gocron.Every(1).Thursday().Do(task)

	// Do a job at a specific time - 'hour:min:sec' - seconds optional
	_ = gocron.Every(1).Day().At("10:30").Do(task)
	_ = gocron.Every(1).Monday().At("18:30").Do(task)
	_ = gocron.Every(1).Tuesday().At("18:30:59").Do(task)

	// Begin job immediately upon start
	_ = gocron.Every(1).Hour().From(gocron.NextTick()).Do(task)

	// Begin job at a specific date/time
	timed := time.Date(2019, time.November, 10, 15, 0, 0, 0, time.Local)
	_ = gocron.Every(1).Hour().From(&timed).Do(task)

	// NextRun gets the next running time
	_, time := gocron.NextRun()
	fmt.Println(time)

	// Remove a specific job
	gocron.Remove(task)

	// Clear all scheduled jobs
	gocron.Clear()

	// Start all the pending jobs
	<-gocron.Start()

	// also, you can create a new scheduler
	// to run two schedulers concurrently
	s := gocron.NewScheduler()
	_ = s.Every(3).Seconds().Do(task)
	<-s.Start()
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}
