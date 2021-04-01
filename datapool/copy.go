package datapool

import (
	"io"
)

func Copy(dst io.Writer, src io.Reader, timer func()) (written int64, err error) {
	buf := make([]byte, 1024)
	for {
		if timer != nil {
			timer()
		}
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err
}
