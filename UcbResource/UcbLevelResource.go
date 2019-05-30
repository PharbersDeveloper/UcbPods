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

type UcbLevelResource struct {
	UcbLevelStorage          	*UcbDataStorage.UcbLevelStorage
	UcbLevelConfigStorage 		*UcbDataStorage.UcbLevelConfigStorage
}

func (c UcbLevelResource) NewLevelResource(args []BmDataStorage.BmStorage) *UcbLevelResource {
	var ls *UcbDataStorage.UcbLevelStorage
	var lcs *UcbDataStorage.UcbLevelConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbLevelStorage" {
			ls = arg.(*UcbDataStorage.UcbLevelStorage)
		} else if tp.Name() == "UcbLevelConfigStorage" {
			lcs = arg.(*UcbDataStorage.UcbLevelConfigStorage)
		}
	}
	return &UcbLevelResource{
		UcbLevelStorage: ls,
		UcbLevelConfigStorage: lcs,
	}
}

// FindAll images
func (c UcbLevelResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	levelConfigsID, lcsOk := r.QueryParams["levelConfigsID"]

	if lcsOk {
		modelRootID := levelConfigsID[0]
		modelRoot, err := c.UcbLevelConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbLevelStorage.GetOne(modelRoot.LevelID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []UcbModel.Level
	result = c.UcbLevelStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbLevelResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbLevelStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbLevelResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Level)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbLevelStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbLevelResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbLevelStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbLevelResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Level)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbLevelStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
