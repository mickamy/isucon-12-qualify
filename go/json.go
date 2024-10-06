package isuports

import (
	"github.com/labstack/echo/v4"
	"github.com/sugawarayuuta/sonnet"
)

type JSONSerializer struct{}

func NewJSONSerializer() *JSONSerializer {
	return &JSONSerializer{}
}

func (js *JSONSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := sonnet.NewEncoder(c.Response())
	return enc.Encode(i)
}

func (js *JSONSerializer) Deserialize(c echo.Context, i interface{}) error {
	dec := sonnet.NewDecoder(c.Request().Body)
	return dec.Decode(i)
}
