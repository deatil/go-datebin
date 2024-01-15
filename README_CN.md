## go-datebin


### 项目介绍

*  go-datebin 是一个简单易用的 go 时间处理库

中文 | [English](README.md)


### 下载安装

~~~go
go get -u github.com/deatil/go-datebin
~~~


### 开始使用

~~~go
package main

import (
    "fmt"

    "github.com/deatil/go-datebin/datebin"
)

func main() {
    // 当前时间
    date := datebin.
        Now().
        ToDatetimeString()
    // 输出: 2024-01-06 12:06:12

    // 解析时间，不带时区
    date2 := datebin.
        Parse("2032-03-15 12:06:17").
        ToDatetimeString(datebin.UTC)
    // 输出: 2032-03-15 12:06:17

    // 解析时间，带时区
    date2 := datebin.
        ParseWithLayout("2032-03-15 12:06:17", datebin.DatetimeFormat, datebin.GMT).
        ToDatetimeString()
    // 输出: 2032-03-15 12:06:17

    // 设置时间并输出格式化时间
    date3 := datebin.
        FromDatetime(2032, 3, 15, 12, 56, 5).
        ToFormatString("Y/m/d H:i:s")
    // 输出: 2032/03/15 12:56:05
}

~~~


### 常用示例

~~~go
// 格式化时间戳
var datetimeString string = datebin.FromTimestamp(1705329727, datebin.Shanghai).ToDatetimeString()
// 输出: 2024-01-15 22:42:07
~~~

~~~go
// 获取当前时间戳
var timestamp int64 = datebin.Now().Timestamp()
// 输出: 1705329727
~~~

更多示例可点击查看 [使用文档](example.md)


### 开源协议

*  本软件包遵循 `Apache2` 开源协议发布，在保留本软件包版权的情况下提供个人及商业免费使用。


### 版权

*  本软件包所属版权归 deatil(https://github.com/deatil) 所有。