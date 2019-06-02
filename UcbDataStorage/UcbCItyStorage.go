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

// UcbCityStorage stores all of the tasty modelleaf, needs to be injected into
// City and City Resource. In the real world, you would use a database for that.
type UcbCityStorage struct {
	Citys map[string]*UcbModel.City
	idCount     int

	db *BmMongodb.BmMongodb
}

func (s UcbCityStorage) NewCityStorage(args []BmDaemons.BmDaemon) *UcbCityStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbCityStorage{make(map[string]*UcbModel.City), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbCityStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.City {
	in := UcbModel.City{}
	var out []*UcbModel.City
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
func (s UcbCityStorage) GetOne(id string) (UcbModel.City, error) {
	in := UcbModel.City{ID: id}
	out := UcbModel.City{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("City for id %s not found", id)
	return UcbModel.City{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbCityStorage) Insert(c UcbModel.City) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbCityStorage) Delete(id string) error {
	in := UcbModel.City{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("City with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbCityStorage) Update(c UcbModel.City) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("City with id does not exist")
	}

	return nil
}
