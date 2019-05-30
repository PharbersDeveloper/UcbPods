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

// UcbScenarioStorage stores all of the tasty modelleaf, needs to be injected into
// Scenario and Scenario Resource. In the real world, you would use a database for that.
type UcbScenarioStorage struct {
	Policies map[string]*UcbModel.Scenario
	idCount  int

	db *BmMongodb.BmMongodb
}

func (s UcbScenarioStorage) NewScenarioStorage(args []BmDaemons.BmDaemon) *UcbScenarioStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbScenarioStorage{make(map[string]*UcbModel.Scenario), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbScenarioStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.Scenario {
	in := UcbModel.Scenario{}
	var out []UcbModel.Scenario
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
func (s UcbScenarioStorage) GetOne(id string) (UcbModel.Scenario, error) {
	in := UcbModel.Scenario{ID: id}
	out := UcbModel.Scenario{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Scenario for id %s not found", id)
	return UcbModel.Scenario{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbScenarioStorage) Insert(c UcbModel.Scenario) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbScenarioStorage) Delete(id string) error {
	in := UcbModel.Scenario{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Scenario with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbScenarioStorage) Update(c UcbModel.Scenario) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Scenario with id does not exist")
	}

	return nil
}
