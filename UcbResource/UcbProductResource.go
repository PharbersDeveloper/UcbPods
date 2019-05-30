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

type UcbProductResource struct {
	UcbProductStorage       *UcbDataStorage.UcbProductStorage
	UcbImageStorage         *UcbDataStorage.UcbImageStorage
	UcbProductConfigStorage *UcbDataStorage.UcbProductConfigStorage
}

func (s UcbProductResource) NewProductResource(args []BmDataStorage.BmStorage) *UcbProductResource {
	var is *UcbDataStorage.UcbImageStorage
	var ps *UcbDataStorage.UcbProductStorage
	var pcs *UcbDataStorage.UcbProductConfigStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbImageStorage" {
			is = arg.(*UcbDataStorage.UcbImageStorage)
		} else if tp.Name() == "UcbProductStorage" {
			ps = arg.(*UcbDataStorage.UcbProductStorage)
		} else if tp.Name() == "UcbProductConfigStorage" {
			pcs = arg.(*UcbDataStorage.UcbProductConfigStorage)
		}
	}
	return &UcbProductResource{
		UcbImageStorage:         is,
		UcbProductStorage:       ps,
		UcbProductConfigStorage: pcs,
	}
}

func (s UcbProductResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	productConfigID, pcok := r.QueryParams["productConfigsID"]
	var result []UcbModel.Product

	if pcok {
		modelRootID := productConfigID[0]
		modelRoot, err := s.UcbProductConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := s.UcbProductStorage.GetOne(modelRoot.ProductID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	models := s.UcbProductStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbProductResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.Product
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
		for _, iter := range s.UcbProductStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbProductStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.Product{}
	count := s.UcbProductStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbProductResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbProductStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	r.QueryParams["ids"] = model.ImagesIDs
	images := s.UcbImageStorage.GetAll(r, -1, -1)
	for _, image := range images {
		model.Imgs = append(model.Imgs, &image)
	}
	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbProductResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Product)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbProductStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbProductResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbProductStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbProductResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Product)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbProductStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
