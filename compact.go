package piscine

func Compact(ptr *[]string) int {
	conta := 0
	aux := *ptr
	final := make([]string, 0)
	for i := range aux {
		if aux[i] != "" {
			conta++
			final = append(final, aux[i])
		}
	}
	*ptr = final
	return conta
}
