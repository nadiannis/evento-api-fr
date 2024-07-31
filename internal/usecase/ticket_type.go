package usecase

import (
	"github.com/nadiannis/evento-api-fr/internal/domain"
	"github.com/nadiannis/evento-api-fr/internal/domain/request"
	"github.com/nadiannis/evento-api-fr/internal/repository"
)

type TicketTypeUsecase struct {
	repository repository.ITicketTypeRepository
}

func NewTicketTypeUsecase(repository repository.ITicketTypeRepository) ITicketTypeUsecase {
	return &TicketTypeUsecase{
		repository: repository,
	}
}

func (u *TicketTypeUsecase) GetAll() ([]*domain.TicketType, error) {
	return u.repository.GetAll()
}

func (u *TicketTypeUsecase) Add(input *request.TicketTypeRequest) (*domain.TicketType, error) {
	ticketType := &domain.TicketType{
		Name:  input.Name,
		Price: input.Price,
	}

	err := u.repository.Add(ticketType)
	if err != nil {
		return nil, err
	}

	return ticketType, nil
}
