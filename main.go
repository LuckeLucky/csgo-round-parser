package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/LuckeLucky/demo-analyser-csgo/analyser"

	"github.com/LuckeLucky/demo-analyser-csgo/utils"
)

func init() {
	utils.ReadConfigFile()
}

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
			utils.CheckError(err)
			defer f.Close()

			fmt.Printf("Analyzing file: %s\n", f.Name())
			an := analyser.NewAnalyser(f)
			an.SimpleRun()
			fmt.Printf("Finished file: %s\n\n", f.Name())
			f.Close()
			/*newName := an.GetDemoNameWithDetails()
			err = os.Rename(path, "analysed-demos/"+newName+".dem")
			if err != nil {
				fmt.Printf("error renaming file: %s", err)
			}*/

			return nil
		})
	if err != nil {
		panic(err)
	}
}
