package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/middlewares"
	"github.com/handarudwiki/models/commons"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/services"
)

type classroomController struct {
	classroomService services.ClassroomService
	JwtService       services.JWTService
}

func NewClassroom(app *fiber.App, classroomService services.ClassroomService, jwtService services.JWTService) {
	classroomController := &classroomController{
		classroomService: classroomService,
		JwtService:       jwtService,
	}

	classrooms := app.Group("/api/v1/classrooms")
	classrooms.Post("/", middlewares.CheckAuth(classroomController.JwtService), classroomController.Create)
	classrooms.Get("/:id", classroomController.GetSingle)
	classrooms.Get("/", classroomController.GetAll)
	classrooms.Put("/:id", middlewares.CheckAuth(classroomController.JwtService), classroomController.Update)
	classrooms.Delete("/:id", middlewares.CheckAuth(classroomController.JwtService), classroomController.Delete)
}

func (c *classroomController) Create(ctx *fiber.Ctx) (err error) {
	var createClassroomDTO dto.CreateClassroomDTO

	if err = ctx.BodyParser(&createClassroomDTO); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	validationErrors := helpers.ValidateRequest(createClassroomDTO)

	if len(validationErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Validation errors", validationErrors),
		)
	}

	classroomResponse, err := c.classroomService.Create(ctx.Context(), createClassroomDTO)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		helpers.ResponseSuccess(classroomResponse),
	)
}

func (c *classroomController) GetSingle(ctx *fiber.Ctx) (err error) {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	classroomResponse, err := c.classroomService.FindByID(ctx.Context(), id)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(classroomResponse),
	)

}

func (c *classroomController) GetAll(ctx *fiber.Ctx) (err error) {
	var queryDTO dto.QueryDTO

	page, size := helpers.GetPaginationParams(ctx, commons.DEFAULTPAGE, commons.DEFAULTSIZE)

	search := ctx.Query("search")
	queryDTO.Page = page
	queryDTO.Size = size

	queryDTO.Search = &search

	classrooms, paginate, err := c.classroomService.FindAll(ctx.Context(), queryDTO)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponsePagination(classrooms, paginate),
	)
}

func (c *classroomController) Update(ctx *fiber.Ctx) (err error) {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	var updateClassroomDTO dto.UpdateClassroomDTO

	if err = ctx.BodyParser(&updateClassroomDTO); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	validationErrors := helpers.ValidateRequest(updateClassroomDTO)

	if len(validationErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Validation errors", validationErrors),
		)
	}

	classroomResponse, err := c.classroomService.Update(ctx.Context(), updateClassroomDTO, id)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(classroomResponse),
	)
}

func (c *classroomController) Delete(ctx *fiber.Ctx) (err error) {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	err = c.classroomService.Delete(ctx.Context(), id)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess("Deleted successfully"),
	)
}
