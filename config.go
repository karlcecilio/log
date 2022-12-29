package log

import (
	"encoding/json"
	"errors"
	"os"
)

type ConfFileWriter struct {
	LogPath string `json:"LogPath"`
	On      bool   `json:"On"`
}

type LogConfig struct {
	Level string         `json:"LogLevel"`
	FW    ConfFileWriter `json:"FileWriter"`
}

func SetupLogWithConf(file string) (err error) {
	var lc LogConfig

	cnt, err := os.ReadFile(file)

	if err = json.Unmarshal(cnt, &lc); err != nil {
		return
	}

	if lc.FW.On {
		w := NewFileWriter()
		w.SetPathPattern(lc.FW.LogPath)
		Register(w)
	}

	switch lc.Level {
	case "debug":
		SetLevel(DEBUG)

	case "info":
		SetLevel(INFO)

	case "warning":
		SetLevel(WARNING)

	case "error":
		SetLevel(ERROR)

	case "fatal":
		SetLevel(FATAL)

	default:
		err = errors.New("Invalid log level")
	}
	return
}
