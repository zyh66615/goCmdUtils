package cmd

import (
	"flag"
	"log"
	"time"
)

const (
	B  ByteUnit = 1
	KB ByteUnit = 1024 * B
	MB ByteUnit = KB * KB
	GB ByteUnit = MB * KB
)

type (
	ByteUnit   uint32
	ArgsParser struct {
		Precision      uint8
		Internal       time.Duration
		MemOutputUnit  ByteUnit
		DiskOutputUnit ByteUnit
	}
)

func ParseFromArgs() ArgsParser {
	var extra = flag.Bool("x", false, "是否打印额外数据")
	var internal = flag.Int64("d", 1, "输出刷新时间间隔")
	flag.Parse()
	timeInternal := time.Duration(*internal)
	if *extra {
		log.Println("extra info!")
	}
	return ArgsParser{
		Precision:      0,
		Internal:       timeInternal,
		MemOutputUnit:  KB,
		DiskOutputUnit: KB,
	}
}
