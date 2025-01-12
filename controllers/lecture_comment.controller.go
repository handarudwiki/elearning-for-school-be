package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/middlewares"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/services"
)

type lectureCommentController struct {
	lectureCommentService services.LectureCommentService
	jwtService            services.JWTService
}

func NewLectureCommentController(app *fiber.App, lectureCommentService services.LectureCommentService, jwtService services.JWTService) *lectureCommentController {
	lectureCommentController := lectureCommentController{
		lectureCommentService: lectureCommentService,
		jwtService:            jwtService,
	}

	lectureComments := app.Group("/api/v1/lecture-comments")
	lectureComments.Post("/", middlewares.CheckAuth(jwtService), lectureCommentController.Create)

	return &lectureCommentController
}

func (c *lectureCommentController) Create(ctx *fiber.Ctx) (err error) {
	var createLectureCommentDTO dto.CreateLectureCommentDTO

	if err = ctx.BodyParser(&createLectureCommentDTO); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	validationErrors := helpers.ValidateRequest(createLectureCommentDTO)

	if len(validationErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Validation errors", validationErrors),
		)
	}

	createLectureCommentDTO.UserID = uint(ctx.Locals("userId").(int))

	lectureCommentResponse, err := c.lectureCommentService.Create(ctx.Context(), &createLectureCommentDTO)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		helpers.ResponseErrorWithData("Success create lecture comment", lectureCommentResponse),
	)
}
