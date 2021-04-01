package utils

import (
	"io"

	"github.com/gditsec/crawlab-go-sdk/datapool"
)

func SaveFileFrom(name string, reader io.Reader, timer func()) error {
	target := datapool.GetTargetConfig()
	if pool := datapool.New(target); pool != nil {
		return pool.UploadFile(name, reader, timer)
	}
	return nil
}
