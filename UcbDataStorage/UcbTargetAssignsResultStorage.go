package UcbDataStorage

import (
	"fmt"
	"errors"
	"Ucb/UcbModel"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
)

// UcbTargetAssignsResultStorage stores all of the tasty modelleaf, needs to be injected into
// TargetAssignsResult and TargetAssignsResult Resource. In the real world, you would use a database for that.
type UcbTargetAssignsResultStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbTargetAssignsResultStorage) NewTargetAssignsResultStorage(args []BmDaemons.BmDaemon) *UcbTargetAssignsResultStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbTargetAssignsResultStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbTargetAssignsResultStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.TargetAssignsResult {
	in := UcbModel.TargetAssignsResult{}
	var out []UcbModel.TargetAssignsResult
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		for i, iter := range out {
			s.db.ResetIdWithId_(&iter)
			out[i] = iter
		}
		return out
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s UcbTargetAssignsResultStorage) GetOne(id string) (UcbModel.TargetAssignsResult, error) {
	in := UcbModel.TargetAssignsResult{ID: id}
	out := UcbModel.TargetAssignsResult{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("TargetAssignsResult for id %s not found", id)
	return UcbModel.TargetAssignsResult{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbTargetAssignsResultStorage) Insert(c UcbModel.TargetAssignsResult) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbTargetAssignsResultStorage) Delete(id string) error {
	in := UcbModel.TargetAssignsResult{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("TargetAssignsResult with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbTargetAssignsResultStorage) Update(c UcbModel.TargetAssignsResult) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("TargetAssignsResult with id does not exist")
	}

	return nil
}
