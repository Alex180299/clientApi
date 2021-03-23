package api

import (
	"clientApi/models"
	"time"
)

var clients []*models.Client

func InitRepository() {
	clients = make([]*models.Client, 0)
}

func saveClient(client *models.Client) int {
	client.FechaCreacion = time.Now()
	client.FechaActualizacion = time.Now()

	if len(clients) == 0 {
		client.ClienteID = 1
	} else {
		lastClient := clients[len(clients)-1]
		client.ClienteID = lastClient.ClienteID + 1
	}

	clients = append(clients, client)
	return client.ClienteID
}

func modifyClient(clientId int, client *models.Client) *models.Client {
	for _, clientInList := range clients {
		if clientInList.ClienteID == clientId {
			clientInList.FechaActualizacion = time.Now()

			if client.NombreUsuario != "" {
				clientInList.NombreUsuario = client.NombreUsuario
			}
			if client.Contrasena != "" {
				clientInList.Contrasena = client.Contrasena
			}
			if client.Nombre != "" {
				clientInList.Nombre = client.Nombre
			}
			if client.Apellidos != "" {
				clientInList.Apellidos = client.Apellidos
			}
			if client.CorreoElectronico != "" {
				clientInList.CorreoElectronico = client.CorreoElectronico
			}
			if client.Edad != 0 {
				clientInList.Edad = client.Edad
			}
			if client.Estatura != 0 {
				clientInList.Estatura = client.Estatura
			}
			if client.Peso != 0 {
				clientInList.Peso = client.Peso
			}
			if client.IMC != 0 {
				clientInList.IMC = client.IMC
			}
			if client.GEB != 0 {
				clientInList.GEB = client.GEB
			}
			if client.ETA != 0 {
				clientInList.ETA = client.ETA
			}

			return clientInList
		}
	}

	return nil
}

func getClients() []*models.Client {
	return clients
}

func getClientById(clientId int) *models.Client {
	for _, client := range clients {
		if client.ClienteID == clientId {
			return client
		}
	}

	return nil
}
