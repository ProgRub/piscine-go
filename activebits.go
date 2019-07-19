package piscine

func ActiveBits(n int) uint {
	binario := []rune(PutNbrBase(n, "01"))
	var ativos uint
	ativos = 0
	for _, s := range binario {
		if s == '1' {
			ativos++
		}
	}
	return ativos
}
