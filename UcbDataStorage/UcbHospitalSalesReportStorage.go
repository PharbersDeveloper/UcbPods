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

// UcbHospitalSalesReportStorage stores all of the tasty modelleaf, needs to be injected into
// HospitalSalesReport and HospitalSalesReport Resource. In the real world, you would use a database for that.
type UcbHospitalSalesReportStorage struct {
	SalesConfigs  map[string]*UcbModel.HospitalSalesReport
	idCount int

	db *BmMongodb.BmMongodb
}

func (s UcbHospitalSalesReportStorage) NewHospitalSalesReportStorage(args []BmDaemons.BmDaemon) *UcbHospitalSalesReportStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbHospitalSalesReportStorage{make(map[string]*UcbModel.HospitalSalesReport), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbHospitalSalesReportStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.HospitalSalesReport {
	in := UcbModel.HospitalSalesReport{}
	var out []UcbModel.HospitalSalesReport
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.HospitalSalesReport
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
func (s UcbHospitalSalesReportStorage) GetOne(id string) (UcbModel.HospitalSalesReport, error) {
	in := UcbModel.HospitalSalesReport{ID: id}
	out := UcbModel.HospitalSalesReport{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("HospitalSalesReport for id %s not found", id)
	return UcbModel.HospitalSalesReport{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbHospitalSalesReportStorage) Insert(c UcbModel.HospitalSalesReport) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbHospitalSalesReportStorage) Delete(id string) error {
	in := UcbModel.HospitalSalesReport{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("HospitalSalesReport with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbHospitalSalesReportStorage) Update(c UcbModel.HospitalSalesReport) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("HospitalSalesReport with id does not exist")
	}

	return nil
}
