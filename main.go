package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/Lngenuus/DeepLearnGo/src/geek/class2"
	"github.com/Lngenuus/DeepLearnGo/src/geek/class3"
)

func geekOption(geek string) {
	switch geek {
	case "class2":
		class2.Run()
	case "class3":
		class3.NewApp()
	case "class4":
		cmd := exec.Command("kratos", "run", "src/geek/class4/coin/cmd/coin")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
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
