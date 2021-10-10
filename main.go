package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/LuckeLucky/demo-analyser-csgo/analyser"
)

func main() {
	err := filepath.Walk("demos/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			if filepath.Ext(path) != ".dem" {
				fmt.Println("Ignoring file: " + path)
				return nil
			}

			f, err := os.Open(path)
			if err != nil {
				panic(err)
			}

			fmt.Println(path, info.Size())

			defer f.Close()

			fmt.Printf("Analyzing file: %s\n", f.Name())
			a := analyser.NewAnalyser(f)
			a.Run()
			fmt.Printf("Finished file: %s\n\n", f.Name())

			return nil
		})
	if err != nil {
		panic(err)
	}
}
