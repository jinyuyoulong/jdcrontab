// file: controller/user_controller.go

package controller

import (
	"context"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"github.com/jinyuyoulong/jdcrontab/src/bootstrap/service"
	"github.com/jinyuyoulong/jdcrontab/src/library/crontab"
	"github.com/jinyuyoulong/jdcrontab/src/library/rpc"
	"github.com/jinyuyoulong/jdcrontab/src/models"
	"log"
	"os"
	"strconv"
	"time"
)

// IndexController index controller
type IndexController struct {
	Ctx iris.Context
	Di  iris.Context
}

// Get url: /
func (c *IndexController) Get() mvc.Result {

	datalist := models.JdcronTasks{}.GetAll()

	return mvc.View{
		Name: "index/index.html",
		Data: iris.Map{
			"Title":    commonTitle,
			"Datalist": datalist,
			"AppOwner": service.AppConfig().App.AppOwner,
		},
	}
}

func (c *IndexController)GetInfo() mvc.Result  {
	return mvc.View{
		Name:"index/info.html",
		Data:iris.Map{
			"Title": commonTitle,
			"Data": "",
			"AppOwner": service.AppConfig().App.AppOwner,
		},
	}
}

// url: /dojob
func (c *IndexController)GetDojob()  {
	jobId, _ := c.Ctx.URLParamInt("jobId")
	c.Ctx.Writef("jobId = %v\n",jobId)
	jobModel := models.JdcronTasks{}.GetBy(jobId)

	if jobModel.Id != 0 {

		c.Ctx.Writef(fmt.Sprintf("job model is %v \n", jobModel))
		c.Ctx.Writef(fmt.Sprintf("job action is %v\n", jobModel.Command))

		if jobModel.Command != "" {

			go models.DoJob(jobModel.Command,jobModel.Handler)
		}else {
			log.Println("command 为空！")
		}
	}else {
		c.Ctx.Writef("taskID 参数值错误")
	}
}

func startJob(task *models.JdcronTasks)  {

	log.Printf("start job name \"%v\"\n",task.Name)
	spec := task.Spec
	entryID := crontab.CycleJob(spec, task)
	log.Println("entry id is ",entryID)

	models.JdcronTasks{}.UpdateTID(task.Id, strconv.FormatInt(int64(entryID),10),1)

}

func (c *IndexController) GetStartjobs() {
	tasks := models.JdcronTasks{}.GetAll()
	//wg := sync.WaitGroup{}
	//wg.Add(1)

	for _,task := range tasks{

		ttask := task
		// 开启任务
		log.Println("start 子协程任务 ")
		go startJob(&ttask)
		log.Println("结束任务")
	}

	//wg.Wait()

	c.Ctx.Writef("start jobs")
}


func (c *IndexController) GetRemovejob() {
	tid, err := c.Ctx.URLParamInt("taskID")
	if err != nil{
		fmt.Println(err)
	}
	log.Println("RemovejobBy")
	removeJob(tid)
	c.Ctx.Writef("终止 task id=",tid)
}

func removeJob(taskID int)  {
	crontab.RemoveJob(cron.EntryID(taskID))
	models.JdcronTasks{}.UpdateTaskStatus(taskID,0)
}

func (c *IndexController)GetRemovealljobs()  {
	tasks := models.JdcronTasks{}.GetAll()
	for _, task := range tasks {
		ttask := task
		if task.Status == 1 {
			removeJob(ttask.TaskId)
		}
	}
	c.Ctx.Writef("终止完毕！")
}

// cron test job
func (c *IndexController) GetGetlog() *APIJson  {
	return APIResult(0,"haha","")
}

// RPC——
// /rpc
// rpc 通过 http 调用启动，或通过tcp调用响应（暂时没有头绪）
func (c *IndexController)GetRpc() {
	command := c.Ctx.URLParam("command")
	if command == "" {
		command = "localhost:8000"
	}
	log.Printf("rpc command is :%v\n",command)
	//c.Ctx.Writef("%v",APIResult(0,fmt.Sprintf("GetRpc command is : ",command),"nil"))

	connectRPC(c.Ctx,command)

}

func connectRPC(ictx iris.Context, command string)  {
	//address := "http://0.0.0.0:8100"
	defaultName := "world!"
	// 建立连接
	conn, err := grpc.Dial(command,grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc connction %v", err)
		return
	}
	defer conn.Close()

	c := rpc.NewGrpcServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 获取 rpc 返回的数据
	r, err := c.GRPCResponse(ctx, &rpc.GRPCRequest{Name:name})
	if err != nil {
		//log.Fatal("can not get :%v", err)
		log.Printf("connectRPC can not get %v\n",err)
		return
	}
	log.Printf("greeting code:%v msg:%v data:%v\n",r.Code, r.Message,r.Data)

	ictx.Writef("greeting code:%v msg:%v data:%v\n",r.Code, r.Message,r.Data)

}

func (c *IndexController)GetAdd() mvc.Result {
	return mvc.View{
		Name: "index/add.html",
		Data:iris.Map{
			"Title": "添加job",
		},

	}
}
func (c *IndexController) PostAdd() {
	//for i,value := range c.Ctx.PostValues() {
	//
	//}
	name := c.Ctx.PostValue("name")
	spec := c.Ctx.PostValueTrim("spec")
	command := c.Ctx.PostValue("command")
	handler := c.Ctx.PostValue("handler")
	addJob(name,spec,command,handler)
	json, err := c.Ctx.JSON(APIResult(0, "添加成功!", ""))
	if err != nil{
		log.Printf("post add json result:\n%v \nerr:\n%v",json,err)
	}
}
func addJob(name,spec,cmd,handler string)  {
	models.JdcronTasks{}.Add(name,spec,cmd,handler)
}