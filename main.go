package main

import (
	"vc.svc/db"
	"vc.svc/services"
	"vc.svc/utils"
)

func main() {
	utils.GetInstance()
	instance := utils.GetInstance()
	config := instance.Config
	db.Init(config.DataProviders.MongodbSmart)
	services.Init(config.MicroServices.MicroMongo)
}
