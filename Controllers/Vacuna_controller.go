package Controllers

import (
	"log"
	"pollitoBackendGolang/ApiHelpers"
	"pollitoBackendGolang/Database/Models"

	"github.com/gin-gonic/gin"
)

func ListVacuna(c *gin.Context) {
	var vacuna []Models.Vacuna
	err := Models.GetAllVacuna(&vacuna)
	if err != nil {
		log.Println("Error on list vacunas")
		ApiHelpers.RespondJSON(c, 404, vacuna, "Error")
	} else {
		log.Println("Success list vacunas")
		ApiHelpers.RespondJSON(c, 200, vacuna, "ok")
	}
}

func AddNewVacuna(c *gin.Context) {
	var vacuna Models.Vacuna
	c.BindJSON(&vacuna)

	ApiHelpers.RespondJSON(c, 200, vacuna, "ok")

	// err := Models.AddNewVacuna(&vacuna)
	// if err != nil {
	// 	log.Println("Error on insert vacuna:", vacuna)
	// 	ApiHelpers.RespondJSON(c, 404, vacuna, "Error")
	// } else {
	// 	log.Println("Success inserted vacuna:", vacuna)
	// 	ApiHelpers.RespondJSON(c, 200, vacuna, "ok")
	// }
}

func GetOneVacuna(c *gin.Context) {
	id := c.Params.ByName("id")
	var vacuna Models.Vacuna
	err := Models.GetOneVacuna(&vacuna, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vacuna, "Error")
	} else {
		ApiHelpers.RespondJSON(c, 200, vacuna, "ok")
	}
}

func PutOneVacuna(c *gin.Context) {
	var vacuna Models.Vacuna
	id := c.Params.ByName("id")
	err := Models.GetOneVacuna(&vacuna, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vacuna, "Error")
	}
	c.BindJSON(&vacuna)
	err = Models.PutOneVacuna(&vacuna, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vacuna, "Error")
	} else {
		ApiHelpers.RespondJSON(c, 200, vacuna, "ok")
	}
}

func DeleteVacuna(c *gin.Context) {
	var vacuna Models.Vacuna
	id := c.Params.ByName("id")
	err := Models.DeleteVacuna(&vacuna, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vacuna, "Error")
	} else {
		ApiHelpers.RespondJSON(c, 200, vacuna, "ok")
	}
}
