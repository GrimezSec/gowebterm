package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	portPtr := flag.String("p", "8080", "port number")
	flag.Parse()

	http.HandleFunc("/", terminalHandler)
	http.HandleFunc("/ws", websocketHandler)

	port := fmt.Sprintf(":%s", *portPtr)
	fmt.Printf("Server listening on port %s...\n", *portPtr)
	log.Fatal(http.ListenAndServe(port, nil))
}

func terminalHandler(w http.ResponseWriter, r *http.Request) {
	renderForm(w, "")
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, cmd, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		output, err := executeCommand(string(cmd))
		if err != nil {
			output = err.Error()
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(output)); err != nil {
			log.Println(err)
			return
		}
	}
}

func executeCommand(cmd string) (string, error) {
	command := exec.Command("bash", "-c", cmd)
	output, err := command.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func renderForm(w http.ResponseWriter, output string) {
	tpl := template.Must(template.ParseFiles("web/index.html"))
	tpl.Execute(w, template.HTML(output))
}
