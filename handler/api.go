package handler

import (
	"student/database"
	"student/models"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello fiber!", "data": nil})
}

func GetAllStudent(ctx *fiber.Ctx) error {
	var students []models.Student

	db := database.DBConn

	db.Find(&students)

	return ctx.JSON(fiber.Map{
		"status":  "ok",
		"message": "GetALL successfully",
		"data":    students})
}

func GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := database.DBConn
	var student models.Student
	db.Find(&student, id)
	if student.Fname == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "No Student found with ID",
			"data":    nil})
	}
	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "User found",
		"data":    student})

}

func CreateStudent(ctx *fiber.Ctx) error {
	db := database.DBConn

	student := new(models.Student)
	if err := ctx.BodyParser(student); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request data",
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

func UpdateStudent(ctx *fiber.Ctx) error {
	db := database.DBConn
	id := ctx.Params("id")

	updatedStudent := new(models.Student)
	if err := ctx.BodyParser(updatedStudent); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request data",
			"data":    nil,
		})
	}

	var existingStudent models.Student
	if err := db.First(&existingStudent, id).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Student not found",
			"data":    nil,
		})
	}

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

func DeleteStudent(ctx *fiber.Ctx) error {
	db := database.DBConn
	id := ctx.Params("id")

	var existingStudent models.Student
	if err := db.First(&existingStudent, id).Error; err != nil {
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
