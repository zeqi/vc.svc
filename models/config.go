package models

import (
	"strconv"
)

type Etcd struct {
	Addrs string
}

type MongodbConfig struct {
	Host           string `json:"Host"`
	Port           int    `json:"Port"`
	DB             string `json:"DB"`
	ReplicaSetName string `json:"ReplicaSetName"`
	User           string `json:"User"`
	Pass           string `json:"Pass"`
}

type DataProviders struct {
	MongodbSmart MongodbConfig `json:"MongodbSmart"`
}

type MicroConfig struct {
	Name string
	Etcd Etcd
}

type MicroServices struct {
	MicroMongo MicroConfig
}

type Config struct {
	MicroServices MicroServices
	DataProviders DataProviders `json:"DataProviders"`
}

type ServiceStatus struct {
	OK    string
	Error string
	Warn  string
}

func (o *ServiceStatus) NewServiceStatus() {
	o.OK = "OK"
	o.Error = "Error"
	o.Warn = "Warn"
}

func (o *ServiceStatus) ParseIntToInt64(num int) int64 {
	str := strconv.Itoa(num)
	num64, _ := strconv.ParseInt(str, 10, 64)
	return num64
}
