package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
	"github.com/gorilla/mux"
)

func DeleteCompany(w http.ResponseWriter, r *http.Request, c pb.CompanyServiceClient) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert Id", http.StatusBadRequest)
	}

	params := &pb.DeleteCompanyRequest{
		Id: int64(id),
	}

	res, err := c.DeleteCompany(r.Context(), params)

	if err != nil {
		fmt.Println("API Gateway :  DeleteCompany - ERROR")
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("API Gateway :  DeleteCompany - SUCCESS")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
