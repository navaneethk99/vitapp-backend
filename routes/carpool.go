package routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"vitapp-backend/database"
	"vitapp-backend/models"
)

type Carpool struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
}

func CreateResponsePool(poolModel models.Carpool) User {
	return Carpool{
		ID:        poolModel.ID,
		FirstName: poolModel.FirstName,
		LastName:  poolModel.LastName,
	}
}

func CreatePool(c *fiber.Ctx) error {
	var pool models.Carpool
	if err := c.BodyParser(&pool); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&pool)
	responsePool := CreateResponsePool(pool)
	return c.Status(200).JSON(responsePool)
}

func GetPools(c *fiber.Ctx) error {
	pools := []models.Carpool{}
	database.Database.Db.Find(&pools)
	responsePools := []Carpool{}
	for _, pool := range pools {
		responsePool := CreateResponsePool(pool)
		responsePools = append(responsePools, responsePool)
	}
	return c.Status(200).JSON(responsePools)
}

//
//func FindUser(id int, user *models.User) error {
//	database.Database.Db.Find(&user, "id = ?", id)
//	if user.ID == 0 {
//		return errors.New("User not found")
//	}
//	return nil
//}
//
//func GetUser(c *fiber.Ctx) error {
//	id, err := c.ParamsInt("id")
//	var user models.User
//
//	if err != nil {
//		return c.Status(400).JSON(err.Error())
//	}
//
//	if err := FindUser(id, &user); err != nil {
//		return c.Status(400).JSON(err.Error())
//	}
//
//	responseUser := CreateResponseUser(user)
//	return c.Status(200).JSON(responseUser)
//}
//
//func UpdateUser(c *fiber.Ctx) error {
//	id, err := c.ParamsInt("id")
//	var user models.User
//
//	if err != nil {
//		return c.Status(400).JSON(err.Error())
//	}
//
//	if err := FindUser(id, &user); err != nil {
//		return c.Status(400).JSON(err.Error())
//	}
//
//	type UpdateUser struct {
//		FirstName string `json:"first_name"`
//		LastName  string `json:"last_name"`
//	}
//
//	var updateData UpdateUser
//
//	if err := c.BodyParser(&updateData); err != nil {
//		return c.Status(500).JSON(err.Error())
//	}
//
//	user.FirstName = updateData.FirstName
//	user.LastName = updateData.LastName
//	database.Database.Db.Save(&user)
//
//	responseUser := CreateResponseUser(user)
//	return c.Status(200).JSON(responseUser)
//}
//
//func DeleteUser(c *fiber.Ctx) error {
//	id, err := c.ParamsInt("id")
//	var user models.User
//
//	if err != nil {
//		return c.Status(400).JSON(err.Error())
//	}
//
//	if err := FindUser(id, &user); err != nil {
//		return c.Status(400).JSON(err.Error())
//	}
//
//	if err := database.Database.Db.Delete(&user).Error; err != nil {
//		return c.Status(404).JSON(err.Error())
//	}
//
//	return c.Status(200).JSON("User deleted")
//}
