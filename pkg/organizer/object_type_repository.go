package organizer

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
	OrganizerRepository interface {
		ListOrganizer(ctx context.Context, search string, offset, limit int, sortBy, sortDir *string) (*[]entities.Organizer, int64, error)
		CreateOrganizer(ctx context.Context, ot *entities.Organizer) error
		GetOrganizer(ctx context.Context, id uint) (*entities.Organizer, error)
		GetOrganizerByName(ctx context.Context, name string) (*entities.Organizer, error)
		UpdateOrganizer(ctx context.Context, ot entities.Organizer) (*entities.Organizer, error)
		DeleteOrganizer(ctx context.Context, id uint) error
	}

	organizerRepository struct {
		db *gorm.DB
	}
)

func NewOrganizerRepository(db *gorm.DB) OrganizerRepository {
	return &organizerRepository{db: db}
}

func (o *organizerRepository) ListOrganizer(ctx context.Context, search string, offset, limit int, sortBy, sortDir *string) (*[]entities.Organizer, int64, error) {
	var organizers []entities.Organizer

	validSortBy := []string{"id", "name", "bank_name", "created_at"}
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

	query := o.db.WithContext(ctx).Model(&entities.Organizer{})

	if search != "" {
		sq := "%" + search + "%"
		query = query.Where("name ILIKE ? or bank_name ILIKE ?", sq, sq) // biar case-insensitive di postgres
	}

	// count harus setelah filter
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// apply pagination & sorting
	if err := query.Offset(offset).Limit(limit).Order(orderStr).Find(&organizers).Error; err != nil {
		return nil, 0, err
	}

	return &organizers, total, nil
}

func (o *organizerRepository) CreateOrganizer(ctx context.Context, ot *entities.Organizer) error {
	if err := o.db.WithContext(ctx).Create(ot).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return gorm.ErrDuplicatedKey
		}
		return err
	}
	return nil
}

func (o *organizerRepository) GetOrganizer(ctx context.Context, id uint) (*entities.Organizer, error) {
	var organizer entities.Organizer
	if err := o.db.WithContext(ctx).First(&organizer, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &organizer, nil
}

func (o *organizerRepository) GetOrganizerByName(ctx context.Context, name string) (*entities.Organizer, error) {
	var organizer entities.Organizer
	if err := o.db.WithContext(ctx).Where("name = ?", name).First(&organizer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &organizer, nil
}

func (o *organizerRepository) UpdateOrganizer(ctx context.Context, ot entities.Organizer) (*entities.Organizer, error) {
	// only update non-zero fields
	if err := o.db.WithContext(ctx).Model(&entities.Organizer{}).
		Where("id = ?", ot.ID).
		Updates(map[string]interface{}{
			"name": ot.Name,
		}).Error; err != nil {
		return nil, err
	}
	return &ot, nil
}

func (o *organizerRepository) DeleteOrganizer(ctx context.Context, id uint) error {
	tx := o.db.WithContext(ctx).Where("id = ?", id).Delete(&entities.Organizer{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
