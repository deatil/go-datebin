package datetimes

import (
    "testing"
    "reflect"

    "github.com/deatil/go-datebin/datebin"
)

func assertEqualT(t *testing.T) func(any, any, string) {
    return func(actual any, expected any, msg string) {
        if !reflect.DeepEqual(actual, expected) {
            t.Errorf("Failed %s: actual: %v, expected: %v", msg, actual, expected)
        }
    }
}

func Test_Intersection(t *testing.T) {
    eq := assertEqualT(t)

    tests := []struct {
        index string
        date1_1 string
        date1_2 string
        date2_1 string
        date2_2 string
        check_1 string
        check_2 string
    } {
        {
            index: "index-1",
            date1_1: "2024-06-01 20:15:12",
            date1_2: "2024-06-03 21:15:12",
            date2_1: "2024-06-02 21:15:12",
            date2_2: "2024-06-05 22:15:12",
            check_1: "2024-06-02 21:15:12",
            check_2: "2024-06-03 21:15:12",
        },
        {
            index: "index-2",
            date1_1: "2024-06-03 21:15:12",
            date1_2: "2024-06-01 20:15:12",
            date2_1: "2024-06-02 21:15:12",
            date2_2: "2024-06-05 22:15:12",
            check_1: "2024-06-02 21:15:12",
            check_2: "2024-06-03 21:15:12",
        },
    }

    for _, td := range tests {
        x := NewDatetimes(
                datebin.ParseWithLayout(td.date2_1, datebin.DatetimeFormat, datebin.UTC),
                datebin.ParseWithLayout(td.date2_2, datebin.DatetimeFormat, datebin.UTC),
            )

        check := NewDatetimes(
                datebin.ParseWithLayout(td.date1_1, datebin.DatetimeFormat, datebin.UTC),
                datebin.ParseWithLayout(td.date1_2, datebin.DatetimeFormat, datebin.UTC),
            ).
            Intersection(x)

        eq(check.Start.ToDatetimeString(datebin.UTC), td.check_1, "failed Intersection Start, index " + td.index)
        eq(check.End.ToDatetimeString(datebin.UTC), td.check_2, "failed Intersection End, index " + td.index)
    }
}

func Test_Union(t *testing.T) {
    eq := assertEqualT(t)

    tests := []struct {
        index string
        date1_1 string
        date1_2 string
        date2_1 string
        date2_2 string
        check_1 string
        check_2 string
        check2_1 string
        check2_2 string
    } {
        {
            index: "index-1",
            date1_1: "2024-06-01 20:15:12",
            date1_2: "2024-06-03 21:15:12",
            date2_1: "2024-06-02 21:15:12",
            date2_2: "2024-06-05 22:15:12",
            check_1: "2024-06-01 20:15:12",
            check_2: "2024-06-05 22:15:12",
        },
        {
            index: "index-2",
            date1_1: "2024-06-03 21:15:12",
            date1_2: "2024-06-05 20:15:12",
            date2_1: "2024-06-06 21:15:12",
            date2_2: "2024-06-07 22:15:12",
            check_1: "2024-06-03 21:15:12",
            check_2: "2024-06-05 20:15:12",
            check2_1: "2024-06-06 21:15:12",
            check2_2: "2024-06-07 22:15:12",
        },
    }

    for _, td := range tests {
        x := NewDatetimes(
                datebin.ParseWithLayout(td.date2_1, datebin.DatetimeFormat, datebin.UTC),
                datebin.ParseWithLayout(td.date2_2, datebin.DatetimeFormat, datebin.UTC),
            )

        check := NewDatetimes(
                datebin.ParseWithLayout(td.date1_1, datebin.DatetimeFormat, datebin.UTC),
                datebin.ParseWithLayout(td.date1_2, datebin.DatetimeFormat, datebin.UTC),
            ).
            Union(x)

        if len(check) == 2 {
            eq(check[0].Start.ToDatetimeString(datebin.UTC), td.check_1, "failed Union Start 1, index " + td.index)
            eq(check[0].End.ToDatetimeString(datebin.UTC), td.check_2, "failed Union End 1, index " + td.index)
            eq(check[1].Start.ToDatetimeString(datebin.UTC), td.check2_1, "failed Union Start 2, index " + td.index)
            eq(check[1].End.ToDatetimeString(datebin.UTC), td.check2_2, "failed Union End 2, index " + td.index)
        } else {
            eq(check[0].Start.ToDatetimeString(datebin.UTC), td.check_1, "failed Union Start, index " + td.index)
            eq(check[0].End.ToDatetimeString(datebin.UTC), td.check_2, "failed Union End, index " + td.index)
        }

    }
}

func Test_IsContain(t *testing.T) {
    eq := assertEqualT(t)

    tests := []struct {
        index string
        date1_1 string
        date1_2 string
        date2_1 string
        date2_2 string
        check bool
    } {
        {
            index: "index-1",
            date1_1: "2024-06-01 20:15:12",
            date1_2: "2024-06-06 21:15:12",
            date2_1: "2024-06-02 21:15:12",
            date2_2: "2024-06-05 22:15:12",
            check: true,
        },
        {
            index: "index-2",
            date1_1: "2024-06-03 21:15:12",
            date1_2: "2024-06-01 20:15:12",
            date2_1: "2024-06-02 21:15:12",
            date2_2: "2024-06-05 22:15:12",
            check: false,
        },
    }

    for _, td := range tests {
        x := NewDatetimes(
                datebin.ParseWithLayout(td.date2_1, datebin.DatetimeFormat, datebin.UTC),
                datebin.ParseWithLayout(td.date2_2, datebin.DatetimeFormat, datebin.UTC),
            )

        check := NewDatetimes(
                datebin.ParseWithLayout(td.date1_1, datebin.DatetimeFormat, datebin.UTC),
                datebin.ParseWithLayout(td.date1_2, datebin.DatetimeFormat, datebin.UTC),
            ).
            IsContain(x)

        eq(check, td.check, "failed IsContain, index " + td.index)
    }
}

