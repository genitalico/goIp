package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"goIp/settings"
	"io"
	"net/http"
	"os"
)

type DataIP struct {
	Ip string `json:"ip"`
}

func saveIp(text string, dataFileName string) {
	var file *os.File
	var err error

	file, err = os.Create(dataFileName)

	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}
}

func getLastIp(dataFileName string) string {
	file, err := os.Open(dataFileName)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return "0.0.0.0"
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return "0.0.0.0"
	}

	return string(content)
}

func getIp(url string) (*DataIP, bool) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al obtener la respuesta:", err)
		return nil, false
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return nil, false
	}

	stringBody := string(body)

	var dataRequest DataIP

	errUnMarshall := json.Unmarshal([]byte(stringBody), &dataRequest)
	if errUnMarshall != nil {
		return nil, false
	}

	return &dataRequest, true
}

func sendTelegramMessage(message string, botUrl, chatId string) {

	data := []byte(`{"chat_id": "` + chatId + `", "text": "` + message + `", "parse_mode": "Markdown"}`)

	// Crea una solicitud POST con los datos
	req, err := http.NewRequest("POST", botUrl, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error al crear la solicitud:", err)
		return
	}

	// Configura la cabecera de contenido
	req.Header.Set("Content-Type", "application/json")

	// Crea un cliente HTTP y envía la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al enviar la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	// Lee la respuesta del servidor
	_, err = io.ReadAll(resp.Body)

	if err != nil {
		return
	}
}

func main() {

	var pathFile string
	flag.StringVar(&pathFile, "p", "nil", "ruta a la configuración")

	flag.Parse()

	dataSettings, err := settings.ReadFileSettings(pathFile)

	if err != nil {
		fmt.Println("Error al leer el archivo de configuración:", err)
	}

	dataIp, isGetIp := getIp(dataSettings.IpUrl)

	if !isGetIp {
		return
	}

	lastIp := getLastIp(dataSettings.DataFile)

	if lastIp != dataIp.Ip {
		messageToTelegram := dataSettings.TelegramMessage + dataIp.Ip
		sendTelegramMessage(messageToTelegram, dataSettings.BotUrl, dataSettings.ChatId)
		saveIp(dataIp.Ip, dataSettings.DataFile)
		fmt.Println(messageToTelegram)
	}
}
