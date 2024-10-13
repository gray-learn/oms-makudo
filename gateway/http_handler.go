package main

import (
	"fmt"
	"net/http"

	"errors"

	common "github.com/sikozonpc/commons"
	pb "github.com/sikozonpc/commons/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	// gateway

	// client
	orderClient pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler { // return type is a pointer to handler
	return &handler{orderClient: client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	// static folder serving
	mux.Handle("/", http.FileServer(http.Dir("public")))

	mux.HandleFunc("POST /api/customers/{custID}/orders", h.handleCreateOrder)
	// mux.HandleFunc("GET /api/customers/{custID}/orders/{orderID}", h.handleGetOrder)
}

func (h *handler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {

	if h.orderClient == nil {
		http.Error(w, "Order service unavailable", http.StatusServiceUnavailable)
		return
	}
	custID := r.PathValue("custID")
	fmt.Println(r)
	fmt.Println(r.Body)

	var items []*pb.ItemWithQuantity                   // point pb.ItemWithQuantity is going to slice[]
	if err := common.ReadJSON(r, &items); err != nil { // consume the request
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := validateItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	o, err := h.orderClient.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustId: custID,
		Items:  items,
	})

	fmt.Println("order items", o)
	rStatus := status.Convert(err)
	if rStatus != nil {
		if rStatus.Code() == codes.InvalidArgument {
			common.WriteError(w, http.StatusBadRequest, rStatus.Message()) // 400
			return
		}

		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSON(w, http.StatusOK, o)
}

func (h *handler) handleGetOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Get Order")
}

func validateItems(items []*pb.ItemWithQuantity) error {
	if len(items) == 0 {
		return common.ErrNoItems
	}

	for _, i := range items {
		fmt.Println("item", i)
		if i.ID == "" {
			return errors.New("item ID is required")
		}
		if i.Quantity <= 0 {
			return errors.New("items must be greater than 0")
		}
	}
	return nil
}
