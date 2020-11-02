package common

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 配置数据
var config = make(map[string]string)

// 配置文件地址
var configPath = "../config/config.config"

// 配置文件注释
var note = "#"

// 配置文件注释
var separator = "="

// 日志前缀
var LOG_INFO = " [Info] "
var LOG_ERROR = " [Error] "

// 配置的变量
var (
	PORT   = "port"
	SCRIPT = "script"
	// info 日志相关
	LOG_INFO_DIR        = "log.info.dir"
	LOG_INFO_COMPRESS   = "log.info.compress"
	LOG_INFO_FILEPREFIX = "log.info.fileprefix"
	LOG_INFO_RESERVEDAY = "log.info.reserveday"
	//  error 日志相关
	LOG_ERROR_DIR        = "log.error.dir"
	LOG_ERROR_COMPRESS   = "log.error.compress"
	LOG_ERROR_FILEPREFIX = "log.error.fileprefix"
	LOG_ERROR_RESERVEDAY = "log.error.ReserveDay"
	SCRIPT_LOG           = "script.log"
)

//读取key=value类型的配置文件
func init() {
	f, err := os.Open(configPath)
	defer f.Close()
	if err != nil {
		fmt.Println("配置文件加载失败", err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		if strings.HasPrefix(s, note) {
			continue
		}
		index := strings.Index(s, separator)
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
}

func GetValue(key string) string {
	return config[key]
}

func GetPort() string {
	return config[PORT]
}

func GetScript() string {
	return config[SCRIPT]
}

// info 日志相关
func GetLogInfoDir() string {
	return config[LOG_INFO_DIR]
}
func GetLogInfoCompress() bool {
	return getBool(LOG_INFO_COMPRESS)
}
func GetLogInfoFilePrefix() string {
	return config[LOG_INFO_FILEPREFIX]
}
func GetLogInfoReserveDay() int {
	return getInt(LOG_INFO_RESERVEDAY)
}

////  error 日志相关
func GetLogErrorDir() string {
	return config[LOG_ERROR_DIR]
}
func GetLogErrorCompress() bool {
	return getBool(LOG_ERROR_COMPRESS)
}
func GetLogErrorFilePrefix() string {
	return config[LOG_ERROR_FILEPREFIX]
}
func GetLogErrorReserveDay() int {
	return getInt(LOG_ERROR_RESERVEDAY)
}

func getInt(key string) int {
	value := config[key]
	res, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return res
}

func getBool(key string) bool {
	value := config[key]
	if len(value) > 0 && strings.ToUpper(value) == "TRUE" {
		return true
	}
	return false
}

func GetScriptLog() string {
	return config[SCRIPT_LOG]
}
