package main

import "fmt"

func main() {
	// 5 kapasiteli ama 3 uzunluklu bir dilim
	original := make([]int, 3, 5)
	original[0], original[1], original[2] = 10, 20, 30

	fmt.Printf("Başlangıç: original -> %v (len=%d, cap=%d)\n", original, len(original), cap(original))

	// s1, original'in bir görünümü. Aynı alttaki diziyi kullanıyorlar.
	s1 := original[:2]
	fmt.Printf("           s1 -> %v (len=%d, cap=%d)\n", s1, len(s1), cap(s1))

	// s1'e eleman ekleyelim. s1'in kapasitesi (5) yeterli olduğu için YENİ DİZİ AYRILMAZ!
	s1 = append(s1, 99)

	fmt.Println("\n--- append(s1, 99) sonrası ---")
	fmt.Printf("SONUÇ: s1 -> %v (len=%d, cap=%d)\n", s1, len(s1), cap(s1))

	// BEKLENMEDİK SONUÇ: s1'e yapılan ekleme original'in 3. elemanının üzerine yazdı!
	fmt.Printf("SÜRPRİZ: original -> %v (len=%d, cap=%d)\n", original, len(original), cap(original))
}

/*
ÇIKTI:
Başlangıç: original -> [10 20 30] (len=3, cap=5)
           s1 -> [10 20] (len=2, cap=5)

--- append(s1, 99) sonrası ---
SONUÇ: s1 -> [10 20 99] (len=3, cap=5)
SÜRPRİZ: original -> [10 20 99] (len=3, cap=5)
*/

//Bu tuzağı önlemek için ya copy fonksiyonu ile bağımsız bir kopya oluşturulmalı ya da
//"full slice expression" (s[a:b:c]) ile yeni dilimin kapasitesi sınırlandırılmalıdır.



//---------------------------------------------------------

/*Dilimin Başına Eleman Ekleme (Prepending)
Go'da prepend diye bir fonksiyon yoktur, ama append kullanarak bunu kolayca yapabilirsiniz. Bu, sona eklemekten daha maliyetlidir çünkü yeni bir dizi ayırmayı ve kopyalamayı garanti eder.
*/

s := []int{1, 2, 3}
newItem := 0

// Idiomatic Go: Geçici bir dilim oluştur ve eskisini onun sonuna ekle
s = append([]int{newItem}, s...)

fmt.Println(s) // Çıktı: [0 1 2 3]