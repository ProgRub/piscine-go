package piscine

func UltimateDivMod(a *int, b *int) {
	auxA := *a
	auxB := *b
	*a = auxA / auxB
	*b = auxA % auxB
}
