package models

import (
	"encoding/json"
	. "mon-api/types"
	"net/http"
)

type AcceptFormat struct {
	Supports []string `json:"supports"`
}

func NewAcceptFormat() AcceptFormat{
	return AcceptFormat{Supports : []string{"image/jpeg", "image/png","image/jpg","image/gif", "image/webp"}}
}

func (support *AcceptFormat) GetAcceptFormat(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Méthode non autorisée"})
		return 
	}
	json.NewEncoder(w).Encode(support)
}



func (AcceptFormatList *AcceptFormat) TestIsSupport(format string) bool{
	for _, f := range AcceptFormatList.Supports {
		if (f == format){
			return true
		}
	}
	return false
}