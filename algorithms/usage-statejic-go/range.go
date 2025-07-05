package main

import "fmt"

type User struct {
	Name     string
	IsActive bool
}

func main() {
	users := []User{
		{Name: "Alice", IsActive: false},
		{Name: "Bob", IsActive: false},
	}

	// Kullanıcıları aktif etmeye ÇALIŞALIM (BAŞARISIZ olacak)
	for _, user := range users {
		user.IsActive = true // BURASI KOPYA ÜZERİNDE ÇALIŞIYOR!
	}
	fmt.Printf("Sonuç: %+v\n", users) // Çıktı: IsActive hala false!

	// DOĞRU YÖNTEM: İndeksi kullanarak orijinal dilimi değiştirmek
	for i := range users {
		users[i].IsActive = true // BURASI ORİJİNAL VERİ ÜZERİNDE ÇALIŞIYOR!
	}
	fmt.Printf("Doğru Sonuç: %+v\n", users)
}

/*
Performans Etkisi ve Optimizasyon:

Değer Kopyalama: Her döngü adımında bir değer kopyalandığı için, eğer çok büyük struct'lar içeren bir dilim üzerinde geziyorsanız, bu kopyalama işlemi bir performans maliyeti oluşturabilir.

Optimizasyon: Bu durumu engellemek ve performansı artırmak için, büyük struct'ların kendilerini değil, işaretçilerini (pointer) tutan bir dilim ([]*User) üzerinde çalışmak çok daha verimlidir. Bu durumda, her döngüde koca struct yerine sadece 8 byte'lık bir pointer kopyalanır, bu da çok daha hızlıdır.

Go
*/
// Optimize edilmiş yapı
users := []*User{
{Name: "Alice", IsActive: false},
{Name: "Bob", IsActive: false},
}
for _, userPtr := range users {
userPtr.IsActive = true // İşaretçi üzerinden yapılan değişiklik orijinal veriyi etkiler.
}
Bu kopyalama davranışı, range'in en temel ve anlaşılması en kritik özelliğidir.