// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Keptn Team",
            "url": "http://www.keptn.sh"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/sequence/project/{project}": {
            "get": {
                "description": "Get all the sequences which are present in a project",
                "tags": [
                    "Sequence"
                ],
                "summary": "Get all sequences for specific project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The name of the project",
                        "name": "project",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/api.GetSequenceExecutionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/sequence/project/{project}/shkeptncontext/{shkeptncontext}": {
            "get": {
                "description": "Get a specific sequence of a project which is identified by the shkeptncontext",
                "tags": [
                    "Sequence"
                ],
                "summary": "Get a sequence with the shkeptncontext",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The name of the project",
                        "name": "project",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The shkeptncontext",
                        "name": "shkeptncontext",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/models.SequenceState"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/sequence/project/{project}/shkeptncontext/{shkeptncontext}/event": {
            "get": {
                "description": "Gets all the events of a project with the given shkeptncontext",
                "tags": [
                    "Sequence"
                ],
                "summary": "Get all the Events",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The name of the project",
                        "name": "project",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The shkeptncontext",
                        "name": "shkeptncontext",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.KeptnContextExtendedCE"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/sequence/project/{project}/shkeptncontext/{shkeptncontext}/event/{eventId}": {
            "get": {
                "description": "Gets a single event of a project with the given shkeptncontext and eventId",
                "tags": [
                    "Sequence"
                ],
                "summary": "Get a single Event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The name of the project",
                        "name": "project",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The shkeptncontext",
                        "name": "shkeptncontext",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The Id of the event",
                        "name": "eventId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/models.KeptnContextExtendedCE"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/sequence/project/{project}/shkeptncontext/{shkeptncontext}/stage/{stage}/blocking": {
            "get": {
                "description": "Get all the sequences that are blocking a sequence from being run",
                "tags": [
                    "Sequence"
                ],
                "summary": "Get all blocking sequences for specific sequence",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The name of the project",
                        "name": "project",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The Context of the sequence",
                        "name": "shkeptncontext",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The Stage of the sequences",
                        "name": "stage",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.SequenceExecution"
                            }
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.GetSequenceExecutionResponse": {
            "type": "object",
            "properties": {
                "nextPageKey": {
                    "description": "NextPageKey is the offset to the next page",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "PageSize is the actual size of returned page",
                    "type": "integer"
                },
                "sequenceExecutions": {
                    "description": "SequenceExecutions array containing the result",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SequenceExecution"
                    }
                },
                "totalCount": {
                    "description": "Total number of matching entries",
                    "type": "integer"
                }
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Error code",
                    "type": "integer"
                },
                "message": {
                    "description": "Error message\nRequired: true",
                    "type": "string"
                }
            }
        },
        "models.EventScope": {
            "type": "object",
            "properties": {
                "eventType": {
                    "type": "string"
                },
                "gitcommitid": {
                    "type": "string"
                },
                "keptnContext": {
                    "type": "string"
                },
                "labels": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                },
                "project": {
                    "type": "string"
                },
                "result": {
                    "type": "string"
                },
                "service": {
                    "type": "string"
                },
                "stage": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "triggeredId": {
                    "type": "string"
                }
            }
        },
        "models.KeptnContextExtendedCE": {
            "type": "object",
            "properties": {
                "contenttype": {
                    "description": "contenttype",
                    "type": "string"
                },
                "data": {
                    "description": "data\nRequired: true"
                },
                "extensions": {
                    "description": "extensions"
                },
                "gitcommitid": {
                    "description": "gitcommitid",
                    "type": "string"
                },
                "id": {
                    "description": "id",
                    "type": "string"
                },
                "shkeptncontext": {
                    "description": "shkeptncontext",
                    "type": "string"
                },
                "shkeptnspecversion": {
                    "description": "shkeptnspecversion",
                    "type": "string"
                },
                "source": {
                    "description": "source\nRequired: true",
                    "type": "string"
                },
                "specversion": {
                    "description": "specversion",
                    "type": "string"
                },
                "time": {
                    "description": "time\nFormat: date-time",
                    "type": "string"
                },
                "triggeredid": {
                    "description": "triggeredid",
                    "type": "string"
                },
                "type": {
                    "description": "type\nRequired: true",
                    "type": "string"
                }
            }
        },
        "models.SequenceExecution": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "inputProperties": {
                    "description": "InputProperties contains properties of the event which triggered the task sequence",
                    "type": "object",
                    "additionalProperties": true
                },
                "schemaVersion": {
                    "description": "SchemaVersion indicates the scheme that is used for the internal representation of the sequence execution",
                    "type": "string"
                },
                "scope": {
                    "$ref": "#/definitions/models.EventScope"
                },
                "sequence": {
                    "description": "Sequence contains the complete sequence definition",
                    "$ref": "#/definitions/v0_2_0.Sequence"
                },
                "status": {
                    "$ref": "#/definitions/models.SequenceExecutionStatus"
                },
                "triggeredAt": {
                    "type": "string"
                }
            }
        },
        "models.SequenceExecutionStatus": {
            "type": "object",
            "properties": {
                "currentTask": {
                    "description": "CurrentTask represents the state of the currently active task",
                    "$ref": "#/definitions/models.TaskExecutionState"
                },
                "previousTasks": {
                    "description": "PreviousTasks contains the results of all completed tasks of the sequence",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TaskExecutionResult"
                    }
                },
                "state": {
                    "description": "triggered, waiting, suspended (approval in progress), paused, finished, cancelled, timedOut",
                    "type": "string"
                },
                "stateBeforePause": {
                    "description": "StateBeforePause is needed to keep track of the state before a sequence has been paused. Example: when a sequence has been paused while being queued, and then resumed, it should not be set to started immediately, but to the state it had before",
                    "type": "string"
                }
            }
        },
        "models.SequenceState": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "problemTitle": {
                    "type": "string"
                },
                "project": {
                    "type": "string"
                },
                "service": {
                    "type": "string"
                },
                "shkeptncontext": {
                    "type": "string"
                },
                "stages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SequenceStateStage"
                    }
                },
                "state": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "models.SequenceStateEvaluation": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "string"
                },
                "score": {
                    "type": "number"
                }
            }
        },
        "models.SequenceStateEvent": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.SequenceStateStage": {
            "type": "object",
            "properties": {
                "image": {
                    "type": "string"
                },
                "latestEvaluation": {
                    "$ref": "#/definitions/models.SequenceStateEvaluation"
                },
                "latestEvent": {
                    "$ref": "#/definitions/models.SequenceStateEvent"
                },
                "latestFailedEvent": {
                    "$ref": "#/definitions/models.SequenceStateEvent"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                }
            }
        },
        "models.TaskEvent": {
            "type": "object",
            "properties": {
                "eventType": {
                    "type": "string"
                },
                "properties": {
                    "type": "object",
                    "additionalProperties": true
                },
                "result": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "models.TaskExecutionResult": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "properties": {
                    "description": "Properties contains the aggregated results of the task's executors",
                    "type": "object",
                    "additionalProperties": true
                },
                "result": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "triggeredID": {
                    "type": "string"
                }
            }
        },
        "models.TaskExecutionState": {
            "type": "object",
            "properties": {
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TaskEvent"
                    }
                },
                "name": {
                    "type": "string"
                },
                "triggeredID": {
                    "type": "string"
                }
            }
        },
        "v0_2_0.Selector": {
            "type": "object",
            "properties": {
                "match": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "v0_2_0.Sequence": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v0_2_0.Task"
                    }
                },
                "triggeredOn": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v0_2_0.Trigger"
                    }
                }
            }
        },
        "v0_2_0.Task": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "properties": {},
                "triggeredAfter": {
                    "type": "string"
                }
            }
        },
        "v0_2_0.Trigger": {
            "type": "object",
            "properties": {
                "event": {
                    "type": "string"
                },
                "selector": {
                    "$ref": "#/definitions/v0_2_0.Selector"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "develop",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Control Plane API",
	Description:      "This is the API documentation of the Shipyard Controller.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
