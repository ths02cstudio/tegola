package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/go-spatial/tegola/internal/log"
	"github.com/go-spatial/tegola/model"
)

type HandlePutLayers struct {
	version model.VArea
	layers  []model.DArea
}

func (req *HandlePutLayers) parseParams(r *http.Request) error {
	params := httptreemux.ContextParams(r.Context())

	log.Debugf("[Layers] [Post] [parseParams] params: %+v", params)

	log.Debugf("[Layers] [Post] [parseParams] form values: %s", r.FormValue("layers"))

	if err := json.Unmarshal([]byte(r.FormValue("version")), &req.version); err != nil {
		log.Errorf("[Layers] [Post] [parseParams] error: %+v", err)
		return fmt.Errorf("convert string to struct failed: {version: %s}", r.FormValue("version"))
	}
	log.Debugf("[Layers] [Post] [parseParams] version: %+v", req.version)

	if err := json.Unmarshal([]byte(r.FormValue("layers")), &req.layers); err != nil {
		log.Errorf("[Layers] [Post] [parseParams] error: %+v", err)
		return fmt.Errorf("convert string to struct failed: {layers: %s}", r.FormValue("layers"))
	}
	log.Debugf("[Layers] [Post] [parseParams] layers: %+v", req.layers)

	return nil
}

func (req HandlePutLayers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("[Layers] [Post] [ServeHTTP] start")
	// parse our URI
	if err := req.parseParams(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := model.GetPgDb()
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// insert
	result := db.Create(req.layers)
	if result.Error != nil {
		log.Error(result.Error)
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	log.Debug("[Layers] [Post] [ServeHTTP] end")
	json.NewEncoder(w).Encode(map[string]int64{
		"RowsAffected": result.RowsAffected,
	})
}
