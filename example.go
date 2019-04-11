package klog

import (
	"log"
	"time"
)

func example() {
	conf := &Conf{
		Path:       "/your/log/path",
		Name:       "log_name",
		MaxAge:     time.Duration(time.Hour * 24 * 7),
		RotateTime: time.Duration(time.Hour * 24),
	}
	if err := conf.Set(); err != nil {
		panic(err)
	}

	log.Println("first line")
	log.Println("second line")
}
