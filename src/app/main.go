package main

import (
	"algrithom"
	"os"
	"runtime/pprof"
)

func main() {
	f, _ := os.Create("profile_file")
	pprof.StartCPUProfile(f)     // 开始cpu profile，结果写到文件f中
	defer pprof.StopCPUProfile() // 结束profile

	algrithom.ForceMethod(100000)
}
