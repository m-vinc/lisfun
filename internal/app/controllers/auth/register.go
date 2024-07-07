package auth

import (
	appcontext "lisfun/internal/app/context"
	"lisfun/internal/app/forms"
	"lisfun/internal/app/views/pages/auth"
	"lisfun/internal/services/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (authController *authController) CompleteRegistrationForm(echoContext echo.Context, requestContext *appcontext.RequestContext) error {
	ok := requestContext.UserContext.User.HasCompleteRegistration()

	if ok {
		return echoContext.Redirect(http.StatusFound, "/")
	}

	ok = requestContext.UserContext.User.HasCompleteRegistration()
	if ok {
		return echoContext.Redirect(http.StatusFound, "/")
	}

	return errors.WithStack(
		auth.CompleteRegistration(requestContext, nil, models.UsersCompleteRegistrationParams{}).
			Render(echoContext.Request().Context(), echoContext.Response().Writer),
	)
}

func (authController *authController) CompleteRegistration(echoContext echo.Context, requestContext *appcontext.RequestContext) error {
	ok := requestContext.UserContext.User.HasCompleteRegistration()

	if ok {
		return echoContext.Redirect(http.StatusFound, "/")
	}

	params := models.UsersCompleteRegistrationParams{
		User: requestContext.UserContext.User,
	}
	err := echoContext.Bind(&params)
	if err != nil {
		return err
	}

	_, err = authController.UsersService.CompleteRegistration(echoContext.Request().Context(), &params)
	form, isFormValidationError := forms.FromValidationsErrors(err, "en")
	if isFormValidationError {
		echoContext.Response().WriteHeader(422)
		return errors.WithStack(
			auth.CompleteRegistrationForm(requestContext, form, params).
				Render(echoContext.Request().Context(), echoContext.Response().Writer),
		)
	}

	if err != nil {
		return err
	}

	return nil
}
