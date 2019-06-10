package UcbFactory

import (
	"Ucb/UcbDataStorage"
	"Ucb/UcbHandler"
	"Ucb/UcbMiddleware"
	"Ucb/UcbResource"
	"Ucb/UcbModel"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
)

type UcbTable struct{}

var NTM_MODEL_FACTORY = map[string]interface{}{
	"UcbImage":                UcbModel.Image{},
	"UcbPolicy":               UcbModel.Policy{},
	"UcbHospital":             UcbModel.Hospital{},
	"UcbDepartment":           UcbModel.Department{},
	"UcbRegion":               UcbModel.Region{},
	"UcbProduct":              UcbModel.Product{},
	"UcbProductConfig":        UcbModel.ProductConfig{},
	"UcbRepresentative":       UcbModel.Representative{},
	"UcbManagerConfig":        UcbModel.ManagerConfig{},
	"UcbRepresentativeConfig": UcbModel.RepresentativeConfig{},
	"UcbRegionConfig":         UcbModel.RegionConfig{},
	"UcbHospitalConfig":       UcbModel.HospitalConfig{},
	"UcbResourceConfig":       UcbModel.ResourceConfig{},
	"UcbGoodsConfig":          UcbModel.GoodsConfig{},
	"UcbBusinessinput":        UcbModel.Businessinput{},
	"UcbRepresentativeinput":  UcbModel.Representativeinput{},
	"UcbManagerinput":         UcbModel.Managerinput{},
	"UcbPaperinput":           UcbModel.Paperinput{},
	"UcbDestConfig":           UcbModel.DestConfig{},
	"UcbScenario":             UcbModel.Scenario{},
	"UcbProposal":             UcbModel.Proposal{},
	"UcbUseableProposal":      UcbModel.UseableProposal{},
	"UcbPaper":                UcbModel.Paper{},
	"UcbSalesConfig":		   UcbModel.SalesConfig{},
	"UcbSalesReport":		   UcbModel.SalesReport{},
	"UcbHospitalSalesReport":  UcbModel.HospitalSalesReport{},
	"UcbProductSalesReport":   UcbModel.ProductSalesReport{},
	"UcbRepresentativeSalesReport":	UcbModel.RepresentativeSalesReport{},
	"UcbTeamConfig":			UcbModel.TeamConfig{},
	"UcbActionKpi":				UcbModel.ActionKpi{},
	"UcbPersonnelAssessment":	UcbModel.PersonnelAssessment{},
	"UcbRepresentativeAbility":	UcbModel.RepresentativeAbility{},

	"UcbLevel":						UcbModel.Level{},
	"UcbLevelConfig":				UcbModel.LevelConfig{},
	"UcbAssessmentReportDescribe":	UcbModel.AssessmentReportDescribe{},
	"UcbRegionalDivisionResult":	UcbModel.RegionalDivisionResult{},
	"UcbTargetAssignsResult":		UcbModel.TargetAssignsResult{},
	"UcbResourceAssignsResult":		UcbModel.ResourceAssignsResult{},
	"UcbManageTimeResult":			UcbModel.ManageTimeResult{},
	"UcbManageTeamResult":			UcbModel.ManageTeamResult{},
	"UcbAssessmentReport":			UcbModel.AssessmentReport{},
	"UcbTitle":						UcbModel.Title{},
	"UcbGeneralPerformanceResult":	UcbModel.GeneralPerformanceResult{},
	"UcbGoodsinput":				UcbModel.Goodsinput{},
	"UcbCity":						UcbModel.City{},
	"UcbManagerGoodsConfig":		UcbModel.ManagerGoodsConfig{},
	"UcbCitySalesReport":			UcbModel.CitySalesReport{},
}

