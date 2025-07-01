package middleware

import (
	"app/internal/config"
	"app/internal/services"
	"app/internal/utils"
	userTokenEntity "app/pkg/database/ent/usertoken"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Auth(userService services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var authHeader string
		authHeader = c.Get("Authorization")
		if authHeader == "" {
			authHeader = c.Cookies("Authorization")
		}
		token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		if token == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Please authenticate")
		}

		payload, err := utils.VerifyToken(token, config.JWTSecret, userTokenEntity.TokenTypeAccessToken)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Please authenticate")
		}

		// user, err := userService.GetUserByID(c, payload.UserID)
		// if err != nil || user == nil {
		// 	return fiber.NewError(fiber.StatusUnauthorized, "Please authenticate")
		// }

		c.Locals("TokenPayload", payload)

		// Không Phân Quyền
		// if len(requiredRights) > 0 {
		// 	userRights, hasRights := config.RoleRights[user.Role]
		// 	if (!hasRights || !hasAllRights(userRights, requiredRights)) && c.Params("userId") != userID {
		// 		return fiber.NewError(fiber.StatusForbidden, "You don't have permission to access this resource")
		// 	}
		// }

		return c.Next()
	}
}

func hasAllRights(userRights, requiredRights []string) bool {
	rightSet := make(map[string]struct{}, len(userRights))
	for _, right := range userRights {
		rightSet[right] = struct{}{}
	}

	for _, right := range requiredRights {
		if _, exists := rightSet[right]; !exists {
			return false
		}
	}
	return true
}
