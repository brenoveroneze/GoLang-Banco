package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)

}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	//clienteBreno := clientes.Titular{Nome: "Breno", CPF: "123.123.123-12", Profissao: "Analista"}
	contaDoBreno := contas.ContaPoupanca{}

	contaDoBreno.Depositar(100)
	PagarBoleto(&contaDoBreno, 60)
	fmt.Println(contaDoBreno.ObterSaldo())

}
