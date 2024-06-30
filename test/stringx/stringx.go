package main

import (
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stringx"
)

func main() {
	orderFieldNames := []string {
		"`id`",
		"`create_at`",
		// "测试1",
		// "测试2",
		// "测试3",
	}
	orderRowsExpectAutoSet := strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	orderRowsWithPlaceHolder := strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
	fmt.Println(orderRowsExpectAutoSet)
	fmt.Println(orderRowsWithPlaceHolder)
}
