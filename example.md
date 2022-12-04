### 使用

常用的一些使用示例，更多示例可以点击 [文档](pkg.go.dev/github.com/deatil/go-datebin) 查看


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

// 全局设置时区，对使用帮助函数有效
datebin.SetTimezone(timezone)

// 固定时间
date := datebin.
    Now().
    // Now(timezone).
    // Today(timezone).
    // Tomorrow(timezone).
    // Yesterday(timezone).
    ToDatetimeString()
~~~


#### 根据具体时间进行格式化

~~~go
import (
    "time"
)

// 时区可不设置
timezone := datebin.UTC

// 添加时间
date := datebin.
    FromTimeTime(time.Now(), timezone).
    // FromTimeUnix(int64(1652587697), int64(0), timezone).
    // FromTimestamp(int64(1652587697), timezone).
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

// 解析时间
date := datebin.
    Parse("2022-10-23 22:18:56").
    // ParseWithLayout("2022-10-23 22:18:56", "2006-01-02 15:04:05", timezone).
    // ParseWithFormat("2022-10-23 22:18:56", "Y-m-d H:i:s", timezone).
    // ParseDatetimeString("2022-10-23 22:18:56", "2006-01-02 15:04:05").
    // ParseDatetimeString("2022-10-23 22:18:56", "Y-m-d H:i:s", "u").
    ToDatetimeString()
~~~


#### 数据输出

~~~go
// 时区可不设置
timezone := datebin.UTC

// 数据输出
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

// 当前日期字符, string
data := datebin.NowDateString(timezone)

// 当前时间字符, string
data := datebin.NowTimeString(timezone)

// 时间戳转为 time.Time
timeData := datebin.TimestampToTime(int64(1652587697), timezone)

// time.Time 转换为时间戳
timestampData := datebin.TimeToTimestamp(timeData, timezone)

// 时间字符转为时间
timeData2 := datebin.StringToTime("2022-10-23 22:18:56", "2006-01-02 15:04:05")
timeData2 := datebin.StringToTime("2022-10-23 22:18:56", "Y-m-d H:i:s", "u")

// 时间字符转为时间戳
timestampData2 := datebin.StringToTimestamp("2022-10-23 22:18:56", "2006-01-02 15:04:05")
timestampData2 := datebin.StringToTimestamp("2022-10-23 22:18:56", "Y-m-d H:i:s", "u")
~~~


#### 比较时间

~~~go
// 准备时间
timeA := datebin.Parse("2022-10-23 22:18:56")
timeB := datebin.Parse("2022-10-25 23:18:56")
timeC := datebin.Parse("2022-10-26 23:18:56")

// timeA 大于 timeB
res := timeA.Gt(timeB)

// timeA 小于 timeB
res := timeA.Lt(timeB)

// timeA 等于 timeB
res := timeA.Eq(timeB)

// timeA 不等于 timeB
res := timeA.Ne(timeB)

// timeA 大于等于 timeB
res := timeA.Gte(timeB)

// timeA 小于等于 timeB
res := timeA.Lte(timeB)

// timeA 是否在两个时间之间(不包括这两个时间)
res := timeA.Between(timeB, timeC)

// timeA 是否在两个时间之间(包括这两个时间)
res := timeA.BetweenIncluded(timeB, timeC)

// timeA 是否在两个时间之间(包括开始时间)
res := timeA.BetweenIncludStart(timeB, timeC)

// timeA 是否在两个时间之间(包括结束时间)
res := timeA.BetweenIncludEnd(timeB, timeC)

// 最小值
res := timeA.Min(timeB)
res := timeA.Minimum(timeB)

// 最大值
res := timeA.Max(timeB)
res := timeA.Maximum(timeB)

// 平均值
res := timeA.Avg(timeB)
res := timeA.Average(timeB)

// 取 a 和 b 中与当前时间最近的一个
res := timeA.Closest(timeB, timeC)

// 取 a 和 b 中与当前时间最远的一个
res := timeA.Farthest(timeB, timeC)

