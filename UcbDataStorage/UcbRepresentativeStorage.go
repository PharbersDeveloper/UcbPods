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

type UcbRepresentativeStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbRepresentativeStorage) NewRepresentativeStorage(args []BmDaemons.BmDaemon) *UcbRepresentativeStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbRepresentativeStorage{mdb}
}

func (s UcbRepresentativeStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.Representative {
	in := UcbModel.Representative{}
	var out []UcbModel.Representative
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.Representative
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

func (s UcbRepresentativeStorage) GetOne(id string) (UcbModel.Representative, error) {
	in := UcbModel.Representative{ID: id}
	out := UcbModel.Representative{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Representative for id %s not found", id)
	return UcbModel.Representative{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *UcbRepresentativeStorage) Insert(c UcbModel.Representative) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbRepresentativeStorage) Delete(id string) error {
	in := UcbModel.Representative{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Representative with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *UcbRepresentativeStorage) Update(c UcbModel.Representative) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Representative with id does not exist")
	}

	return nil
}

func (s *UcbRepresentativeStorage) Count(req api2go.Request, c UcbModel.Representative) int {
	r, _ := s.db.Count(req, &c)
	return r
}
