package schemas

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Sequence struct {
	Id          bson.ObjectId `json:"Id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"Name" bson:"name"`
	Value       int64         `json:"Value" bson:"value"`
	Comments    string        `json:"Comments,omitempty" bson:"comments,omitempty"`
	Status      string        `json:"Status,omitempty" bson:"status,omitempty"`
	Archived    bool          `json:"Archived,omitempty" bson:"archived,omitempty"`
	CreatedTime time.Time     `json:"CreatedTime,omitempty" bson:"createdTime"`
	LastModTime time.Time     `json:"LastModTime,omitempty" bson:"lastModTime"`
	Tracking    []Tracking    `json:"Tracking,omitempty" bson:"tracking"`
}

func (o *Sequence) GetCollectionName() string {
	return "sys.sequence"
}

func (o *Sequence) GetIndexs() []mgo.Index {
	var indexs []mgo.Index
	indexs = append(indexs, mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}, mgo.Index{
		Key: []string{"value"},
	}, mgo.Index{
		Key: []string{"comments"},
	}, mgo.Index{
		Key: []string{"status"},
	}, mgo.Index{
		Key: []string{"archived"},
	}, mgo.Index{
		Key: []string{"createdTime"},
	}, mgo.Index{
		Key: []string{"lastModTime"},
	})
	return indexs
}

func (o *Sequence) NewSequence(name string, comments string) *Sequence {
	o.Name = name
	o.Value = 0
	o.Comments = comments
	o.Status = "settled"
	o.Archived = false
	o.CreatedTime = time.Now()
	o.LastModTime = time.Now()
	// o.Tracking = append(o.Tracking, Tracking{Operator: "zeqi", OptTime: time.Now()})
	return o
}
