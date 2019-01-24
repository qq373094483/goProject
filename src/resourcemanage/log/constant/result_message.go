package constant

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type message struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	Success bool `json:"success"`
}

const(
	Success="200"
	Fail="201"
	)
var Messages = make(map[string]*message)

func init() {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(workPath, "conf", "message.conf")
	f, err := os.Open(appConfigPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil {
			if (io.EOF == err && len(line) > 0) {
				goto end
			}
			break
		}
	end:
		var m message
		for index, field := range strings.Split(line, "=") {
			if index == 0 {
				m.Code = field
			} else if index == 1 {
				m.Msg = strings.Split(field,"\r\n")[0]
			} else {
			}
		}
		Messages[m.Code] = &m
	}
	Messages["200"].Success = true
}
