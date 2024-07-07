package auth

import (
	"lisfun/internal/app/context"
	"lisfun/internal/services/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

const LisfunSessionKey = "lisfun_auth"

// The goal of this middleware is to authenticate the request using anything we want and set a new
func (authController *authController) Authenticate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoContext echo.Context) error {
			requestContext := context.RequestContextFromEcho(echoContext)

			session, err := authController.store.Get(echoContext.Request(), LisfunSessionKey)
			if err != nil {
				return next(echoContext)
			}

			if session.IsNew {
				return next(echoContext)
			}

			userID, ok := session.Values["user_context"].(string)
			if !ok {
				authController.clearSession(session, echoContext.Request(), echoContext.Response().Writer)
				return next(echoContext)
			}

			userUUID, err := uuid.Parse(userID)
			if err != nil {
				authController.clearSession(session, echoContext.Request(), echoContext.Response().Writer)
				return next(echoContext)
			}

			user, err := authController.UsersService.Get(
				echoContext.Request().Context(),
				&models.UsersGetParams{ID: &userUUID},
			)
			if err != nil {
				authController.clearSession(session, echoContext.Request(), echoContext.Response().Writer)
				return next(echoContext)
			}

			requestContext.SetUserContext(&context.UserContext{ID: userUUID, User: user})
			return next(echoContext)
		}
	}
}

var hasCompleteRegistrationSkipKey = "has_complete_registration_skip"

func (authController *authController) HasCompleteRegistration(skipper func(echo.Context) bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoContext echo.Context) error {
			requestContext := context.RequestContextFromEcho(echoContext)

			skip := skipper != nil && skipper(echoContext)

			if requestContext.UserContext != nil && !skip {
				if requestContext.UserContext.User.HasCompleteRegistration() {
					return next(echoContext)
				}

				return echoContext.Redirect(http.StatusTemporaryRedirect, "/auth/complete_registration")
			}

			return next(echoContext)
		}
	}
}

func (authController *authController) clearSession(session *sessions.Session, req *http.Request, w http.ResponseWriter) error {
	session.Options.MaxAge = -1

	err := session.Save(req, w)
	if err != nil {
		return err
	}

	return nil
}
