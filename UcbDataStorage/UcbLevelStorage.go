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

// UcbLevelStorage stores all of the tasty modelleaf, needs to be injected into
// Level and Level Resource. In the real world, you would use a database for that.
type UcbLevelStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbLevelStorage) NewLevelStorage(args []BmDaemons.BmDaemon) *UcbLevelStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbLevelStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbLevelStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.Level {
	in := UcbModel.Level{}
	var out []UcbModel.Level
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
func (s UcbLevelStorage) GetOne(id string) (UcbModel.Level, error) {
	in := UcbModel.Level{ID: id}
	out := UcbModel.Level{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Level for id %s not found", id)
	return UcbModel.Level{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbLevelStorage) Insert(c UcbModel.Level) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbLevelStorage) Delete(id string) error {
	in := UcbModel.Level{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Level with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbLevelStorage) Update(c UcbModel.Level) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Level with id does not exist")
	}

	return nil
}
