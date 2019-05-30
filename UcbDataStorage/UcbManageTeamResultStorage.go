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

// UcbManageTeamResultStorage stores all of the tasty modelleaf, needs to be injected into
// ManageTeamResult and ManageTeamResult Resource. In the real world, you would use a database for that.
type UcbManageTeamResultStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbManageTeamResultStorage) NewManageTeamResultStorage(args []BmDaemons.BmDaemon) *UcbManageTeamResultStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbManageTeamResultStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbManageTeamResultStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.ManageTeamResult {
	in := UcbModel.ManageTeamResult{}
	var out []UcbModel.ManageTeamResult
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
func (s UcbManageTeamResultStorage) GetOne(id string) (UcbModel.ManageTeamResult, error) {
	in := UcbModel.ManageTeamResult{ID: id}
	out := UcbModel.ManageTeamResult{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ManageTeamResult for id %s not found", id)
	return UcbModel.ManageTeamResult{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbManageTeamResultStorage) Insert(c UcbModel.ManageTeamResult) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbManageTeamResultStorage) Delete(id string) error {
	in := UcbModel.ManageTeamResult{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ManageTeamResult with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbManageTeamResultStorage) Update(c UcbModel.ManageTeamResult) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ManageTeamResult with id does not exist")
	}

	return nil
}
