package logger

import (
	"log"
	"os"
)

var (
	outfile, _ = os.OpenFile("/var/log/news-bot.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	LogFile    = log.New(outfile, "", 0)
)

func ForError(er error) {
	if er != nil {
		LogFile.Fatalln(er)
	}
}
