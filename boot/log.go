package boot

import (
	"fmt"
	"golang.org/x/xerrors"
	"ihtest/library/global"
)

type oathLog struct {
	Type    int64
	content string
}

var lo oathLog

func InitLogger() error {
	path := global.Config.Log.Dir

	if "" == path {
		return xerrors.New("log.dir invalid")
	}
	return nil
}

func (l *oathLog) Info(content string) {
	fmt.Println(content)
}
func (l *oathLog) Error(content string) {
	fmt.Println(content)
}
func (l *oathLog) Debug(content string) {
	fmt.Println(content)
}
