package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	port = ":8090"
)

type Persona struct {
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Edad      int    `json:"edad"`
	Direccion string `json:"dirrecion"`
	Telefono  string `json:"telefono"`
	Activo    bool   `json:"activo"`
}

func main() {
	persona := Persona{
		Nombre:    "Juan",
		Apellido:  "Perez",
		Edad:      20,
		Direccion: "Av. Siempre Viva",
		Telefono:  "09876523",
		Activo:    true,
	}

	json, err := json.Marshal(persona) //serializa a json
	if err != nil {
		log.Fatal(err)
	}
	data := string(json)
	fmt.Println(data)

	//go get -u github.com/gin-gonic/gin

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) { // Para saber que levanto el servicio
		ctx.JSON(http.StatusOK, gin.H{
			"mensaje": "pong",
		})

	})

	router.GET("/persona", func(ctx *gin.Context) {
		personaEnpoint := Persona{
			Nombre:    "Elkin",
			Apellido:  "Fuentes",
			Direccion: "Tocancipa",
			Telefono:  "123345345",
			Activo:    true,
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":  personaEnpoint, // seriliza directo a json
			"data2": persona,
		})

	})

	router.Run(port)
}
