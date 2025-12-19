package handlers

import (
	"fmt"
	"net/http"
)

// HandleFirstCallJS serves the first call JavaScript
func HandleFirstCallJS(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	//fmt.Fprint(w, templates.FirstCallJS) // todo fix this
	fmt.Println("entered 1st call")
	http.ServeFile(w, r, "storage/js/firstcall.js")
}
