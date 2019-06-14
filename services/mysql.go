package services

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Engine *xorm.Engine
	Err error
)

func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func InitMysql() {

	cfg,Err := ini.Load("./conf/database.ini")
	if Err != nil {
		log.Fatal(Err.Error())
	}
	Engine, Err = xorm.NewEngine(cfg.Section("mysql").Key("DbType").String(), fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		cfg.Section("mysql").Key("DbUser").String(),
		cfg.Section("mysql").Key("DbPassword").String(),
		cfg.Section("mysql").Key("DbHost").String(),
		cfg.Section("mysql").Key("DbPort").String(),
		cfg.Section("mysql").Key("DbName").String(),
	))
	// f, err := os.Create(config.Log.Sqllog)
	// if err != nil {
	//  println(err.Error())
	//  return
	// }
	// engine.SetLogger(xorm.NewSimpleLogger(f))
	Engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai") //上海时区

	if Err = Engine.Ping(); Err != nil { //数据库 ping
		log.Fatalf("database ping err: %s", Err.Error())
	}

	if cfg.Section("ServerConfig").Key("ServerRunEnv").String() == "master" { //根据运行环境判断日志模式
		Engine.ShowSQL(false)
	} else {
		Engine.ShowSQL(true)
	}
	Engine.ShowSQL(true)
	MaxIdleConn,Err := cfg.Section("mysql").Key("DbMaxIdleConn").Int()
	if Err != nil {
		log.Fatalf("DbMaxIdleConn Get err: %s", Err.Error())
	}
	Engine.SetMaxIdleConns(MaxIdleConn) //连接池的空闲数大小

	MaxOpenConn,Err := cfg.Section("mysql").Key("DbMaxOpenConn").Int()
	if Err != nil {
		log.Fatalf("DbMaxOpenConn Get err: %s", Err.Error())
	}
	Engine.SetMaxOpenConns(MaxOpenConn) //最大打开连接数

	timer := time.NewTicker(time.Minute * 30) //定时器 ping
	go func(x *xorm.Engine) {
		for range timer.C {
			Err = x.Ping()
			if Err != nil {
				log.Fatalf("数据库连接错误: %#v\n", Err.Error())
			}
		}
	}(Engine)

}








