package serverweb

import (
	"encoding/json"
	"fmt"
	"formation/api"
	"io"
	"net/http"
)

func Delete(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query()["id"]
	if len(id) > 0  {
		fmt.Println("create", len(api.List()))
		//io.WriteString(w, "Afficher ma Liste" )
		//fmt.Println( api.List())
		json.NewEncoder(w).Encode(api.List())
	} else {
		io.WriteString(w, "Il manque les paramètres" )
	}
}