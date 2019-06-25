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

// UcbEvaluationStorage stores all of the tasty modelleaf, needs to be injected into
// Evaluation and Evaluation Resource. In the real world, you would use a database for that.
type UcbEvaluationStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbEvaluationStorage) NewEvaluationStorage(args []BmDaemons.BmDaemon) *UcbEvaluationStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbEvaluationStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbEvaluationStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.Evaluation {
	in := UcbModel.Evaluation{}
	var out []UcbModel.Evaluation
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
func (s UcbEvaluationStorage) GetOne(id string) (UcbModel.Evaluation, error) {
	in := UcbModel.Evaluation{ID: id}
	out := UcbModel.Evaluation{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Evaluation for id %s not found", id)
	return UcbModel.Evaluation{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbEvaluationStorage) Insert(c UcbModel.Evaluation) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbEvaluationStorage) Delete(id string) error {
	in := UcbModel.Evaluation{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Evaluation with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbEvaluationStorage) Update(c UcbModel.Evaluation) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Evaluation with id does not exist")
	}

	return nil
}
