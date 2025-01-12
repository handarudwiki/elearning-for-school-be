package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/middlewares"
	"github.com/handarudwiki/models/commons"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/services"
)

type userController struct {
	userService services.UserService
}

func NewUser(app *fiber.App, userService services.UserService) {
	controller := userController{
		userService: userService,
	}

	api := app.Group("/api/v1")

	api.Post("/login", controller.Login)

	users := api.Group("/users")

	users.Get("/me", middlewares.CheckAuth(userService.GetJwtService()), controller.Me)
	users.Post("/teacher", middlewares.CheckAuth(userService.GetJwtService()), controller.CreateTeacher)
	users.Get("/teacher", middlewares.CheckAuth(userService.GetJwtService()), controller.GetAllTeacher)
	users.Get("/student", middlewares.CheckAuth(userService.GetJwtService()), controller.GetAllStudent)
	users.Get("/single/:id", middlewares.CheckAuth(userService.GetJwtService()), controller.GetSingle)
	users.Delete("/:id", middlewares.CheckAuth(userService.GetJwtService()), controller.Delete)
	users.Put("/:id", middlewares.CheckAuth(userService.GetJwtService()), controller.Update)

}

func (c *userController) Login(ctx *fiber.Ctx) (err error) {
	var loginRequest dto.LoginDTO
	if err = ctx.BodyParser(&loginRequest); err != nil {

		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErrors := helpers.ValidateRequest(loginRequest)
	if len(validationErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation errors",
			"errors":  validationErrors,
		})
	}

	loginResponse, err := c.userService.Login(ctx.Context(), loginRequest)
	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		loginResponse,
	)
}

func (c *userController) Me(ctx *fiber.Ctx) (err error) {
	id := ctx.Locals("userId").(int)
	fmt.Print(id)
	userResponse, err := c.userService.Me(ctx.Context(), id)
	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(userResponse),
	)
}

func (c *userController) CreateTeacher(ctx *fiber.Ctx) (err error) {
	var createUserDto dto.CreateUserDTO
	if err = ctx.BodyParser(&createUserDto); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	validationErrors := helpers.ValidateRequest(createUserDto)
	if len(validationErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Validation errors", validationErrors),
		)
	}

	userResponse, err := c.userService.CreateTeacher(ctx.Context(), createUserDto)
	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		helpers.ResponseSuccess(userResponse),
	)
}

func (c *userController) GetAllStudent(ctx *fiber.Ctx) (err error) {
	var queryDto dto.QueryDTO

	page, size := helpers.GetPaginationParams(ctx, commons.DEFAULTPAGE, commons.DEFAULTSIZE)
	fmt.Println(page, size)
	queryDto.Page = page
	queryDto.Size = size

	search := ctx.Query("search")

	queryDto.Search = &search

	isActive := ctx.Query("is_active")

	if isActive != "" {
		isActiveBool, err := strconv.ParseBool(isActive)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(
				helpers.ResponseError(err.Error()),
			)
		}
		queryDto.Is_active = &isActiveBool

	}

	users, paginate, err := c.userService.GetAllStudent(ctx.Context(), queryDto)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponsePagination(users, paginate),
	)

}

func (c *userController) GetAllTeacher(ctx *fiber.Ctx) (err error) {
	var queryDto dto.QueryDTO

	page, size := helpers.GetPaginationParams(ctx, commons.DEFAULTPAGE, commons.DEFAULTSIZE)

	queryDto.Page = page
	queryDto.Size = size

	search := ctx.Query("search")

	queryDto.Search = &search

	isActive := ctx.Query("is_active")

	if isActive != "" {
		isActiveBool, err := strconv.ParseBool(isActive)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(
				helpers.ResponseError(err.Error()),
			)
		}
		queryDto.Is_active = &isActiveBool

	}
	users, paginate, err := c.userService.GetAllTeacher(ctx.Context(), queryDto)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponsePagination(users, paginate),
	)

}

func (c *userController) GetSingle(ctx *fiber.Ctx) (err error) {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	userResponse, err := c.userService.GetUser(ctx.Context(), id)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(userResponse),
	)
}

func (c *userController) Delete(ctx *fiber.Ctx) (err error) {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	err = c.userService.Delete(ctx.Context(), id)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess("User deleted"),
	)
}

func (c *userController) Update(ctx *fiber.Ctx) (err error) {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	var updateUserDto dto.UpdateUserDTO
	if err = ctx.BodyParser(&updateUserDto); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	validationErrors := helpers.ValidateRequest(updateUserDto)
	if len(validationErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Validation errors", validationErrors),
		)
	}

	userResponse, err := c.userService.Update(ctx.Context(), id, updateUserDto)
	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(userResponse),
	)
}
