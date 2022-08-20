package utils

import (
	"context"
	"fmt"
	"promise/internal/errorCode"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
)

func MyCopy(ctx context.Context, toValue interface{}, fromValue interface{}) (err error) {
	if err = copier.Copy(toValue, fromValue); err != nil {
		return errorCode.NewMyErr(ctx, errorCode.MyInternalError, err)
	} else {
		return nil
	}
}

func Interface2String(inter interface{}) string {
	switch inter.(type) {
	case string:
		return inter.(string)
	case int:
		return strconv.Itoa(inter.(int))
	case float64:
		return strconv.FormatFloat(inter.(float64), 'f', -1, 32)
	case float32:
		return strconv.FormatFloat(float64(inter.(float32)), 'f', -1, 32)
	default:
		fmt.Println("Interface2String unknown type", inter)
	}
	return ""
}

func Interface2Int64(inter interface{}) int64 {
	switch inter.(type) {
	case string:
		if i, ok := strconv.ParseInt(inter.(string), 0, 64); ok == nil {
			return i
		}
		return 0
	case int:
		return int64(inter.(int))
	case float64:
		return int64(inter.(float64))
	case float32:
		return int64(inter.(float32))
	default:
		fmt.Println("Interface2Int unknown type", inter)
	}
	return 0
}

func Interface2Int32(inter interface{}) int {
	switch inter.(type) {
	case string:
		if i, ok := strconv.Atoi(inter.(string)); ok == nil {
			return i
		}
		return 0
	case int:
		return inter.(int)
	case float64:
		return int(inter.(float64))
	case float32:
		return int(inter.(float32))
	default:
		fmt.Println("Interface2Int32 unknown type", inter)
	}
	return 0
}

func String2Time(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	tm := time.Unix(0, msInt*int64(time.Millisecond))

	// fmt.Println(tm.Format("2006-02-01 15:04:05.000"))

	return tm, nil
}

func Int2Time(msInt int64) time.Time {
	return time.Unix(0, msInt*int64(time.Millisecond))
}
