package v1

import (
	"embed"
	"github.com/swaggo/swag"
)

//go:embed user/*
var swaggerFS embed.FS

func swaggerContent() string {
	bytes, _ := swaggerFS.ReadFile("user/service.swagger.json")
	return string(bytes)
}

var SwaggerInfo = &swag.Spec{
	InfoInstanceName: swag.Name,
	SwaggerTemplate:  swaggerContent(),
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
