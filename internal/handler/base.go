package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func statusBadRequestMap(errno int, message string) map[string]any {
	return map[string]any{
		"errno":   errno,
		"message": message,
	}
}

func render(c *gin.Context, v any, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, statusBadRequestMap(999, err.Error()))
		return
	}
	data := map[string]any{
		"errno":   0,
		"message": "success",
	}
	if ft, ok := c.Get("format"); ok {
		data[ft.(string)] = v
	} else {
		data["data"] = v
	}
	c.JSON(http.StatusOK, data)
}

func NoParam[R any](handlerFunc func() (R, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		v, err := handlerFunc()
		render(c, v, err)
	}
}

type (
	THandlerFunc[T any]         func(ctx *gin.Context, t *T) error
	TRHandlerFunc[T any, R any] func(ctx *gin.Context, t *T) (R, error)
	GHandlerFunc                func(ctx *gin.Context) error

	TRPathParamHandlerFunc[T any, R any] func(ctx *gin.Context, t *T) (R, error)
	TPathParamHandlerFunc[T any]         func(ctx *gin.Context, t *T) error

	PathParamHandlerFunc func(ctx *gin.Context, params gin.Params) error
)

func THandler[T any](
	handler THandlerFunc[T],
) gin.HandlerFunc {
	return handlerWithContext[T, any](handler)
}

func TRHandler[T any, R any](
	handler TRHandlerFunc[T, R],
) gin.HandlerFunc {
	return handlerWithContext[T, R](handler)
}

func Handler(
	handler GHandlerFunc,
) gin.HandlerFunc {
	return handlerWithContext[any, any](handler)
}

func TRPathParamHandler[T any, R any](
	handler TRPathParamHandlerFunc[T, R],
) gin.HandlerFunc {
	return handlerWithContext[T, R](handler)
}

func TPathParamHandler[T any](
	handler TPathParamHandlerFunc[T],
) gin.HandlerFunc {
	return handlerWithContext[T, any](handler)
}

func PathParamHandler(
	handler PathParamHandlerFunc,
) gin.HandlerFunc {
	return handlerWithContext[any, any](handler)
}

func handlerWithContext[T any, R any](
	handlerFunc any,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := new(T)
		if err := c.ShouldBind(&t); err != nil {
			c.JSON(http.StatusBadRequest, statusBadRequestMap(999, "请求格式错误"))
			return
		}

		switch handler := handlerFunc.(type) {
		case THandlerFunc[T]:
			render(c, nil, handler(c, t))
		case TRHandlerFunc[T, R]:
			v, err := handler(c, t)
			render(c, v, err)
		case GHandlerFunc:
			err := handler(c)
			render(c, nil, err)
		case TRPathParamHandlerFunc[T, R]:
			v, err := handler(c, t)
			render(c, v, err)
		case TPathParamHandlerFunc[T]:
			render(c, nil, handler(c, t))
		case PathParamHandlerFunc:
			render(c, nil, handler(c, c.Params))
		default:
			panic("unsupported handler")
		}
	}
}
