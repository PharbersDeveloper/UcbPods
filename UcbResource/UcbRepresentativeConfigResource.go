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

type UcbRepresentativeConfigResource struct {
	UcbRepresentativeConfigStorage *UcbDataStorage.UcbRepresentativeConfigStorage
	UcbResourceConfigStorage       *UcbDataStorage.UcbResourceConfigStorage
	UcbRepresentativeStorage       *UcbDataStorage.UcbRepresentativeStorage
}

func (s UcbRepresentativeConfigResource) NewRepresentativeConfigResource(args []BmDataStorage.BmStorage) *UcbRepresentativeConfigResource {
	var repcs *UcbDataStorage.UcbRepresentativeConfigStorage
	var rcs *UcbDataStorage.UcbResourceConfigStorage
	var reps *UcbDataStorage.UcbRepresentativeStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbRepresentativeConfigStorage" {
			repcs = arg.(*UcbDataStorage.UcbRepresentativeConfigStorage)
		} else if tp.Name() == "UcbResourceConfigStorage" {
			rcs = arg.(*UcbDataStorage.UcbResourceConfigStorage)
		} else if tp.Name() == "UcbRepresentativeStorage" {
			reps = arg.(*UcbDataStorage.UcbRepresentativeStorage)
		}
	}
	return &UcbRepresentativeConfigResource{
		UcbRepresentativeConfigStorage: repcs,
		UcbResourceConfigStorage:       rcs,
		UcbRepresentativeStorage:       reps,
	}
}

func (s UcbRepresentativeConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	resourceConfigsID, rcok := r.QueryParams["resourceConfigsID"]
	var result []UcbModel.RepresentativeConfig

	if rcok {
		modelRootID := resourceConfigsID[0]
		modelRoot, err := s.UcbResourceConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := s.UcbRepresentativeConfigStorage.GetOne(modelRoot.ResourceID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	models := s.UcbRepresentativeConfigStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbRepresentativeConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.RepresentativeConfig
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
		for _, iter := range s.UcbRepresentativeConfigStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbRepresentativeConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.RepresentativeConfig{}
	count := s.UcbRepresentativeConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbRepresentativeConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	modelRoot, err := s.UcbRepresentativeConfigStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	if modelRoot.RepresentativeID != "" {
		model, err := s.UcbRepresentativeStorage.GetOne(modelRoot.RepresentativeID)
		if err != nil {
			return &Response{}, err
		}
		modelRoot.Representative = &model
	}
	return &Response{Res: modelRoot}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbRepresentativeConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.RepresentativeConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbRepresentativeConfigStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbRepresentativeConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbRepresentativeConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbRepresentativeConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.RepresentativeConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbRepresentativeConfigStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
