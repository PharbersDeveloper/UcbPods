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

// UcbImageStorage stores all of the tasty modelleaf, needs to be injected into
// Image and Image Resource. In the real world, you would use a database for that.
type UcbImageStorage struct {
	images  map[string]*UcbModel.Image
	idCount int

	db *BmMongodb.BmMongodb
}

func (s UcbImageStorage) NewImageStorage(args []BmDaemons.BmDaemon) *UcbImageStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbImageStorage{make(map[string]*UcbModel.Image), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbImageStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.Image {
	in := UcbModel.Image{}
	var out []UcbModel.Image
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		for i, iter := range out {
			s.db.ResetIdWithId_(&iter)
			out[i] = iter
		}
		return out
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s UcbImageStorage) GetOne(id string) (UcbModel.Image, error) {
	in := UcbModel.Image{ID: id}
	out := UcbModel.Image{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Image for id %s not found", id)
	return UcbModel.Image{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbImageStorage) Insert(c UcbModel.Image) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbImageStorage) Delete(id string) error {
	in := UcbModel.Image{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Image with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbImageStorage) Update(c UcbModel.Image) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Image with id does not exist")
	}

	return nil
}
