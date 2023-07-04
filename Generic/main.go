// 泛型
package main

import (
	_case "go-study/Generic/case"
)

func main() {
	_case.SimpleCase()
}

/*
go中没有项目的说法，只有包
GOROOT：GOROOT就是Go的安装目录
GOPATH：GOPATH是我们的工作空间,保存go项目代码和第三方依赖包
1.保存编译后的二进制文件。
2.go get和go install命令会下载go代码到GOPATH。
3.import包时的搜索路径
*/
// gomod,work
