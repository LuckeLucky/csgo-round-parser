package main

import (
	"fmt"
	"os"

	"github.com/LuckeLucky/demo-analyser-csgo/analyser"

	"github.com/LuckeLucky/demo-analyser-csgo/utils"
)

func main() {
	fmt.Println(os.Args[0])

	f, err := os.Open(os.Args[1])
	utils.CheckError(err)
	defer f.Close()

	fmt.Printf("Analyzing file: %s\n", f.Name())
	an := analyser.NewAnalyser(f)
	an.SetDefaultConvarConfig()
	an.SimpleRun()
	fmt.Printf("Finished file: %s\n\n", f.Name())
	f.Close()

	fmt.Scanf("oi")
}
