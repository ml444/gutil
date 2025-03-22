package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func findDuplicateStructsInPackage(packagePath string) (map[string]int, error) {
	structCounter := make(map[string]int)

	// 获取包下的所有Go文件
	filepaths, err := getGoFilesInPackage(packagePath)
	if err != nil {
		return nil, err
	}

	// 遍历每个文件并解析AST
	for _, filePath := range filepaths {
		fset := token.NewFileSet()

		// 解析源文件
		f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		// 使用visitor遍历AST
		ast.Inspect(f, func(n ast.Node) bool {
			// 检查是否是结构体声明
			if typeSpec, ok := n.(*ast.TypeSpec); ok {
				if _, ok := typeSpec.Type.(*ast.StructType); ok {
					// 获取结构体的名称
					structName := typeSpec.Name.Name
					// 记录结构体的出现次数
					structCounter[structName]++
					// 如果出现次数大于1，表示重复声明
					if structCounter[structName] > 1 {
						fmt.Printf("重复声明的结构体：%s 在文件 %s\n", structName, filePath)
					}
					//fmt.Println(structType)
				}
			}
			return true
		})
	}

	return structCounter, nil
}

func getGoFilesInPackage(packagePath string) ([]string, error) {
	var filepaths []string

	err := filepath.Walk(packagePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".go") {
			filepaths = append(filepaths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filepaths, nil
}

func main() {
	//if len(os.Args) != 2 {
	//	fmt.Println("Usage: go run main.go <package_path>")
	//	os.Exit(1)
	//}
	//
	//packagePath := os.Args[1]
	packagePath := "/Users/cml/github.com/ml444/gkit/cmd/protoc-gen-go-validate/tests/go/cases"
	structCounter, err := findDuplicateStructsInPackage(packagePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("结构体声明次数：")
	for structName, count := range structCounter {
		fmt.Printf("%s: %d\n", structName, count)
	}
}
