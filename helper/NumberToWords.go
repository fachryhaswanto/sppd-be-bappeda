package helper

func NumberToWords(number int64) string {
	var satuan = []string{"", "Satu", "Dua", "Tiga", "Empat", "Lima", "Enam", "Tujuh", "Delapan", "Sembilan", "Sepuluh", "Sebelas"}

	var puluhan = []string{"", "Sepuluh", "Dua Puluh", "Tiga Puluh", "Empat Puluh", "Lima Puluh", "Enam Puluh", "Tujuh Puluh", "Delapan Puluh", "Sembilan Puluh"}

	if number < 12 {
		return satuan[number]
	} else if number < 20 {
		return satuan[number%10] + " Belas"
	} else if number < 100 {
		return puluhan[number/10] + " " + satuan[number%10]
	} else if number < 200 {
		return "Seratus " + NumberToWords(number%100)
	} else if number < 1000 {
		return satuan[number/100] + " Ratus " + NumberToWords(number%100)
	} else if number < 2000 {
		return "Seribu " + NumberToWords(number%1000)
	} else if number < 1000000 {
		return NumberToWords(number/1000) + " Ribu " + NumberToWords(number%1000)
	} else if number < 1000000000 {
		return NumberToWords(number/1000000) + " Juta " + NumberToWords(number%1000000)
	} else if number < 1000000000000 {
		return NumberToWords(number/1000000000) + " Milyar " + NumberToWords(number%1000000000)
	}

	return "Angka terlalu besar"
}
