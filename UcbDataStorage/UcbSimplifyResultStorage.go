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

// UcbSimplifyResultStorage stores all of the tasty modelleaf, needs to be injected into
// SimplifyResult and SimplifyResult Resource. In the real world, you would use a database for that.
type UcbSimplifyResultStorage struct {
	Policies map[string]*UcbModel.SimplifyResult
	idCount  int

	db *BmMongodb.BmMongodb
}

func (s UcbSimplifyResultStorage) NewSimplifyResultStorage(args []BmDaemons.BmDaemon) *UcbSimplifyResultStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbSimplifyResultStorage{make(map[string]*UcbModel.SimplifyResult), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbSimplifyResultStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.SimplifyResult {
	in := UcbModel.SimplifyResult{}
	var out []UcbModel.SimplifyResult
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
func (s UcbSimplifyResultStorage) GetOne(id string) (UcbModel.SimplifyResult, error) {
	in := UcbModel.SimplifyResult{ID: id}
	out := UcbModel.SimplifyResult{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("SimplifyResult for id %s not found", id)
	return UcbModel.SimplifyResult{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbSimplifyResultStorage) Insert(c UcbModel.SimplifyResult) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbSimplifyResultStorage) Delete(id string) error {
	in := UcbModel.SimplifyResult{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("SimplifyResult with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbSimplifyResultStorage) Update(c UcbModel.SimplifyResult) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("SimplifyResult with id does not exist")
	}

	return nil
}
