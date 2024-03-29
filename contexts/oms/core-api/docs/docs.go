// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/orders/": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "body",
                        "name": "c",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.CreateOrderCommand"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/orders.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/orders/orderNumber/{orderNumber}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order Number",
                        "name": "orderNumber",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/orders/{orderNumber}/reject-payment": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order Number",
                        "name": "orderNumber",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "c",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.RejectOrderPaymentCommand"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "order.CreateOrderCommand": {
            "type": "object",
            "properties": {
                "orderLines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/order.OrderLine"
                    }
                },
                "orderNumber": {
                    "type": "string"
                },
                "price": {
                    "$ref": "#/definitions/shared.Money"
                }
            }
        },
        "order.OrderLine": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "price": {
                    "$ref": "#/definitions/shared.Money"
                },
                "sku": {
                    "type": "string"
                }
            }
        },
        "order.RejectOrderPaymentCommand": {
            "type": "object",
            "properties": {
                "orderNumber": {
                    "type": "string"
                },
                "rejectReason": {
                    "type": "string"
                }
            }
        },
        "orders.Order": {
            "type": "object",
            "properties": {
                "finalPrice": {
                    "$ref": "#/definitions/shared.Money"
                },
                "id": {
                    "type": "string"
                },
                "orderNumber": {
                    "type": "string"
                },
                "paymentRejectReason": {
                    "type": "string"
                },
                "paymentStatus": {
                    "type": "string"
                },
                "price": {
                    "$ref": "#/definitions/shared.Money"
                }
            }
        },
        "shared.Money": {
            "type": "object",
            "properties": {
                "currencyCode": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func Initialize() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
