package main

import (
	"context"
	"log"

	common "github.com/sikozonpc/commons"
	pb "github.com/sikozonpc/commons/api"
)

type service struct {
	store OrderStore
}

func NewService(store OrderStore) *service { // return type is a pointer to service
	return &service{store}
}

func (s *service) CreateOrder(context.Context) error {
	return nil
}

func (s *service) validateOrder(ctx context.Context, p *pb.CreateOrderRequest) error {
	if len(p.Items) == 0 {
		return common.ErrNoItems
	}
	mergeItems := mergeItemsQuantities(p.Items)
	log.Print((mergeItems))
	// validate with stock service
	return nil
}

func mergeItemsQuantities(items []*pb.ItemWithQuantity) []*pb.ItemWithQuantity {
	merged := make([]*pb.ItemWithQuantity, 0)

	for _, item := range items {
		found := false
		for _, finalItem := range merged {
			if finalItem.ID == item.ID {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}
		}
		if !found {
			merged = append(merged, item)
		}
	}
	return merged
}
