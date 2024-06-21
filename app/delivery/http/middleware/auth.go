package middleware

import (
	"WeekendPOS/app/model"
	"WeekendPOS/app/service"
	"github.com/gofiber/fiber/v2"
)

func NewAuth(userService *service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &model.VerifyUserRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		userService.Log.Debugf("Authorization : %s", request.Token)

		auth, err := userService.Verify(ctx.UserContext(), request)
		if err != nil {
			userService.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		userService.Log.Debugf("User : %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
