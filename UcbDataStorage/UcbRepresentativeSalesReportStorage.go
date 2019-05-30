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

// UcbRepresentativeSalesReportStorage stores all of the tasty modelleaf, needs to be injected into
// RepresentativeSalesReport and RepresentativeSalesReport Resource. In the real world, you would use a database for that.
type UcbRepresentativeSalesReportStorage struct {
	SalesConfigs  map[string]*UcbModel.RepresentativeSalesReport
	idCount int

	db *BmMongodb.BmMongodb
}

func (s UcbRepresentativeSalesReportStorage) NewRepresentativeSalesReportStorage(args []BmDaemons.BmDaemon) *UcbRepresentativeSalesReportStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbRepresentativeSalesReportStorage{make(map[string]*UcbModel.RepresentativeSalesReport), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbRepresentativeSalesReportStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.RepresentativeSalesReport {
	in := UcbModel.RepresentativeSalesReport{}
	var out []*UcbModel.RepresentativeSalesReport
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
func (s UcbRepresentativeSalesReportStorage) GetOne(id string) (UcbModel.RepresentativeSalesReport, error) {
	in := UcbModel.RepresentativeSalesReport{ID: id}
	out := UcbModel.RepresentativeSalesReport{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("RepresentativeSalesReport for id %s not found", id)
	return UcbModel.RepresentativeSalesReport{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbRepresentativeSalesReportStorage) Insert(c UcbModel.RepresentativeSalesReport) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbRepresentativeSalesReportStorage) Delete(id string) error {
	in := UcbModel.RepresentativeSalesReport{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("RepresentativeSalesReport with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbRepresentativeSalesReportStorage) Update(c UcbModel.RepresentativeSalesReport) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("RepresentativeSalesReport with id does not exist")
	}

	return nil
}
