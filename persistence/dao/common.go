package dao

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"vc.svc/db"
)

type Dao struct {
	Session        *mgo.Session
	DBname         string
	CollectionName string
	Collection     *mgo.Collection
	Indexs         []mgo.Index
}

func (o *Dao) EnsureIndex() {
	for i := 0; i < len(o.Indexs); i++ {
		o.Collection.EnsureIndex(o.Indexs[i])
	}
}

func (o *Dao) GetCollection() *mgo.Collection {
	c := o.Session.DB(o.DBname).C(o.CollectionName)
	o.Collection = c
	return c
}

func (o *Dao) Init(collectionName string, indexs []mgo.Index) *Dao {
	o.DBname = db.DialInfo.Database
	o.Session = db.MongoSession
	o.CollectionName = collectionName
	o.Indexs = indexs
	o.GetCollection()
	o.EnsureIndex()
	return o
}

func (o *Dao) Insert(docs ...interface{}) error {
	err := o.Collection.Insert(docs...)
	if err != nil {
		fmt.Errorf("Dao Insert Error %s", err)
	}
	return err
}

func (o *Dao) Update(condition bson.M, update bson.M) error {
	err := o.Collection.Update(condition, update)
	if err != nil {
		fmt.Errorf("Dao Update Error %s", err)
	}
	return err
}
func (o *Dao) UpdateId(id string, update bson.M) error {
	err := o.Collection.UpdateId(bson.ObjectIdHex(id), update)
	if err != nil {
		fmt.Errorf("Dao Update Error %s", err)
	}
	return err
}

func (o *Dao) UpdateAll(condition bson.M, update bson.M) error {
	info, err := o.Collection.UpdateAll(condition, update)
	if err != nil {
		fmt.Errorf("Dao UpdateAll Error %s", err)
	}
	fmt.Println("Dao UpdateAll Count", info.Matched, "records")
	return err
}

func (o *Dao) Upsert(condition bson.M, update bson.M) error {
	info, err := o.Collection.Upsert(condition, update)
	fmt.Println("Dao Upsert info %#v", info.UpsertedId)
	if err != nil {
		fmt.Errorf("Dao Upsert Error %#v", err)
	}
	return err
}

func (o *Dao) UpsertId(id string, update bson.M) error {
	info, err := o.Collection.UpsertId(bson.ObjectIdHex(id), update)
	fmt.Println("Dao UpsertId info %#v", info)
	if err != nil {
		fmt.Errorf("Dao UpsertId Error %#v", err)
	}
	return err
}

func (o *Dao) Remove(condition bson.M) error {
	err := o.Collection.Remove(condition)
	if err != nil {
		fmt.Errorf("Dao Error %s", err)
	}
	return err
}

func (o *Dao) Count(condition bson.M) (int, error) {
	i, err := o.Collection.Find(condition).Count()
	if err != nil {
		fmt.Errorf("Dao Error %s", err)
	}
	return i, err
}
