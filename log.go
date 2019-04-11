package klog

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"log"
	"time"
)

type Conf struct {
	Path       string        // log path
	Name       string        // log name
	MaxAge     time.Duration // log expire time
	RotateTime time.Duration // log split time
}

func (c *Conf) Set() error {
	rl, err := rotatelogs.New(
		fmt.Sprintf("%s/%s", c.Path, c.Name)+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(fmt.Sprintf("%s/%s", c.Path, c.Name)),
		rotatelogs.WithMaxAge(c.MaxAge),
		rotatelogs.WithRotationTime(c.RotateTime),
	)
	if err != nil {
		return err
	}
	log.SetOutput(rl)
	return nil
}
