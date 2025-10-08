package main

import (
	"errors"
	"fmt"
)

// Notificador é o nosso CONTRATO de comportamento.
//
// Estamos declarando que, para ser considerado um "Notificador" no nosso sistema,
// um tipo de dado DEVE OBRIGATORIAMENTE possuir um método chamado `Enviar`.
//
// Este método precisa ter exatamente esta assinatura:
// - Nome: Enviar
// - Parâmetros: receber uma (1) string
// - Retorno: retornar um (1) error
//
// A interface não diz COMO o método deve funcionar, apenas que ele deve existir.
type Notificador interface {
	Enviar(mensagem string) error
}

// Email é um tipo de dado concreto que armazena as informações
// necessárias para enviar um email.
//
// Note que neste ponto, `Email` ainda NÃO é um `Notificador`.
// É apenas uma estrutura de dados.
type Email struct {
	EnderecoDestino string
	Assunto         string
}

// O método Enviar está "anexado" ao struct Email (func (e Email)...).
// A assinatura deste método `Enviar(mensagem string) error` é IDÊNTICA
// à definida na interface Notificador.
//
// Ao fazer isso, o Go passa a considerar, automática e implicitamente,
// que o tipo `Email` SATISFAZ a interface `Notificador`. Não há palavra-chave "implements".
func (e Email) Enviar(mensagem string) error {
	// Lógica de negócio específica para enviar um email.
	// Primeiro, validamos os dados do próprio struct.
	if e.EnderecoDestino == "" {
		// Se a validação falhar, retornamos um erro, cumprindo o contrato.
		return errors.New("validação falhou: endereço de destino do email está vazio")
	}

	// Simulamos o envio do email imprimindo na tela.
	fmt.Printf(
		"--- EMAIL ENVIADO ---\nPara: %s\nAssunto: %s\nMensagem: %s\n---------------------\n",
		e.EnderecoDestino,
		e.Assunto,
		mensagem,
	)

	// Se tudo deu certo, retornamos 'nil', que em Go significa "sem erros".
	return nil
}

// SMS é outro tipo de dado concreto. Seus campos são
// completamente diferentes dos campos de Email.
type SMS struct {
	NumeroDestino int
}

// SMS também implementa o método Enviar, com a MESMA assinatura.
// A lógica interna, no entanto, é totalmente diferente da do Email.
// É aqui que reside a beleza da abstração.
func (s SMS) Enviar(mensagem string) error {
	// Lógica de negócio específica para enviar um SMS.
	if s.NumeroDestino < 100000000 { // Uma validação simples de número de telefone.
		return errors.New("validação falhou: número de telefone do SMS parece inválido")
	}

	// Simulamos o envio do SMS.
	fmt.Printf(
		"--- SMS ENVIADO ---\nPara o número: %d\nMensagem: %s\n-------------------\n",
		s.NumeroDestino,
		mensagem,
	)

	return nil
}

// ProcessarNotificacao é o nosso orquestrador de alto nível.
//
// O ponto mais importante está no parâmetro: `n Notificador`.
// Esta função NÃO SABE e NÃO SE IMPORTA se o `n` que ela recebe é
// um Email, um SMS, ou qualquer outra coisa.
//
// A única garantia que o compilador do Go nos dá é que a variável `n`
// certamente terá um método `Enviar(mensagem string) error`, pois
// para ser do tipo `Notificador`, ela é obrigada a ter.
//
// Este é o ponto de desacoplamento que torna nosso sistema flexível.
func ProcessarNotificacao(n Notificador, mensagem string) {
	fmt.Printf(">>> Processando uma nova notificação...\n")

	// Chamamos o método Enviar. O Go se encarrega de executar a
	// implementação correta (a do Email, se 'n' for um Email; a do SMS, se for um SMS).
	// Isso é chamado de "polimorfismo".
	err := n.Enviar(mensagem)

	// A lógica de checagem de erro é a mesma para qualquer notificador.
	if err != nil {
		fmt.Printf(">>> ERRO: Falha ao processar notificação. Motivo: %s\n\n", err)
	} else {
		fmt.Printf(">>> SUCESSO: Notificação processada.\n\n")
	}
}

func main() {
	// Mensagem padrão que enviaremos.
	mensagemDePromocao := "Aproveite nossa promoção de 50% de desconto! Use o cupom GO125."

	// Criamos uma lista de Notificadores.
	// Este slice pode conter QUALQUER tipo que satisfaça a interface Notificador.
	// Estamos misturando Email e SMS na mesma lista!
	notificadores := []Notificador{
		Email{EnderecoDestino: "cliente1@exemplo.com", Assunto: "Oferta Imperdível!"},
		SMS{NumeroDestino: 999887766},
		Email{EnderecoDestino: "parceiro-vip@exemplo.com", Assunto: "Convite Exclusivo"},
		SMS{NumeroDestino: 123}, // Este SMS vai falhar na validação.
	}

	// Agora, iteramos sobre a lista.
	// Para cada item da lista, a variável 'notificador' será um Email ou um SMS.
	// Mas o nosso loop não se importa com o tipo concreto. Ele só vê "Notificadores".
	for _, notificador := range notificadores {
		// Passamos cada item para a MESMA função de processamento.
		ProcessarNotificacao(notificador, mensagemDePromocao)
	}

	// A PROVA FINAL: E se quisermos adicionar Slack?
	// Basta criar o struct `Slack` com seu método `Enviar` e adicioná-lo
	// à nossa lista. A função `ProcessarNotificacao` e a interface `Notificador`
	// NÃO PRECISAM DE NENHUMA ALTERAÇÃO. O sistema é extensível.
}
