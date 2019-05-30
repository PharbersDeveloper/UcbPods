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

// UcbPaperinputStorage stores all of the tasty modelleaf, needs to be injected into
// Paperinput and Paperinput Resource. In the real world, you would use a database for that.
type UcbPaperinputStorage struct {
	images  map[string]*UcbModel.Paperinput
	idCount int

	db *BmMongodb.BmMongodb
}

func (s UcbPaperinputStorage) NewPaperinputStorage(args []BmDaemons.BmDaemon) *UcbPaperinputStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbPaperinputStorage{make(map[string]*UcbModel.Paperinput), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbPaperinputStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.Paperinput {
	in := UcbModel.Paperinput{}
	var out []UcbModel.Paperinput
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.Paperinput
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
func (s UcbPaperinputStorage) GetOne(id string) (UcbModel.Paperinput, error) {
	in := UcbModel.Paperinput{ID: id}
	out := UcbModel.Paperinput{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Paperinput for id %s not found", id)
	return UcbModel.Paperinput{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbPaperinputStorage) Insert(c UcbModel.Paperinput) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbPaperinputStorage) Delete(id string) error {
	in := UcbModel.Paperinput{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Paperinput with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbPaperinputStorage) Update(c UcbModel.Paperinput) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Paperinput with id does not exist")
	}

	return nil
}

func (s *UcbPaperinputStorage) Count(req api2go.Request, c UcbModel.Paperinput) int {
	r, _ := s.db.Count(req, &c)
	return r
}