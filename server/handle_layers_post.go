package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dimfeld/httptreemux"
	"github.com/go-spatial/tegola/internal/log"
	"github.com/go-spatial/tegola/internal/ttools"
	"github.com/go-spatial/tegola/model"
	"gorm.io/gorm"
)

type HandlePostLayers struct {
	// layer id, eg. ?id=123
	id int
}

func (req *HandlePostLayers) parseParams(r *http.Request) error {
	params := httptreemux.ContextParams(r.Context())

	log.Debugf("[Layers] [Get] [parseParams] params: %+v", params)

	log.Debugf("[Layers] [Get] [parseParams] search params: %s", r.URL.RawQuery)

	id_str := r.FormValue("id")
	if id_str == "" {
		req.id = -1
		return nil
	}
	s, e := strconv.Atoi(id_str)
	if e != nil {
		return e
	}
	req.id = s

	return nil
}

// urlExample := "postgres://username:password@localhost:5432/database_name"
var POSTGRES_DSN string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
	ttools.GetEnvDefault("PG_DB_USER", "postgres"),
	ttools.GetEnvDefault("PG_DB_PASSWORD", "postgres"),
	ttools.GetEnvDefault("PG_DB_HOST", "localhost"),
	ttools.GetEnvDefault("PG_DB_PORT", "5432"),
	ttools.GetEnvDefault("PG_DB_NAME", "sobo"),
)

func (req HandlePostLayers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("[Layers] [Get] [ServeHTTP] start")
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

	var result *gorm.DB
	data := model.MArea{}
	datas := []model.MArea{}
	// query
	if req.id == -1 {
		result = db.Find(&datas)
	} else {
		result = db.First(&data, req.id)
	}

	if result.Error != nil {
		log.Error(result.Error)
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	log.Debug("[Layers] [Get] [ServeHTTP] end")
	if req.id == -1 {
		json.NewEncoder(w).Encode(datas)
	} else {
		json.NewEncoder(w).Encode(data)
	}

}
