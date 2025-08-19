package objectType

import (
	"context"
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
	return o.repo.CreateObjectType(ctx, ot)
}

func (o objectTypeService) GetObjectType(ctx context.Context, id uint) (*entities.ObjectType, error) {
	ot, err := o.repo.GetObjectType(ctx, id)
	if err != nil {
		return nil, err
	}

	if ot == nil {
		return nil, gorm.ErrRecordNotFound
	}

	return ot, nil
}

func (o objectTypeService) UpdateObjectType(ctx context.Context, ot entities.ObjectType) (*entities.ObjectType, error) {
	return o.repo.UpdateObjectType(ctx, ot)
}

func (o objectTypeService) DeleteObjectType(ctx context.Context, id uint) error {
	return o.repo.DeleteObjectType(ctx, id)
}
