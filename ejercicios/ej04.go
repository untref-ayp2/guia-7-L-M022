package ejercicios

// Escriba un método recursivo que tome una cadena y
// devuelva como resultado la cadena invertida.
func Invertir(cadena string) string {
	if len(cadena) == 0 {
		return cadena
	} else {
		return Invertir(string(cadena[1:])) + string(cadena[0])
	}
}
