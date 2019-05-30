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

// UcbResourceConfigStorage stores all of the tasty chocolate, needs to be injected into
// ResourceConfig Resource. In the real world, you would use a database for that.
type UcbResourceConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbResourceConfigStorage) NewResourceConfigStorage(args []BmDaemons.BmDaemon) *UcbResourceConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbResourceConfigStorage{mdb}
}

// GetAll of the chocolate
func (s UcbResourceConfigStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.ResourceConfig {
	in := UcbModel.ResourceConfig{}
	var out []UcbModel.ResourceConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.ResourceConfig
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
func (s UcbResourceConfigStorage) GetOne(id string) (UcbModel.ResourceConfig, error) {
	in := UcbModel.ResourceConfig{ID: id}
	out := UcbModel.ResourceConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ResourceConfig for id %s not found", id)
	return UcbModel.ResourceConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbResourceConfigStorage) Insert(c UcbModel.ResourceConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbResourceConfigStorage) Delete(id string) error {
	in := UcbModel.ResourceConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ResourceConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *UcbResourceConfigStorage) Update(c UcbModel.ResourceConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ResourceConfig with id does not exist")
	}

	return nil
}

func (s *UcbResourceConfigStorage) Count(req api2go.Request, c UcbModel.ResourceConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}
