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

// UcbManageTimeResultStorage stores all of the tasty modelleaf, needs to be injected into
// ManageTimeResult and ManageTimeResult Resource. In the real world, you would use a database for that.
type UcbManageTimeResultStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbManageTimeResultStorage) NewManageTimeResultStorage(args []BmDaemons.BmDaemon) *UcbManageTimeResultStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbManageTimeResultStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbManageTimeResultStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.ManageTimeResult {
	in := UcbModel.ManageTimeResult{}
	var out []UcbModel.ManageTimeResult
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
func (s UcbManageTimeResultStorage) GetOne(id string) (UcbModel.ManageTimeResult, error) {
	in := UcbModel.ManageTimeResult{ID: id}
	out := UcbModel.ManageTimeResult{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ManageTimeResult for id %s not found", id)
	return UcbModel.ManageTimeResult{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbManageTimeResultStorage) Insert(c UcbModel.ManageTimeResult) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbManageTimeResultStorage) Delete(id string) error {
	in := UcbModel.ManageTimeResult{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ManageTimeResult with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbManageTimeResultStorage) Update(c UcbModel.ManageTimeResult) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ManageTimeResult with id does not exist")
	}

	return nil
}