// 年龄，可为负数
res := timeA.Age()

// 用于查找将规定的持续时间 'd' 舍入为 'm' 持续时间的最接近倍数的结果
res := timeA.Round(d time.Duration)

// 用于查找将规定的持续时间 'd' 朝零舍入到 'm' 持续时间的倍数的结果
res := timeA.Truncate(d time.Duration)
~~~


#### 获取时间

~~~go
// 准备时间
time := datebin.Parse("2022-10-23 22:18:56")

// 获取当前世纪
res := time.Century()

// 获取当前年代
res := time.Decade()

// 获取当前年
res := time.Year()

// 获取当前季度
res := time.Quarter()

// 获取当前月
res := time.Month()

// 星期几数字
res := time.Weekday()

// 获取当前日
res := time.Day()

// 获取当前小时
res := time.Hour()

// 获取当前分钟数
res := time.Minute()

// 获取当前秒数
res := time.Second()

// 获取当前毫秒数，范围[0, 999]
res := time.Millisecond()

// 获取当前微秒数，范围[0, 999999]
res := time.Microsecond()

// 获取当前纳秒数，范围[0, 999999999]
res := time.Nanosecond()

// 秒级时间戳，10位
res := time.Timestamp()
res := time.TimestampWithSecond()

// 毫秒级时间戳，13位
res := time.TimestampWithMillisecond()

// 微秒级时间戳，16位
res := time.TimestampWithMicrosecond()

// 纳秒级时间戳，19位
res := time.TimestampWithNanosecond()

// 返回年月日数据
year, month, day := time.Date()

// 返回时分秒数据
hour, minute, second := time.Time()

// 返回年月日时分秒数据
year, month, day, hour, minute, second := time.Datetime()

// 返回年月日时分秒数据带纳秒
year, month, day, hour, minute, second, nanosecond := time.DatetimeWithNanosecond()

// 返回年月日时分秒数据带微秒
year, month, day, hour, minute, second, microsecond := time.DatetimeWithMicrosecond()

// 返回年月日时分秒数据带毫秒
year, month, day, hour, minute, second, millisecond := time.DatetimeWithMillisecond()

// 获取本月的总天数
res := time.DaysInMonth()

// 获取本年的第几月
res := time.MonthOfYear()

// 获取本年的第几天
res := time.DayOfYear()

// 获取本月的第几天
res := time.DayOfMonth()

// 获取本周的第几天
res := time.DayOfWeek()

// 获取本年的第几周
res := time.WeekOfYear()
~~~


#### 求两个时间差值

~~~go
// 准备时间
timeA := datebin.Parse("2022-10-23 22:18:56")
timeB := datebin.Parse("2022-10-25 23:18:56")

diffTime := timeA.Diff(timeB)

// 相差秒
data := diffTime.Seconds()

// 相差秒，绝对值
data := diffTime.SecondsAbs()

// 其他
data := diffTime.Minutes()
data := diffTime.MinutesAbs()
data := diffTime.Hours()
data := diffTime.HoursAbs()
data := diffTime.Days()
data := diffTime.DaysAbs()
data := diffTime.Weeks()
data := diffTime.WeeksAbs()
data := diffTime.Months()
data := diffTime.MonthsAbs()
data := diffTime.Years()
data := diffTime.YearsAbs()

// 格式化输出
data := diffTime.Format("时间相差 {Y} 年")
data := diffTime.Format("时间相差 {m} 月")
data := diffTime.Format("时间相差 {d} 天")
data := diffTime.Format("时间相差 {H} 小时")
data := diffTime.Format("时间相差 {i} 分钟")
data := diffTime.Format("时间相差 {s} 秒")
~~~


#### 常用加减时间

~~~go
// 准备时间
time := datebin.Parse("2022-10-23 22:18:56")

