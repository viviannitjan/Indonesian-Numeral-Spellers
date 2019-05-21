package main

import "fmt"

func satuan(i int) string {
	switch i {
	case 1:
		return "satu"
	case 2:
		return "dua"
	case 3:
		return "tiga"
	case 4:
		return "empat"
	case 5:
		return "lima"
	case 6:
		return "enam"
	case 7:
		return "tujuh"
	case 8:
		return "delapan"
	case 9:
		return "sembilan"
	}
	return ""
}

func puluhan(i int) string {
	if i/10 == 1 {
		if i%10 == 0 {
			return "sepuluh "
		} else if i%10 == 1 {
			return "sebelas "
		} else {
			return satuan(i%10) + " belas "
		}
	} else if i/10 > 1 {
		return satuan(i/10) + " puluh " + satuan(i%10)
	} else {
		return satuan(i)
	}
}

func ratusan(i int) string {
	if i/100 == 1 {
		return "seratus "
	} else if i/100 > 1 {
		return satuan(i/100) + " ratus " + puluhan(i%100)
	} else {
		return puluhan(i)
	}
}

func spellAngka(i int) string {
	var result = ""
	for i > 0 {
		if i/1000000000 > 0 {
			result = result + ratusan(i/1000000000) + " milyar "
			i = i % 1000000000
		} else if i/1000000 > 0 {
			result = result + ratusan(i/1000000) + " juta "
			i = i % 1000000
		} else if i/1000 > 0 {
			if i/1000 == 1 {
				result = result + "seribu "
				// i=i%1000
			} else {
				result = result + ratusan(i/1000) + " ribu "
			}
			i = i % 1000
		} else {
			result = result + ratusan(i)
			i = 0
		}
	}
	return result
}

func main() {
	fmt.Println(spellAngka(1465431392))
}
