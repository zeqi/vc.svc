package sequence

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/micro/go-micro"
	"gopkg.in/mgo.v2/bson"
	pb "vc.pb/sequence"
	"vc.svc/models"
	"vc.svc/persistence/dao"
	"vc.svc/persistence/schemas"
)

type Service struct {
	models.ServiceStatus
	dao *dao.SequenceDao
}

func (o *Service) Create(ctx context.Context, req *pb.ReqCreate, rsp *pb.ResDoc) error {
	s := schemas.Sequence{}
	rsp.Status = o.OK
	return o.dao.Insert(s.NewSequence(req.Name, req.Comments))
}

func (o *Service) getComdition(req *pb.ReqFind) (bson.M, []string, int, int) {
	skip, err := strconv.Atoi(strconv.FormatInt(req.Skip, 10))
	if err != nil {
		fmt.Errorf("SequenceService Error %s", err)
	}
	limit, err := strconv.Atoi(strconv.FormatInt(req.Limit, 10))
	if err != nil {
		fmt.Errorf("SequenceService Error %s", err)
	}
	sorts := req.Sorts
	items := []bson.M{bson.M{}}
	if req.Condition.Name != "" {
		items = append(items, bson.M{"name": req.Condition.Name})
	}
	condition := bson.M{"$and": items}
	return condition, sorts, skip, limit
}

func (o *Service) Find(ctx context.Context, req *pb.ReqFind, rsp *pb.ResDocs) error {
	fmt.Println("Sequence Find")
	condition, sorts, skip, limit := o.getComdition(req)
	docs, err := o.dao.Find(condition, sorts, skip, limit)
	rsp.Status = o.OK
	for i := 0; i < len(docs); i++ {
		var doc = docs[i]
		rsp.Data = append(rsp.Data, &pb.Model{Name: doc.Name, Id: doc.Id.Hex()})
	}
	return err
}

func (o *Service) FindOne(ctx context.Context, req *pb.ReqFind, rsp *pb.ResDoc) error {
	condition, sorts, skip, limit := o.getComdition(req)
	doc, err := o.dao.FindOne(condition, sorts, skip, limit)
	rsp.Status = o.OK
	rsp.Data = &pb.Model{Name: doc.Name, Id: doc.Id.Hex()}
	return err
}

func (o *Service) FindById(ctx context.Context, req *pb.Model, rsp *pb.ResDoc) error {
	doc, err := o.dao.FindById(req.Id)
	rsp.Status = o.OK
	rsp.Data = &pb.Model{Name: doc.Name, Id: doc.Id.Hex()}
	return err
}

func (o *Service) FindDocsAndCount(ctx context.Context, req *pb.ReqFind, rsp *pb.ResDocsAndCount) error {
	condition, sorts, skip, limit := o.getComdition(req)
	docs, err := o.dao.Find(condition, sorts, skip, limit)
	rsp.Status = o.OK
	for i := 0; i < len(docs); i++ {
		var doc = docs[i]
		rsp.Data.Docs = append(rsp.Data.Docs, &pb.Model{Name: doc.Name, Id: doc.Id.Hex()})
	}
	count, err := o.dao.Count(condition)
	rsp.Data.Count = o.ParseIntToInt64(count)
	return err
}

func (o *Service) IncByName(ctx context.Context, req *pb.Model, rsp *pb.ResDoc) error {
	err := o.dao.Upsert(bson.M{"name": req.Name}, bson.M{"$inc": bson.M{"value": 1}, "$set": bson.M{"lastModTime": time.Now()}})
	rsp.Status = o.OK
	// rsp.Data = &pb.Model{Name: doc.Name, Id: doc.Id.Hex()}
	return err
}

func Register(service micro.Service) {
	// Register handler
	handler := new(Service)
	handler.NewServiceStatus()
	d := dao.SequenceDao{}
	handler.dao = d.NewDao()
	pb.RegisterSequenceHandler(service.Server(), handler)
	fmt.Println("Sequence service Registered")
}
