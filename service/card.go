package service

import (
	"context"
	"finance/api/models"
	"finance/storage"
	"fmt"
)

type cardService struct {
	storage storage.IStorage
}

func NewCardService(storage storage.IStorage) cardService {
	return cardService{storage: storage}
}

func (s cardService) Create(ctx context.Context, card models.CreateCard) (string, error) {

	CardId, err := s.storage.CardStorage().CreateCard(ctx, card)
	if err != nil {
		fmt.Println("error while creating card, err: ", err)
		return "", err
	}

	return CardId, nil
}

func (s cardService) Delete(ctx context.Context,id string) error {
	err:=s.storage.CardStorage().Delete(ctx,id)
	if err != nil {
		return err
	}
	
	return nil
}