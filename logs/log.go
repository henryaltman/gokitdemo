package logs

import "fmt"

const (
	LevelDebug = iota
	LevelInfo
	LevelError
	LogFileMode = iota
)

func Log(level int, info string, mode int) {
	logInfo := ""
	switch level {
	case LevelDebug:
		logInfo = fmt.Sprintf("debug-%s", info)
	case LevelInfo:
		logInfo = fmt.Sprintf("info-%s", info)
	case LevelError:
		logInfo = fmt.Sprintf("error-%s", info)
	default:
		logInfo = fmt.Sprintf("default-%s", info)
	}
	fmt.Println("%s", logInfo)
}
