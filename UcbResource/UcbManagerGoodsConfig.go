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

type UcbManagerGoodsConfigResource struct {
	UcbManagerGoodsConfigStorage *UcbDataStorage.UcbManagerGoodsConfigStorage
	UcbManagerConfigStorage *UcbDataStorage.UcbManagerConfigStorage
}

func (c UcbManagerGoodsConfigResource) NewManagerGoodsConfigResource(args []BmDataStorage.BmStorage) *UcbManagerGoodsConfigResource {
	var cs *UcbDataStorage.UcbManagerGoodsConfigStorage
	var mc *UcbDataStorage.UcbManagerConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbManagerGoodsConfigStorage" {
			cs = arg.(*UcbDataStorage.UcbManagerGoodsConfigStorage)
		} else if tp.Name() == "UcbManagerConfigStorage" {
			mc = arg.(*UcbDataStorage.UcbManagerConfigStorage)
		}
	}
	return &UcbManagerGoodsConfigResource{
		UcbManagerGoodsConfigStorage: cs,
		UcbManagerConfigStorage: mc,
	}
}

// FindAll ManagerGoodsConfigs
func (c UcbManagerGoodsConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	managerConfigsID , mok := r.QueryParams["managerConfigsID"]

	if mok {
		modelRootID := managerConfigsID[0]
		modelRoot, err := c.UcbManagerConfigStorage.GetOne(modelRootID)

		if err != nil {
			return  &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.ManagerGoodsConfigIds

		result := c.UcbManagerGoodsConfigStorage.GetAll(r,  -1,-1)
		return &Response{Res: result}, nil

	}

	result := c.UcbManagerGoodsConfigStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbManagerGoodsConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbManagerGoodsConfigStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbManagerGoodsConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ManagerGoodsConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbManagerGoodsConfigStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbManagerGoodsConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbManagerGoodsConfigStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbManagerGoodsConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ManagerGoodsConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbManagerGoodsConfigStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
