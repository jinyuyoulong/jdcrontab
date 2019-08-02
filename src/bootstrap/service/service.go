package service

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	"github.com/kataras/iris/view"
	"github.com/xormplus/xorm"

	// "github.com/pelletier/go-toml"
	"github.com/BurntSushi/toml"
	"go.uber.org/dig"
)

type Di struct {
	Container *dig.Container
}

var di *Di

// GetDi get
func GetDi() *Di {
	if di == nil {
		di = &Di{
			Container: dig.New(),
		}
	}
	return di
}

var conf *Config

func AppConfig() *Config {
	if conf == nil {
		conf = new(Config)
		file := "../config/config.toml"
		_, err := toml.DecodeFile(file, conf)
		if err != nil {
			fmt.Println("Toml Error!", err.Error())
		}
	}

	return conf
}

func viewEngine() *view.HTMLEngine {
	viewPath := "../view"
	layoutPath := "layout/layout.html"

	var htmlEngine *view.HTMLEngine
	htmlEngine = iris.HTML(viewPath, ".html").Layout(layoutPath)

	return htmlEngine.Reload(true)
}

func db() *xorm.Engine {
	//  读取配置文件的数据
	tomlC := AppConfig()
	driver := tomlC.Database.Dirver
	configTree := tomlC.Mysql
	userName := configTree.Username
	password := configTree.Password
	dbname := configTree.Dbname
	connet := fmt.Sprintf("%s:%s%s", userName, password, dbname)
	engine, err := xorm.NewEngine(driver, connet)
	if err != nil {
		log.Fatal("database connet failed : %s", err)
	}
	return engine
}

func createSessions() *sessions.Sessions {

	sessionConf := sessions.Config{
		Cookie:  conf.Session.Cookie,
		Expires: conf.Session.Expires * time.Minute,
	}
	sess := sessions.New(sessionConf)

	if conf.Session.Dirver == "redis" {
		timeout := conf.Redis.IdleTimeout
		if timeout == 0 {
			timeout = service.DefaultRedisIdleTimeout
		} else {
			timeout = conf.Redis.IdleTimeout
		}
		redisConf := service.Config{
			Network:     conf.Redis.Network,
			Addr:        conf.Redis.Addr,
			Password:    conf.Redis.Password,
			Database:    conf.Redis.Database,
			MaxIdle:     conf.Redis.MaxIdle,
			MaxActive:   conf.Redis.MaxActive,
			IdleTimeout: timeout,
			Prefix:      conf.Redis.Prefix,
		}
		db := redis.New(redisConf) // optionally configure the bridge between your redis server

		// close connection when control+C/cmd+C
		iris.RegisterOnInterrupt(func() {
			db.Close()
		})
		sess.UseDatabase(db)
	} else {
		hero.Register(sess.Start)
	}

	return sess
}

func init() {
	BuildContainer()
}
// BuildContainer 容器创建&注入
func BuildContainer() {
	container := GetDi().Container

	container.Provide(AppConfig)
	container.Provide(viewEngine)
	container.Provide(db)
	container.Provide(createSessions)
}
