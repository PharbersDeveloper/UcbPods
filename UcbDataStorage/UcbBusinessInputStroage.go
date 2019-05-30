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

// UcbBusinessinputStorage stores all of the tasty modelleaf, needs to be injected into
// Businessinput and Businessinput Resource. In the real world, you would use a database for that.
type UcbBusinessinputStorage struct {
	images  map[string]*UcbModel.Businessinput
	idCount int

	db *BmMongodb.BmMongodb
}

func (s UcbBusinessinputStorage) NewBusinessinputStorage(args []BmDaemons.BmDaemon) *UcbBusinessinputStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbBusinessinputStorage{make(map[string]*UcbModel.Businessinput), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbBusinessinputStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.Businessinput {
	in := UcbModel.Businessinput{}
	var out []UcbModel.Businessinput
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.Businessinput
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s UcbBusinessinputStorage) GetOne(id string) (UcbModel.Businessinput, error) {
	in := UcbModel.Businessinput{ID: id}
	out := UcbModel.Businessinput{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Businessinput for id %s not found", id)
	return UcbModel.Businessinput{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbBusinessinputStorage) Insert(c UcbModel.Businessinput) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbBusinessinputStorage) Delete(id string) error {
	in := UcbModel.Businessinput{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Businessinput with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbBusinessinputStorage) Update(c UcbModel.Businessinput) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Businessinput with id does not exist")
	}

	return nil
}

func (s *UcbBusinessinputStorage) Count(req api2go.Request, c UcbModel.Businessinput) int {
	r, _ := s.db.Count(req, &c)
	return r
}