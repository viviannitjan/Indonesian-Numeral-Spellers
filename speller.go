package main

import (
	"strings"
)

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
			if i > 100 {
				i = i % 100
			} else {
				i = 0
			}
		}
	}
	return result
}

func toAngka(s string, i int) int {
	switch s {
	case "satu":
		return i + 1
	case "dua":
		return i + 2
	case "tiga":
		return i + 3
	case "empat":
		return i + 4
	case "lima":
		return i + 5
	case "enam":
		return i + 6
	case "tujuh":
		return i + 7
	case "delapan":
		return i + 8
	case "sembilan":
		return i + 9
	case "sepuluh":
		return i + 10
	case "sebelas":
		return i + 11
	case "seratus":
		return i + 100
	case "belas":
		return i + 10
	case "puluh":
		j := i % 10
		i = i - j
		i = i + j*10
		return i
	case "ratus":
		j := i % 10
		i = i - j
		i = i + j*100
		return i
	case "seribu":
		return i + 1000
	case "ribu":
		j := i % 1000
		i = i - j
		i = i + j*1000
		return i
	case "juta":
		j := i % 1000000
		i = i - j
		i = i + j*1000000
		return i
	case "milyar":
		j := i % 1000000
		i = i - j
		i = i + j*1000000000
		return i
	}
	return 0
}

func prosesString(s string) int {
	result := 0
	i := 0
	arr := strings.Split(s, " ")
	for i < len(arr) {
		// fmt.Printf(arr[i])
		// fmt.Printf("%d", result)
		result = toAngka(arr[i], result)
		i++
	}
	return result
}

// func main() {
// 	fmt.Println(spellAngka(54321122))
// 	fmt.Printf("%d", prosesString("dua ribu sembilan belas"))
// }
