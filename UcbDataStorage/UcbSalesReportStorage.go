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

// UcbSalesReportStorage stores all of the tasty modelleaf, needs to be injected into
// SalesReport and SalesReport Resource. In the real world, you would use a database for that.
type UcbSalesReportStorage struct {
	SalesConfigs  map[string]*UcbModel.SalesReport
	idCount int

	db *BmMongodb.BmMongodb
}

func (s UcbSalesReportStorage) NewSalesReportStorage(args []BmDaemons.BmDaemon) *UcbSalesReportStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbSalesReportStorage{make(map[string]*UcbModel.SalesReport), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbSalesReportStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.SalesReport {
	in := UcbModel.SalesReport{}
	var out []UcbModel.SalesReport
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.SalesReport
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil //make(map[string]*BmModel.Student)
	}
}

// GetOne tasty modelleaf
func (s UcbSalesReportStorage) GetOne(id string) (UcbModel.SalesReport, error) {
	in := UcbModel.SalesReport{ID: id}
	out := UcbModel.SalesReport{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("SalesReport for id %s not found", id)
	return UcbModel.SalesReport{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbSalesReportStorage) Insert(c UcbModel.SalesReport) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbSalesReportStorage) Delete(id string) error {
	in := UcbModel.SalesReport{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("SalesReport with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbSalesReportStorage) Update(c UcbModel.SalesReport) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("SalesReport with id does not exist")
	}

	return nil
}

func (s *UcbSalesReportStorage) Count(req api2go.Request, c UcbModel.SalesReport) int {
	r, _ := s.db.Count(req, &c)
	return r
}
