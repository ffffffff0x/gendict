package main

import (
	"flag"
	"fmt"
	"gendict/mod"
	"gendict/utils"
	"os"
)

func main() {
	flag.Parse()

	if mod.V {
		fmt.Println(mod.Version)
		os.Exit(0) // 输出版本后停止
	}

	utils.OutputProcess()
}
