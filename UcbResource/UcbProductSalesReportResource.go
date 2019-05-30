package UcbResource

import (
	"errors"
	"Ucb/UcbDataStorage"
	"Ucb/UcbModel"
	"reflect"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
)

type UcbProductSalesReportResource struct {
	UcbProductSalesReportStorage       *UcbDataStorage.UcbProductSalesReportStorage
	UcbSalesReportStorage               *UcbDataStorage.UcbSalesReportStorage
}

func (c UcbProductSalesReportResource) NewProductSalesReportResource(args []BmDataStorage.BmStorage) *UcbProductSalesReportResource {
	var psr  *UcbDataStorage.UcbProductSalesReportStorage
	var sr *UcbDataStorage.UcbSalesReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbProductSalesReportStorage" {
			psr = arg.(*UcbDataStorage.UcbProductSalesReportStorage)
		} else if tp.Name() == "UcbSalesReportStorage" {
			sr = arg.(*UcbDataStorage.UcbSalesReportStorage)
		}
	}
	return &UcbProductSalesReportResource{
		UcbProductSalesReportStorage: psr,
		UcbSalesReportStorage: sr,
	}
}

// FindAll SalesConfigs
func (c UcbProductSalesReportResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	salesReportsID, dcok := r.QueryParams["salesReportsID"]

	if dcok {
		modelRootID := salesReportsID[0]
		modelRoot, err := c.UcbSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		r.QueryParams["ids"] = modelRoot.ProductSalesReportIDs

		model := c.UcbProductSalesReportStorage.GetAll(r, -1,-1)


		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []*UcbModel.ProductSalesReport
	result = c.UcbProductSalesReportStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbProductSalesReportResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbProductSalesReportStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbProductSalesReportResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ProductSalesReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbProductSalesReportStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbProductSalesReportResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbProductSalesReportStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbProductSalesReportResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ProductSalesReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbProductSalesReportStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
