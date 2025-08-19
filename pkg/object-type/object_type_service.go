package objectType

import (
	"context"
	"errors"

	"github.com/yayayapluto/api-ukk-online/domain"
	"github.com/yayayapluto/api-ukk-online/entities"
	"gorm.io/gorm"
)

type (
	ObjectTypeService interface {
		ListObjectType(ctx context.Context, search string, offset, limit int, sortBy, sortDir *string) (*[]entities.ObjectType, int64, error)
		CreateObjectType(ctx context.Context, ot *entities.ObjectType) error
		GetObjectType(ctx context.Context, id uint) (*entities.ObjectType, error)
		UpdateObjectType(ctx context.Context, ot entities.ObjectType) (*entities.ObjectType, error)
		DeleteObjectType(ctx context.Context, id uint) error
	}

	objectTypeService struct {
		repo ObjectTypeRepository
	}
)

func NewObjectTypeService(r ObjectTypeRepository) ObjectTypeService {
	return &objectTypeService{repo: r}
}

func (o objectTypeService) ListObjectType(ctx context.Context, search string, offset, limit int, sortBy, sortDir *string) (*[]entities.ObjectType, int64, error) {
	return o.repo.ListObjectType(ctx, search, offset, limit, sortBy, sortDir)
}

func (o objectTypeService) CreateObjectType(ctx context.Context, ot *entities.ObjectType) error {
	if err := o.repo.CreateObjectType(ctx, ot); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.ErrObjectTypeAlreadyExists
		}
		return err
	}
	return nil
}

func (o objectTypeService) GetObjectType(ctx context.Context, id uint) (*entities.ObjectType, error) {
	ot, err := o.repo.GetObjectType(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrObjectTypeNotFound
		}
		return nil, err
	}
	return ot, nil
}

func (o objectTypeService) UpdateObjectType(ctx context.Context, ot entities.ObjectType) (*entities.ObjectType, error) {
	if ot.Name == "" {
		return nil, domain.ErrInvalidPayload // jangan return nil, nil
	}
	updated, err := o.repo.UpdateObjectType(ctx, ot)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrObjectTypeNotFound
		}
		return nil, err
	}
	return updated, nil
}

func (o objectTypeService) DeleteObjectType(ctx context.Context, id uint) error {
	if err := o.repo.DeleteObjectType(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrObjectTypeNotFound
		}
		return err
	}
	return nil
}

