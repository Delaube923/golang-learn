package utils

import (
	"context"
	"encoding/json"
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

func GetInterfaceToString(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

func GetInterfaceToInt(t1 interface{}) int {
	var t2 int
	switch t1.(type) {
	case uint:
		t2 = int(t1.(uint))
		break
	case int8:
		t2 = int(t1.(int8))
		break
	case uint8:
		t2 = int(t1.(uint8))
		break
	case int16:
		t2 = int(t1.(int16))
		break
	case uint16:
		t2 = int(t1.(uint16))
		break
	case int32:
		t2 = int(t1.(int32))
		break
	case uint32:
		t2 = int(t1.(uint32))
		break
	case int64:
		t2 = int(t1.(int64))
		break
	case uint64:
		t2 = int(t1.(uint64))
		break
	case float32:
		t2 = int(t1.(float32))
		break
	case float64:
		t2 = int(t1.(float64))
		break
	case string:
		t2, _ = strconv.Atoi(t1.(string))
		break
	default:
		t2 = t1.(int)
		break
	}
	return t2
}
