package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"promise/internal/errorCode"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"github.com/minio/minio-go/v6"
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

// 上传文件
func FileUpload(objName string) (Ossurl string) {
	//admin
	// accessKeyID := "zXOwV7hYd1faTvbw"
	// secretAccessKey := "MsUyyMXqaD7zGKkJgI3vPAgCAgDNOldT"
	endpoint := "oss.cowarobot.work"
	accessKeyID := "m0DUSIi31Xh6h4rR"
	secretAccessKey := "cmAnWPaF5kion13aZqX1FrS690WsjuXr"
	useSSL := true

	// 初使化minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// 创建一个叫promise的存储桶。
	bucketName := "promise"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)

	// 上传一个zip文件。
	objectName := objName
	//cmd文件下的相对路径
	filePath := "./data/saveFile.7z"
	contentType := "application/7z"
	var buf strings.Builder
	buf.WriteString("http://172.16.0.28:30099/minio/promise/")
	buf.WriteString(objectName)
	Ossurl = buf.String()
	// 使用FPutObject上传一个zip文件。
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
	return Ossurl
	//上传文件默认路径	http://172.16.0.28:30099/minio/promise/文件名
}

// 下载文件
func DownloadFile(url string) {
	//Get the file
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//创建文件用于保存
	//cmd文件下的相对路径
	out, err := os.Create("./data/saveFile.7z")
	// fmt.Println(os.Getwd())

	if err != nil {
		// fmt.Println(os.Getwd())
		fmt.Println(err, "创建文件err.....................")
		// panic(err)
	}
	defer out.Close()
	//对接相应流和文件流
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println(err, "接受流写入err....................")
		panic(err)

	}

}