var NTM_STORAGE_FACTORY = map[string]interface{}{
	"UcbImageStorage":                UcbDataStorage.UcbImageStorage{},
	"UcbPolicyStorage":               UcbDataStorage.UcbPolicyStorage{},
	"UcbHospitalStorage":             UcbDataStorage.UcbHospitalStorage{},
	"UcbDepartmentStorage":           UcbDataStorage.UcbDepartmentStorage{},
	"UcbRegionStorage":               UcbDataStorage.UcbRegionStorage{},
	"UcbProductStorage":              UcbDataStorage.UcbProductStorage{},
	"UcbProductConfigStorage":        UcbDataStorage.UcbProductConfigStorage{},
	"UcbRepresentativeStorage":       UcbDataStorage.UcbRepresentativeStorage{},
	"UcbManagerConfigStorage":        UcbDataStorage.UcbManagerConfigStorage{},
	"UcbRepresentativeConfigStorage": UcbDataStorage.UcbRepresentativeConfigStorage{},
	"UcbRegionConfigStorage":         UcbDataStorage.UcbRegionConfigStorage{},
	"UcbHospitalConfigStorage":       UcbDataStorage.UcbHospitalConfigStorage{},
	"UcbResourceConfigStorage":       UcbDataStorage.UcbResourceConfigStorage{},
	"UcbGoodsConfigStorage":          UcbDataStorage.UcbGoodsConfigStorage{},
	"UcbBusinessinputStorage":        UcbDataStorage.UcbBusinessinputStorage{},
	"UcbRepresentativeinputStorage":  UcbDataStorage.UcbRepresentativeinputStorage{},
	"UcbManagerinputStorage":         UcbDataStorage.UcbManagerinputStorage{},
	"UcbPaperinputStorage":           UcbDataStorage.UcbPaperinputStorage{},
	"UcbDestConfigStorage":           UcbDataStorage.UcbDestConfigStorage{},
	"UcbScenarioStorage":             UcbDataStorage.UcbScenarioStorage{},
	"UcbProposalStorage":             UcbDataStorage.UcbProposalStorage{},
	"UcbUseableProposalStorage":      UcbDataStorage.UcbUseableProposalStorage{},
	"UcbPaperStorage":                UcbDataStorage.UcbPaperStorage{},
	"UcbSalesConfigStorage":		  UcbDataStorage.UcbSalesConfigStorage{},

	"UcbSalesReportStorage":		  UcbDataStorage.UcbSalesReportStorage{},
	"UcbHospitalSalesReportStorage":  UcbDataStorage.UcbHospitalSalesReportStorage{},
	"UcbProductSalesReportStorage":   UcbDataStorage.UcbProductSalesReportStorage{},
	"UcbRepresentativeSalesReportStorage":	UcbDataStorage.UcbRepresentativeSalesReportStorage{},
	"UcbTeamConfigStorage":			  UcbDataStorage.UcbTeamConfigStorage{},
	"UcbActionKpiStorage":			  UcbDataStorage.UcbActionKpiStorage{},
	"UcbPersonnelAssessmentStorage":  UcbDataStorage.UcbPersonnelAssessmentStorage{},
	"UcbRepresentativeAbilityStorage":  UcbDataStorage.UcbRepresentativeAbilityStorage{},

	"UcbLevelStorage":						UcbDataStorage.UcbLevelStorage{},
	"UcbLevelConfigStorage":				UcbDataStorage.UcbLevelConfigStorage{},
	"UcbAssessmentReportDescribeStorage":	UcbDataStorage.UcbAssessmentReportDescribeStorage{},
	"UcbRegionalDivisionResultStorage":		UcbDataStorage.UcbRegionalDivisionResultStorage{},
	"UcbTargetAssignsResultStorage":		UcbDataStorage.UcbTargetAssignsResultStorage{},
	"UcbResourceAssignsResultStorage":		UcbDataStorage.UcbResourceAssignsResultStorage{},
	"UcbManageTimeResultStorage":			UcbDataStorage.UcbManageTimeResultStorage{},
	"UcbManageTeamResultStorage":			UcbDataStorage.UcbManageTeamResultStorage{},
	"UcbAssessmentReportStorage":			UcbDataStorage.UcbAssessmentReportStorage{},
	"UcbTitleStorage":						UcbDataStorage.UcbTitleStorage{},
	"UcbGeneralPerformanceResultStorage":	UcbDataStorage.UcbGeneralPerformanceResultStorage{},
	"UcbGoodsinputStorage":					UcbDataStorage.UcbGoodsinputStorage{},
	"UcbCityStorage":						UcbDataStorage.UcbCityStorage{},
	"UcbManagerGoodsConfigStorage":			UcbDataStorage.UcbManagerGoodsConfigStorage{},
	"UcbCitySalesReportStorage":			UcbDataStorage.UcbCitySalesReportStorage{},
}

