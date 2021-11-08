package consumer

import (
	"testing"
	"time"

	"github.com/ozonmp/rcn-service-api/internal/app/repo"
	"github.com/ozonmp/rcn-service-api/internal/model"
)

var testMap = map[uint64]*model.ServiceEvent{
	1:  {ID: 1, Type: model.Created, Status: model.Deferred, Entity: &model.Service{Id: 1, Title: "Service_1", Description: "Created"}},
	2:  {ID: 2, Type: model.Updated, Status: model.Deferred, Entity: &model.Service{Id: 1, Title: "Service_1", Description: "Updated"}},
	3:  {ID: 3, Type: model.Removed, Status: model.Deferred, Entity: &model.Service{Id: 1, Title: "Service_1", Description: "Deleted"}},
	4:  {ID: 4, Type: model.Created, Status: model.Deferred, Entity: &model.Service{Id: 2, Title: "Service_2", Description: "Created"}},
	5:  {ID: 5, Type: model.Created, Status: model.Deferred, Entity: &model.Service{Id: 3, Title: "Service_3", Description: "Created"}},
	6:  {ID: 6, Type: model.Created, Status: model.Deferred, Entity: &model.Service{Id: 2, Title: "Service_2", Description: "Updated"}},
	7:  {ID: 7, Type: model.Created, Status: model.Deferred, Entity: &model.Service{Id: 3, Title: "Service_3", Description: "Updated"}},
	8:  {ID: 8, Type: model.Created, Status: model.Deferred, Entity: &model.Service{Id: 2, Title: "Service_2", Description: "Updated"}},
	9:  {ID: 9, Type: model.Created, Status: model.Deferred, Entity: &model.Service{Id: 3, Title: "Service_3", Description: "Deleted"}},
	10: {ID: 10, Type: model.Created, Status: model.Deferred, Entity: &model.Service{Id: 2, Title: "Service_2", Description: "Deleted"}},
}

func TestStart(t *testing.T) {
	repo := repo.NewDbEventRepo(testMap)
	ch := make(chan model.ServiceEvent)
	cons := NewDbConsumer(2, 1, time.Second, repo, ch)
	cons.Start()
	for i := uint64(1); i <= 10; i++ {
		evnt := <-ch
		if evnt.ID != i {
			t.Error("Events order is broken")
		}
	}
}
