package UcbDataStorage

import (
	"errors"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"Ucb/UcbModel"
	"github.com/manyminds/api2go"
	"net/http"
)

// UcbGoodsinputStorage stores all of the tasty chocolate, needs to be injected into
// Goodsinput Goods. In the real world, you would use a database for that.
type UcbGoodsinputStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbGoodsinputStorage) NewGoodsinputStorage(args []BmDaemons.BmDaemon) *UcbGoodsinputStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbGoodsinputStorage{mdb}
}

// GetAll of the chocolate
func (s UcbGoodsinputStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.Goodsinput {
	in := UcbModel.Goodsinput{}
	var out []UcbModel.Goodsinput
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.Goodsinput
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Student)
	}
}

// GetOne
func (s UcbGoodsinputStorage) GetOne(id string) (UcbModel.Goodsinput, error) {
	in := UcbModel.Goodsinput{ID: id}
	out := UcbModel.Goodsinput{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Goodsinput for id %s not found", id)
	return UcbModel.Goodsinput{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbGoodsinputStorage) Insert(c UcbModel.Goodsinput) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbGoodsinputStorage) Delete(id string) error {
	in := UcbModel.Goodsinput{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Goodsinput with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *UcbGoodsinputStorage) Update(c UcbModel.Goodsinput) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Goodsinput with id does not exist")
	}

	return nil
}

func (s *UcbGoodsinputStorage) Count(req api2go.Request, c UcbModel.Goodsinput) int {
	r, _ := s.db.Count(req, &c)
	return r
}
