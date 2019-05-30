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

// UcbTitleStorage stores all of the tasty modelleaf, needs to be injected into
// Title and Title Resource. In the real world, you would use a database for that.
type UcbTitleStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbTitleStorage) NewTitleStorage(args []BmDaemons.BmDaemon) *UcbTitleStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbTitleStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbTitleStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.Title {
	in := UcbModel.Title{}
	var out []UcbModel.Title
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
func (s UcbTitleStorage) GetOne(id string) (UcbModel.Title, error) {
	in := UcbModel.Title{ID: id}
	out := UcbModel.Title{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Title for id %s not found", id)
	return UcbModel.Title{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbTitleStorage) Insert(c UcbModel.Title) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbTitleStorage) Delete(id string) error {
	in := UcbModel.Title{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Title with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbTitleStorage) Update(c UcbModel.Title) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Title with id does not exist")
	}

	return nil
}
