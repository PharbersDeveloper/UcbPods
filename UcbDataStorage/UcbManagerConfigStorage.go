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

// UcbManagerConfigStorage stores all of the tasty modelleaf, needs to be injected into
// ManagerConfig and ManagerConfig Resource. In the real world, you would use a database for that.
type UcbManagerConfigStorage struct {
	images  map[string]*UcbModel.ManagerConfig
	idCount int

	db *BmMongodb.BmMongodb
}

func (s UcbManagerConfigStorage) NewManagerConfigStorage(args []BmDaemons.BmDaemon) *UcbManagerConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbManagerConfigStorage{make(map[string]*UcbModel.ManagerConfig), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbManagerConfigStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.ManagerConfig {
	in := UcbModel.ManagerConfig{}
	out := []UcbModel.ManagerConfig{}
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
func (s UcbManagerConfigStorage) GetOne(id string) (UcbModel.ManagerConfig, error) {
	in := UcbModel.ManagerConfig{ID: id}
	out := UcbModel.ManagerConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ManagerConfig for id %s not found", id)
	return UcbModel.ManagerConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbManagerConfigStorage) Insert(c UcbModel.ManagerConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbManagerConfigStorage) Delete(id string) error {
	in := UcbModel.ManagerConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ManagerConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbManagerConfigStorage) Update(c UcbModel.ManagerConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ManagerConfig with id does not exist")
	}

	return nil
}
