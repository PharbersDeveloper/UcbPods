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

type UcbSalesReportResource struct {
	UcbSalesReportStorage               *UcbDataStorage.UcbSalesReportStorage
	UcbHospitalSalesReportStorage       *UcbDataStorage.UcbHospitalSalesReportStorage
	UcbRepresentativeSalesReportStorage *UcbDataStorage.UcbRepresentativeSalesReportStorage
	UcbProductSalesReportStorage        *UcbDataStorage.UcbProductSalesReportStorage
	UcbPaperStorage						*UcbDataStorage.UcbPaperStorage
	UcbSalesConfigStorage				*UcbDataStorage.UcbSalesConfigStorage
	UcbScenarioStorage					*UcbDataStorage.UcbScenarioStorage
	UcbProposalStorage					*UcbDataStorage.UcbProposalStorage
	UcbCitySalesReportStorage			*UcbDataStorage.UcbCitySalesReportStorage
}

func (c UcbSalesReportResource) NewSalesReportResource(args []BmDataStorage.BmStorage) *UcbSalesReportResource {
	var sr  *UcbDataStorage.UcbSalesReportStorage
	var hsr *UcbDataStorage.UcbHospitalSalesReportStorage
	var rsp *UcbDataStorage.UcbRepresentativeSalesReportStorage
	var psr *UcbDataStorage.UcbProductSalesReportStorage
	var ps	*UcbDataStorage.UcbPaperStorage
	var sc	*UcbDataStorage.UcbSalesConfigStorage
	var ss	*UcbDataStorage.UcbScenarioStorage
	var pss *UcbDataStorage.UcbProposalStorage
	var csrs *UcbDataStorage.UcbCitySalesReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbSalesReportStorage" {
			sr = arg.(*UcbDataStorage.UcbSalesReportStorage)
		} else if tp.Name() == "UcbHospitalSalesReportStorage" {
			hsr = arg.(*UcbDataStorage.UcbHospitalSalesReportStorage)
		} else if tp.Name() == "UcbRepresentativeSalesReportStorage" {
			rsp = arg.(*UcbDataStorage.UcbRepresentativeSalesReportStorage)
		} else if tp.Name() == "UcbProductSalesReportStorage" {
			psr = arg.(*UcbDataStorage.UcbProductSalesReportStorage)
		} else if tp.Name() == "UcbPaperStorage" {
			ps = arg.(*UcbDataStorage.UcbPaperStorage)
		} else if tp.Name() == "UcbSalesConfigStorage" {
			sc = arg.(*UcbDataStorage.UcbSalesConfigStorage)
		} else if tp.Name() == "UcbScenarioStorage" {
			ss =arg.(*UcbDataStorage.UcbScenarioStorage)
		} else if tp.Name() == "UcbProposalStorage" {
			pss =arg.(*UcbDataStorage.UcbProposalStorage)
		} else if tp.Name() == "UcbCitySalesReportStorage" {
			csrs =arg.(*UcbDataStorage.UcbCitySalesReportStorage)
		}
	}
	return &UcbSalesReportResource{
		UcbSalesReportStorage : sr,
		UcbHospitalSalesReportStorage: hsr,
		UcbRepresentativeSalesReportStorage: rsp,
		UcbProductSalesReportStorage: psr,
		UcbPaperStorage: ps,
		UcbSalesConfigStorage: sc,
		UcbScenarioStorage: ss,
		UcbProposalStorage: pss,
		UcbCitySalesReportStorage: csrs,
	}
}

// FindAll SalesConfigs
func (c UcbSalesReportResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	papersID, dcok := r.QueryParams["papersID"]
	proposalsID, pok := r.QueryParams["proposalsID"]

	var result []UcbModel.SalesReport


	if dcok {
		modelRootID := papersID[0]
		modelRoot, err := c.UcbPaperStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.SalesReportIDs

		result := c.UcbSalesReportStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	if pok {
		modelRootID := proposalsID[0]
		modelRoot, err := c.UcbProposalStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.SalesReportIDs

		result := c.UcbSalesReportStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	models := c.UcbSalesReportStorage.GetAll(r, -1, -1)

	for _, model := range models {
		result = append(result, *model)
	}
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbSalesReportResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	modelRoot, err := c.UcbSalesReportStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	modelRoot.HospitalSalesReport = []*UcbModel.HospitalSalesReport{}
	r.QueryParams["ids"] = modelRoot.HospitalSalesReportIDs
	for _, hospitalSalesReport := range c.UcbHospitalSalesReportStorage.GetAll(r, -1,-1) {
		modelRoot.HospitalSalesReport = append(modelRoot.HospitalSalesReport, hospitalSalesReport)
	}

	modelRoot.RepresentativeSalesReport = []*UcbModel.RepresentativeSalesReport{}
	r.QueryParams["ids"] = modelRoot.RepresentativeSalesReportIDs
	for _, representativeSalesReport := range c.UcbRepresentativeSalesReportStorage.GetAll(r, -1,-1) {
		modelRoot.RepresentativeSalesReport = append(modelRoot.RepresentativeSalesReport, representativeSalesReport)
	}

	modelRoot.ProductSalesReport = []*UcbModel.ProductSalesReport{}
	r.QueryParams["ids"] = modelRoot.ProductSalesReportIDs
	for _, productSalesReport := range c.UcbProductSalesReportStorage.GetAll(r, -1,-1) {
		modelRoot.ProductSalesReport = append(modelRoot.ProductSalesReport, productSalesReport)
	}

	modelRoot.CitySalesReport = []*UcbModel.CitySalesReport{}
	r.QueryParams["ids"] = modelRoot.CitySalesReportIDs
	for _, citySalesReport := range c.UcbCitySalesReportStorage.GetAll(r, -1,-1) {
		modelRoot.CitySalesReport = append(modelRoot.CitySalesReport, citySalesReport)
	}

	if modelRoot.ScenarioID != "" {
		ScenarioResult, _ := c.UcbScenarioStorage.GetOne(modelRoot.ScenarioID)
		modelRoot.Scenario = &ScenarioResult
	}
	return &Response{Res: modelRoot}, err
}

// Create a new choc
func (c UcbSalesReportResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.SalesReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbSalesReportStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbSalesReportResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbSalesReportStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbSalesReportResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.SalesReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbSalesReportStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
