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

type UcbPaperStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbPaperStorage) NewPaperStorage(args []BmDaemons.BmDaemon) *UcbPaperStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbPaperStorage{mdb}
}

func (s UcbPaperStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.Paper {
	in := UcbModel.Paper{}
	var out []UcbModel.Paper
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.Paper
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

func (s UcbPaperStorage) GetOne(id string) (UcbModel.Paper, error) {
	in := UcbModel.Paper{ID: id}
	out := UcbModel.Paper{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Paper for id %s not found", id)
	return UcbModel.Paper{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *UcbPaperStorage) Insert(c UcbModel.Paper) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbPaperStorage) Delete(id string) error {
	in := UcbModel.Paper{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Paper with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *UcbPaperStorage) Update(c UcbModel.Paper) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Paper with id does not exist")
	}

	return nil
}

func (s *UcbPaperStorage) Count(req api2go.Request, c UcbModel.Paper) int {
	r, _ := s.db.Count(req, &c)
	return r
}