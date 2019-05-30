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

type UcbProductConfigResource struct {
	UcbProductConfigStorage *UcbDataStorage.UcbProductConfigStorage
	UcbGoodsConfigStorage   *UcbDataStorage.UcbGoodsConfigStorage
	UcbProductStorage       *UcbDataStorage.UcbProductStorage
}

func (s UcbProductConfigResource) NewProductConfigResource(args []BmDataStorage.BmStorage) *UcbProductConfigResource {
	var pcs *UcbDataStorage.UcbProductConfigStorage
	var gcs *UcbDataStorage.UcbGoodsConfigStorage
	var pc *UcbDataStorage.UcbProductStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbProductConfigStorage" {
			pcs = arg.(*UcbDataStorage.UcbProductConfigStorage)
		} else if tp.Name() == "UcbGoodsConfigStorage" {
			gcs = arg.(interface{}).(*UcbDataStorage.UcbGoodsConfigStorage)
		} else if tp.Name() == "UcbProductStorage" {
			pc = arg.(interface{}).(*UcbDataStorage.UcbProductStorage)
		}
	}
	return &UcbProductConfigResource{
		UcbProductConfigStorage: pcs,
		UcbGoodsConfigStorage:   gcs,
		UcbProductStorage:       pc,
	}
}

func (s UcbProductConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	goodsConfigID, gcok := r.QueryParams["goodsConfigsID"]
	var result []UcbModel.ProductConfig

	if gcok {
		modelRootID := goodsConfigID[0]
		modelRoot, err := s.UcbGoodsConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := s.UcbProductConfigStorage.GetOne(modelRoot.GoodsID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	models := s.UcbProductConfigStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbProductConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.ProductConfig
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
		for _, iter := range s.UcbProductConfigStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbProductConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.ProductConfig{}
	count := s.UcbProductConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbProductConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	modelRoot, err := s.UcbProductConfigStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	if modelRoot.ProductID != "" {
		model, err := s.UcbProductStorage.GetOne(modelRoot.ProductID)
		if err != nil {
			return &Response{}, err
		}
		modelRoot.Product = &model
	}
	return &Response{Res: modelRoot}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbProductConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.ProductConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := s.UcbProductConfigStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbProductConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbProductConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbProductConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.ProductConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := s.UcbProductConfigStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
