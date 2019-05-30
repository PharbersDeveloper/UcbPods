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

type UcbAssessmentReportDescribeResource struct {
	UcbAssessmentReportDescribeStorage          *UcbDataStorage.UcbAssessmentReportDescribeStorage
	UcbRegionalDivisionResultStorage	*UcbDataStorage.UcbRegionalDivisionResultStorage
	UcbTargetAssignsResultStorage		*UcbDataStorage.UcbTargetAssignsResultStorage
	UcbResourceAssignsResultStorage		*UcbDataStorage.UcbResourceAssignsResultStorage
	UcbManageTimeResultStorage			*UcbDataStorage.UcbManageTimeResultStorage
	UcbManageTeamResultStorage			*UcbDataStorage.UcbManageTeamResultStorage

}

func (c UcbAssessmentReportDescribeResource) NewAssessmentReportDescribeResource(args []BmDataStorage.BmStorage) *UcbAssessmentReportDescribeResource {
	var ard *UcbDataStorage.UcbAssessmentReportDescribeStorage
	var rdr	*UcbDataStorage.UcbRegionalDivisionResultStorage
	var tar	*UcbDataStorage.UcbTargetAssignsResultStorage
	var rar *UcbDataStorage.UcbResourceAssignsResultStorage
	var mtr *UcbDataStorage.UcbManageTimeResultStorage
	var mtrs *UcbDataStorage.UcbManageTeamResultStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbAssessmentReportDescribeStorage" {
			ard = arg.(*UcbDataStorage.UcbAssessmentReportDescribeStorage)
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
		}
	}
	return &UcbAssessmentReportDescribeResource{
		UcbAssessmentReportDescribeStorage: ard,
		UcbRegionalDivisionResultStorage: rdr,
		UcbTargetAssignsResultStorage: tar,
		UcbResourceAssignsResultStorage: rar,
		UcbManageTimeResultStorage: mtr,
		UcbManageTeamResultStorage: mtrs,
	}
}

// FindAll images
func (c UcbAssessmentReportDescribeResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []UcbModel.AssessmentReportDescribe

	regionalDivisionResultsID, rdrOk := r.QueryParams["regionalDivisionResultsID"]
	targetAssignsResultsID, tarOk := r.QueryParams["targetAssignsResultsID"]
	resourceAssignsResultsID, rarOk := r.QueryParams["resourceAssignsResultsID"]
	manageTimeResultsID, mtrOk := r.QueryParams["manageTimeResultsID"]
	manageTeamResultsID, mtrsOk := r.QueryParams["manageTeamResultsID"]

	if rdrOk {
		modelRootID := regionalDivisionResultsID[0]
		modelRoot, err := c.UcbRegionalDivisionResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.AssessmentReportDescribeIDs
		model := c.UcbAssessmentReportDescribeStorage.GetAll(r, -1,-1)

		return &Response{Res: model}, nil
	}

	if tarOk {
		modelRootID := targetAssignsResultsID[0]
		modelRoot, err := c.UcbTargetAssignsResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.AssessmentReportDescribeIDs
		model := c.UcbAssessmentReportDescribeStorage.GetAll(r, -1,-1)

		return &Response{Res: model}, nil
	}

	if rarOk {
		modelRootID := resourceAssignsResultsID[0]
		modelRoot, err := c.UcbResourceAssignsResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.AssessmentReportDescribeIDs
		model := c.UcbAssessmentReportDescribeStorage.GetAll(r, -1,-1)

		return &Response{Res: model}, nil
	}

	if mtrOk {
		modelRootID := manageTimeResultsID[0]
		modelRoot, err := c.UcbManageTimeResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.AssessmentReportDescribeIDs
		model := c.UcbAssessmentReportDescribeStorage.GetAll(r, -1,-1)

		return &Response{Res: model}, nil
	}

	if mtrsOk {
		modelRootID := manageTeamResultsID[0]
		modelRoot, err := c.UcbManageTeamResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.AssessmentReportDescribeIDs
		model := c.UcbAssessmentReportDescribeStorage.GetAll(r, -1,-1)

		return &Response{Res: model}, nil
	}


	result = c.UcbAssessmentReportDescribeStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbAssessmentReportDescribeResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbAssessmentReportDescribeStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbAssessmentReportDescribeResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.AssessmentReportDescribe)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbAssessmentReportDescribeStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbAssessmentReportDescribeResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbAssessmentReportDescribeStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbAssessmentReportDescribeResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.AssessmentReportDescribe)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbAssessmentReportDescribeStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
