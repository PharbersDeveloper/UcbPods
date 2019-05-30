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

// UcbDestConfigStorage stores all of the tasty chocolate, needs to be injected into
// DestConfig Dest. In the real world, you would use a database for that.
type UcbDestConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbDestConfigStorage) NewDestConfigStorage(args []BmDaemons.BmDaemon) *UcbDestConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbDestConfigStorage{mdb}
}

// GetAll of the chocolate
func (s UcbDestConfigStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.DestConfig {
	in := UcbModel.DestConfig{}
	var out []UcbModel.DestConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.DestConfig
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
func (s UcbDestConfigStorage) GetOne(id string) (UcbModel.DestConfig, error) {
	in := UcbModel.DestConfig{ID: id}
	out := UcbModel.DestConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("DestConfig for id %s not found", id)
	return UcbModel.DestConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbDestConfigStorage) Insert(c UcbModel.DestConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbDestConfigStorage) Delete(id string) error {
	in := UcbModel.DestConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("DestConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *UcbDestConfigStorage) Update(c UcbModel.DestConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("DestConfig with id does not exist")
	}

	return nil
}

func (s *UcbDestConfigStorage) Count(req api2go.Request, c UcbModel.DestConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}
