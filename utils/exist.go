package utils

import (
	"github.com/gditsec/crawlab-go-sdk/constants"
	"github.com/gditsec/crawlab-go-sdk/database"
	"github.com/gditsec/crawlab-go-sdk/entity"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func IsExistItem(item entity.Item) (bool, error) {
	dsType := GetDataSourceType()
	if dsType == constants.DataSourceTypeMongo {
		return IsExistItemMongo(item)
	}
	return false, nil
}

func IsExistItemMongo(item entity.Item) (bool, error) {
	ds := GetDataSource()
	_, c, err := database.GetMongoCol(ds)
	if err != nil {
		return false, err
	}

	isDedup := GetIsDedup()

	if isDedup == "1" {
		dedupField := GetDedupField()

		var res interface{}
		if err := c.Find(bson.M{dedupField: item[dedupField]}).One(&res); err != nil {
			if err == mgo.ErrNotFound {
				// 不存在
				return false, nil
			} else {
				// 出错
				return false, err
			}
		}
		// 已存在
		return true, nil
	}
	return false, nil
}
