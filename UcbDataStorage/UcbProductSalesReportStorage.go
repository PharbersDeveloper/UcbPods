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

// UcbProductSalesReportStorage stores all of the tasty modelleaf, needs to be injected into
// ProductSalesReport and ProductSalesReport Resource. In the real world, you would use a database for that.
type UcbProductSalesReportStorage struct {
	SalesConfigs  map[string]*UcbModel.ProductSalesReport
	idCount int

	db *BmMongodb.BmMongodb
}

func (s UcbProductSalesReportStorage) NewProductSalesReportStorage(args []BmDaemons.BmDaemon) *UcbProductSalesReportStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbProductSalesReportStorage{make(map[string]*UcbModel.ProductSalesReport), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbProductSalesReportStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.ProductSalesReport {
	in := UcbModel.ProductSalesReport{}
	var out []*UcbModel.ProductSalesReport
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
func (s UcbProductSalesReportStorage) GetOne(id string) (UcbModel.ProductSalesReport, error) {
	in := UcbModel.ProductSalesReport{ID: id}
	out := UcbModel.ProductSalesReport{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("ProductSalesReport for id %s not found", id)
	return UcbModel.ProductSalesReport{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbProductSalesReportStorage) Insert(c UcbModel.ProductSalesReport) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbProductSalesReportStorage) Delete(id string) error {
	in := UcbModel.ProductSalesReport{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("ProductSalesReport with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbProductSalesReportStorage) Update(c UcbModel.ProductSalesReport) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("ProductSalesReport with id does not exist")
	}

	return nil
}
