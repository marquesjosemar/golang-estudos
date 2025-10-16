package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Estrutura do Usuário
type Usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome" binding:"required"`
	Email string `json:"email" binding:"required"`
	Idade int    `json:"idade" binding:"required"`
}

// Banco de dados em memória simulando MySQL
var baseDeDados = []Usuario{}
var proximoID = 1

func main() {
	router := gin.Default()

	// Middleware para logs
	router.Use(gerarLogs())

	configurarRotas(router)

	// Adicionar alguns usuários de exemplo
	adicionarUsuariosExemplo()

	router.Run(":8080")
}

func configurarRotas(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		usuarios := api.Group("/usuarios")
		{
			usuarios.GET("", listarUsuarios)
			usuarios.GET("/:id", buscarUsuarioPorID)
			usuarios.POST("", criarUsuario)
			usuarios.PUT("/:id", atualizarUsuario)
			usuarios.DELETE("/:id", deletarUsuario)
		}
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensagem":      "Bem-vindo ao CRUD de Usuários",
			"instrucoes":    "Use /api/v1/usuarios para gerenciar usuários",
			"usuarios_cadastrados": len(baseDeDados),
		})
	})

	router.GET("/health", healthCheck)
}

// Middleware para logs
func gerarLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		println("Requisição:", c.Request.Method, c.Request.URL.Path, "Status:", c.Writer.Status())
	}
}

// Health check
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "servidor_online",
		"mensagem":  "Sistema funcionando corretamente",
		"banco_de_dados": "conectado",
	})
}

// Adicionar usuários de exemplo
func adicionarUsuariosExemplo() {
	usuariosExemplo := []Usuario{
		{ID: proximoID, Nome: "João Silva", Email: "joao@email.com", Idade: 30},
		{ID: proximoID + 1, Nome: "Maria Santos", Email: "maria@email.com", Idade: 25},
	}
	
	baseDeDados = append(baseDeDados, usuariosExemplo...)
	proximoID += 2
}

// Listar todos os usuários
func listarUsuarios(c *gin.Context) {
	if len(baseDeDados) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"mensagem": "Nenhum usuário cadastrado",
			"dados":    []Usuario{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Usuários listados com sucesso",
		"dados":    baseDeDados,
		"total":    len(baseDeDados),
	})
}

// Buscar usuário por ID
func buscarUsuarioPorID(c *gin.Context) {
	idParametro := c.Param("id")
	id, err := strconv.Atoi(idParametro)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID deve ser um número válido",
		})
		return
	}

	for _, usuario := range baseDeDados {
		if usuario.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"mensagem": "Usuário encontrado",
				"dados":    usuario,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"erro": "Usuário não encontrado",
	})
}

// Criar novo usuário
func criarUsuario(c *gin.Context) {
	var novoUsuario Usuario

	// Validar dados de entrada
	if err := c.ShouldBindJSON(&novoUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos: " + err.Error(),
		})
		return
	}

	// Validar email único
	for _, usuario := range baseDeDados {
		if usuario.Email == novoUsuario.Email {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro": "Email já cadastrado",
			})
			return
		}
	}

	// Atribuir ID e salvar
	novoUsuario.ID = proximoID
	proximoID++
	baseDeDados = append(baseDeDados, novoUsuario)

	c.JSON(http.StatusCreated, gin.H{
		"mensagem": "Usuário criado com sucesso",
		"dados":    novoUsuario,
	})
}

// Atualizar usuário
func atualizarUsuario(c *gin.Context) {
	idParametro := c.Param("id")
	id, err := strconv.Atoi(idParametro)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID deve ser um número válido",
		})
		return
	}

	var dadosAtualizados Usuario
	if err := c.ShouldBindJSON(&dadosAtualizados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos: " + err.Error(),
		})
		return
	}

	// Buscar usuário e atualizar
	for indice, usuario := range baseDeDados {
		if usuario.ID == id {
			// Validar email único (exceto para o próprio usuário)
			for i, u := range baseDeDados {
				if u.Email == dadosAtualizados.Email && i != indice {
					c.JSON(http.StatusBadRequest, gin.H{
						"erro": "Email já está em uso por outro usuário",
					})
					return
				}
			}

			// Manter o ID original e atualizar outros campos
			dadosAtualizados.ID = usuario.ID
			baseDeDados[indice] = dadosAtualizados

			c.JSON(http.StatusOK, gin.H{
				"mensagem": "Usuário atualizado com sucesso",
				"dados":    dadosAtualizados,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"erro": "Usuário não encontrado",
	})
}

// Deletar usuário
func deletarUsuario(c *gin.Context) {
	idParametro := c.Param("id")
	id, err := strconv.Atoi(idParametro)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID deve ser um número válido",
		})
		return
	}

	for indice, usuario := range baseDeDados {
		if usuario.ID == id {
			// Remover usuário do slice
			baseDeDados = append(baseDeDados[:indice], baseDeDados[indice+1:]...)
			
			c.JSON(http.StatusOK, gin.H{
				"mensagem": "Usuário deletado com sucesso",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"erro": "Usuário não encontrado",
	})
}

// ... (as funções listarUsuarios, buscarUsuarioPorID, criarUsuario, 
// atualizarUsuario e deletarUsuario permanecem as mesmas)