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

type UcbGoodsinputResource struct {
	UcbGoodsinputStorage   			*UcbDataStorage.UcbGoodsinputStorage
	UcbBusinessinputStorage			*UcbDataStorage.UcbBusinessinputStorage
}

func (s UcbGoodsinputResource) NewGoodsinputResource(args []BmDataStorage.BmStorage) *UcbGoodsinputResource {
	var gcs *UcbDataStorage.UcbGoodsinputStorage
	var bis *UcbDataStorage.UcbBusinessinputStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbGoodsinputStorage" {
			gcs = arg.(*UcbDataStorage.UcbGoodsinputStorage)
		} else if tp.Name() == "UcbBusinessinputStorage" {
			bis = arg.(*UcbDataStorage.UcbBusinessinputStorage)
		}
	}
	return &UcbGoodsinputResource{
		UcbGoodsinputStorage:   gcs,
		UcbBusinessinputStorage: bis,
	}
}

func (s UcbGoodsinputResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	businessinputsID, bok := r.QueryParams["businessinputsID"]

	if bok {
		modelRootID := businessinputsID[0]
		modelRoot, err := s.UcbBusinessinputStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.GoodsInputIds

		result := s.UcbGoodsinputStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	var result []*UcbModel.Goodsinput
	models := s.UcbGoodsinputStorage.GetAll(r, -1, -1)

	for _, model := range models {
		result = append(result, model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbGoodsinputResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.Goodsinput
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
		for _, iter := range s.UcbGoodsinputStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbGoodsinputStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.Goodsinput{}
	count := s.UcbGoodsinputStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbGoodsinputResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbGoodsinputStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbGoodsinputResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Goodsinput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}
	id := s.UcbGoodsinputStorage.Insert(model)
	model.ID = id
	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbGoodsinputResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbGoodsinputStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbGoodsinputResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Goodsinput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}
	err := s.UcbGoodsinputStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
