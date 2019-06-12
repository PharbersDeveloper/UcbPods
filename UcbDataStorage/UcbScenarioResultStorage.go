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

// UcbScenarioResultStorage stores all of the tasty modelleaf, needs to be injected into
// ScenarioResult and ScenarioResult Resource. In the real world, you would use a database for that.
type UcbScenarioResultStorage struct {
	Policies map[string]*UcbModel.ScenarioResult
	idCount  int

	db *BmMongodb.BmMongodb
}

func (s UcbScenarioResultStorage) NewScenarioResultStorage(args []BmDaemons.BmDaemon) *UcbScenarioResultStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbScenarioResultStorage{make(map[string]*UcbModel.ScenarioResult), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbScenarioResultStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.ScenarioResult {
	in := UcbModel.ScenarioResult{}
	var out []UcbModel.ScenarioResult
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
func (s UcbScenarioResultStorage) GetOne(id string) (UcbModel.ScenarioResult, error) {
	in := UcbModel.ScenarioResult{ID: id}
	out := UcbModel.ScenarioResult{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ScenarioResult for id %s not found", id)
	return UcbModel.ScenarioResult{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbScenarioResultStorage) Insert(c UcbModel.ScenarioResult) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbScenarioResultStorage) Delete(id string) error {
	in := UcbModel.ScenarioResult{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ScenarioResult with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbScenarioResultStorage) Update(c UcbModel.ScenarioResult) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ScenarioResult with id does not exist")
	}

	return nil
}
