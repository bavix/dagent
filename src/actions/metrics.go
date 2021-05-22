package actions

import (
	"github.com/bavix/dagent/src/dto"
	"github.com/bavix/dagent/src/store"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Metrics struct {
	Store store.Store
}

type formObj struct {
	Data []dto.MetricDto `json:"data" validate:"required"`
}

func (a *Metrics) Get(c echo.Context) error {
	actSp := jaegertracing.CreateChildSpan(c, "metrics.Get")
	defer actSp.Finish()

	values := jaegertracing.TraceFunction(c, a.Store.ReadAll)
	result := values[0].Interface().([]string)

	return c.String(http.StatusOK, strings.Join(result, "\n"))
}

func (a *Metrics) Post(c echo.Context) error {
	actSp := jaegertracing.CreateChildSpan(c, "metrics.Post")
	defer actSp.Finish()

	m := new(formObj)
	if err := c.Bind(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for _, s := range m.Data {
		if err := c.Validate(s); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	for _, s := range m.Data {
		jaegertracing.TraceFunction(c,
			a.Store.Set,
			s.UniqueId(),
			s.ToString(),
			s.Duration)
	}

	return c.JSON(http.StatusOK, struct {
		Success bool `json:"success"`
	}{
		Success: len(m.Data) > 0,
	})
}
