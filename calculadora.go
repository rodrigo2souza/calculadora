package calculadora

import (
	"fmt"
)

// Teste calculadora
func Calculadora(a, b int) string {
	return fmt.Sprintf("A soma de %d mais %d é igual a: %d" , a, b, a+b)
}