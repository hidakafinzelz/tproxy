package protocol

import (
	"bufio"
	"github.com/kevwan/tproxy/display"
	"io"
	"time"
)

type redisInterop struct {
}

func (red *redisInterop) Dump(r io.Reader, source string, id int, quiet bool) {
	buf := bufio.NewReader(r)
	for {
		var buffer = make([]byte, 1<<20)
		// read raw data
		n, _ := buf.Read(buffer)
		if n > 0 {
			if source == "SERVER" {
				go display.PrintfWithTime("\n[Server -> Client] : \n%s\n", string(buffer[:n:n]))
			} else {
				go display.PrintfWithTime("\n[Client -> Server] : \n%s\n", string(buffer[:n:n]))
			}
		}
		time.Sleep(time.Millisecond * 30)
	}
}
