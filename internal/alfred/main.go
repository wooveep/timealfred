/*
 * @Author: cloudyi.li
 * @Date: 2021-12-22 16:20:28
 * @LastEditTime: 2022-08-24 07:06:34
 * @LastEditors: cloudyi.li
 * @FilePath: /timealfred/internal/alfred/main.go
 */
package alfred

import (
	"encoding/xml"
	"fmt"
)

type Items struct {
	XMLName xml.Name `xml:"items"`
	Item    []Item   `xml:"item"`
}
type Item struct {
	Uid      string `xml:"uid,attr"`
	Arg      string `xml:"arg,attr"`
	Title    string `xml:"title"`
	SubTitle string `xml:"subtitle"`
	Icon     string `xml:"icon"`
}

func OutputXml(item_list []map[string]string) (resultstring string) {
	var items Items
	item_count := len(item_list)
	items.Item = make([]Item, item_count)
	item := items.Item
	for i := 0; i < item_count; i++ {
		m := &item[i]
		m.Arg = item_list[i]["arg"]
		m.Uid = item_list[i]["uid"]
		m.Title = item_list[i]["title"]
		m.SubTitle = item_list[i]["subtitle"]
		m.Icon = "icon"
	}

	data, err := xml.MarshalIndent(&items, " ", "\t")
	if err != nil {
		fmt.Print(err)
	}
	resultstring = fmt.Sprintf("%s%s", xml.Header, string(data))
	return resultstring
}
