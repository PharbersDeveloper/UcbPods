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

type UcbRegionStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbRegionStorage) NewRegionStorage(args []BmDaemons.BmDaemon) *UcbRegionStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbRegionStorage{mdb}
}

func (s UcbRegionStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.Region {
	in := UcbModel.Region{}
	var out []UcbModel.Region
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.Region
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

func (s UcbRegionStorage) GetOne(id string) (UcbModel.Region, error) {
	in := UcbModel.Region{ID: id}
	out := UcbModel.Region{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Region for id %s not found", id)
	return UcbModel.Region{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *UcbRegionStorage) Insert(c UcbModel.Region) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbRegionStorage) Delete(id string) error {
	in := UcbModel.Region{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Region with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *UcbRegionStorage) Update(c UcbModel.Region) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Region with id does not exist")
	}

	return nil
}

func (s *UcbRegionStorage) Count(req api2go.Request, c UcbModel.Region) int {
	r, _ := s.db.Count(req, &c)
	return r
}