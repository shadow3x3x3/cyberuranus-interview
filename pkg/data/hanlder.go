package data

import (
	"cyberuranus-interview/pkg/common"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// DataHandler represents the HTTP handler for resource of data.
type DataHandler struct {
	db *sql.DB
}

// ServeHTTP implements the http Handler interface for DataHandler
func (h DataHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set(common.ContentType, common.JSONType)

	switch req.Method {
	case common.GetMethod: // GET /data
		getAllData(resp, req, h.db)
	case common.PostMethod: // POST /data
		postData(resp, req, h.db)
	default:
		jsonResult := common.MakeGeneralResp(http.StatusMethodNotAllowed)
		resp.Write(jsonResult)
	}
}

func getAllData(resp http.ResponseWriter, req *http.Request, db *sql.DB) {
	isError := false
	defer func() {
		if isError {
			http.Error(resp, "Get Data Failed", http.StatusInternalServerError)
		}
	}()

	allData, err := GetData(db)
	if err != nil {
		log.Println(err)
		isError = true
		return
	}

	jsonResult, err := common.MakeDataResp(http.StatusOK, "Get All Data Success", allData)
	if err != nil {
		log.Println(err)
		isError = true
		return
	}

	resp.Write(jsonResult)
}

func postData(resp http.ResponseWriter, req *http.Request, db *sql.DB) {
	if req.Body == nil {
		http.Error(resp, "Empty Request Body", http.StatusBadRequest)
		return
	}

	var d Data
	// TODO: Do user friendly response without raw internal error
	if err := json.NewDecoder(req.Body).Decode(&d); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
	}
	d.DateAdded = time.Now()

	if err := InsertData(db, d); err != nil {
		log.Println(err)
		jsonResult := common.MakeMessageResp(http.StatusBadRequest, err.Error())
		resp.Write(jsonResult)
		return
	}

	jsonResult := common.MakeGeneralResp(http.StatusOK)
	resp.Write(jsonResult)
}

// Register registers data routers to http.ServeMux with the db instance
func Register(mux *http.ServeMux, db *sql.DB) {
	mux.Handle("/data", &DataHandler{db})
}
