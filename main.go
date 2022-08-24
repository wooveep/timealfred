/*
 * @Author: cloudyi.li
 * @Date: 2022-08-23 19:18:14
 * @LastEditTime: 2022-08-24 07:39:08
 * @LastEditors: cloudyi.li
 * @FilePath: /timealfred/main.go
 */
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"timealfred/internal/alfred"
	"timealfred/internal/times"
)

func main() {
	var t time.Time
	args := strings.Join(os.Args[1:], " ")
	tsu, _ := strconv.ParseInt(args, 10, 64)
	// // fmt.Println(strings.Join(args, " "))
	if len(args) == 10 {
		t = time.Unix(tsu, 0)
	} else if len(args) == 13 {
		t = time.UnixMilli(tsu)
	} else {
		t = time.Now()
	}
	//
	// args := "1661297475"

	// t := time.Now()
	// fmt.Println(t.Unix())
	// fmt.Println(strconv.FormatInt(t.Unix(), 10))
	fmt.Printf("%d年%d月%d日%d时%d分%d秒", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	output := alfred.OutputXml(times.Times(t))
	fmt.Println(output)
}
