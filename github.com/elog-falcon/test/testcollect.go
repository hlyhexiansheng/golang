package main

import (
	"github.com/elog-falcon/agent/funcs"
	"github.com/elog-falcon/agent/cron"
)

func main() {


	funcs.BuildMappers() //初始化各种监控点

	go cron.InitDataHistory()  //先初始化需要历史记录的metric点.

	cron.Collect()  //执行收集.

	select {}
}
