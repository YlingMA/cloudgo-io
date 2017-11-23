package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func tableHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		name := req.FormValue("name")
		q1 := req.FormValue("q1")
		q2 := req.FormValue("q2")
		q3 := req.FormValue("q3")
		q3last := req.FormValue("q3last")
		formatter.HTML(w, http.StatusOK, "table_in", struct {
			NAME   string
			Q1     string
			Q2     string
			Q3     string
			Q3last string
		}{NAME: name, Q1: q1, Q2: q2, Q3: q3, Q3last: q3last})
	}
}
