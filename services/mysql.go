package services

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
	"fmt"
	"github.com/go-ini/ini"
	"log"
)

var (
	engine *xorm.Engine
	err error
)

func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func InitMysql() {

	cfg,err := ini.Load("/conf/database.ini")
	if err != nil {
		log.Fatal(err.Error())
	}
	engine, err = xorm.NewEngine(cfg.Section("mysql").Key("DbType").String(), fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True",
		cfg.Section("mysql").Key("DbUser").String(),
		cfg.Section("mysql").Key("DbPassword").String(),
		cfg.Section("mysql").Key("DbHost").String(),
		cfg.Section("mysql").Key("DbName").String(),
	))
	// f, err := os.Create(config.Log.Sqllog)
	// if err != nil {
	//  println(err.Error())
	//  return
	// }
	// engine.SetLogger(xorm.NewSimpleLogger(f))
	/*engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai") //上海时区

	if err = engine.Ping(); err != nil { //数据库 ping
		log.Fatalf("database ping err: %s", err.Error())
	}

	if config.ServerConfig.ServerRunEnv == "master" { //根据运行环境判断日志模式
		engine.ShowSQL(false)
	} else {
		engine.ShowSQL(true)
	}

	engine.SetMaxIdleConns(config.DbMysqlConfig.DbMaxIdleConn) //连接池的空闲数大小
	engine.SetMaxOpenConns(config.DbMysqlConfig.DbMaxOpenConn) //最大打开连接数

	timer := time.NewTicker(time.Minute * 30) //定时器 ping
	go func(x *xorm.Engine) {
		for range timer.C {
			err = x.Ping()
			if err != nil {
				log.Fatalf("数据库连接错误: %#v\n", err.Error())
			}
		}
	}(engine)*/

}








