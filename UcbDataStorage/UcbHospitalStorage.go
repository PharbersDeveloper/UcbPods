package UcbDataStorage

import (
	"errors"
	"fmt"
	"Ucb/UcbModel"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
	"net/http"
)

type UcbHospitalStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbHospitalStorage) NewHospitalStorage(args []BmDaemons.BmDaemon) *UcbHospitalStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbHospitalStorage{mdb}
}

func (s UcbHospitalStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.Hospital {
	in := UcbModel.Hospital{}
	var out []UcbModel.Hospital
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.Hospital
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil
	}
}

func (s UcbHospitalStorage) GetOne(id string) (UcbModel.Hospital, error) {
	in := UcbModel.Hospital{ID: id}
	out := UcbModel.Hospital{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Hospital for id %s not found", id)
	return UcbModel.Hospital{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *UcbHospitalStorage) Insert(c UcbModel.Hospital) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbHospitalStorage) Delete(id string) error {
	in := UcbModel.Hospital{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Hospital with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *UcbHospitalStorage) Update(c UcbModel.Hospital) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Hospital with id does not exist")
	}

	return nil
}

func (s *UcbHospitalStorage) Count(req api2go.Request, c UcbModel.Hospital) int {
	r, _ := s.db.Count(req, &c)
	return r
}