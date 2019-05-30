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

type UcbTitleResource struct {
	UcbTitleStorage	*UcbDataStorage.UcbTitleStorage
	UcbLevelConfigStorage 		*UcbDataStorage.UcbLevelConfigStorage
}

func (c UcbTitleResource) NewTitleResource(args []BmDataStorage.BmStorage) *UcbTitleResource {
	var rdr *UcbDataStorage.UcbTitleStorage
	var lcs *UcbDataStorage.UcbLevelConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbTitleStorage" {
			rdr = arg.(*UcbDataStorage.UcbTitleStorage)
		} else if tp.Name() == "UcbLevelConfigStorage" {
			lcs = arg.(*UcbDataStorage.UcbLevelConfigStorage)
		}
	}
	return &UcbTitleResource{
		UcbTitleStorage: rdr,
		UcbLevelConfigStorage: lcs,
	}
}

// FindAll images
func (c UcbTitleResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	levelConfigsID, lcsOk := r.QueryParams["levelConfigsID"]

	if lcsOk {
		modelRootID := levelConfigsID[0]
		modelRoot, err := c.UcbLevelConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbTitleStorage.GetOne(modelRoot.TitleID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []UcbModel.Title
	result = c.UcbTitleStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbTitleResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbTitleStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbTitleResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Title)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbTitleStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbTitleResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbTitleStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbTitleResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Title)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbTitleStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
