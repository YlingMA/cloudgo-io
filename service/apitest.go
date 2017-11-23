package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			NAME   string `json:"name"`
			Q1     string `json:"q1"`
			Q2     string `json:"q2"`
			Q3     string `json:"q3"`
			Q3last string `json:"q3last"`
		}{NAME: "xiaomi", Q1: "0", Q2: "0", Q3: "0", Q3last: "0"})
	}
}
