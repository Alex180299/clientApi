package api

import (
	"clientApi/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func newClient(c echo.Context) error {
	client := new(models.Client)

	if err := c.Bind(client); err != nil {
		response := &models.Response{
			CveError:   -1,
			CveMensaje: "No se pudo obtener la informacion del cliente",
		}

		return c.JSON(http.StatusBadRequest, response)
	}

	isDuplicated := verifyDuplicated(client)

	if isDuplicated != "" {
		response := &models.Response{
			CveError:   0,
			CveMensaje: isDuplicated,
		}

		return c.JSON(http.StatusNotAcceptable, response)
	}

	idSaved := saveClient(client)

	response := &models.Response{
		CveError:   0,
		CveMensaje: fmt.Sprintf("Client guardado con el id %d", idSaved),
	}
	return c.JSON(http.StatusCreated, response)
}

func verifyDuplicated(client *models.Client) string {
	clientList := getClients()

	for _, clientInList := range clientList {
		if client.NombreUsuario == clientInList.NombreUsuario {
			return "Nombre de usuario ya utilizado"
		}
		if client.CorreoElectronico == clientInList.CorreoElectronico {
			return "Correo electronico ya utilizado"
		}
	}

	return ""
}

func updateClient(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	client := new(models.Client)

	if err := c.Bind(client); err != nil {
		response := &models.Response{
			CveError:   -1,
			CveMensaje: "No se pudo obtener la informacion del cliente",
		}

		return c.JSON(http.StatusBadRequest, response)
	}

	clientModified := modifyClient(idInt, client)

	if clientModified == nil {
		response := &models.Response{
			CveError:   -1,
			CveMensaje: fmt.Sprintf("No existe el cliente con id %d", idInt),
		}

		return c.JSON(http.StatusNotFound, response)
	}

	response := &models.Response{
		CveError:   0,
		CveMensaje: clientModified,
	}
	return c.JSON(http.StatusAccepted, response)
}

func getClient(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	client := getClientById(idInt)

	if client == nil {
		response := &models.Response{
			CveError:   -1,
			CveMensaje: fmt.Sprintf("No existe el cliente con id %d", idInt),
		}

		return c.JSON(http.StatusNotFound, response)
	}

	response := &models.Response{
		CveError:   0,
		CveMensaje: client,
	}
	return c.JSON(http.StatusFound, response)
}

func getAllClients(c echo.Context) error {
	clientList := getClients()
	response := &models.Response{
		CveError:   0,
		CveMensaje: clientList,
	}
	return c.JSON(http.StatusFound, response)
}
