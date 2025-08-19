package handlers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/yayayapluto/api-ukk-online/domain"
	"github.com/yayayapluto/api-ukk-online/entities"
	"github.com/yayayapluto/api-ukk-online/internal/api/presenters"
	"github.com/yayayapluto/api-ukk-online/internal/utils/pagination"
	organizer "github.com/yayayapluto/api-ukk-online/pkg/organizer"
	"gorm.io/gorm"
	"math"
)

type (
	OrganizerHandler interface {
		List(ctx *fiber.Ctx) error
		Create(ctx *fiber.Ctx) error
		Get(ctx *fiber.Ctx) error
		Update(ctx *fiber.Ctx) error
		Delete(ctx *fiber.Ctx) error
	}

	organizerHandler struct {
		s organizer.OrganizerService
		v *validator.Validate
	}
)

func NewOrganizerHandler(service organizer.OrganizerService, validator *validator.Validate) OrganizerHandler {
	return &organizerHandler{s: service, v: validator}
}

func (o *organizerHandler) List(ctx *fiber.Ctx) error {
	search := ctx.Query("search")

	page := ctx.QueryInt("page", 1)
	size := ctx.QueryInt("size", 10)

	// minimal 10, maximal 100
	size = int(math.Min(math.Max(float64(size), 10), 100))
	offset := (page - 1) * size

	sortBy := ctx.Query("sortBy", "id")
	sortDir := ctx.Query("sortBy", "asc")

	organizers, total, err := o.s.ListOrganizer(ctx.UserContext(), search, offset, size, &sortBy, &sortDir)
	if err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve organizers list", err)
	}

	currentPageUrl := pagination.BuildPageURL(ctx, search, page, size, sortBy, sortDir)
	firstPageUrl := pagination.BuildPageURL(ctx, search, 1, size, sortBy, sortDir)

	var nextPageUrl, prevPageUrl *string
	if offset+len(*organizers) < int(total) {
		url := pagination.BuildPageURL(ctx, search, page+1, size, sortBy, sortDir)
		nextPageUrl = &url
	}

	if page > 1 {
		url := pagination.BuildPageURL(ctx, search, page-1, size, sortBy, sortDir)
		prevPageUrl = &url
	}

	paginationRes := pagination.NewResponseMetaData[entities.Organizer](page, currentPageUrl, *organizers, firstPageUrl, nextPageUrl, size, prevPageUrl)

	return presenters.SuccessResponse[pagination.ResponseMetaData[entities.Organizer]](ctx, fiber.StatusOK, "Success retrieve organizers list", &paginationRes)
}

func (o *organizerHandler) Create(ctx *fiber.Ctx) error {
	var cr domain.OrganizerCreateRequest
	if err := ctx.BodyParser(&cr); err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid body request: (name: not null, min=4)", nil)
	}

	if err := o.v.Struct(cr); err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	ot := &entities.Organizer{
		Name:          cr.Name,
		Address:       cr.Address,
		BankName:      cr.BankName,
		AccountNumber: cr.AccountNumber,
		AccountName:   cr.AccountName,
	}

	if err := o.s.CreateOrganizer(ctx.UserContext(), ot); err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "failed to create new organizer", err)
	}

	return presenters.SuccessResponse[entities.Organizer](ctx, fiber.StatusOK, "successfully create new organizer", ot)
}

func (o *organizerHandler) Get(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "invalid id param", err)
	}

	ot, err := o.s.GetOrganizer(ctx.UserContext(), uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "cannot find organizer detail with provided id", nil)
		}
		return presenters.ErrorResponse(ctx, fiber.StatusInternalServerError, "failed get organizer detail", err)
	}

	return presenters.SuccessResponse[entities.Organizer](ctx, fiber.StatusOK, "successfully get organizer detail", ot)
}

func (o *organizerHandler) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "invalid id param", err)
	}

	var ur domain.OrganizerUpdateRequest
	if err := ctx.BodyParser(&ur); err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid body request: (name: min=4)", nil)
	}

	if err := o.v.Struct(ur); err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	ot := &entities.Organizer{
		ID: uint(id),
	}

	if ur.Name != nil {
		ot.Name = *ur.Name
	}

	if ur.Address != nil {
		ot.Address = *ur.Address
	}

	if ur.BankName != nil {
		ot.BankName = *ur.BankName
	}

	if ur.AccountNumber != nil {
		ot.AccountNumber = *ur.AccountNumber
	}

	if ur.AccountName != nil {
		ot.AccountName = *ur.AccountName
	}

	res, err := o.s.UpdateOrganizer(ctx.UserContext(), *ot)
	if err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "failed to update organizer", err)
	}

	return presenters.SuccessResponse[entities.Organizer](ctx, fiber.StatusOK, "successfully update organizer", res)
}

func (o *organizerHandler) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "invalid id param", err)
	}
	if err := o.s.DeleteOrganizer(ctx.UserContext(), uint(id)); err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "failed to delete organizer", err)
	}

	return presenters.SuccessResponse[any](ctx, fiber.StatusOK, "successfully remove organizer", nil)
}
