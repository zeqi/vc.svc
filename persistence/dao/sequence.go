package dao

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
	"vc.svc/persistence/schemas"
)

type SequenceDao struct {
	Dao
	Schema schemas.Sequence
}

func (o *SequenceDao) NewDao() *SequenceDao {
	s := o.Schema
	o.Init(s.GetCollectionName(), s.GetIndexs())
	return o
}

func (o *SequenceDao) Find(condition bson.M, sorts []string, skip int, limit int) ([]schemas.Sequence, error) {
	result := []schemas.Sequence{}
	err := o.Collection.Find(condition).Sort(sorts...).Skip(skip).Limit(limit).All(&result)
	if err != nil {
		fmt.Errorf("SequenceDao Find Error %s", err)
		return nil, err
	}
	return result, nil
}

func (o *SequenceDao) FindOne(condition bson.M, sorts []string, skip int, limit int) (schemas.Sequence, error) {
	result := schemas.Sequence{}
	err := o.Collection.Find(condition).Sort(sorts...).Skip(skip).Limit(limit).One(&result)
	if err != nil {
		fmt.Errorf("SequenceDao FindOne Error %s", err)
	}
	return result, err
}

func (o *SequenceDao) FindById(id string) (schemas.Sequence, error) {
	result := schemas.Sequence{}
	err := o.Collection.FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		fmt.Errorf("SequenceDao FindById Error %s", err)
	}
	return result, err
}
