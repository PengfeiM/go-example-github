// Package cmd
/*
Copyright © 2024 MaPengfei <PengfeiM@outlook.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"go-example/pkg/goexpl"
	"os"
	"reflect"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run a go example",
	Long:  `specify a go example to run`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("run called %s\n", args[0])
		if len(args) < 1 {
			fmt.Println("no example specified. exit")
			os.Exit(1)
		}
		goExampleName := args[0]
		goExample, ok := goExampleMap[goExampleName]
		if !ok {
			fmt.Println("Error: not supported Go Example: ", goExampleName)
			fmt.Println("you may use list(l) to look up the go examples can be run")
		} else {
			inputParams := goexpl.NewInputParams(
				len(args)-1,
				args[1:],
				os.File{}, // 尚未使用，预留字段
			)
			fmt.Printf("# Run Example [%s]...\n", reflect.TypeOf(goExample).Elem().Name())
			fmt.Printf("- Input:\n%v\n", *inputParams)
			fmt.Printf("- Output:\n")
			fmt.Printf("--------------------------------\n")
			if err := goExample.RunExample(inputParams); err != nil {
				fmt.Println("Run example failed! ", err)
			}
			fmt.Printf("--------------------------------\n")
			fmt.Printf("# Example [%s] Over! \n", reflect.TypeOf(goExample).Elem().Name())
		}
	},
	Aliases: []string{"r"},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
