package handler

import (
	"fiber-joglo-dev/database"
	"fiber-joglo-dev/models/entity"
	"fiber-joglo-dev/models/request"
	"fiber-joglo-dev/models/response"
	"fiber-joglo-dev/utils"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(c *fiber.Ctx) error {
	var users []entity.User
	err := database.DB.Debug().Find(&users).Error
	if err != nil {
		c.Status(http.StatusNotFound)
		return c.JSON(err)
	}

	return c.JSON(users)
}

func UserHandlerCreate(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}
	newUser := entity.User{
		Nama:    user.Nama,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}
	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	newUser.Password = hashedPassword
	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed To Store",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Succesfuly",
		"data":    newUser,
	})
}

func UserHandlerGetById(c *fiber.Ctx) error {
	userId := c.Params("nik")
	var user entity.User
	err := database.DB.First(&user, "nik = ?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}
	userResponse := response.UserResponse{
		NIK:       user.NIK,
		Nama:      user.Nama,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return c.JSON(fiber.Map{
		"message": "Succesfuly",
		"data":    userResponse,
	})
}

func UserHandlerUpdate(c *fiber.Ctx) error {
	var user entity.User
	userRequest := new(request.UserUpdateRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	// Check Avaiable User
	userId := c.Params("nik")
	err := database.DB.First(&user, "nik = ?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}

	// Update User
	if userRequest.Nama != "" {
		user.Nama = userRequest.Nama
	}
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone
	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Succesfuly",
		"data":    user,
	})
}

func UserHandlerDelete(c *fiber.Ctx) error {
	var user entity.User
	// Check Avaiable User
	userId := c.Params("nik")
	err := database.DB.First(&user, "nik = ?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}
	// Delete User
	errDelete := database.DB.Delete(&user).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(fiber.Map{
		"message": "User Was Deleted",
	})
}
