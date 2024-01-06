package datetimes

import (
    "github.com/deatil/go-datebin/datebin"
)

/**
 * 时间范围 / Datetimes
 *
 * @create 2024-1-6
 * @author deatil
 */
type Datetimes struct {
    // 开始时间
    Start datebin.Datebin

    // 结束时间
    End datebin.Datebin
}

// 构造函数 / NewDatetimes
func NewDatetimes(start, end datebin.Datebin) Datetimes {
    if start.Gt(end) {
        start, end = end, start
    }

    return Datetimes{
        Start: start,
        End:   end,
    }
}

// 构造函数 / New
func New(start, end datebin.Datebin) Datetimes {
    return NewDatetimes(start, end)
}

// 交集 / Intersection
func (this Datetimes) Intersection(x Datetimes) Datetimes {
    ds := Datetimes{}

    left, right := this.swap(this, x)

    if left.Start.Lt(right.Start) {
        ds.Start = right.Start
    } else {
        ds.Start = left.Start
    }

    if left.End.Gt(right.End) {
        ds.End = right.End
    } else {
        ds.End = left.End
    }

    return ds
}

// 并集 / Union
func (this Datetimes) Union(x Datetimes) []Datetimes {
    ds := make([]Datetimes, 0)

    left, right := this.swap(this, x)

    if left.End.Lt(right.Start) {
        ds = append(ds, left)
        ds = append(ds, right)
    } else {
        ds = append(ds, Datetimes{
            Start: left.Start,
            End:   right.End,
        })
    }

    return ds
}

// 是否包含 x / if Contain x
func (this Datetimes) IsContain(x Datetimes) bool {
    if this.Start.Gt(x.Start) {
        return false
    }

    if this.End.Lt(x.End) {
        return false
    }

    return true
}

// 交换大小 / swap x, y
func (this Datetimes) swap(x, y Datetimes) (Datetimes, Datetimes) {
    left, right := x, y

    if left.Start.Gt(right.Start) {
        left, right = right, left
    }

    return left, right
}
