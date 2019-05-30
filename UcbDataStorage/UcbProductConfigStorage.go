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

// UcbProductConfigStorage stores all of the tasty chocolate, needs to be injected into
// ProductConfig Resource. In the real world, you would use a database for that.
type UcbProductConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbProductConfigStorage) NewProductConfigStorage(args []BmDaemons.BmDaemon) *UcbProductConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbProductConfigStorage{mdb}
}

// GetAll of the chocolate
func (s UcbProductConfigStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.ProductConfig {
	in := UcbModel.ProductConfig{}
	var out []UcbModel.ProductConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.ProductConfig
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
func (s UcbProductConfigStorage) GetOne(id string) (UcbModel.ProductConfig, error) {
	in := UcbModel.ProductConfig{ID: id}
	out := UcbModel.ProductConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ProductConfig for id %s not found", id)
	return UcbModel.ProductConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbProductConfigStorage) Insert(c UcbModel.ProductConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbProductConfigStorage) Delete(id string) error {
	in := UcbModel.ProductConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ProductConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *UcbProductConfigStorage) Update(c UcbModel.ProductConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ProductConfig with id does not exist")
	}

	return nil
}

func (s *UcbProductConfigStorage) Count(req api2go.Request, c UcbModel.ProductConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}
