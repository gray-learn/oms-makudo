package main

import (
	"log"
	"net/http"

	common "github.com/sikozonpc/commons"
)

type handler struct {
	// gateway

	// client
	client pb.NewOrderServiceClient
}

func NewHandler() *handler { // return type is a pointer to handler
	return &handler{}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/cusotmers/{custID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello")

	custID := r.PathValue("custID")
	var items []*pb.ItemsWithQuantity // point to slice[]
	if err:= common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())	
		return
	}


	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustID: custID,
		Items: 
	})

}
