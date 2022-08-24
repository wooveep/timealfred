/*
 * @Author: cloudyi.li
 * @Date: 2022-08-23 19:18:14
 * @LastEditTime: 2022-08-24 10:59:58
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
	t, err := time.Parse("2006-01-02 15:04:05", args)
	if err != nil {
		t, err = time.Parse("20060102150405", args)
		if err != nil {
			t, err = time.Parse("2006-01-02T15:04:05Z07:00", args)
			if err != nil {
				tsu, err := strconv.ParseInt(args, 10, 64)
				if len(args) == 10 {
					t = time.Unix(tsu, 0)
				} else if len(args) == 13 {
					t = time.UnixMilli(tsu)
				} else {
					t = time.Now()
				}
				if err != nil {
					t = time.Now()
				}
			}
		}
	}

	output := alfred.OutputXml(times.Times(t))
	fmt.Println(output)
}
