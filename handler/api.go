package studenthandler

import (
	"student/database"
	"student/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type IStudentHandler interface {
	GetAllStudent(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	CreateStudent(ctx *fiber.Ctx) error
	UpdateStudent(ctx *fiber.Ctx) error
	DeleteStudent(ctx *fiber.Ctx) error
}

type StudentHandler struct{}

func NewStudentHandler() IStudentHandler {
	return &StudentHandler{}
}

func (s *StudentHandler) GetAllStudent(ctx *fiber.Ctx) error {
	var students []models.Student
	db := database.DBConn
	db.Find(&students)
	return ctx.JSON(fiber.Map{
		"status":  "ok",
		"message": "Get all students successfully",
		"data":    students,
	})
}

func (s *StudentHandler) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("student_id")
	db := database.DBConn
	var student models.Student
	if err := db.Where("student_id = ?", id).First(&student).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No Student found with ID",
			"data":    nil,
		})
	}
	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Student found",
		"data":    student,
	})
}

func (s StudentHandler) CreateStudent(ctx *fiber.Ctx) error {
	db := database.DBConn
	student := new(models.Student)
	if err := ctx.BodyParser(student); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request data",
			"data":    nil,
		})
	}
	if student.Fname == "" || student.Lname == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "First and Last name are required",
			"data":    nil,
		})
	}
	var existingStudent models.Student
	if err := db.Where("fname = ? AND lname = ?", student.Fname, student.Lname).First(&existingStudent).Error; err == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "Student already exists",
			"data":    nil,
		})
	}
	if err := db.Where("student_id = ?", student.StudentID).First(&existingStudent).Error; err == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "Student ID already exists",
			"data":    nil,
		})
	}
	student.ID = uuid.New().String()
	student.CreatedDate = time.Now()
	if err := db.Create(&student).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create student",
			"data":    nil,
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "ok",
		"message": "Student created successfully",
		"data":    student,
	})
}

func (s *StudentHandler) UpdateStudent(ctx *fiber.Ctx) error {
	db := database.DBConn
	id := ctx.Params("student_id")
	updatedStudent := new(models.Student)
	if err := ctx.BodyParser(updatedStudent); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request data",
			"data":    nil,
		})
	}
	var existingStudent models.Student
	if err := db.Where("student_id = ?", id).First(&existingStudent).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Student not found",
			"data":    nil,
		})
	}
	updatedStudent.UpdatedDate = time.Now()
	if err := db.Model(&existingStudent).Updates(updatedStudent).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to update student",
			"data":    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "Student updated successfully",
		"data":    existingStudent,
	})
}

func (s *StudentHandler) DeleteStudent(ctx *fiber.Ctx) error {
	db := database.DBConn
	id := ctx.Params("student_id")
	var existingStudent models.Student
	if err := db.Where("student_id = ?", id).First(&existingStudent).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Student not found",
			"data":    nil,
		})
	}
	if err := db.Delete(&existingStudent).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete student",
			"data":    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "Student deleted successfully",
		"data":    nil,
	})
}
