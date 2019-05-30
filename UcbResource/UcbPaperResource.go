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

type UcbPaperResource struct {
	UcbPaperStorage			*UcbDataStorage.UcbPaperStorage
	UcbPaperinputStorage	*UcbDataStorage.UcbPaperinputStorage
	UcbSalesReportStorage	*UcbDataStorage.UcbSalesReportStorage
	UcbPersonnelAssessmentStorage	*UcbDataStorage.UcbPersonnelAssessmentStorage
}

func (s UcbPaperResource) NewPaperResource (args []BmDataStorage.BmStorage) *UcbPaperResource {
	var ps *UcbDataStorage.UcbPaperStorage
	var pis *UcbDataStorage.UcbPaperinputStorage
	var srs *UcbDataStorage.UcbSalesReportStorage
	var pas *UcbDataStorage.UcbPersonnelAssessmentStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbPaperStorage" {
			ps = arg.(*UcbDataStorage.UcbPaperStorage)
		} else if tp.Name() == "UcbPaperinputStorage" {
			pis = arg.(*UcbDataStorage.UcbPaperinputStorage)
		} else if tp.Name() == "UcbSalesReportStorage" {
			srs = arg.(*UcbDataStorage.UcbSalesReportStorage)
		} else if tp.Name() == "UcbPersonnelAssessmentStorage" {
			pas = arg.(*UcbDataStorage.UcbPersonnelAssessmentStorage)
		}
	}
	return &UcbPaperResource{
		UcbPaperinputStorage: pis,
		UcbPaperStorage: ps,
		UcbSalesReportStorage: srs,
		UcbPersonnelAssessmentStorage: pas,
	}
}

func (s UcbPaperResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	//var result []UcbModel.Paper
	result := s.UcbPaperStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbPaperResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.Paper
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
		for _, iter := range s.UcbPaperStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbPaperStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.Paper{}
	count := s.UcbPaperStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbPaperResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbPaperStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	model.Paperinputs = []*UcbModel.Paperinput{}
	r.QueryParams["ids"] = model.InputIDs
	paperInputModels := s.UcbPaperinputStorage.GetAll(r, -1,-1)
	model.Paperinputs = paperInputModels

	model.SalesReports = []*UcbModel.SalesReport{}
	r.QueryParams["ids"] = model.SalesReportIDs
	salesReportModels := s.UcbSalesReportStorage.GetAll(r, -1,-1)
	model.SalesReports = salesReportModels

	model.PersonnelAssessments = []*UcbModel.PersonnelAssessment{}
	r.QueryParams["ids"] = model.PersonnelAssessmentIDs
	personnelAssessmentModels := s.UcbPersonnelAssessmentStorage.GetAll(r, -1,-1)
	model.PersonnelAssessments = personnelAssessmentModels

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbPaperResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Paper)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbPaperStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbPaperResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbPaperStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbPaperResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Paper)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbPaperStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}