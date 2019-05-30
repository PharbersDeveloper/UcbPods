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

// UcbRepresentativeinputStorage stores all of the tasty modelleaf, needs to be injected into
// Representativeinput and Representativeinput Resource. In the real world, you would use a database for that.
type UcbRepresentativeinputStorage struct {
	images  map[string]*UcbModel.Representativeinput
	idCount int

	db *BmMongodb.BmMongodb
}

func (s UcbRepresentativeinputStorage) NewRepresentativeinputStorage(args []BmDaemons.BmDaemon) *UcbRepresentativeinputStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbRepresentativeinputStorage{make(map[string]*UcbModel.Representativeinput), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbRepresentativeinputStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.Representativeinput {
	in := UcbModel.Representativeinput{}
	var out []UcbModel.Representativeinput
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.Representativeinput
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
func (s UcbRepresentativeinputStorage) GetOne(id string) (UcbModel.Representativeinput, error) {
	in := UcbModel.Representativeinput{ID: id}
	out := UcbModel.Representativeinput{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Representativeinput for id %s not found", id)
	return UcbModel.Representativeinput{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbRepresentativeinputStorage) Insert(c UcbModel.Representativeinput) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbRepresentativeinputStorage) Delete(id string) error {
	in := UcbModel.Representativeinput{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Representativeinput with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbRepresentativeinputStorage) Update(c UcbModel.Representativeinput) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Representativeinput with id does not exist")
	}

	return nil
}

func (s *UcbRepresentativeinputStorage) Count(req api2go.Request, c UcbModel.Representativeinput) int {
	r, _ := s.db.Count(req, &c)
	return r
}