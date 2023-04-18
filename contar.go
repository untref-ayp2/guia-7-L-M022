package main

func main() {
	contar("Hola como te va", 'a')
}
func contar(text string, buscado byte) int {
	if text == " " {
		return 0
	}
	if text[0] == buscado {
		return 1 + contar(text[1:], buscado)
	} else {
		return 0 + contar(text[1:], buscado)
	}
}
