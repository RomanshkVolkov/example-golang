// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "email": "jose@guz-studio.dev"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/change-password": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "This endpoint will change the password of authenticated the user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Change password",
                "parameters": [
                    {
                        "description": "New password",
                        "name": "NewPassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ChangePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Operation information",
                        "schema": {
                            "$ref": "#/definitions/domain.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Unhandled error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/forgot-password": {
            "patch": {
                "description": "This endpoint will send an email with the OTP code",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Send an email with the OTP code",
                "parameters": [
                    {
                        "description": "Requires the username to identify the user",
                        "name": "UserIdentity",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.PasswordResetRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Return just a message",
                        "schema": {
                            "$ref": "#/definitions/domain.APIResponse"
                        }
                    },
                    "401": {
                        "description": "Unhandled error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/forgot-password/reset": {
            "patch": {
                "description": "This endpoint will reset the password of the user with the OTP code",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Change password with the OTP code",
                "parameters": [
                    {
                        "description": "New credentials by OTP",
                        "name": "NewCredentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ResetForgottenPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Operation information",
                        "schema": {
                            "$ref": "#/definitions/domain.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Unhandled error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/forgot-password/verify": {
            "post": {
                "description": "Returns data about the code",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Verify the code is valid",
                "parameters": [
                    {
                        "description": "Require the username and the OTP code",
                        "name": "UserIdentity",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ForgottenPasswordCode"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Information about the code",
                        "schema": {
                            "$ref": "#/definitions/domain.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Unhandled error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "Sign in to the application",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Just Sign In",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "UserCredentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful sign in",
                        "schema": {
                            "$ref": "#/definitions/domain.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Unhandled error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Just Sign Up",
                "parameters": [
                    {
                        "description": "Just the user data",
                        "name": "UserData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.NewUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Return message",
                        "schema": {
                            "$ref": "#/definitions/domain.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Unhandled error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/mail/test": {
            "post": {
                "security": [
                    {
                        "none": []
                    }
                ],
                "description": "This endpoint send a test email",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mail"
                ],
                "summary": "Test email sending",
                "responses": {
                    "200": {
                        "description": "Operation information",
                        "schema": {
                            "$ref": "#/definitions/domain.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Unhandled error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/profile": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get user profile by token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Just User Profile by token",
                "responses": {
                    "200": {
                        "description": "User profile",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Unhandled error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error (report it)",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.APIResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "message": {
                    "$ref": "#/definitions/domain.Message"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "domain.ChangePassword": {
            "type": "object",
            "required": [
                "confirmPassword",
                "currentPassword",
                "password"
            ],
            "properties": {
                "confirmPassword": {
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 6
                },
                "currentPassword": {
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 6
                },
                "password": {
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 6
                }
            }
        },
        "domain.ForgottenPasswordCode": {
            "type": "object",
            "required": [
                "otp",
                "username"
            ],
            "properties": {
                "otp": {
                    "type": "string",
                    "maxLength": 5,
                    "minLength": 5
                },
                "username": {
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 6
                }
            }
        },
        "domain.Message": {
            "type": "object",
            "properties": {
                "en": {
                    "type": "string"
                },
                "es": {
                    "type": "string"
                }
            }
        },
        "domain.NewUser": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "role",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 300
                },
                "name": {
                    "type": "string",
                    "maxLength": 300,
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "role": {
                    "$ref": "#/definitions/domain.Role"
                },
                "username": {
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 6
                }
            }
        },
        "domain.PasswordResetRequest": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.ResetForgottenPassword": {
            "type": "object",
            "required": [
                "confirmPassword",
                "otp",
                "password",
                "username"
            ],
            "properties": {
                "confirmPassword": {
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 6
                },
                "otp": {
                    "type": "string",
                    "maxLength": 5,
                    "minLength": 5
                },
                "password": {
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 6
                }
            }
        },
        "domain.Role": {
            "type": "string",
            "enum": [
                "root",
                "admin",
                "customer",
                "employer"
            ],
            "x-enum-varnames": [
                "Root",
                "Admin",
                "Customer",
                "Employer"
            ]
        },
        "domain.SignInRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "3.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "GO API",
	Description:      "Created by @RomanshkVolkov.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
