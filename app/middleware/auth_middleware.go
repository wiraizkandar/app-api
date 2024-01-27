package middleware

import (
	"errors"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/officemaid/app-api/app/data/response"
)

type AccessTokenClaims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var (
	ErrTokenExpiredMesage = "Token is expired"
	ErrUnauthorizedMesage = "Unauthorized"
)

/**	AuthMiddleware
 * @param c *fiber.Ctx
 * @return error
 */
func AuthMiddleware(c *fiber.Ctx) error {

	authorizationToken := getAuthorizationToken(c)

	if authorizationToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ApiResponse{
			Status:  response.ErrorStatus,
			Message: "Unauthorized",
		})
	}

	// parse token
	_, err := jwt.ParseWithClaims(authorizationToken, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ApiResponse{
			Status:  response.ErrorStatus,
			Message: getErrorMessage(err),
		})
	}

	return c.Next()
}

/** getAuthorizationToken
 * @param ctx *fiber.Ctx
 * @return string
 */
func getAuthorizationToken(ctx *fiber.Ctx) string {

	authorizationToken := string(ctx.Request().Header.Peek("Authorization"))
	authorizationToken = strings.Replace(authorizationToken, "Bearer ", "", 1)

	return authorizationToken
}

/** getErrorMessage
 * @param err error
 * @return string
 */
func getErrorMessage(err error) string {

	var message string

	switch {
	case errors.Is(err, jwt.ErrTokenExpired):
		message = ErrTokenExpiredMesage
	default:
		message = ErrUnauthorizedMesage
	}

	return message
}
