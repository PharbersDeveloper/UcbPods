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

// UcbManagerinputStorage stores all of the tasty modelleaf, needs to be injected into
// Managerinput and Managerinput Resource. In the real world, you would use a database for that.
type UcbManagerinputStorage struct {
	images  map[string]*UcbModel.Managerinput
	idCount int

	db *BmMongodb.BmMongodb
}

func (s UcbManagerinputStorage) NewManagerinputStorage(args []BmDaemons.BmDaemon) *UcbManagerinputStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbManagerinputStorage{make(map[string]*UcbModel.Managerinput), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbManagerinputStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.Managerinput {
	in := UcbModel.Managerinput{}
	var out []UcbModel.Managerinput
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.Managerinput
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
func (s UcbManagerinputStorage) GetOne(id string) (UcbModel.Managerinput, error) {
	in := UcbModel.Managerinput{ID: id}
	out := UcbModel.Managerinput{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Managerinput for id %s not found", id)
	return UcbModel.Managerinput{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbManagerinputStorage) Insert(c UcbModel.Managerinput) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbManagerinputStorage) Delete(id string) error {
	in := UcbModel.Managerinput{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Managerinput with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbManagerinputStorage) Update(c UcbModel.Managerinput) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Managerinput with id does not exist")
	}

	return nil
}

func (s *UcbManagerinputStorage) Count(req api2go.Request, c UcbModel.Managerinput) int {
	r, _ := s.db.Count(req, &c)
	return r
}