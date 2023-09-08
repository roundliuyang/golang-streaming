package ossops

import (
	"github.com/alanhou/golang-streaming/video_server/web/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

var EP string
var AK string
var SK string

func init() {
	AK = "LTAI4GKTkNmshPHuFku28PXi"
	SK = "Um6lYxAU1wKHr8gGSj005ZSJdjf0J3"
	//EP = "oss-cn-beijing.aliyuncs.com"
	EP = config.GetOssAddr()
}

func UploadToOss(filename string, path string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss service error: %s", err)
		return false
	}
	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("Getting bucket error: %s", err)
		return false
	}
	bucket.UploadFile(filename, path, 500*1024, oss.Routines(3))
	if err != nil {
		log.Printf("Uploading object error: %s", err)
		return false
	}
	return true
}

func DeleteObject(filename string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss service error: %s", err)
		return false
	}
	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("Getting bucket error: %s", err)
		return false
	}
	bucket.DeleteObject(filename)
	if err != nil {
		log.Printf("Deleting object error: %s", err)
		return false
	}
	return true
}
