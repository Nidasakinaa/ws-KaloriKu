// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/Nidasakinaa",
            "email": "714220040@std.ulbi.ac.id"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/delete/{id}": {
            "delete": {
                "description": "Hapus data pasien.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Delete data pasien.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/insert": {
            "post": {
                "description": "Input data pasien.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Insert data pasien.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqPasien"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Biodata"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/pasien": {
            "get": {
                "description": "Mengambil semua data pasien.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Get All Data Pasien.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Biodata"
                        }
                    }
                }
            }
        },
        "/pasien/{id}": {
            "get": {
                "description": "Ambil per ID data pasien.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Get By ID Data Pasien.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Biodata"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/update/{id}": {
            "put": {
                "description": "Ubah data pasien.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pasien"
                ],
                "summary": "Update data pasien.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqPasien"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Biodata"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Biodata": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "alamat": {
                    "type": "string"
                },
                "doctor": {
                    "$ref": "#/definitions/controller.Doctor"
                },
                "gender": {
                    "type": "string"
                },
                "medicalRecord": {
                    "$ref": "#/definitions/controller.MedicalRecord"
                },
                "pasienName": {
                    "type": "string"
                },
                "phonenumber": {
                    "type": "string"
                },
                "usia": {
                    "type": "string"
                }
            }
        },
        "controller.Doctor": {
            "type": "object",
            "properties": {
                "contact": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "specialty": {
                    "type": "string"
                }
            }
        },
        "controller.MedicalRecord": {
            "type": "object",
            "properties": {
                "diagnosis": {
                    "type": "string"
                },
                "m_id": {
                    "type": "string"
                },
                "treatment": {
                    "type": "string"
                },
                "visitdate": {
                    "type": "string"
                }
            }
        },
        "controller.ReqDoctor": {
            "type": "object",
            "properties": {
                "contact": {
                    "type": "string",
                    "example": "022987"
                },
                "name": {
                    "type": "string",
                    "example": "Sari"
                },
                "specialty": {
                    "type": "string",
                    "example": "Oncology"
                }
            }
        },
        "controller.ReqMedicalRecord": {
            "type": "object",
            "properties": {
                "diagnosis": {
                    "type": "string",
                    "example": "Flu"
                },
                "treatment": {
                    "type": "string",
                    "example": "istirahat dan banyak mengonsumsi air putih"
                },
                "visitdate": {
                    "type": "string",
                    "example": "12 Juni 2026"
                }
            }
        },
        "controller.ReqPasien": {
            "type": "object",
            "properties": {
                "alamat": {
                    "type": "string",
                    "example": "Sariasih 25, Bandung"
                },
                "doctor": {
                    "$ref": "#/definitions/controller.ReqDoctor"
                },
                "gender": {
                    "type": "string",
                    "example": "Laki-laki"
                },
                "medicalRecord": {
                    "$ref": "#/definitions/controller.ReqMedicalRecord"
                },
                "pasienName": {
                    "type": "string",
                    "example": "Doni"
                },
                "phonenumber": {
                    "type": "string",
                    "example": "08567432"
                },
                "usia": {
                    "type": "string",
                    "example": "20 Tahun"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "ws-rumahsakit-2e1eb71f14e2.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "TES SWAGGER PASIEN",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
