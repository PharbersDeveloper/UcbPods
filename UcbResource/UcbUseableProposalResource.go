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

type UcbUseableProposalResource struct {
	UcbUseableProposalStorage *UcbDataStorage.UcbUseableProposalStorage
	UcbProposalStorage        *UcbDataStorage.UcbProposalStorage
}

func (s UcbUseableProposalResource) NewUseableProposalResource(args []BmDataStorage.BmStorage) *UcbUseableProposalResource {
	var rcs *UcbDataStorage.UcbUseableProposalStorage
	var rs *UcbDataStorage.UcbProposalStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbUseableProposalStorage" {
			rcs = arg.(*UcbDataStorage.UcbUseableProposalStorage)
		} else if tp.Name() == "UcbProposalStorage" {
			rs = arg.(*UcbDataStorage.UcbProposalStorage)
		}
	}
	return &UcbUseableProposalResource{
		UcbUseableProposalStorage: rcs,
		UcbProposalStorage:        rs,
	}
}

func (s UcbUseableProposalResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []UcbModel.UseableProposal
	models := s.UcbUseableProposalStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}
	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbUseableProposalResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.UseableProposal
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
		for _, iter := range s.UcbUseableProposalStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbUseableProposalStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.UseableProposal{}
	count := s.UcbUseableProposalStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbUseableProposalResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbUseableProposalStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	if model.ProposalID != "" {
		response, err := s.UcbProposalStorage.GetOne(model.ProposalID)
		if err != nil {
			return &Response{}, err
		}
		model.Proposal = &response
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbUseableProposalResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.UseableProposal)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbUseableProposalStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbUseableProposalResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbUseableProposalStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbUseableProposalResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.UseableProposal)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbUseableProposalStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
