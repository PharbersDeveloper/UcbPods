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

// UcbRegionalDivisionResultStorage stores all of the tasty modelleaf, needs to be injected into
// RegionalDivisionResult and RegionalDivisionResult Resource. In the real world, you would use a database for that.
type UcbRegionalDivisionResultStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbRegionalDivisionResultStorage) NewRegionalDivisionResultStorage(args []BmDaemons.BmDaemon) *UcbRegionalDivisionResultStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbRegionalDivisionResultStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbRegionalDivisionResultStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.RegionalDivisionResult {
	in := UcbModel.RegionalDivisionResult{}
	var out []UcbModel.RegionalDivisionResult
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
func (s UcbRegionalDivisionResultStorage) GetOne(id string) (UcbModel.RegionalDivisionResult, error) {
	in := UcbModel.RegionalDivisionResult{ID: id}
	out := UcbModel.RegionalDivisionResult{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("RegionalDivisionResult for id %s not found", id)
	return UcbModel.RegionalDivisionResult{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbRegionalDivisionResultStorage) Insert(c UcbModel.RegionalDivisionResult) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbRegionalDivisionResultStorage) Delete(id string) error {
	in := UcbModel.RegionalDivisionResult{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("RegionalDivisionResult with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbRegionalDivisionResultStorage) Update(c UcbModel.RegionalDivisionResult) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("RegionalDivisionResult with id does not exist")
	}

	return nil
}
