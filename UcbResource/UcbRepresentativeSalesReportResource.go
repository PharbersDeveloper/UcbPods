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

type UcbRepresentativeSalesReportResource struct {
	UcbRepresentativeSalesReportStorage *UcbDataStorage.UcbRepresentativeSalesReportStorage
	UcbSalesReportStorage               *UcbDataStorage.UcbSalesReportStorage
}

func (c UcbRepresentativeSalesReportResource) NewRepresentativeSalesReportResource(args []BmDataStorage.BmStorage) *UcbRepresentativeSalesReportResource {
	var psr  *UcbDataStorage.UcbRepresentativeSalesReportStorage
	var sr *UcbDataStorage.UcbSalesReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbRepresentativeSalesReportStorage" {
			psr = arg.(*UcbDataStorage.UcbRepresentativeSalesReportStorage)
		} else if tp.Name() == "UcbSalesReportStorage" {
			sr = arg.(*UcbDataStorage.UcbSalesReportStorage)
		}
	}
	return &UcbRepresentativeSalesReportResource{
		UcbRepresentativeSalesReportStorage: psr,
		UcbSalesReportStorage: sr,
	}
}

// FindAll SalesConfigs
func (c UcbRepresentativeSalesReportResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	salesReportsID, dcok := r.QueryParams["salesReportsID"]

	if dcok {
		modelRootID := salesReportsID[0]
		modelRoot, err := c.UcbSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		r.QueryParams["ids"] = modelRoot.RepresentativeSalesReportIDs

		model := c.UcbRepresentativeSalesReportStorage.GetAll(r, -1,-1)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []*UcbModel.RepresentativeSalesReport
	result = c.UcbRepresentativeSalesReportStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbRepresentativeSalesReportResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbRepresentativeSalesReportStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbRepresentativeSalesReportResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.RepresentativeSalesReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbRepresentativeSalesReportStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbRepresentativeSalesReportResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbRepresentativeSalesReportStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbRepresentativeSalesReportResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.RepresentativeSalesReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbRepresentativeSalesReportStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
