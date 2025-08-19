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

	query := o.db.WithContext(ctx).Model(&entities.ObjectType{})

	if search != "" {
		sq := "%" + search + "%"
		query = query.Where("name ILIKE ?", sq) // biar case-insensitive di postgres
	}

	// count harus setelah filter
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// apply pagination & sorting
	if err := query.Offset(offset).Limit(limit).Order(orderStr).Find(&objectTypes).Error; err != nil {
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
	if err := o.db.WithContext(ctx).First(&objectType, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &objectType, nil
}

func (o *objectTypeRepository) GetObjectTypeByName(ctx context.Context, name string) (*entities.ObjectType, error) {
	var objectType entities.ObjectType
	if err := o.db.WithContext(ctx).Where("name = ?", name).First(&objectType).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &objectType, nil
}

func (o *objectTypeRepository) UpdateObjectType(ctx context.Context, ot entities.ObjectType) (*entities.ObjectType, error) {
	// only update non-zero fields
	if err := o.db.WithContext(ctx).Model(&entities.ObjectType{}).
		Where("id = ?", ot.ID).
		Updates(map[string]interface{}{
			"name": ot.Name,
		}).Error; err != nil {
		return nil, err
	}
	return &ot, nil
}

func (o *objectTypeRepository) DeleteObjectType(ctx context.Context, id uint) error {
	tx := o.db.WithContext(ctx).Where("id = ?", id).Delete(&entities.ObjectType{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