// 常用加减时间
date := time.SubYears(uint(2)). # 年
    // SubYears(uint(2))
    // SubYearsNoOverflow(uint(2))
    // SubYear()
    // SubYearNoOverflow()

    // AddYears(uint(2))
    // AddYearsNoOverflow(uint(2))
    // AddYear()
    // AddYearNoOverflow()

    // 月份
    // SubMonths(uint(2))
    // SubMonthsNoOverflow(uint(2))
    // SubMonth()
    // SubMonthNoOverflow()

    // AddMonths(uint(2))
    // AddMonthsNoOverflow(uint(2))
    // AddMonth()
    // AddMonthNoOverflow()

    // 周
    // SubWeekdays(uint(2))
    // SubWeekday()

    // AddWeekdays(uint(2))
    // AddWeekday()

    // 天
    // SubDays(uint(2)) # 前 n 天
    // SubDay() # 前一天

    // AddDays(uint(2)) # 后 n 天
    // AddDay() # 后一天

    // 小时
    // SubHours(uint(2))
    // SubHour()

    // AddHours(uint(2))
    // AddHour()
    
    // 分钟
    // SubMinutes(uint(2))
    // SubMinute()

    // AddMinutes(uint(2))
    // AddMinute()

    // 秒
    // SubSeconds(uint(2))
    // SubSecond()

    // AddSeconds(uint(2))
    // AddSecond()

    // 毫秒
    // SubMilliseconds(uint(2))
    // SubMillisecond()

    // AddMilliseconds(uint(2))
    // AddMillisecond()

    // 微妙
    // SubMicroseconds(uint(2))
    // SubMicrosecond()

    // AddMicroseconds(uint(2))
    // AddMicrosecond()

    // 纳秒
    // SubNanoseconds(uint(2))
    // SubNanosecond()

    // AddNanoseconds(uint(2))
    // AddNanosecond()

    ToDatetimeString()
~~~


#### 判断是否是

~~~go
// 准备时间
time := datebin.Parse("2022-10-23 22:18:56")

// 是否是零值时间
res := time.IsZero()

// 是否是无效时间
res := time.IsInvalid()

// 是否是 UTC 时区
res := time.IsUTC()

// 是否是本地时区
res := time.IsLocal()

// 是否是当前时间
res := time.IsNow()

// 是否是未来时间
res := time.IsFuture()

// 是否是过去时间
res := time.IsPast()

// 是否是闰年
res := time.IsLeapYear()

// 是否是长年
res := time.IsLongYear()

// 是否是今天
res := time.IsToday()

// 是否是昨天
res := time.IsYesterday()

// 是否是明天
res := time.IsTomorrow()

// 是否是当年
res := time.IsCurrentYear()

// 是否是当月
res := time.IsCurrentMonth()

// 时间是否是当前最近的一周
res := time.IsLatelyWeek()

// 时间是否是当前最近的一个月
res := time.IsLatelyMonth()

// 是否是当前月最后一天
res := time.IsLastOfMonth()

// 是否当天开始
res := time.IsStartOfDay()

// 是否当天开始
res := time.IsStartOfDayWithMicrosecond()

// 是否当天结束
res := time.IsEndOfDay()

// 是否当天结束
res := time.IsEndOfDayWithMicrosecond()

// 是否是半夜
res := time.IsMidnight()

// 是否是中午
res := time.IsMidday()

// 是否是春季
res := time.IsSpring()

// 是否是夏季
res := time.IsSummer()

// 是否是秋季
res := time.IsAutumn()

// 是否是冬季
res := time.IsWinter()

// 是否是一月
res := time.IsJanuary()

// 是否是二月
res := time.IsFebruary()

// 是否是三月
res := time.IsMarch()

// 是否是四月
res := time.IsApril()

// 是否是五月
res := time.IsMay()

// 是否是六月
res := time.IsJune()

// 是否是七月
res := time.IsJuly()

// 是否是八月
res := time.IsAugust()

// 是否是九月
res := time.IsSeptember()

// 是否是十月
res := time.IsOctober()

// 是否是十一月
res := time.IsNovember()

// 是否是十二月
res := time.IsDecember()
~~~
