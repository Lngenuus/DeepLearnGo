package main

import (
	"flag"
	"fmt"

	"githu.com/Lngenuus/DeepLearnGO/src/geek/class2"
)

func geekOption(geek string) {
	switch geek {
	case "class2":
		class2.Run()
	default:
		fmt.Printf("选择执行单元[%s]不存在\n", geek)
	}
}

func main() {
	var (
		geek string
		test string
	)
	flag.StringVar(&geek, "geek", "", "选择执行src/geek下属模块,默认为空")
	flag.StringVar(&test, "test", "", "占位")
	flag.Parse()

	if len(geek) > 0 {
		geekOption(geek)
	}
}
