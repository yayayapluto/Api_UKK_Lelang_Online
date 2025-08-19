package handlers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/yayayapluto/api-ukk-online/domain"
	"github.com/yayayapluto/api-ukk-online/entities"
	"github.com/yayayapluto/api-ukk-online/internal/api/presenters"
	"github.com/yayayapluto/api-ukk-online/internal/utils/pagination"
	objectType "github.com/yayayapluto/api-ukk-online/pkg/object-type"
	"gorm.io/gorm"
	"math"
)

type (
	ObjectTypeHandler interface {
		List(ctx *fiber.Ctx) error
		Create(ctx *fiber.Ctx) error
		Get(ctx *fiber.Ctx) error
		Update(ctx *fiber.Ctx) error
		Delete(ctx *fiber.Ctx) error
	}

	objectTypeHandler struct {
		s objectType.ObjectTypeService
		v *validator.Validate
	}
)

func NewObjectTypeHandler(service objectType.ObjectTypeService, validator *validator.Validate) ObjectTypeHandler {
	return &objectTypeHandler{s: service, v: validator}
}

func (o *objectTypeHandler) List(ctx *fiber.Ctx) error {
	search := ctx.Query("search")

	page := ctx.QueryInt("page", 1)
	size := ctx.QueryInt("size", 10)

	// minimal 10, maximal 100
	size = int(math.Min(math.Max(float64(size), 10), 100))
	offset := (page - 1) * size

	sortBy := ctx.Query("sortBy", "id")
	sortDir := ctx.Query("sortBy", "asc")

	objectTypes, total, err := o.s.ListObjectType(ctx.UserContext(), search, offset, size, &sortBy, &sortDir)
	if err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve object types list", err)
	}

	currentPageUrl := pagination.BuildPageURL(ctx, search, page, size, sortBy, sortDir)
	firstPageUrl := pagination.BuildPageURL(ctx, search, 1, size, sortBy, sortDir)

	var nextPageUrl, prevPageUrl *string
	if offset+len(*objectTypes) < int(total) {
		url := pagination.BuildPageURL(ctx, search, page+1, size, sortBy, sortDir)
		nextPageUrl = &url
	}

	if page > 1 {
		url := pagination.BuildPageURL(ctx, search, page-1, size, sortBy, sortDir)
		prevPageUrl = &url
	}

	paginationRes := pagination.NewResponseMetaData[entities.ObjectType](page, currentPageUrl, *objectTypes, firstPageUrl, nextPageUrl, size, prevPageUrl)

	return presenters.SuccessResponse[pagination.ResponseMetaData[entities.ObjectType]](ctx, fiber.StatusOK, "Success retrieve object types list", &paginationRes)
}

func (o *objectTypeHandler) Create(ctx *fiber.Ctx) error {
	var cr domain.ObjectTypeCreateRequest
	if err := ctx.BodyParser(&cr); err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid body request: (name: not null, min=4)", nil)
	}

	if err := o.v.Struct(cr); err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	ot := &entities.ObjectType{
		Name: cr.Name,
	}

	if err := o.s.CreateObjectType(ctx.UserContext(), ot); err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "failed to create new object type", err)
	}

	return presenters.SuccessResponse[entities.ObjectType](ctx, fiber.StatusOK, "successfully create new object type", ot)
}

func (o *objectTypeHandler) Get(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "invalid id param", err)
	}

	ot, err := o.s.GetObjectType(ctx.UserContext(), uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "cannot find object type detail with provided id", nil)
		}
		return presenters.ErrorResponse(ctx, fiber.StatusInternalServerError, "failed get object type detail", err)
	}

	return presenters.SuccessResponse[entities.ObjectType](ctx, fiber.StatusOK, "successfully get object type detail", ot)
}

func (o *objectTypeHandler) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "invalid id param", err)
	}

	var ur domain.ObjectTypeUpdateRequest
	if err := ctx.BodyParser(&ur); err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid body request: (name: min=4)", nil)
	}

	if err := o.v.Struct(ur); err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	ot := &entities.ObjectType{
		ID:   uint(id),
		Name: *ur.Name,
	}

	res, err := o.s.UpdateObjectType(ctx.UserContext(), *ot)
	if err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "failed to update object type", err)
	}

	return presenters.SuccessResponse[entities.ObjectType](ctx, fiber.StatusOK, "successfully update object type", res)
}

func (o *objectTypeHandler) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "invalid id param", err)
	}
	if err := o.s.DeleteObjectType(ctx.UserContext(), uint(id)); err != nil {
		return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, "failed to delete object type", err)
	}

	return presenters.SuccessResponse[any](ctx, fiber.StatusOK, "successfully remove object type", nil)
}
