package ejercicios

// Escriba un método recursivo que tome una cadena y
// devuelva como resultado la cadena invertida.
func Invertir(cadena string) string {
	//panic("Not implemented")
	if len(cadena) < 2 {
		return cadena
	}
	return Invertir(cadena[1:]) + cadena[:1]
}
