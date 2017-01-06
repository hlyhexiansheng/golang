package main

import (
	"fmt"
	"github.com/elog-falcon/agent/funcs"
)

func main() {
	fmt.Println("hello,world1112224")

	funcs.UpdateCpuStat()
	funcs.UpdateCpuStat()


	items := funcs.CpuMetrics();

	for _, mv := range items {
		fmt.Println(mv.String());
	}
}
