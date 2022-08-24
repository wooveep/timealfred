/*
 * @Author: cloudyi.li
 * @Date: 2022-08-24 07:06:24
 * @LastEditTime: 2022-08-24 07:12:05
 * @LastEditors: cloudyi.li
 * @FilePath: /timealfred/internal/times/main.go
 */
package times

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var WeekDayMap = map[string]string{
	"Monday":    "周一",
	"Tuesday":   "周二",
	"Wednesday": "周三",
	"Thursday":  "周四",
	"Friday":    "周五",
	"Saturday":  "周六",
	"Sunday":    "周日",
}

type Result struct {
	Name  string
	Times string
}

func Times(t time.Time) (alfreddatas []map[string]string) {
	var tmplist []Result
	tmplist = append(tmplist, Result{"日期时间RFC3339", t.Local().Format("2006-01-02T15:04:05Z07:00")},
		Result{"日期时间", t.Format("2006-01-02 15:04:05")},
		Result{"日期时间毫秒", t.Format("2006-01-02 15:04:05.000000")},
		Result{"日期中文", fmt.Sprintf("%d年%d月%d日 %s", t.Year(), t.Month(), t.Day(), WeekDayMap[t.Weekday().String()])},
		Result{"日期时间中文", fmt.Sprintf("%d年%d月%d日%d时%d分%d秒", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())},
		Result{"时间戳", strconv.FormatInt(t.Unix(), 10)},
		Result{"毫秒时间戳", strconv.FormatInt(t.UnixNano()/1e6, 10)})

	for _, v := range tmplist {
		result := make(map[string]string)
		result["uid"] = uuid.NewString()
		result["arg"] = v.Times
		result["title"] = v.Name
		result["subtitle"] = v.Times
		alfreddatas = append(alfreddatas, result)
	}
	return
}
