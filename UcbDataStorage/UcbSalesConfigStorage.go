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

// UcbSalesConfigStorage stores all of the tasty modelleaf, needs to be injected into
// SalesConfig and SalesConfig Resource. In the real world, you would use a database for that.
type UcbSalesConfigStorage struct {
	SalesConfigs  map[string]*UcbModel.SalesConfig
	idCount int

	db *BmMongodb.BmMongodb
}

func (s UcbSalesConfigStorage) NewSalesConfigStorage(args []BmDaemons.BmDaemon) *UcbSalesConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbSalesConfigStorage{make(map[string]*UcbModel.SalesConfig), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbSalesConfigStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.SalesConfig {
	in := UcbModel.SalesConfig{}
	var out []*UcbModel.SalesConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		for i, iter := range out {
			s.db.ResetIdWithId_(iter)
			out[i] = iter
		}
		return out
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s UcbSalesConfigStorage) GetOne(id string) (UcbModel.SalesConfig, error) {
	in := UcbModel.SalesConfig{ID: id}
	out := UcbModel.SalesConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("SalesConfig for id %s not found", id)
	return UcbModel.SalesConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbSalesConfigStorage) Insert(c UcbModel.SalesConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbSalesConfigStorage) Delete(id string) error {
	in := UcbModel.SalesConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("SalesConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbSalesConfigStorage) Update(c UcbModel.SalesConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("SalesConfig with id does not exist")
	}

	return nil
}
