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

// UcbResourceAssignsResultStorage stores all of the tasty modelleaf, needs to be injected into
// ResourceAssignsResult and ResourceAssignsResult Resource. In the real world, you would use a database for that.
type UcbResourceAssignsResultStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbResourceAssignsResultStorage) NewResourceAssignsResultStorage(args []BmDaemons.BmDaemon) *UcbResourceAssignsResultStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbResourceAssignsResultStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbResourceAssignsResultStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.ResourceAssignsResult {
	in := UcbModel.ResourceAssignsResult{}
	var out []UcbModel.ResourceAssignsResult
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
func (s UcbResourceAssignsResultStorage) GetOne(id string) (UcbModel.ResourceAssignsResult, error) {
	in := UcbModel.ResourceAssignsResult{ID: id}
	out := UcbModel.ResourceAssignsResult{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ResourceAssignsResult for id %s not found", id)
	return UcbModel.ResourceAssignsResult{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbResourceAssignsResultStorage) Insert(c UcbModel.ResourceAssignsResult) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbResourceAssignsResultStorage) Delete(id string) error {
	in := UcbModel.ResourceAssignsResult{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ResourceAssignsResult with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbResourceAssignsResultStorage) Update(c UcbModel.ResourceAssignsResult) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ResourceAssignsResult with id does not exist")
	}

	return nil
}
