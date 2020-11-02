package log

import (
	"ci_client/common"
	"io"
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	common.GetLogInfoDir()
	infoWriter := &TimeWriter{
		Dir:        common.GetLogInfoDir(),
		Compress:   common.GetLogInfoCompress(),
		FilePrefix: common.GetLogInfoFilePrefix(),
		ReserveDay: common.GetLogInfoReserveDay(),
	}
	errorWriter := &TimeWriter{
		Dir:        common.GetLogErrorDir(),
		Compress:   common.GetLogErrorCompress(),
		FilePrefix: common.GetLogErrorFilePrefix(),
		ReserveDay: common.GetLogErrorReserveDay(),
	}
	Info = log.New(io.MultiWriter(os.Stdout, infoWriter), common.LOG_INFO, log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errorWriter), common.LOG_ERROR, log.Ldate|log.Ltime|log.Lshortfile)
}
