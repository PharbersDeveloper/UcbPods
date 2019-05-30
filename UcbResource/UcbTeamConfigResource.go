package UcbResource

import (
	"errors"
	"Ucb/UcbDataStorage"
	"Ucb/UcbModel"
	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
)

type UcbTeamConfigResource struct {
	UcbTeamConfigStorage			*UcbDataStorage.UcbTeamConfigStorage
	UcbResourceConfigStorage		*UcbDataStorage.UcbResourceConfigStorage
}

func (s UcbTeamConfigResource) NewTeamConfigResource (args []BmDataStorage.BmStorage) *UcbTeamConfigResource {
	var ps *UcbDataStorage.UcbTeamConfigStorage
	var rc *UcbDataStorage.UcbResourceConfigStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbTeamConfigStorage" {
			ps = arg.(*UcbDataStorage.UcbTeamConfigStorage)
		} else if tp.Name() == "UcbResourceConfigStorage" {
			rc = arg.(*UcbDataStorage.UcbResourceConfigStorage)
		}
	}
	return &UcbTeamConfigResource{UcbResourceConfigStorage: rc, UcbTeamConfigStorage: ps}
}

func (s UcbTeamConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []UcbModel.TeamConfig
	models := s.UcbTeamConfigStorage.GetAll(r, -1, -1)

	for _, model := range models {
		result = append(result, *model)
	}
	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbTeamConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.TeamConfig
		number, size, offset, limit string
	)

	numberQuery, ok := r.QueryParams["page[number]"]
	if ok {
		number = numberQuery[0]
	}
	sizeQuery, ok := r.QueryParams["page[size]"]
	if ok {
		size = sizeQuery[0]
	}
	offsetQuery, ok := r.QueryParams["page[offset]"]
	if ok {
		offset = offsetQuery[0]
	}
	limitQuery, ok := r.QueryParams["page[limit]"]
	if ok {
		limit = limitQuery[0]
	}

	if size != "" {
		sizeI, err := strconv.ParseInt(size, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		numberI, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		start := sizeI * (numberI - 1)
		for _, iter := range s.UcbTeamConfigStorage.GetAll(r, int(start), int(sizeI)) {
			result = append(result, *iter)
		}

	} else {
		limitI, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		offsetI, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		for _, iter := range s.UcbTeamConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.TeamConfig{}
	count := s.UcbTeamConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbTeamConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbTeamConfigStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbTeamConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.TeamConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbTeamConfigStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbTeamConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbTeamConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbTeamConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.TeamConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbTeamConfigStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}