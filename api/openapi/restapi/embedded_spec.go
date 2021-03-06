// Code generated by go-swagger; DO NOT EDIT.

package restapi

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
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "# ...\n## List of all custom errors\nFirst number is HTTP Status code, second is value of \"code\" field in returned JSON object, text description may or may not match \"message\" field in returned JSON object.\n- 409.1000: contact already exists\n",
    "title": "paint",
    "version": "0.2.0"
  },
  "basePath": "/",
  "paths": {
    "/health-check": {
      "get": {
        "security": [],
        "description": "Returns 200 if service works okay.",
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "Extra details about service status, if any.",
            "schema": {
              "type": "object",
              "additionalProperties": true
            }
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    },
    "/v1/pyrMeanShiftFilter": {
      "get": {
        "security": [],
        "description": "pyrMeanShiftFiltering",
        "operationId": "pyrMeanShiftFilter",
        "parameters": [
          {
            "type": "string",
            "name": "picture",
            "in": "query"
          },
          {
            "type": "number",
            "name": "sp",
            "in": "query"
          },
          {
            "type": "number",
            "name": "sr",
            "in": "query"
          },
          {
            "type": "integer",
            "name": "maxLevel",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "response",
            "schema": {
              "type": "object",
              "required": [
                "result"
              ],
              "properties": {
                "result": {
                  "type": "object"
                }
              }
            }
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600 with HTTP Status Code 422.",
          "type": "integer",
          "format": "int32",
          "x-order": 0
        },
        "message": {
          "type": "string",
          "x-order": 1
        }
      }
    }
  },
  "responses": {
    "GenericError": {
      "description": "General errors using same model as used by go-swagger for validation errors.",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "security": [
    {
      "api_key": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "# ...\n## List of all custom errors\nFirst number is HTTP Status code, second is value of \"code\" field in returned JSON object, text description may or may not match \"message\" field in returned JSON object.\n- 409.1000: contact already exists\n",
    "title": "paint",
    "version": "0.2.0"
  },
  "basePath": "/",
  "paths": {
    "/health-check": {
      "get": {
        "security": [],
        "description": "Returns 200 if service works okay.",
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "Extra details about service status, if any.",
            "schema": {
              "type": "object",
              "additionalProperties": true
            }
          },
          "default": {
            "description": "General errors using same model as used by go-swagger for validation errors.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/pyrMeanShiftFilter": {
      "get": {
        "security": [],
        "description": "pyrMeanShiftFiltering",
        "operationId": "pyrMeanShiftFilter",
        "parameters": [
          {
            "type": "string",
            "name": "picture",
            "in": "query"
          },
          {
            "type": "number",
            "name": "sp",
            "in": "query"
          },
          {
            "type": "number",
            "name": "sr",
            "in": "query"
          },
          {
            "type": "integer",
            "name": "maxLevel",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "response",
            "schema": {
              "type": "object",
              "required": [
                "result"
              ],
              "properties": {
                "result": {
                  "type": "object"
                }
              }
            }
          },
          "default": {
            "description": "General errors using same model as used by go-swagger for validation errors.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600 with HTTP Status Code 422.",
          "type": "integer",
          "format": "int32",
          "x-order": 0
        },
        "message": {
          "type": "string",
          "x-order": 1
        }
      }
    }
  },
  "responses": {
    "GenericError": {
      "description": "General errors using same model as used by go-swagger for validation errors.",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "security": [
    {
      "api_key": []
    }
  ]
}`))
}