var NTM_RESOURCE_FACTORY = map[string]interface{}{
	"UcbImageResource":                UcbResource.UcbImageResource{},
	"UcbPolicyResource":               UcbResource.UcbPolicyResource{},
	"UcbHospitalResource":             UcbResource.UcbHospitalResource{},
	"UcbDepartmentResource":           UcbResource.UcbDepartmentResource{},
	"UcbRegionResource":               UcbResource.UcbRegionResource{},
	"UcbProductResource":              UcbResource.UcbProductResource{},
	"UcbProductConfigResource":        UcbResource.UcbProductConfigResource{},
	"UcbRepresentativeResource":       UcbResource.UcbRepresentativeResource{},
	"UcbManagerConfigResource":        UcbResource.UcbManagerConfigResource{},
	"UcbRepresentativeConfigResource": UcbResource.UcbRepresentativeConfigResource{},
	"UcbRegionConfigResource":         UcbResource.UcbRegionConfigResource{},
	"UcbHospitalConfigResource":       UcbResource.UcbHospitalConfigResource{},
	"UcbResourceConfigResource":       UcbResource.UcbResourceConfigResource{},
	"UcbGoodsConfigResource":          UcbResource.UcbGoodsConfigResource{},
	"UcbBusinessinputResource":        UcbResource.UcbBusinessinputResource{},
	"UcbRepresentativeinputResource":  UcbResource.UcbRepresentativeinputResource{},
	"UcbManagerinputResource":         UcbResource.UcbManagerinputResource{},
	"UcbPaperinputResource":           UcbResource.UcbPaperinputResource{},
	"UcbDestConfigResource":           UcbResource.UcbDestConfigResource{},
	"UcbScenarioResource":             UcbResource.UcbScenarioResource{},
	"UcbProposalResource":             UcbResource.UcbProposalResource{},
	"UcbUseableProposalResource":      UcbResource.UcbUseableProposalResource{},
	"UcbPaperResource":                UcbResource.UcbPaperResource{},
	"UcbSalesConfigResource":		   UcbResource.UcbSalesConfigResource{},

	"UcbSalesReportResource":		   UcbResource.UcbSalesReportResource{},
	"UcbHospitalSalesReportResource":  UcbResource.UcbHospitalSalesReportResource{},
	"UcbProductSalesReportResource":   UcbResource.UcbProductSalesReportResource{},
	"UcbRepresentativeSalesReportResource":	UcbResource.UcbRepresentativeSalesReportResource{},
	"UcbTeamConfigResource":		   UcbResource.UcbTeamConfigResource{},
	"UcbActionKpiResource":		   	   UcbResource.UcbActionKpiResource{},
	"UcbPersonnelAssessmentResource":  UcbResource.UcbPersonnelAssessmentResource{},
	"UcbRepresentativeAbilityResource":  UcbResource.UcbRepresentativeAbilityResource{},



	"UcbLevelResource":						UcbResource.UcbLevelResource{},
	"UcbLevelConfigResource":				UcbResource.UcbLevelConfigResource{},
	"UcbAssessmentReportDescribeResource":	UcbResource.UcbAssessmentReportDescribeResource{},
	"UcbRegionalDivisionResultResource":	UcbResource.UcbRegionalDivisionResultResource{},
	"UcbTargetAssignsResultResource":		UcbResource.UcbTargetAssignsResultResource{},
	"UcbResourceAssignsResultResource":		UcbResource.UcbResourceAssignsResultResource{},
	"UcbManageTimeResultResource":			UcbResource.UcbManageTimeResultResource{},
	"UcbManageTeamResultResource":			UcbResource.UcbManageTeamResultResource{},
	"UcbAssessmentReportResource":			UcbResource.UcbAssessmentReportResource{},
	"UcbTitleResource":						UcbResource.UcbTitleResource{},
	"UcbGeneralPerformanceResultResource":	UcbResource.UcbGeneralPerformanceResultResource{},
	"UcbGoodsinputResource":				UcbResource.UcbGoodsinputResource{},
	"UcbCityResource":						UcbResource.UcbCityResource{},
	"UcbManagerGoodsConfigResource":		UcbResource.UcbManagerGoodsConfigResource{},
	"UcbCitySalesReportResource":			UcbResource.UcbCitySalesReportResource{},
}

var NTM_FUNCTION_FACTORY = map[string]interface{}{
	"UcbCommonPanicHandle":         	UcbHandler.CommonPanicHandle{},
	"UcbGeneratePaperHandler": 			UcbHandler.UcbGeneratePaperHandler{},
	"UcbCallRHandler":					UcbHandler.UcbCallRHandler{},
	"UcbRResulrHandler":				UcbHandler.UcbCallRHandler{},
}
var NTM_MIDDLEWARE_FACTORY = map[string]interface{}{
	"UcbCheckTokenMiddleware": UcbMiddleware.UcbCheckTokenMiddleware{},
}

var NTM_DAEMON_FACTORY = map[string]interface{}{
	"BmMongodbDaemon": BmMongodb.BmMongodb{},
	"BmRedisDaemon":   BmRedis.BmRedis{},
}

func (t UcbTable) GetModelByName(name string) interface{} {
	return NTM_MODEL_FACTORY[name]
}

func (t UcbTable) GetResourceByName(name string) interface{} {
	return NTM_RESOURCE_FACTORY[name]
}

func (t UcbTable) GetStorageByName(name string) interface{} {
	return NTM_STORAGE_FACTORY[name]
}

func (t UcbTable) GetDaemonByName(name string) interface{} {
	return NTM_DAEMON_FACTORY[name]
}

func (t UcbTable) GetFunctionByName(name string) interface{} {
	return NTM_FUNCTION_FACTORY[name]
}

func (t UcbTable) GetMiddlewareByName(name string) interface{} {
	return NTM_MIDDLEWARE_FACTORY[name]
}
