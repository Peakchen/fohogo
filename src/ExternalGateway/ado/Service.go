package ado

import (
	"github.com/Peakchen/xgameCommon/ado/dbCache"
	"github.com/Peakchen/xgameCommon/ado/service"
	"github.com/Peakchen/xgameCommon/aktime"
)

var (
	_dbProvider *service.TDBProvider
)

func StartDBSerice(server string) {
	_dbProvider.StartDBService(server, dbCache.UpdateDBCache)
	dbCache.InitDBCache(_dbProvider)
	aktime.InitAkTime(_dbProvider.GetRedisConn())
}

func init() {
	_dbProvider = &service.TDBProvider{}
}
