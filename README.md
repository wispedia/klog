## 简介

`klog`是一个轻量级的日志package，有如下特性

- 日志输出至文件
- 支持日志按时间自动切分
- 支持日志自动清理
- 基于golang标注库`log`以及`file-rotatelogs`构建，很简单😀

## 使用

`go get github.com/wispedia/klog`


```golang
package main

import (
    "github.com/wispedia/klog"
    "time"
    "log"
)

func init() {
    conf := &klog.Conf{
        Path:       "/your/log/path",
        Name:       "log_name",
        MaxAge:     time.Duration(time.Hour * 24 * 7), // log expire time
        RotateTime: time.Duration(time.Hour * 24), // log split time
    }
    if err := conf.Set(); err != nil {
        panic(err)
    }
}

func main() {
    log.Println("first line")
    log.Println("second line")
}

```
