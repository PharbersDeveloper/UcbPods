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

type UcbProposalResource struct {
	UcbProposalStorage        *UcbDataStorage.UcbProposalStorage
	UcbUseableProposalStorage *UcbDataStorage.UcbUseableProposalStorage
}

func (c UcbProposalResource) NewProposalResource(args []BmDataStorage.BmStorage) *UcbProposalResource {
	var ps *UcbDataStorage.UcbProposalStorage
	var ups *UcbDataStorage.UcbUseableProposalStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbProposalStorage" {
			ps = arg.(*UcbDataStorage.UcbProposalStorage)
		} else if tp.Name() == "UcbUseableProposalStorage" {
			ups = arg.(*UcbDataStorage.UcbUseableProposalStorage)
		}
	}
	return &UcbProposalResource{
		UcbProposalStorage:        ps,
		UcbUseableProposalStorage: ups,
	}
}

// FindAll Proposals
func (c UcbProposalResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	useableProposalsID, upiok := r.QueryParams["useableProposalsID"]

	if upiok {
		modelRootID := useableProposalsID[0]
		modelRoot, err := c.UcbUseableProposalStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := c.UcbProposalStorage.GetOne(modelRoot.ProposalID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	result := c.UcbProposalStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbProposalResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbProposalStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbProposalResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Proposal)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbProposalStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbProposalResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbProposalStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbProposalResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Proposal)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbProposalStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
