package UcbHandler

import (
	"Ucb/UcbDataStorage"
	"Ucb/UcbModel"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type UcbGeneratePaperHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
}

func (h UcbGeneratePaperHandler) NewGeneratePaperHandler(args ...interface{}) UcbGeneratePaperHandler {
	var m *BmMongodb.BmMongodb
	var r *BmRedis.BmRedis
	var hm string
	var md string
	var ag []string
	for i, arg := range args {
		if i == 0 {
			sts := arg.([]BmDaemons.BmDaemon)
			for _, dm := range sts {
				tp := reflect.ValueOf(dm).Interface()
				tm := reflect.ValueOf(tp).Elem().Type()
				if tm.Name() == "BmMongodb" {
					m = dm.(*BmMongodb.BmMongodb)
				}
				if tm.Name() == "BmRedis" {
					r = dm.(*BmRedis.BmRedis)
				}
			}
		} else if i == 1 {
			md = arg.(string)
		} else if i == 2 {
			hm = arg.(string)
		} else if i == 3 {
			lst := arg.([]string)
			for _, str := range lst {
				ag = append(ag, str)
			}
		} else {
		}
	}

	return UcbGeneratePaperHandler{Method: md, HttpMethod: hm, Args: ag, db: m, rd: r}
}

func (h UcbGeneratePaperHandler) GeneratePaper(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	mdb := []BmDaemons.BmDaemon{h.db}
	w.Header().Add("Content-Type", "application/json")

	//_, err := UcbMiddleware.UcbCheckToken.CheckTokenFormFunction(w, r)
	//if err != nil {
	//	panic(fmt.Sprintf(err.Error()))
	//}

	proposalId := r.FormValue("proposal-id")
	accountId := r.FormValue("account-id")

	if len(proposalId) == 0 || proposalId == "undefined" || len(accountId) == 0 || accountId == "undefined" {
		panic("生成Paper的参数不完整")
		return 1
	}
	proposalModel, _ := UcbDataStorage.UcbProposalStorage{}.NewProposalStorage(mdb).GetOne(proposalId)

	var (
		//out UcbModel.Paper
		paperId string
		)
	//cond := bson.M{"proposal-id": proposalId, "account-id": accountId}

	//err = h.db.FindOneByCondition(&UcbModel.Paper{}, &out, cond)

	reqs := api2go.Request{
		PlainRequest: r,
		Header: w.Header(),
		QueryParams: map[string][]string{},
	}

	reqs.QueryParams["account-id"] = []string{accountId}
	reqs.QueryParams["proposal-id"] = []string{proposalId}
	reqs.QueryParams["orderby"] = []string{"time"}

	papers := UcbDataStorage.UcbPaperStorage{}.NewPaperStorage(mdb).GetAll(reqs, -1, -1)

	paper := papers[len(papers)-1]

	if paper.InputState == 3 {
		paperModel := UcbModel.Paper{
			AccountID: accountId, //token.UserID,
			ProposalID: proposalModel.ID,
			Name: proposalModel.Name,
			Describe: proposalModel.Describe,
			TotalPhase: proposalModel.TotalPhase,
			StartTime: time.Now().UnixNano(),
			EndTime: 0,
			InputState: 0,
			InputIDs: proposalModel.InputIDs,
			SalesReportIDs: proposalModel.SalesReportIDs,
			PersonnelAssessmentIDs: proposalModel.PersonnelAssessmentIDs,
			Time: time.Now().UnixNano(),
		}

		paperId = UcbDataStorage.UcbPaperStorage{}.NewPaperStorage(mdb).Insert(paperModel)
	}

	//if err != nil && err.Error() != "not found" {
	//	panic(fmt.Sprintf(err.Error()))
	//} else if len(out.ID) > 0 {
	//	paperId = out.ID
	//} else {
	//	paperModel := UcbModel.Paper{
	//		AccountID: accountId, //token.UserID,
	//		ProposalID: proposalModel.ID,
	//		Name: proposalModel.Name,
	//		Describe: proposalModel.Describe,
	//		TotalPhase: proposalModel.TotalPhase,
	//		StartTime: time.Now().UnixNano(),
	//		EndTime: 0,
	//		InputState: 0,
	//		InputIDs: proposalModel.InputIDs,
	//		SalesReportIDs: proposalModel.SalesReportIDs,
	//		PersonnelAssessmentIDs: proposalModel.PersonnelAssessmentIDs,
	//	}
	//
	//	paperId = UcbDataStorage.UcbPaperStorage{}.NewPaperStorage(mdb).Insert(paperModel)
	//}

	//拼接转发的URL
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	toUrl := strings.Replace(r.URL.Path, "GeneratePaper", h.Args[0], -1) + "/" + paperId
	paperURL := strings.Join([]string{scheme, r.Host, toUrl}, "")

	// 转发
	client := &http.Client{}
	req, _ := http.NewRequest("GET", paperURL, nil)
	for k, v := range r.Header {
		req.Header.Add(k, v[0])
	}
	response, _ := client.Do(req)

	responseBody, _ := ioutil.ReadAll(response.Body)

	w.Header().Add("Content-Type", "application/json")
	w.Write(responseBody)

	return 0
}

func (h UcbGeneratePaperHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h UcbGeneratePaperHandler) GetHandlerMethod() string {
	return h.Method
}