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

type UcbGoodsConfigResource struct {
	UcbGoodsConfigStorage   			*UcbDataStorage.UcbGoodsConfigStorage
	UcbProductConfigStorage 			*UcbDataStorage.UcbProductConfigStorage
	UcbProductSalesReportStorage   		*UcbDataStorage.UcbProductSalesReportStorage
	UcbHospitalSalesReportStorage		*UcbDataStorage.UcbHospitalSalesReportStorage
	UcbRepresentativeSalesReportStorage	*UcbDataStorage.UcbRepresentativeSalesReportStorage
	UcbSalesConfigStorage 				*UcbDataStorage.UcbSalesConfigStorage
	UcbGoodsinputStorage				*UcbDataStorage.UcbGoodsinputStorage
}

func (s UcbGoodsConfigResource) NewGoodsConfigResource(args []BmDataStorage.BmStorage) *UcbGoodsConfigResource {
	var gcs *UcbDataStorage.UcbGoodsConfigStorage
	var pcs *UcbDataStorage.UcbProductConfigStorage
	var psr *UcbDataStorage.UcbProductSalesReportStorage
	var hsr *UcbDataStorage.UcbHospitalSalesReportStorage
	var rsr *UcbDataStorage.UcbRepresentativeSalesReportStorage
	var sc *UcbDataStorage.UcbSalesConfigStorage
	var gis *UcbDataStorage.UcbGoodsinputStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbGoodsConfigStorage" {
			gcs = arg.(*UcbDataStorage.UcbGoodsConfigStorage)
		} else if tp.Name() == "UcbProductConfigStorage" {
			pcs = arg.(interface{}).(*UcbDataStorage.UcbProductConfigStorage)
		} else if tp.Name() == "UcbProductSalesReportStorage" {
			psr = arg.(*UcbDataStorage.UcbProductSalesReportStorage)
		} else if tp.Name() == "UcbHospitalSalesReportStorage" {
			hsr = arg.(*UcbDataStorage.UcbHospitalSalesReportStorage)
		} else if tp.Name() == "UcbRepresentativeSalesReportStorage" {
			rsr = arg.(*UcbDataStorage.UcbRepresentativeSalesReportStorage)
		} else if tp.Name() == "UcbSalesConfigStorage" {
			sc = arg.(*UcbDataStorage.UcbSalesConfigStorage)
		} else if tp.Name() == "UcbGoodsinputStorage" {
			gis = arg.(*UcbDataStorage.UcbGoodsinputStorage)
		}
	}
	return &UcbGoodsConfigResource{
		UcbGoodsConfigStorage:   gcs,
		UcbProductConfigStorage: pcs,
		UcbProductSalesReportStorage: psr,
		UcbHospitalSalesReportStorage : hsr,
		UcbRepresentativeSalesReportStorage: rsr,
		UcbSalesConfigStorage: sc,
		UcbGoodsinputStorage: gis,
	}
}

func (s UcbGoodsConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	productSalesReportsID, psrok := r.QueryParams["productSalesReportsID"]
	hospitalSalesReportsID, hsrok := r.QueryParams["hospitalSalesReportsID"]
	representativeSalesReportsID, rsrok := r.QueryParams["representativeSalesReportsID"]
	salesConfigsID, scok := r.QueryParams["salesConfigsID"]
	goodsinputsID, gok := r.QueryParams["goodsinputsID"]

	if psrok {
		modelRootID := productSalesReportsID[0]
		modelRoot, err := s.UcbProductSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		model, err:= s.UcbGoodsConfigStorage.GetOne(modelRoot.GoodsConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if hsrok {
		modelRootID := hospitalSalesReportsID[0]
		modelRoot, err := s.UcbHospitalSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		model, err:= s.UcbGoodsConfigStorage.GetOne(modelRoot.GoodsConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if rsrok {
		modelRootID := representativeSalesReportsID[0]
		modelRoot, err := s.UcbRepresentativeSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		model, err:= s.UcbGoodsConfigStorage.GetOne(modelRoot.GoodsConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if scok {
		modelRootID := salesConfigsID[0]
		modelRoot, err := s.UcbSalesConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		model, err:= s.UcbGoodsConfigStorage.GetOne(modelRoot.GoodsConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if gok {
		modelRootID := goodsinputsID[0]
		modelRoot, err := s.UcbGoodsinputStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		result, err := s.UcbGoodsConfigStorage.GetOne(modelRoot.GoodsConfigId)

		if err != nil {
			return &Response{}, nil
		}

		return &Response{Res: result}, nil
	}

	var result []*UcbModel.GoodsConfig
	models := s.UcbGoodsConfigStorage.GetAll(r, -1, -1)

	for _, model := range models {
		result = append(result, model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbGoodsConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.GoodsConfig
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
		for _, iter := range s.UcbGoodsConfigStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbGoodsConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.GoodsConfig{}
	count := s.UcbGoodsConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbGoodsConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbGoodsConfigStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	if model.GoodsType == 0 {
		resp, err := s.UcbProductConfigStorage.GetOne(model.GoodsID)
		if err != nil {
			return &Response{}, err
		}
		model.ProductConfig = &resp
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbGoodsConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.GoodsConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}
	id := s.UcbGoodsConfigStorage.Insert(model)
	model.ID = id
	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbGoodsConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbGoodsConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbGoodsConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.GoodsConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}
	err := s.UcbGoodsConfigStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
