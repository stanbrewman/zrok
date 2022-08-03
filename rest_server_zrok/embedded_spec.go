// Code generated by go-swagger; DO NOT EDIT.

package rest_server_zrok

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/zrok.v1+json"
  ],
  "produces": [
    "application/zrok.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "zrok client access",
    "title": "zrok",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/account": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "createAccount",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/accountRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "account created",
            "schema": {
              "$ref": "#/definitions/accountResponse"
            }
          },
          "400": {
            "description": "account not created (already exists)",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/enable": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "identity"
        ],
        "operationId": "enable",
        "responses": {
          "201": {
            "description": "environment enabled",
            "schema": {
              "$ref": "#/definitions/enableResponse"
            }
          },
          "401": {
            "description": "invalid api key"
          },
          "404": {
            "description": "account not found"
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/listEnvironments": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "metadata"
        ],
        "operationId": "listEnvironments",
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/environments"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/login": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "login",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/loginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "login successful",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "401": {
            "description": "invalid login"
          }
        }
      }
    },
    "/tunnel": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "tunnel"
        ],
        "operationId": "tunnel",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/tunnelRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "tunnel created",
            "schema": {
              "$ref": "#/definitions/tunnelResponse"
            }
          },
          "401": {
            "description": "invalid environment identity",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/untunnel": {
      "delete": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "tunnel"
        ],
        "operationId": "untunnel",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/untunnelRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "tunnel removed"
          },
          "404": {
            "description": "not found",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/version": {
      "get": {
        "tags": [
          "metadata"
        ],
        "operationId": "version",
        "responses": {
          "200": {
            "description": "retrieve the current server version",
            "schema": {
              "$ref": "#/definitions/version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "accountRequest": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "accountResponse": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "enableResponse": {
      "type": "object",
      "properties": {
        "cfg": {
          "type": "string"
        },
        "identity": {
          "type": "string"
        }
      }
    },
    "environment": {
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        },
        "createdAt": {
          "type": "string"
        },
        "updatedAt": {
          "type": "string"
        },
        "zitiId": {
          "type": "string"
        }
      }
    },
    "environments": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/environment"
      }
    },
    "errorMessage": {
      "type": "string"
    },
    "loginRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "loginResponse": {
      "type": "string"
    },
    "principal": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer"
        },
        "token": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "tunnelRequest": {
      "type": "object",
      "properties": {
        "endpoint": {
          "type": "string"
        },
        "identity": {
          "type": "string"
        }
      }
    },
    "tunnelResponse": {
      "type": "object",
      "properties": {
        "service": {
          "type": "string"
        }
      }
    },
    "untunnelRequest": {
      "type": "object",
      "properties": {
        "service": {
          "type": "string"
        }
      }
    },
    "version": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "apiKey",
      "name": "x-token",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/zrok.v1+json"
  ],
  "produces": [
    "application/zrok.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "zrok client access",
    "title": "zrok",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/account": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "createAccount",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/accountRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "account created",
            "schema": {
              "$ref": "#/definitions/accountResponse"
            }
          },
          "400": {
            "description": "account not created (already exists)",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/enable": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "identity"
        ],
        "operationId": "enable",
        "responses": {
          "201": {
            "description": "environment enabled",
            "schema": {
              "$ref": "#/definitions/enableResponse"
            }
          },
          "401": {
            "description": "invalid api key"
          },
          "404": {
            "description": "account not found"
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/listEnvironments": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "metadata"
        ],
        "operationId": "listEnvironments",
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/environments"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/login": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "login",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/loginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "login successful",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "401": {
            "description": "invalid login"
          }
        }
      }
    },
    "/tunnel": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "tunnel"
        ],
        "operationId": "tunnel",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/tunnelRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "tunnel created",
            "schema": {
              "$ref": "#/definitions/tunnelResponse"
            }
          },
          "401": {
            "description": "invalid environment identity",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/untunnel": {
      "delete": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "tunnel"
        ],
        "operationId": "untunnel",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/untunnelRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "tunnel removed"
          },
          "404": {
            "description": "not found",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/version": {
      "get": {
        "tags": [
          "metadata"
        ],
        "operationId": "version",
        "responses": {
          "200": {
            "description": "retrieve the current server version",
            "schema": {
              "$ref": "#/definitions/version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "accountRequest": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "accountResponse": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "enableResponse": {
      "type": "object",
      "properties": {
        "cfg": {
          "type": "string"
        },
        "identity": {
          "type": "string"
        }
      }
    },
    "environment": {
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        },
        "createdAt": {
          "type": "string"
        },
        "updatedAt": {
          "type": "string"
        },
        "zitiId": {
          "type": "string"
        }
      }
    },
    "environments": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/environment"
      }
    },
    "errorMessage": {
      "type": "string"
    },
    "loginRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "loginResponse": {
      "type": "string"
    },
    "principal": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer"
        },
        "token": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "tunnelRequest": {
      "type": "object",
      "properties": {
        "endpoint": {
          "type": "string"
        },
        "identity": {
          "type": "string"
        }
      }
    },
    "tunnelResponse": {
      "type": "object",
      "properties": {
        "service": {
          "type": "string"
        }
      }
    },
    "untunnelRequest": {
      "type": "object",
      "properties": {
        "service": {
          "type": "string"
        }
      }
    },
    "version": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "apiKey",
      "name": "x-token",
      "in": "header"
    }
  }
}`))
}