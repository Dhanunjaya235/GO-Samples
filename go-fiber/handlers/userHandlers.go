package handlers

import (
	"encoding/json"
	"fmt"
	"gofiber/apis/database"
	"gofiber/apis/modals"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(ctx *fiber.Ctx) error {

	queryParams := ctx.Queries()
	var (
		keys   []string
		values []interface{}
		query  string
	)

	for key, value := range queryParams {
		keys = append(keys, fmt.Sprintf("%v = ?", key))
		if intval, err := strconv.ParseInt(value, 10, 32); err != nil {
			values = append(values, intval)
			continue
		}
		values = append(values, value)

	}
	if len(keys) == 0 {
		query = "SELECT * FROM users"
	} else {
		query = fmt.Sprintf("SELECT * FROM users WHERE %v", strings.Join(keys, " and "))
	}

	rows, err := database.DB.Query(query, values...)

	if err != nil {
		ctx.Status(500).JSON(&fiber.Map{
			"isSuccess": false,
			"error":     err.Error(),
		})

		return nil
	}

	defer rows.Close()

	users := []modals.User{}

	for rows.Next() {
		user := modals.User{}

		err := rows.Scan(&user.ID, &user.Age, &user.Name)

		if err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"isSuccess": false,
				"error":     err.Error(),
			})
			return nil
		}

		users = append(users, user)
	}

	if err = ctx.JSON(&fiber.Map{
		"isSuccess": true,
		"users":     users,
	}); err != nil {
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return nil
	}
	return nil
}

func GetUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	rows, err := database.DB.Query("SELECT * FROM users WHERE id = ? ", id)

	if err != nil {
		ctx.Status(500).JSON(&fiber.Map{
			"isSuccess": false,
			"error":     err.Error(),
		})

		return nil
	}

	defer rows.Close()

	user := modals.User{}

	for rows.Next() {

		err := rows.Scan(&user.ID, &user.Age, &user.Name)

		if err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"isSuccess": false,
				"error":     err.Error(),
			})
			return nil
		}
	}

	if user.ID == 0 {
		ctx.JSON(&fiber.Map{
			"isSuccess": true,
			"user":      map[string]interface{}{},
		})
		return nil
	}

	if err = ctx.JSON(&fiber.Map{
		"isSuccess": true,
		"user":      user,
	}); err != nil {
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return nil
	}
	return nil
}

func AddNewUser(ctx *fiber.Ctx) error {
	body := ctx.Body()
	var user modals.User
	json.Unmarshal(body, &user)

	response, err := database.DB.Exec("INSERT INTO users (name, age) values (?, ?)", user.Name, user.Age)
	if err != nil {
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return nil
	}

	if id, err := response.LastInsertId(); err == nil {
		user.ID = id
	} else {
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return nil
	}

	ctx.Status(200).JSON(
		&fiber.Map{
			"success": true,
			"user":    user,
		})

	return nil
}

func DeleteUserByID(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	if id == "" {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error":   "ID is required to delete a user",
		})
		return nil
	}

	response, err := database.DB.Exec("DELETE FROM users WHERE id = ?", id)

	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})

		return nil
	}

	if rows, err := response.RowsAffected(); err == nil && rows == 0 {
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   fmt.Sprintf("User With ID %v is not present", id),
		})
		return nil
	} else if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})

		return nil
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"user": &fiber.Map{
			"id": id}})

	return nil
}

func UpdateUserById(ctx *fiber.Ctx) error {

	body := ctx.Body()
	var user map[string]interface{}
	json.Unmarshal(body, &user)
	var updateFields []string
	var values []interface{}
	id := ctx.Params("id")

	for key, value := range user {
		if key != "id" {
			updateFields = append(updateFields, fmt.Sprintf("%s = ?", key))
			values = append(values, value)
		}
	}

	if len(updateFields) == 0 {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error":   "Cannot Perform Update Without Data In Body",
		})
		return nil
	}
	values = append(values, id)

	query := fmt.Sprintf("UPDATE users SET %v WHERE id = ? ", strings.Join(updateFields, ", "))
	response, err := database.DB.Exec(query, values...)
	if err != nil {
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return nil
	}

	if rows, err := response.RowsAffected(); err == nil && rows == 0 {
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   fmt.Sprintf("User With ID %v is not present to update", rows),
		})
		return nil
	} else if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})

		return nil
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"user":    user})

	return nil
}
