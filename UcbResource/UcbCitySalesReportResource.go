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

type UcbCitySalesReportResource struct {
	UcbCitySalesReportStorage *UcbDataStorage.UcbCitySalesReportStorage
	UcbSalesReportStorage     *UcbDataStorage.UcbSalesReportStorage
}

func (c UcbCitySalesReportResource) NewCitySalesReportResource(args []BmDataStorage.BmStorage) *UcbCitySalesReportResource {
	var cs *UcbDataStorage.UcbCitySalesReportStorage
	var sr 	*UcbDataStorage.UcbSalesReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbCitySalesReportStorage" {
			cs = arg.(*UcbDataStorage.UcbCitySalesReportStorage)
		} else if tp.Name() == "UcbSalesReportStorage" {
			sr = arg.(*UcbDataStorage.UcbSalesReportStorage)
		}
	}
	return &UcbCitySalesReportResource{
		UcbCitySalesReportStorage: cs,
		UcbSalesReportStorage: sr,
	}
}

// FindAll CitySalesReports
func (c UcbCitySalesReportResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	salesReportsID, dcok := r.QueryParams["salesReportsID"]

	if dcok {
		modelRootID := salesReportsID[0]
		modelRoot, err := c.UcbSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		r.QueryParams["ids"] = modelRoot.HospitalSalesReportIDs

		model := c.UcbCitySalesReportStorage.GetAll(r, -1,-1)

		return &Response{Res: model}, nil
	}

	result := c.UcbCitySalesReportStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbCitySalesReportResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbCitySalesReportStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbCitySalesReportResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.CitySalesReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbCitySalesReportStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbCitySalesReportResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbCitySalesReportStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbCitySalesReportResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.CitySalesReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbCitySalesReportStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
