package logger

import "log"

func WriteLogs(s ...string) {
	log.Println(s)
}
