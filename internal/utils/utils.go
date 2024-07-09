// package utils

// import (
// 	"encoding/json"
// 	"net/http"
// )

// func (r *http.Request) jsonDecode(w http.ResponseWriter, target *any) {
// 	err := json.NewDecoder(r.Body).Decode(&target)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// }