package UcbResource

import (
	"errors"
	"Ucb/UcbDataStorage"
	"Ucb/UcbModel"
	"net/http"
	"reflect"
	"strconv"

	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
)

type UcbScenarioResource struct {
	UcbScenarioStorage 		*UcbDataStorage.UcbScenarioStorage
	UcbProposalStorage		*UcbDataStorage.UcbProposalStorage
	UcbPaperStorage			*UcbDataStorage.UcbPaperStorage
	UcbPaperinputStorage	*UcbDataStorage.UcbPaperinputStorage
	UcbPersonnelAssessmentStorage *UcbDataStorage.UcbPersonnelAssessmentStorage
	UcbSalesReportStorage		*UcbDataStorage.UcbSalesReportStorage
	UcbScenarioResultStorage	*UcbDataStorage.UcbScenarioResultStorage
}

func (c UcbScenarioResource) NewScenarioResource(args []BmDataStorage.BmStorage) *UcbScenarioResource {
	var cs *UcbDataStorage.UcbScenarioStorage
	var ps *UcbDataStorage.UcbProposalStorage
	var pas *UcbDataStorage.UcbPaperStorage
	var pis *UcbDataStorage.UcbPaperinputStorage
	var pass *UcbDataStorage.UcbPersonnelAssessmentStorage
	var srs *UcbDataStorage.UcbSalesReportStorage
	var srss *UcbDataStorage.UcbScenarioResultStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbScenarioStorage" {
			cs = arg.(*UcbDataStorage.UcbScenarioStorage)
		} else if tp.Name() == "UcbProposalStorage" {
			ps = arg.(*UcbDataStorage.UcbProposalStorage)
		} else if tp.Name() == "UcbPaperStorage" {
			pas = arg.(*UcbDataStorage.UcbPaperStorage)
		} else if tp.Name() == "UcbPaperinputStorage" {
			pis = arg.(*UcbDataStorage.UcbPaperinputStorage)
		} else if tp.Name() == "UcbPersonnelAssessmentStorage" {
			pass = arg.(*UcbDataStorage.UcbPersonnelAssessmentStorage)
		} else if tp.Name() == "UcbSalesReportStorage" {
			srs = arg.(*UcbDataStorage.UcbSalesReportStorage)
		} else if tp.Name() == "UcbScenarioResultStorage" {
			srss = arg.(*UcbDataStorage.UcbScenarioResultStorage)
		}
	}
	return &UcbScenarioResource{
		UcbScenarioStorage: cs,
		UcbProposalStorage: ps,
		UcbPaperStorage: pas,
		UcbPaperinputStorage: pis,
		UcbPersonnelAssessmentStorage: pass,
		UcbSalesReportStorage: srs,
		UcbScenarioResultStorage: srss,
	}
}

// FindAll Scenarios
// TODO @Alex 这边后续必须重构，太难看了 自己留
func (c UcbScenarioResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []UcbModel.Scenario
	proposalsID, psok := r.QueryParams["proposal-id"]
	_, acok := r.QueryParams["account-id"]

	paperinputsID, piok := r.QueryParams["paperinputsID"]
	personnelAssessmentsID, paok := r.QueryParams["personnelAssessmentsID"]
	salesReportsID, srok := r.QueryParams["salesReportsID"]

	scenarioResultsID, sOk := r.QueryParams["scenarioResultsID"]


	if psok && acok {

		proposalModel, _ := c.UcbProposalStorage.GetOne(proposalsID[0])
		paperModel := c.UcbPaperStorage.GetAll(r, -1,-1)[0]
		r.QueryParams["ids"] = paperModel.InputIDs
		r.QueryParams["orderby"] = []string{"time"}
		paperInputModel := c.UcbPaperinputStorage.GetAll(r, -1,-1)
		var (
			lastPhase int

		)
		if paperInputModel != nil {
			lastPaperInputModel := paperInputModel[len(paperInputModel)-1:][0]
			lastPhase = lastPaperInputModel.Phase
		} else {
			lastPhase = 1
		}
		totalPhase := proposalModel.TotalPhase

		if paperModel.InputState == 1 ||paperModel.InputState == 4 {
			r.QueryParams["phase"] = []string{strconv.Itoa(lastPhase)}
			result = c.UcbScenarioStorage.GetAll(r, -1, -1)
		} else if paperModel.InputState == 2 && lastPhase != totalPhase {
			r.QueryParams["phase"] = []string{strconv.Itoa(lastPhase + 1)}
			result = c.UcbScenarioStorage.GetAll(r, -1, -1)
		} else {
			r.QueryParams["phase"] = []string{"1"}
			result = c.UcbScenarioStorage.GetAll(r, -1, -1)
		}
		return &Response{Res: result}, nil
	}

	if piok {
		modelRootID := paperinputsID[0]
		modelRoot, err := c.UcbPaperinputStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := c.UcbScenarioStorage.GetOne(modelRoot.ScenarioID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	if paok {
		modelRootID := personnelAssessmentsID[0]
		modelRoot, err := c.UcbPersonnelAssessmentStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := c.UcbScenarioStorage.GetOne(modelRoot.ScenarioID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	if srok {
		modelRootID := salesReportsID[0]
		modelRoot, err := c.UcbSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := c.UcbScenarioStorage.GetOne(modelRoot.ScenarioID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	if sOk {
		modelRootID := scenarioResultsID[0]
		modelRoot, err := c.UcbScenarioResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := c.UcbScenarioStorage.GetOne(modelRoot.ScenarioID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	result = c.UcbScenarioStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbScenarioResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbScenarioStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbScenarioResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Scenario)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbScenarioStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbScenarioResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbScenarioStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbScenarioResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Scenario)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbScenarioStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
