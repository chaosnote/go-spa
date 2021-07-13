package web

import (
	"build/log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

//@see https://github.com/gorilla/mux#serving-single-page-applications

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)
	log.Console("path", zap.String("static", path))

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

// Register 設置靜態資料夾
func Register(r *mux.Router) {

	dir, e := os.Getwd() // 取得當前執行路徑
	if e != nil {
		log.Console("os", zap.Error(e))
	}

	var p = filepath.Join(dir, "./front")
	log.Console("local folder", zap.String("path", p))

	spa := spaHandler{staticPath: p, indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

}
