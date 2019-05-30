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

type UcbProductStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbProductStorage) NewProductStorage(args []BmDaemons.BmDaemon) *UcbProductStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbProductStorage{mdb}
}

func (s UcbProductStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.Product {
	in := UcbModel.Product{}
	var out []UcbModel.Product
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.Product
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

func (s UcbProductStorage) GetOne(id string) (UcbModel.Product, error) {
	in := UcbModel.Product{ID: id}
	out := UcbModel.Product{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Product for id %s not found", id)
	return UcbModel.Product{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *UcbProductStorage) Insert(c UcbModel.Product) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbProductStorage) Delete(id string) error {
	in := UcbModel.Product{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Product with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *UcbProductStorage) Update(c UcbModel.Product) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Product with id does not exist")
	}

	return nil
}

func (s *UcbProductStorage) Count(req api2go.Request, c UcbModel.Product) int {
	r, _ := s.db.Count(req, &c)
	return r
}
