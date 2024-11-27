package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"ray-d-song.com/echo-rss/model"
	"ray-d-song.com/echo-rss/utils"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login godoc
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param loginForm body LoginForm true "Login credentials"
// @Success 200 {object} LoginRes
// @Failure 400 {object} utils.ErrRes
// @Failure 401 {object} utils.ErrRes
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	form := new(LoginForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError("Invalid request"))
	}

	user := new(model.User)
	if err := user.GetByUsername(form.Username); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.LogError("Invalid username or password"))
	}

	if user.Password != form.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.LogError("Invalid username or password"))
	}

	accessToken, err := model.GenerateAccessToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError("Failed to generate token pair"))
	}

	utils.Logger.Info("login", zap.String("user_id", user.ID))
	return c.JSON(fiber.Map{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"token":    accessToken,
	})
}

type LoginRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
