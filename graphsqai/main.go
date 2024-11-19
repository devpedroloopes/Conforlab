package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type Card struct {
	Title   string
	Values  [4]int
	Total   int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cards := []Card{
			{Title: "AVACR", Values: generateRandomValues()},
			{Title: "Conservação e Limpeza", Values: generateRandomValues()},
			{Title: "Controle de Pragas", Values: generateRandomValues()},
			{Title: "Obras e Reformas", Values: generateRandomValues()},
			{Title: "Recursos Humanos", Values: generateRandomValues()},
			{Title: "Comunicação", Values: generateRandomValues()},
		}

		for i := range cards {
			cards[i].Total = calculateTotal(cards[i].Values)
		}

		tpl := template.Must(template.New("index").Parse(indexTemplate))
		err := tpl.Execute(w, cards)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}

func generateRandomValues() [4]int {
	var values [4]int
	for i := range values {
		values[i] = rand.Intn(100)
	}
	return values
}

func calculateTotal(values [4]int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

const indexTemplate = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Grid de Cards</title>
	<link rel="stylesheet" type="text/css" href="/static/styles.css">
	<script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
	<style>
    * {
        font-family: 'Arial', sans-serif;
    }

    body {
        background-color: #f4f4f9;
        margin: 0;
        padding: 0;
    }

    .grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
        gap: 10px;
        padding: 20px;
        justify-content: center;
    }

    @media (min-width: 768px) {
        .grid {
            grid-template-columns: repeat(3, 1fr);
        }
    }

    .card {
        background-color: #ffffff;
        border-radius: 8px;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        padding: 20px;
        transition: transform 0.2s ease-in-out;
        display: flex;
        flex-direction: column;
        align-items: center;
        text-align: center;
        width: 100%;
        max-width: 350px;
        margin: 0 auto;
		margin-bottom: 20px;
    }

    .card:hover {
        transform: translateY(-5px);
        box-shadow: 0 6px 10px rgba(0, 0, 0, 0.15);
    }

    .card h3 {
        margin-top: 0;
        font-size: 20px;
        font-weight: 600;
        color: #333333;
        margin-bottom: 10px;
    }

    .chart-container {
        width: 100%;
        max-width: 200px; 
        text-align: center;
        margin-top: 20px;
        margin-bottom: 20px;
    }

    .card p {
        margin-top: 10px;
        font-size: 16px;
        font-weight: 500;
        color: #666666;
    }

    .title {
        text-align: center;
        font-size: 28px;
        font-weight: 700;
        color: #333333;
        margin: 20px 0;
    }
</style>
</head>
<body>
	<div class="title">Zephyros Resumo dos Formulários</div>
	<div class="grid">
		{{ range . }}
			<div class="card">
				<h3>{{ .Title }}</h3>
				<div class="chart-container">
					<canvas id="chart-{{ .Title }}"></canvas>
				</div>
				<p>Total: {{ .Total }}</p>
			</div>
		{{ end }}
	</div>

	<script>
		{{ range . }}
			var ctx = document.getElementById('chart-{{ .Title }}').getContext('2d');
			var chart = new Chart(ctx, {
				type: 'doughnut',
				data: {
					labels: [
						'Ok: {{ index .Values 0 }}',
						'Recomendação: {{ index .Values 1 }}',
						'Obrigatório: {{ index .Values 2 }}',
						'Não Aplicável: {{ index .Values 3 }}'
					],
					datasets: [{
						label: 'Porcentagem',
						data: {{ .Values }},
						backgroundColor: ['#4CAF50', '#FFC107', '#F44336', '#607D8B'], 
						borderWidth: 1 
					}]
				},
				options: {
					cutoutPercentage: 80, 
					responsive: true,
					maintainAspectRatio: false,
					plugins: {
						legend: {
							display: true,
							position: 'bottom',
							align: 'start',
							labels: {
								usePointStyle: true,
								pointStyle: 'rectRounded',
								textAlign: 'left',
								padding: 20,  
								font: {
									size: 14,
									family: 'Arial',
									weight: '500'
								}
							}
						}
					},
					animation: {
						animateScale: true,
						animateRotate: true
					}
				}
			});
		{{ end }}
	</script>
</body>
</html>
`
