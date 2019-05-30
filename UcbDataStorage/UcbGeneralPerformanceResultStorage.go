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

// UcbGeneralPerformanceResultStorage stores all of the tasty modelleaf, needs to be injected into
// GeneralPerformanceResult and GeneralPerformanceResult Resource. In the real world, you would use a database for that.
type UcbGeneralPerformanceResultStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbGeneralPerformanceResultStorage) NewGeneralPerformanceResultStorage(args []BmDaemons.BmDaemon) *UcbGeneralPerformanceResultStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbGeneralPerformanceResultStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbGeneralPerformanceResultStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.GeneralPerformanceResult {
	in := UcbModel.GeneralPerformanceResult{}
	var out []UcbModel.GeneralPerformanceResult
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
func (s UcbGeneralPerformanceResultStorage) GetOne(id string) (UcbModel.GeneralPerformanceResult, error) {
	in := UcbModel.GeneralPerformanceResult{ID: id}
	out := UcbModel.GeneralPerformanceResult{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("GeneralPerformanceResult for id %s not found", id)
	return UcbModel.GeneralPerformanceResult{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbGeneralPerformanceResultStorage) Insert(c UcbModel.GeneralPerformanceResult) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbGeneralPerformanceResultStorage) Delete(id string) error {
	in := UcbModel.GeneralPerformanceResult{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("GeneralPerformanceResult with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbGeneralPerformanceResultStorage) Update(c UcbModel.GeneralPerformanceResult) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("GeneralPerformanceResult with id does not exist")
	}

	return nil
}
