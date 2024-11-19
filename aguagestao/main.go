package main

import (
	"fmt"
	"html/template"
	"net/http"
	"views/components"
)

func main() {
	// Rota principal (home)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.New("home").Parse(components.Home))
		err := tpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Rota /novoCliente que renderiza a página "Novo Cliente"
	http.HandleFunc("/novoCliente", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.New("novoCliente").Parse(components.NovoCliente))
		err := tpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Servir arquivos estáticos
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Porta do servidor
	port := ":9242"
	serverURL := fmt.Sprintf("http://localhost%s", port)

	// Mensagem no console ao iniciar o servidor
	fmt.Printf("Servidor iniciado em %s\n", serverURL)

	// Iniciar o servidor
	http.ListenAndServe(port, nil)
}
