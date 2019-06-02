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

type UcbManagerConfigResource struct {
	UcbManagerConfigStorage  *UcbDataStorage.UcbManagerConfigStorage
	UcbResourceConfigStorage *UcbDataStorage.UcbResourceConfigStorage
}

func (c UcbManagerConfigResource) NewManagerConfigResource(args []BmDataStorage.BmStorage) *UcbManagerConfigResource {
	var mcs *UcbDataStorage.UcbManagerConfigStorage
	var rcs *UcbDataStorage.UcbResourceConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbManagerConfigStorage" {
			mcs = arg.(*UcbDataStorage.UcbManagerConfigStorage)
		} else if tp.Name() == "UcbResourceConfigStorage" {
			rcs = arg.(*UcbDataStorage.UcbResourceConfigStorage)
		}
	}
	return &UcbManagerConfigResource{
		UcbManagerConfigStorage:  mcs,
		UcbResourceConfigStorage: rcs,
	}
}

func (c UcbManagerConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	resourceConfigsID, rcok := r.QueryParams["resourceConfigsID"]

	if rcok {
		modelRootID := resourceConfigsID[0]
		modelRoot, err := c.UcbResourceConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := c.UcbManagerConfigStorage.GetOne(modelRoot.ResourceID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	result := c.UcbManagerConfigStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

func (c UcbManagerConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbManagerConfigStorage.GetOne(ID)
	return &Response{Res: res}, err
}

func (c UcbManagerConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ManagerConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbManagerConfigStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

func (c UcbManagerConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbManagerConfigStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

func (c UcbManagerConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ManagerConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbManagerConfigStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
