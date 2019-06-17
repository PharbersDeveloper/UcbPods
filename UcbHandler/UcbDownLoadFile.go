package UcbHandler

import (
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
)

type UcbDownLoadFileHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
}

func (h UcbDownLoadFileHandler) NewDownLoadFileHandler(args ...interface{}) UcbDownLoadFileHandler {
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

	return UcbDownLoadFileHandler{Method: md, HttpMethod: hm, Args: ag, db: m, rd: r}
}

func (h UcbDownLoadFileHandler) DownLoad(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	query := r.URL.Query()
	q, ok := query["filename"]

	if ok {
		filename := q[0]
		env := os.Getenv("DOWNLOAD")
		localFile := fmt.Sprint(env, filename)
		out, err := ioutil.ReadFile(localFile)
		if err != nil {
			fmt.Println("error")
			fmt.Println(err.Error())
		}
		w.Header().Set("Content-Disposition", "attachment; filename=" + filename)
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		w.Header().Set("charset", "utf-8")
		w.Write(out)
	}

	return 0
}

func (h UcbDownLoadFileHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h UcbDownLoadFileHandler) GetHandlerMethod() string {
	return h.Method
}
