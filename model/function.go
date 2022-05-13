package model

import (
	"grid/readconfig"
)

func ReadGrid() *GridInfo {
	gridinfo := new(GridInfo)
	gridinfo.Cap = readconfig.ReadConfig("config.trade.ini_cap").(int)
	gridinfo.Leverage = readconfig.ReadConfig("config.trade.leverage").(int)

	if _, ok := readconfig.ReadConfig("config.trade.grid_low").(int); ok {
		gridinfo.Low = float64(readconfig.ReadConfig("config.trade.grid_low").(int))
	} else {
		gridinfo.Low = readconfig.ReadConfig("config.trade.grid_low").(float64)
	}

	if _, ok := readconfig.ReadConfig("config.trade.grid_high").(int); ok {
		gridinfo.High = float64(readconfig.ReadConfig("config.trade.grid_high").(int))
	} else {
		gridinfo.High = readconfig.ReadConfig("config.trade.grid_high").(float64)
	}

	gridinfo.GridNum = readconfig.ReadConfig("config.trade.grid_count").(int)
	gridinfo.Limit = readconfig.ReadConfig("config.trade.stop_limit").(float64)
	gridinfo.MaxRate = readconfig.ReadConfig("config.trade.max_rate").(float64)
	return gridinfo
}
