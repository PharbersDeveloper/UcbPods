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

// UcbManagerGoodsConfigStorage stores all of the tasty modelleaf, needs to be injected into
// ManagerGoodsConfig and ManagerGoodsConfig Resource. In the real world, you would use a database for that.
type UcbManagerGoodsConfigStorage struct {
	ManagerGoodsConfigs map[string]*UcbModel.ManagerGoodsConfig
	idCount     int

	db *BmMongodb.BmMongodb
}

func (s UcbManagerGoodsConfigStorage) NewManagerGoodsConfigStorage(args []BmDaemons.BmDaemon) *UcbManagerGoodsConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbManagerGoodsConfigStorage{make(map[string]*UcbModel.ManagerGoodsConfig), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbManagerGoodsConfigStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.ManagerGoodsConfig {
	in := UcbModel.ManagerGoodsConfig{}
	var out []*UcbModel.ManagerGoodsConfig
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
func (s UcbManagerGoodsConfigStorage) GetOne(id string) (UcbModel.ManagerGoodsConfig, error) {
	in := UcbModel.ManagerGoodsConfig{ID: id}
	out := UcbModel.ManagerGoodsConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ManagerGoodsConfig for id %s not found", id)
	return UcbModel.ManagerGoodsConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbManagerGoodsConfigStorage) Insert(c UcbModel.ManagerGoodsConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbManagerGoodsConfigStorage) Delete(id string) error {
	in := UcbModel.ManagerGoodsConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ManagerGoodsConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbManagerGoodsConfigStorage) Update(c UcbModel.ManagerGoodsConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ManagerGoodsConfig with id does not exist")
	}

	return nil
}
