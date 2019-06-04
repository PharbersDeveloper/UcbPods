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

// UcbCitySalesReportStorage stores all of the tasty modelleaf, needs to be injected into
// CitySalesReport and CitySalesReport Resource. In the real world, you would use a database for that.
type UcbCitySalesReportStorage struct {
	CitySalesReports map[string]*UcbModel.CitySalesReport
	idCount     int

	db *BmMongodb.BmMongodb
}

func (s UcbCitySalesReportStorage) NewCitySalesReportStorage(args []BmDaemons.BmDaemon) *UcbCitySalesReportStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbCitySalesReportStorage{make(map[string]*UcbModel.CitySalesReport), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbCitySalesReportStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.CitySalesReport {
	in := UcbModel.CitySalesReport{}
	var out []*UcbModel.CitySalesReport
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
func (s UcbCitySalesReportStorage) GetOne(id string) (UcbModel.CitySalesReport, error) {
	in := UcbModel.CitySalesReport{ID: id}
	out := UcbModel.CitySalesReport{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("CitySalesReport for id %s not found", id)
	return UcbModel.CitySalesReport{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbCitySalesReportStorage) Insert(c UcbModel.CitySalesReport) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbCitySalesReportStorage) Delete(id string) error {
	in := UcbModel.CitySalesReport{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("CitySalesReport with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbCitySalesReportStorage) Update(c UcbModel.CitySalesReport) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("CitySalesReport with id does not exist")
	}

	return nil
}
