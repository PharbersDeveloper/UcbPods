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

type UcbHospitalSalesReportResource struct {
	UcbHospitalSalesReportStorage       *UcbDataStorage.UcbHospitalSalesReportStorage
	UcbSalesReportStorage               *UcbDataStorage.UcbSalesReportStorage
}

func (c UcbHospitalSalesReportResource) NewHospitalSalesReportResource(args []BmDataStorage.BmStorage) *UcbHospitalSalesReportResource {
	var hsr  *UcbDataStorage.UcbHospitalSalesReportStorage
	var sr *UcbDataStorage.UcbSalesReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbHospitalSalesReportStorage" {
			hsr = arg.(*UcbDataStorage.UcbHospitalSalesReportStorage)
		} else if tp.Name() == "UcbSalesReportStorage" {
			sr = arg.(*UcbDataStorage.UcbSalesReportStorage)
		}
	}
	return &UcbHospitalSalesReportResource{
		UcbHospitalSalesReportStorage: hsr,
		UcbSalesReportStorage: sr,
	}
}

// FindAll SalesConfigs
func (c UcbHospitalSalesReportResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	salesReportsID, dcok := r.QueryParams["salesReportsID"]

	if dcok {
		modelRootID := salesReportsID[0]
		modelRoot, err := c.UcbSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		r.QueryParams["ids"] = modelRoot.HospitalSalesReportIDs

		model := c.UcbHospitalSalesReportStorage.GetAll(r, -1,-1)


		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []UcbModel.HospitalSalesReport

	models := c.UcbHospitalSalesReportStorage.GetAll(r, -1, -1)

	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbHospitalSalesReportResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbHospitalSalesReportStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbHospitalSalesReportResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.HospitalSalesReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbHospitalSalesReportStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbHospitalSalesReportResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbHospitalSalesReportStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbHospitalSalesReportResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.HospitalSalesReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbHospitalSalesReportStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
