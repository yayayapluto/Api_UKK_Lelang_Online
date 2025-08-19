package organizer

import (
	"context"
	"errors"

	"github.com/yayayapluto/api-ukk-online/domain"
	"github.com/yayayapluto/api-ukk-online/entities"
	"gorm.io/gorm"
)

type (
	OrganizerService interface {
		ListOrganizer(ctx context.Context, search string, offset, limit int, sortBy, sortDir *string) (*[]entities.Organizer, int64, error)
		CreateOrganizer(ctx context.Context, ot *entities.Organizer) error
		GetOrganizer(ctx context.Context, id uint) (*entities.Organizer, error)
		UpdateOrganizer(ctx context.Context, ot entities.Organizer) (*entities.Organizer, error)
		DeleteOrganizer(ctx context.Context, id uint) error
	}

	organizerService struct {
		repo OrganizerRepository
	}
)

func NewOrganizerService(r OrganizerRepository) OrganizerService {
	return &organizerService{repo: r}
}

func (o organizerService) ListOrganizer(ctx context.Context, search string, offset, limit int, sortBy, sortDir *string) (*[]entities.Organizer, int64, error) {
	return o.repo.ListOrganizer(ctx, search, offset, limit, sortBy, sortDir)
}

func (o organizerService) CreateOrganizer(ctx context.Context, ot *entities.Organizer) error {
	if err := o.repo.CreateOrganizer(ctx, ot); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.ErrOrganizerAlreadyExists
		}
		return err
	}
	return nil
}

func (o organizerService) GetOrganizer(ctx context.Context, id uint) (*entities.Organizer, error) {
	ot, err := o.repo.GetOrganizer(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrOrganizerNotFound
		}
		return nil, err
	}
	return ot, nil
}

func (o organizerService) UpdateOrganizer(ctx context.Context, ot entities.Organizer) (*entities.Organizer, error) {
	if ot.Name == "" {
		return nil, domain.ErrInvalidPayload // jangan return nil, nil
	}
	updated, err := o.repo.UpdateOrganizer(ctx, ot)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrOrganizerNotFound
		}
		return nil, err
	}
	return updated, nil
}

func (o organizerService) DeleteOrganizer(ctx context.Context, id uint) error {
	if err := o.repo.DeleteOrganizer(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrOrganizerNotFound
		}
		return err
	}
	return nil
}
