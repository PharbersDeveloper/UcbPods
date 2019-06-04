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

type UcbCityResource struct {
	UcbCityStorage *UcbDataStorage.UcbCityStorage
	UcbRegionStorage *UcbDataStorage.UcbRegionStorage
	UcbCitySalesReportStorage *UcbDataStorage.UcbCitySalesReportStorage
}

func (c UcbCityResource) NewCityResource(args []BmDataStorage.BmStorage) *UcbCityResource {
	var cs *UcbDataStorage.UcbCityStorage
	var rs  *UcbDataStorage.UcbRegionStorage
	var csrs *UcbDataStorage.UcbCitySalesReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbCityStorage" {
			cs = arg.(*UcbDataStorage.UcbCityStorage)
		} else if tp.Name() == "UcbRegionStorage" {
			rs = arg.(*UcbDataStorage.UcbRegionStorage)
		} else if tp.Name() == "UcbCitySalesReportStorage" {
			csrs = arg.(*UcbDataStorage.UcbCitySalesReportStorage)
		}
	}
	return &UcbCityResource{
		UcbCityStorage: cs,
		UcbRegionStorage: rs,
		UcbCitySalesReportStorage: csrs,
	}
}

// FindAll Citys
func (c UcbCityResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	regionsID, rok := r.QueryParams["regionsID"]
	citySalesReportsID, cok := r.QueryParams["citySalesReportsID"]

	if rok {
		modelRootID := regionsID[0]

		modelRoot, err := c.UcbRegionStorage.GetOne(modelRootID)

		if err != nil {
			return  &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.CityIds

		result := c.UcbCityStorage.GetAll(r, -1, -1)

		return &Response{Res: result}, nil
	}

	if cok {
		modelRootID := citySalesReportsID[0]

		modelRoot, err := c.UcbCitySalesReportStorage.GetOne(modelRootID)

		if err != nil {
			return  &Response{}, nil
		}

		result, _ := c.UcbCityStorage.GetOne(modelRoot.CityId)

		return &Response{Res: result}, nil
	}


	result := c.UcbCityStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbCityResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbCityStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbCityResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.City)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbCityStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbCityResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbCityStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbCityResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.City)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbCityStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
