package context

import "github.com/labstack/echo/v4"

type RequestContextType string

var RequestContextKey RequestContextType = "request_context"

func DefaultRequestContext(app *AppContext) *RequestContext {
	return &RequestContext{
		AppContext: app,
	}
}

type RequestContext struct {
	*AppContext

	UserContext *UserContext
	ViewContext *ViewContext
}

func (reqContext *RequestContext) SetUserContext(userContext *UserContext) *RequestContext {
	reqContext.UserContext = userContext

	return reqContext
}

func (reqContext *RequestContext) SetViewContext(viewContext *ViewContext) *RequestContext {
	reqContext.ViewContext = viewContext

	return reqContext
}

func RequestContextFromEcho(echoContext echo.Context) *RequestContext {
	requestContextValue, ok := echoContext.Get(string(RequestContextKey)).(*RequestContext)

	if !ok {
		return nil
	}

	return requestContextValue
}

func RequestContextWrap(handler func(echoContext echo.Context, requestContext *RequestContext) error) echo.HandlerFunc {
	return func(echoContext echo.Context) error {
		requestContext := RequestContextFromEcho(echoContext)

		return handler(echoContext, requestContext)
	}
}
