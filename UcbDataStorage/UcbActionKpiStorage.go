package UcbDataStorage

import (
	"Ucb/UcbModel"
	"fmt"
	"errors"

	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
)

// UcbActionKpiStorage stores all of the tasty modelleaf, needs to be injected into
// ActionKpi and ActionKpi Resource. In the real world, you would use a database for that.
type UcbActionKpiStorage struct {

	db *BmMongodb.BmMongodb
}

func (s UcbActionKpiStorage) NewActionKpiStorage(args []BmDaemons.BmDaemon) *UcbActionKpiStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbActionKpiStorage{ mdb}
}

// GetAll of the modelleaf
func (s UcbActionKpiStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.ActionKpi {
	in := UcbModel.ActionKpi{}
	var out []*UcbModel.ActionKpi
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
func (s UcbActionKpiStorage) GetOne(id string) (UcbModel.ActionKpi, error) {
	in := UcbModel.ActionKpi{ID: id}
	out := UcbModel.ActionKpi{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ActionKpi for id %s not found", id)
	return UcbModel.ActionKpi{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbActionKpiStorage) Insert(c UcbModel.ActionKpi) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbActionKpiStorage) Delete(id string) error {
	in := UcbModel.ActionKpi{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ActionKpi with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbActionKpiStorage) Update(c UcbModel.ActionKpi) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ActionKpi with id does not exist")
	}

	return nil
}
