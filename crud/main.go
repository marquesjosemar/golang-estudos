package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Estrutura que representa uma Tarefa
type Tarefa struct {
	ID          int    `json:"id"`
	Titulo      string `json:"titulo"`
	Descricao   string `json:"descricao"`
	Concluida   bool   `json:"concluida"`
}

// Nosso “banco de dados” em memória
var tarefas []Tarefa
var proximoID = 1

// ------------------------------
// FUNÇÕES DE CADA OPERAÇÃO CRUD
// ------------------------------

//  1. Listar todas as tarefas
func listarTarefas(c *gin.Context) {
	c.JSON(http.StatusOK, tarefas)
}

//  2. Buscar tarefa por ID
func buscarTarefa(c *gin.Context) {
	id := c.Param("id")

	for _, t := range tarefas {
		if id == intParaString(t.ID) {
			c.JSON(http.StatusOK, t)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"mensagem": "Tarefa não encontrada"})
}

//  3. Criar nova tarefa
func criarTarefa(c *gin.Context) {
	var nova Tarefa

	// Lê o JSON enviado no corpo da requisição
	if err := c.ShouldBindJSON(&nova); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	nova.ID = proximoID
	proximoID++
	tarefas = append(tarefas, nova)

	c.JSON(http.StatusCreated, nova)
}

//  4. Atualizar tarefa existente
func atualizarTarefa(c *gin.Context) {
	id := c.Param("id")
	var atualizada Tarefa

	if err := c.ShouldBindJSON(&atualizada); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for i, t := range tarefas {
		if id == intParaString(t.ID) {
			atualizada.ID = t.ID
			tarefas[i] = atualizada
			c.JSON(http.StatusOK, atualizada)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"mensagem": "Tarefa não encontrada"})
}

//  5. Deletar tarefa
func deletarTarefa(c *gin.Context) {
	id := c.Param("id")

	for i, t := range tarefas {
		if id == intParaString(t.ID) {
			tarefas = append(tarefas[:i], tarefas[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"mensagem": "Tarefa deletada"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"mensagem": "Tarefa não encontrada"})
}

// ------------------------------
// Função auxiliar
// ------------------------------
func intParaString(num int) string {
	return fmt.Sprintf("%d", num)
}

// ------------------------------
// Função principal
// ------------------------------
func main() {
	r := gin.Default()

	// Rotas da API
	r.GET("/tarefas", listarTarefas)
	r.GET("/tarefas/:id", buscarTarefa)
	r.POST("/tarefas", criarTarefa)
	r.PUT("/tarefas/:id", atualizarTarefa)
	r.DELETE("/tarefas/:id", deletarTarefa)

	r.Run(":8080")
}
