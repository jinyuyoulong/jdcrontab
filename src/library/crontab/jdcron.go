package crontab

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

type Tesk struct {
	Id string
}
// 实现cron.Job{}接口
func (t Tesk)Run()  {
	fmt.Println("tesk job ",t.Id," running...")
}

var serviceCron *cron.Cron

func init() {
	serviceCron = cron.New()
	serviceCron.Start()
}

// 周期任务
func CycleCron(spec string)  {
	tesk2 := Tesk{Id:"name2"}

	//serviceCron.AddFunc(spec1,func1)

	entryID, err := serviceCron.AddJob(spec, tesk2)
	if err != nil {
		fmt.Println(err)
	}
	serviceCron.Start()

	//即将执行的任务
	go getNextJobs()

	time.Sleep(3600*time.Second)
	serviceCron.Remove(entryID)
	//select {}
}

func CycleJob(spec string,job cron.Job)  cron.EntryID {
	entryID, err := serviceCron.AddJob(spec, job)
	if err != nil {
		fmt.Println(err)
	}


	serviceCron.Start()

	//即将执行的任务
	//go getNextJobs()
	return entryID
}

//定时任务
func TimeCountJob(spec string)  {
	//spec := "1 * * * * *" 1秒后执行，执行一次
	entryID := addJob(spec, Tesk{Id:"job1"})

	getNextJobs()

	// 2秒后移除 entry
	c := time.After(2*time.Minute+1)
	<- c
	RemoveJob(entryID)


}

func addJob(spec string,job cron.Job)  cron.EntryID{
	//paser := cron.Parser{}
	//schedule, err := paser.Parse(spec)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//return serviceCron.Schedule(schedule,job)

	entryID, err := serviceCron.AddJob(spec, job)
	if err != nil {
		log.Println("TimeCountJob:",err)
		return -1
	}
	fmt.Println("entry id=",entryID)

	return entryID
}

func RemoveJob(id  cron.EntryID) {
	serviceCron.Remove(id)
}

func getNextJobs()  {
	fmt.Println("next jobs start... now ===>",time.Now())

	entries := getEntries(2)
	for k,v := range entries{
		job := v.Job
		nextTime := v.Next.Unix()
		nextType := fmt.Sprintf("%T\n",job)

		if nextType == "main.Test"{
			fmt.Println(k,"job type ",job.(Tesk).Id,"nextTime ===>",time.Unix(nextTime,0))
		}else {
			fmt.Println(k,"job type ","func","nextTime ===>",time.Unix(nextTime,0))
		}
	}
	fmt.Println("next job end")
}

func getEntries(size int) []cron.Entry  {

	result := serviceCron.Entries()
	if len(result) > size {
		result  = result[:size]
	}
	return result
}

func func1()  {
	fmt.Println("func1 name1 running")
}