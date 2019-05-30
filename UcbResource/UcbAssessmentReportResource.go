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

type UcbAssessmentReportResource struct {
	UcbAssessmentReportStorage          *UcbDataStorage.UcbAssessmentReportStorage
	UcbRegionalDivisionResultStorage	*UcbDataStorage.UcbRegionalDivisionResultStorage
	UcbTargetAssignsResultStorage		*UcbDataStorage.UcbTargetAssignsResultStorage
	UcbResourceAssignsResultStorage		*UcbDataStorage.UcbResourceAssignsResultStorage
	UcbManageTimeResultStorage			*UcbDataStorage.UcbManageTimeResultStorage
	UcbManageTeamResultStorage			*UcbDataStorage.UcbManageTeamResultStorage
	UcbGeneralPerformanceResultStorage	*UcbDataStorage.UcbGeneralPerformanceResultStorage
	UcbPaperStorage 					*UcbDataStorage.UcbPaperStorage

}

func (c UcbAssessmentReportResource) NewAssessmentReportResource(args []BmDataStorage.BmStorage) *UcbAssessmentReportResource {
	var ard *UcbDataStorage.UcbAssessmentReportStorage
	var rdr	*UcbDataStorage.UcbRegionalDivisionResultStorage
	var tar	*UcbDataStorage.UcbTargetAssignsResultStorage
	var rar *UcbDataStorage.UcbResourceAssignsResultStorage
	var mtr *UcbDataStorage.UcbManageTimeResultStorage
	var mtrs *UcbDataStorage.UcbManageTeamResultStorage
	var gpr *UcbDataStorage.UcbGeneralPerformanceResultStorage
	var p *UcbDataStorage.UcbPaperStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbAssessmentReportStorage" {
			ard = arg.(*UcbDataStorage.UcbAssessmentReportStorage)
		} else if tp.Name() == "UcbRegionalDivisionResultStorage" {
			rdr = arg.(*UcbDataStorage.UcbRegionalDivisionResultStorage)
		} else if tp.Name() == "UcbTargetAssignsResultStorage" {
			tar = arg.(*UcbDataStorage.UcbTargetAssignsResultStorage)
		} else if tp.Name() == "UcbResourceAssignsResultStorage" {
			rar = arg.(*UcbDataStorage.UcbResourceAssignsResultStorage)
		} else if tp.Name() == "UcbManageTimeResultStorage" {
			mtr = arg.(*UcbDataStorage.UcbManageTimeResultStorage)
		} else if tp.Name() == "UcbManageTeamResultStorage" {
			mtrs = arg.(*UcbDataStorage.UcbManageTeamResultStorage)
		} else if tp.Name() == "UcbPaperStorage" {
			p = arg.(*UcbDataStorage.UcbPaperStorage)
		} else if tp.Name() == "UcbGeneralPerformanceResultStorage" {
			gpr = arg.(*UcbDataStorage.UcbGeneralPerformanceResultStorage)
		}
	}
	return &UcbAssessmentReportResource{
		UcbAssessmentReportStorage: ard,
		UcbRegionalDivisionResultStorage: rdr,
		UcbTargetAssignsResultStorage: tar,
		UcbResourceAssignsResultStorage: rar,
		UcbManageTimeResultStorage: mtr,
		UcbManageTeamResultStorage: mtrs,
		UcbPaperStorage: p,
		UcbGeneralPerformanceResultStorage: gpr,
	}
}

// FindAll images
func (c UcbAssessmentReportResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []UcbModel.AssessmentReport

	papersID, pOk := r.QueryParams["papersID"]

	if pOk {
		modelRootID := papersID[0]
		modelRoot, err := c.UcbPaperStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.AssessmentReportIDs

		models := c.UcbAssessmentReportStorage.GetAll(r, -1,-1)

		return &Response{Res: models}, nil
	}

	result = c.UcbAssessmentReportStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbAssessmentReportResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbAssessmentReportStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbAssessmentReportResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.AssessmentReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbAssessmentReportStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbAssessmentReportResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbAssessmentReportStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbAssessmentReportResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.AssessmentReport)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbAssessmentReportStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
