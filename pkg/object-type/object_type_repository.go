package objectType

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/yayayapluto/api-ukk-online/domain"
	"github.com/yayayapluto/api-ukk-online/entities"
	"gorm.io/gorm"
	"slices"
)

type (
	ObjectTypeRepository interface {
		ListObjectType(ctx context.Context, search string, offset, limit int, sortBy, sortDir *string) (*[]entities.ObjectType, int64, error)
		CreateObjectType(ctx context.Context, ot *entities.ObjectType) error
		GetObjectType(ctx context.Context, id uint) (*entities.ObjectType, error)
		GetObjectTypeByName(ctx context.Context, name string) (*entities.ObjectType, error)
		UpdateObjectType(ctx context.Context, ot entities.ObjectType) (*entities.ObjectType, error)
		DeleteObjectType(ctx context.Context, id uint) error
	}

	objectTypeRepository struct {
		db *gorm.DB
	}
)

func NewObjectTypeRepository(db *gorm.DB) ObjectTypeRepository {
	return &objectTypeRepository{db: db}
}

func (o *objectTypeRepository) ListObjectType(ctx context.Context, search string, offset, limit int, sortBy, sortDir *string) (*[]entities.ObjectType, int64, error) {
	var objectTypes []entities.ObjectType

	validSortBy := []string{"id", "name", "created_at"}
	defSortBy := validSortBy[0]
	if sortBy != nil {
		if !slices.Contains(validSortBy, *sortBy) {
			return nil, 0, domain.ErrInvalidSortByColumn
		}
		defSortBy = *sortBy
	}

	validSortDir := []string{"asc", "desc"}
	defSortDir := validSortDir[0]
	if sortDir != nil {
		if !slices.Contains(validSortDir, *sortDir) {
			return nil, 0, domain.ErrInvalidSortDir
		}
		defSortDir = *sortDir
	}

	orderStr := fmt.Sprintf("%s %s", defSortBy, defSortDir)

	query := o.db.WithContext(ctx).Limit(limit).Offset(offset).Order(orderStr)

	if search != "" && len(search) != 0 {
		sq := "%" + search + "%"
		query = query.Where("name LIKE ?", sq)
	}

	if err := query.Find(&objectTypes).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := o.db.WithContext(ctx).Model(&entities.ObjectType{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return &objectTypes, total, nil
}

func (o *objectTypeRepository) CreateObjectType(ctx context.Context, ot *entities.ObjectType) error {
	if err := o.db.WithContext(ctx).Create(ot).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return gorm.ErrDuplicatedKey
		}
		return err
	}
	return nil
}

func (o *objectTypeRepository) GetObjectType(ctx context.Context, id uint) (*entities.ObjectType, error) {
	var objectType entities.ObjectType
	if err := o.db.WithContext(ctx).Where("id = ?", id).First(&objectType).Error; err != nil {
		return nil, err
	}
	return &objectType, nil
}

func (o *objectTypeRepository) GetObjectTypeByName(ctx context.Context, name string) (*entities.ObjectType, error) {
	var objectType entities.ObjectType
	if err := o.db.WithContext(ctx).Where("name = ?", name).First(&objectType).Error; err != nil {
		return nil, err
	}
	return &objectType, nil
}

func (o *objectTypeRepository) UpdateObjectType(ctx context.Context, ot entities.ObjectType) (*entities.ObjectType, error) {
	if err := o.db.WithContext(ctx).Where("id = ?", ot.ID).Updates(&ot).Error; err != nil {
		return nil, err
	}
	return &ot, nil
}

func (o *objectTypeRepository) DeleteObjectType(ctx context.Context, id uint) error {
	if err := o.db.WithContext(ctx).Delete("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
