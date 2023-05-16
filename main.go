package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteMariano := clientes.Titular{
		Nome:      "Mariano",
		CPF:       "123.213.312-21",
		Profissao: "Desenvolverdor Go",
	}

	contaDoMariano := contas.ContaCorrente{
		Titular:       clienteMariano,
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}

	poupancaDoMariano := contas.ContaPoupanca{
		Titular:       clienteMariano,
		NumeroAgencia: 123,
		NumeroConta:   654321,
		Operacao:      015,
	}

	contaDoMariano.Depositar(100)
	poupancaDoMariano.Depositar(60)
	PagarBoleto(&poupancaDoMariano, 60)
	fmt.Println(poupancaDoMariano.ObterSaldo())
}

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}
