package crawlab

import (
	"bytes"
	"io"

	"github.com/gditsec/crawlab-go-sdk/entity"
	"github.com/gditsec/crawlab-go-sdk/utils"
)

func SaveItem(item entity.Item) (err error) {
	if err := utils.SaveItem(item); err != nil {
		return err
	}
	return nil
}

func IsExistItem(item entity.Item) (bool, error) {
	return utils.IsExistItem(item)
}

func SaveFile(name string, data []byte) error {
	reader := bytes.NewReader(data)
	return SaveFileFrom(name, reader, nil)
}

func SaveFileFrom(name string, reader io.Reader, timer func()) error {
	return utils.SaveFileFrom(name, reader, timer)
}

// 接口
func NotifyTarget(item entity.Item, files []string) error {
	return utils.NotifyTarget(item, files)
}
