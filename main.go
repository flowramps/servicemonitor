package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "sync"
)

var (
    mu          sync.Mutex
    clicksView  int
)

func main() {
    // Configuração do roteador e tratamento da rota /metrics
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/metrics", metricsHandler)

    // Inicia o servidor na porta 8080
    fmt.Println("Servidor escutando na porta 8080...")
    http.ListenAndServe(":8080", nil)
}

// Manipulador para página "hello world"
func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Incrementa o contador de clicks_view
    mu.Lock()
    clicksView++
    mu.Unlock()

    // Escreve a mensagem na resposta HTTP
    fmt.Fprintln(w, "Hello, world!")
}

// Manipulador para rota /metrics
func metricsHandler(w http.ResponseWriter, r *http.Request) {
    // Obtém o valor atual de clicks_view
    mu.Lock()
    currentClicksView := clicksView
    mu.Unlock()

    // Gera métricas de exemplo
    cpuUsage := rand.Float64() * 100
    memoryUsage := rand.Float64() * 100
    diskUsage := rand.Float64() * 100

    // Formata as métricas como uma string no formato Prometheus
    metrics := fmt.Sprintf(`# HELP cpu_usage CPU usage in percentage
# TYPE cpu_usage gauge
cpu_usage %f
# HELP memory_usage Memory usage in percentage
# TYPE memory_usage gauge
memory_usage %f
# HELP disk_usage Disk usage in percentage
# TYPE disk_usage gauge
disk_usage %f
# HELP clicks_view Clicks view on hello world page
# TYPE clicks_view counter
clicks_view %d`, cpuUsage, memoryUsage, diskUsage, currentClicksView)

    // Escreve as métricas na resposta HTTP
    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintln(w, metrics)
}
