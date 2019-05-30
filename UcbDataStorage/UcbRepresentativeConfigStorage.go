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

// UcbRepresentativeConfigStorage stores all of the tasty chocolate, needs to be injected into
// RepresentativeConfig Resource. In the real world, you would use a database for that.
type UcbRepresentativeConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbRepresentativeConfigStorage) NewRepresentativeConfigStorage(args []BmDaemons.BmDaemon) *UcbRepresentativeConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbRepresentativeConfigStorage{mdb}
}

// GetAll of the chocolate
func (s UcbRepresentativeConfigStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.RepresentativeConfig {
	in := UcbModel.RepresentativeConfig{}
	var out []UcbModel.RepresentativeConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.RepresentativeConfig
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
func (s UcbRepresentativeConfigStorage) GetOne(id string) (UcbModel.RepresentativeConfig, error) {
	in := UcbModel.RepresentativeConfig{ID: id}
	out := UcbModel.RepresentativeConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("RepresentativeConfig for id %s not found", id)
	return UcbModel.RepresentativeConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbRepresentativeConfigStorage) Insert(c UcbModel.RepresentativeConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbRepresentativeConfigStorage) Delete(id string) error {
	in := UcbModel.RepresentativeConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("RepresentativeConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *UcbRepresentativeConfigStorage) Update(c UcbModel.RepresentativeConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("RepresentativeConfig with id does not exist")
	}

	return nil
}

func (s *UcbRepresentativeConfigStorage) Count(req api2go.Request, c UcbModel.RepresentativeConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}
