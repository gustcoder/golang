package main

import "fmt"

func diaDaSemana(dia int) string {
	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Dia inexistente"
	}
}

func diaDaSemanaInternacional(dia int, lang string) string {
	var diaSemana string

	switch {
	case (dia == 1 && lang == "pt"):
		diaSemana = "Domingo"
	case (dia == 1 && lang == "fr"):
		diaSemana = "Dimanche"
		fallthrough
	case (dia == 1 && lang == "en"):
		diaSemana = "Sunday"
	default:
		diaSemana = "Não encontrado"
	}

	return diaSemana
}

func main() {
	dia := diaDaSemana(11)
	fmt.Println(dia)

	diaInternacional := diaDaSemanaInternacional(1, "en")
	fmt.Println(diaInternacional)
	diaInternacionalFr := diaDaSemanaInternacional(1, "fr")
	fmt.Println(diaInternacionalFr)
}
