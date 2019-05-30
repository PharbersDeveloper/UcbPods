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

type UcbRepresentativeinputResource struct {
	UcbRepresentativeinputStorage *UcbDataStorage.UcbRepresentativeinputStorage
	UcbPaperinputStorage          *UcbDataStorage.UcbPaperinputStorage
}

func (s UcbRepresentativeinputResource) NewRepresentativeinputResource(args []BmDataStorage.BmStorage) *UcbRepresentativeinputResource {
	var bis *UcbDataStorage.UcbRepresentativeinputStorage
	var pis *UcbDataStorage.UcbPaperinputStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbRepresentativeinputStorage" {
			bis = arg.(*UcbDataStorage.UcbRepresentativeinputStorage)
		} else if tp.Name() == "UcbPaperinputStorage" {
			pis = arg.(*UcbDataStorage.UcbPaperinputStorage)
		}
	}
	return &UcbRepresentativeinputResource{
		UcbRepresentativeinputStorage: bis,
		UcbPaperinputStorage:          pis,
	}
}

func (s UcbRepresentativeinputResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	PaperinputsID, piok := r.QueryParams["paperinputsID"]
	var result []*UcbModel.Representativeinput

	if piok {
		modelRootID := PaperinputsID[0]

		modelRoot, err := s.UcbPaperinputStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		r.QueryParams["ids"] = modelRoot.RepresentativeinputIDs

		result = s.UcbRepresentativeinputStorage.GetAll(r, -1, -1)

		return &Response{Res: result}, nil
	}

	result = s.UcbRepresentativeinputStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbRepresentativeinputResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.Representativeinput
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
		for _, iter := range s.UcbRepresentativeinputStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbRepresentativeinputStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.Representativeinput{}
	count := s.UcbRepresentativeinputStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbRepresentativeinputResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbRepresentativeinputStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbRepresentativeinputResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Representativeinput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbRepresentativeinputStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbRepresentativeinputResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbRepresentativeinputStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbRepresentativeinputResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Representativeinput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbRepresentativeinputStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
