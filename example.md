### 使用

#### 引入包

~~~go
import (
    "github.com/deatil/go-datebin/datebin"
)
~~~


#### 固定时间使用

~~~go
// 时区可不设置
timezone := datebin.UTC

// 固定时间
date := datebin.
    Now(timezone).
    // Now().
    // Today(timezone).
    // Tomorrow(timezone).
    // Yesterday(timezone).
    ToDatetimeString()
~~~


#### 设置时间

~~~go
import (
    "time"
)

// 时区可不设置
timezone := datebin.UTC

// 固定时间
date := datebin.
    FromTimeTime(time.Now(), timezone).
    // FromTimeUnix(int64(1652587697), int64(0), timezone).
    // FromTimestamp(int64(1652587697), timezone.
    // FromDatetimeWithNanosecond(2022, 10, 23, 22, 18, 56, 123, timezone).
    // FromDatetimeWithMicrosecond(2022, 10, 23, 22, 18, 56, 123, timezone).
    // FromDatetimeWithMillisecond(2022, 10, 23, 22, 18, 56, 123, timezone).
    // FromDatetime(2022, 10, 23, 22, 18, 56, timezone).
    // FromDate(2022, 10, 23, timezone).
    // FromTime(22, 18, 56, timezone).
    ToDatetimeString()
~~~


#### 解析使用

~~~go
// 时区可不设置
timezone := datebin.UTC

// 固定时间
date := datebin.
    Parse("2022-10-23 22:18:56").
    // ParseWithLayout("2022-10-23 22:18:56", "2006-01-02 15:04:05", timezone).
    // ParseWithFormat("2022-10-23 22:18:56", "Y-m-d H:i:s", timezone.
    // ParseDatetimeString("2022-10-23 22:18:56", "2006-01-02 15:04:05").
    // ParseDatetimeString("2022-10-23 22:18:56", "Y-m-d H:i:s", "u").
    ToDatetimeString()
~~~


#### 数据输出

~~~go
// 时区可不设置
timezone := datebin.UTC

// 固定时间
date := datebin.
    Parse("2022-10-23 22:18:56").
    // String()
    // ToString(timezone) # 返回字符
    // ToStarString(timezone) # 返回星座名称
    // ToSeasonString(timezone) # 返回当前季节，以气象划分
    // ToWeekdayString(timezone) # 周几
    // Layout("2006-01-02 15:04:05", timezone) # 原始格式
    // ToLayoutString("2006-01-02 15:04:05", timezone) # 原始格式
    // Format("Y-m-d H:i:s", timezone) # 输出指定布局的时间字符串
    // ToFormatString("Y-m-d H:i:s", timezone) # 输出指定布局的时间字符串
    // ToAnsicString()
    // ToUnixDateString()
    // ToRubyDateString()
    // ToRFC822String()
    // ToRFC822ZString()
    // ToRFC850String()
    // ToRFC1123String()
    // ToRFC1123ZString()
    // ToRssString()
    // ToRFC2822String()
    // ToRFC3339String()
    // ToKitchenString()
    // ToCookieString()
    // ToISO8601String()
    // ToRFC1036String()
    // ToRFC7231String()
    // ToAtomString()
    // ToW3CString()
    // ToDayDateTimeString()
    // ToFormattedDateString()
    // ToDatetimeString()
    // ToDateString()
    // ToTimeString()
    // ToShortDatetimeString()
    // ToShortDateString()
    // ToShortTimeString()
    ToDatetimeString()
~~~


#### 快捷方式

~~~go
// 时区可不设置
timezone := datebin.UTC

// 当前时间，单位：秒, int64
data := datebin.NowTime(timezone)

// 当前日期时间字符, string
data := datebin.NowDatetimeString(timezone)

// 当前日期, string
data := datebin.NowDateString(timezone)

// 当前时间字符, string
data := datebin.NowTimeString(timezone)

// 时间戳转为 time.Time
timeData := datebin.TimestampToTime(int64(1652587697), timezone)

// 时间转换为时间戳
timestampData := datebin.TimeToTimestamp(timeData, timezone)

// 时间字符转为时间
timeData2 := datebin.StringToTime("2022-10-23 22:18:56", "2006-01-02 15:04:05")
timeData2 := datebin.StringToTime("2022-10-23 22:18:56", "Y-m-d H:i:s", "u")

// 时间字符转为时间戳
timestampData2 := datebin.StringToTimestamp("2022-10-23 22:18:56", "2006-01-02 15:04:05")
timestampData2 := datebin.StringToTimestamp("2022-10-23 22:18:56", "Y-m-d H:i:s", "u")
~~~
