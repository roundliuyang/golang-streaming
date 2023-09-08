package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	LBAddr  string `json:"lb_addr"`
	OssAddr string `json:"oss__addr"`
}

var configuration *Configuration

func init() {
	file, _ := os.Open("D:\\xm\\golang-streaming\\video_server\\conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = &Configuration{}

	err := decoder.Decode(configuration)
	if err != nil {
		panic(err)
	}
}

func GetLbAddr() string {
	return configuration.LBAddr
}

func GetOssAddr() string {
	return configuration.OssAddr
}
