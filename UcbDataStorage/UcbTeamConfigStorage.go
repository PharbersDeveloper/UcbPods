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

type UcbTeamConfigStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbTeamConfigStorage) NewTeamConfigStorage(args []BmDaemons.BmDaemon) *UcbTeamConfigStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbTeamConfigStorage{mdb}
}

func (s UcbTeamConfigStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.TeamConfig {
	in := UcbModel.TeamConfig{}
	var out []UcbModel.TeamConfig
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.TeamConfig
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

func (s UcbTeamConfigStorage) GetOne(id string) (UcbModel.TeamConfig, error) {
	in := UcbModel.TeamConfig{ID: id}
	out := UcbModel.TeamConfig{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("TeamConfig for id %s not found", id)
	return UcbModel.TeamConfig{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *UcbTeamConfigStorage) Insert(c UcbModel.TeamConfig) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbTeamConfigStorage) Delete(id string) error {
	in := UcbModel.TeamConfig{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("TeamConfig with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *UcbTeamConfigStorage) Update(c UcbModel.TeamConfig) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("TeamConfig with id does not exist")
	}

	return nil
}

func (s *UcbTeamConfigStorage) Count(req api2go.Request, c UcbModel.TeamConfig) int {
	r, _ := s.db.Count(req, &c)
	return r
}