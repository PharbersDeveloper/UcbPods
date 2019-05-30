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

// UcbHospitalConfigStorage stores all of the tasty chocolate, needs to be injected into
// HospitalConfig Resource. In the real world, you would use a database for that.
type UcbHospitalConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbHospitalConfigStorage) NewHospitalConfigStorage(args []BmDaemons.BmDaemon) *UcbHospitalConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbHospitalConfigStorage{mdb}
}

// GetAll of the chocolate
func (s UcbHospitalConfigStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.HospitalConfig {
	in := UcbModel.HospitalConfig{}
	var out []UcbModel.HospitalConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.HospitalConfig
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
func (s UcbHospitalConfigStorage) GetOne(id string) (UcbModel.HospitalConfig, error) {
	in := UcbModel.HospitalConfig{ID: id}
	out := UcbModel.HospitalConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("HospitalConfig for id %s not found", id)
	return UcbModel.HospitalConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbHospitalConfigStorage) Insert(c UcbModel.HospitalConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbHospitalConfigStorage) Delete(id string) error {
	in := UcbModel.HospitalConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("HospitalConfig with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing chocolate
func (s *UcbHospitalConfigStorage) Update(c UcbModel.HospitalConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("HospitalConfig with id does not exist")
	}

	return nil
}

func (s *UcbHospitalConfigStorage) Count(req api2go.Request, c UcbModel.HospitalConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}
