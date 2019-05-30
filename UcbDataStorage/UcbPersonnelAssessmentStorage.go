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

// UcbPersonnelAssessmentStorage stores all of the tasty modelleaf, needs to be injected into
// PersonnelAssessment and PersonnelAssessment Resource. In the real world, you would use a database for that.
type UcbPersonnelAssessmentStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbPersonnelAssessmentStorage) NewPersonnelAssessmentStorage(args []BmDaemons.BmDaemon) *UcbPersonnelAssessmentStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbPersonnelAssessmentStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbPersonnelAssessmentStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.PersonnelAssessment {
	in := UcbModel.PersonnelAssessment{}
	var out []*UcbModel.PersonnelAssessment
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
func (s UcbPersonnelAssessmentStorage) GetOne(id string) (UcbModel.PersonnelAssessment, error) {
	in := UcbModel.PersonnelAssessment{ID: id}
	out := UcbModel.PersonnelAssessment{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("PersonnelAssessment for id %s not found", id)
	return UcbModel.PersonnelAssessment{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbPersonnelAssessmentStorage) Insert(c UcbModel.PersonnelAssessment) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbPersonnelAssessmentStorage) Delete(id string) error {
	in := UcbModel.PersonnelAssessment{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("PersonnelAssessment with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbPersonnelAssessmentStorage) Update(c UcbModel.PersonnelAssessment) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("PersonnelAssessment with id does not exist")
	}

	return nil
}

