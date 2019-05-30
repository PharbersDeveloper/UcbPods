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

// UcbGoodsConfigStorage stores all of the tasty chocolate, needs to be injected into
// GoodsConfig Goods. In the real world, you would use a database for that.
type UcbGoodsConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbGoodsConfigStorage) NewGoodsConfigStorage(args []BmDaemons.BmDaemon) *UcbGoodsConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbGoodsConfigStorage{mdb}
}

// GetAll of the chocolate
func (s UcbGoodsConfigStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.GoodsConfig {
	in := UcbModel.GoodsConfig{}
	var out []UcbModel.GoodsConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.GoodsConfig
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
func (s UcbGoodsConfigStorage) GetOne(id string) (UcbModel.GoodsConfig, error) {
	in := UcbModel.GoodsConfig{ID: id}
	out := UcbModel.GoodsConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("GoodsConfig for id %s not found", id)
	return UcbModel.GoodsConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbGoodsConfigStorage) Insert(c UcbModel.GoodsConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbGoodsConfigStorage) Delete(id string) error {
	in := UcbModel.GoodsConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("GoodsConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *UcbGoodsConfigStorage) Update(c UcbModel.GoodsConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("GoodsConfig with id does not exist")
	}

	return nil
}

func (s *UcbGoodsConfigStorage) Count(req api2go.Request, c UcbModel.GoodsConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}
