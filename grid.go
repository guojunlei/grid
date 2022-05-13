package main

import (
	"fmt"
	"grid/model"
)

func main() {

	// 	// 计算生成网格信息
	gridinfo := model.ReadGrid()
	fmt.Println(*gridinfo)

}
