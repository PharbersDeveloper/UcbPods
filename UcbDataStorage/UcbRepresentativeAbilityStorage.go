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

// UcbRepresentativeAbilityStorage stores all of the tasty modelleaf, needs to be injected into
// RepresentativeAbility and RepresentativeAbility Resource. In the real world, you would use a database for that.
type UcbRepresentativeAbilityStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbRepresentativeAbilityStorage) NewRepresentativeAbilityStorage(args []BmDaemons.BmDaemon) *UcbRepresentativeAbilityStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbRepresentativeAbilityStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbRepresentativeAbilityStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.RepresentativeAbility {
	in := UcbModel.RepresentativeAbility{}
	var out []*UcbModel.RepresentativeAbility
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
func (s UcbRepresentativeAbilityStorage) GetOne(id string) (UcbModel.RepresentativeAbility, error) {
	in := UcbModel.RepresentativeAbility{ID: id}
	out := UcbModel.RepresentativeAbility{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("RepresentativeAbility for id %s not found", id)
	return UcbModel.RepresentativeAbility{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbRepresentativeAbilityStorage) Insert(c UcbModel.RepresentativeAbility) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbRepresentativeAbilityStorage) Delete(id string) error {
	in := UcbModel.RepresentativeAbility{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("RepresentativeAbility with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbRepresentativeAbilityStorage) Update(c UcbModel.RepresentativeAbility) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("RepresentativeAbility with id does not exist")
	}

	return nil
}
