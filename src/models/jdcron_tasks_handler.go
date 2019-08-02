package models

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func (j JdcronTasks)GetBy(id int) *JdcronTasks {
	var model JdcronTasks
	sql := "select * from jdcron_tasks where id=?"
	ok, err := engine.SQL(sql, id).Get(&model)
	//if err != nil {
	//	log.Printf("jdcron_tasks get by id:%v err:\n",id,err)
	//}

	log.Printf("GetBy:%v",model)

	//ok, err := engine.Get(model)
	if ok && err == nil {
		return &model
	}else {
		model.Id = 0
		return &model
	}
}
func (j JdcronTasks) GetAll() []JdcronTasks {
	data := []JdcronTasks{}
	sql := "select * from jdcron_tasks"
	engine.SQL(sql).Find(&data)
	return data
}

func (j *JdcronTasks)Insert(task *JdcronTasks)  {
	_, err := engine.Insert(task)
	if err != nil {
		log.Println("JdcronTasks Insert",err)
	}
}

func (j JdcronTasks)UpdateTID(id int,tid string,status int)  {
	fmt.Printf("UpdateTID id:tid is %v:%v \n",id,tid)
	sql := "update jdcron_tasks set task_id = ? ,status = ? where id= ?"
	engine.SQL(sql,tid,status,id).Execute()
}

func (j JdcronTasks) UpdateTaskStatus(taskId int,status int) {
	sql := "update jdcron_tasks set status = ? where task_id = ?"
	engine.Exec(sql,status,taskId)
}

func (j JdcronTasks) Add(name,spec,cmd,handler string) {
	sql := "insert jdcron_tasks (name,spec,command,handler,create_time) values (?,?,?,?,?)"
	_, err := engine.SQL(sql, name, spec, cmd,handler,time.Now()).Execute()
	if err != nil {
		log.Println("jdcron_tasks sql insert err:",err)
	}
}
// ------
// job cron interface function
func (j *JdcronTasks) Run() {
	DoJob(j.Command,j.Handler)
}

// interface run job
func DoJob(cmd,handler string,)  {
	if handler == "shall" {
		execShall(cmd)
	}else {
		callHTTP(cmd)
	}
}

func callHTTP(url string)  {
	log.Println("JdcronTasks Run do job's cmd is :",url)
	response, err := http.Get(url)

	if err != nil {
		log.Println("do job interface err: ",err)
	}

	fmt.Printf("doJob response : %v\n",response.Body)
}

func execShall(command string)  {
	var cmd *exec.Cmd
	cmd = exec.Command(command)
	_, err := cmd.Output()
	if err != nil {
		log.Printf("命令行执行错误：cmd:%v \nerr:%v", command,err)
	}
}