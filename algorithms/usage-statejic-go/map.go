//package main

import "fmt"

func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a - b }

var operations = map[string]func(int, int) int{
	"+": add,
	"-": subtract,
}

func main() {
	result := operations["+"](10, 5) // add(10, 5) çağrılır
	fmt.Println("10 + 5 =", result)  // 15

	result = operations["-"](10, 5) // subtract(10, 5) çağrılır
	fmt.Println("10 - 5 =", result) // 5
}

/*

Çalışma Mantığı Adım Adım:

Hashing: Bir haritaya m["anahtar"] = deger gibi bir eleman eklediğinizde, Go önce "anahtar"ı alır ve onu bir hash fonksiyonundan geçirir. Bu fonksiyon, anahtarı alıp uint64 gibi büyük bir sayısal değere (hash değeri) dönüştürür. Go, her anahtar türü için (string, int vb.) önceden tanımlanmış, hızlı ve iyi dağılım sağlayan hash fonksiyonları kullanır.

Bucket (Kova) Seçimi: Elde edilen bu hash değeri, elemanın hangi "kova"ya (bucket) yerleştirileceğini belirlemek için kullanılır. Kovalar, anahtar-değer çiftlerini tutan küçük dizilerdir. Hangi kovaya gidileceği genellikle hash_degeri % kova_sayısı gibi bir modulo işlemiyle belirlenir.

Çakışma (Collision): Farklı iki anahtarın hash fonksiyonundan geçtikten sonra aynı kova indeksini üretmesi mümkündür. Buna çakışma (collision) denir. Go bunu, her kovanın birden fazla (Go'da genelde 8) anahtar-değer çifti tutabilmesiyle çözer. Yeni gelen çift, aynı kovadaki boş bir slota yerleştirilir.

Overflow (Taşma): Eğer bir kova tamamen dolarsa (8 çiftin hepsi de kullanılırsa), Go bu kovaya bağlı bir "taşma kovası" (overflow bucket) oluşturur ve yeni elemanı oraya koyar. Bu, kovaları birer bağlı liste (linked list) gibi davranmaya iter.

Diyagram:



4. İşlemci ve Bellek Seviyesinde Ne Oluyor?
Yapı: Bir map değişkeni, aslında heap'te bulunan hmap adlı bir struct'a işaret eden bir pointer'dır. Bu hmap struct'ı, haritanın boyutu, kova sayısı gibi meta verileri ve en önemlisi, asıl kova dizisine işaret eden başka bir pointer'ı içerir. Yani bir map elemanına erişim en az iki pointer takibi (pointer dereferencing) gerektirir.

Yazma (m["key"] = val):

İşlemci, anahtarın hash'ini hesaplar.

Hash'e göre kova indeksini hesaplar.

hmap pointer'ını takip ederek kova dizisine ulaşır, oradan da ilgili kovanın bellek adresine gider.

Kova içindeki 8 slotu tek tek gezer. Her slotta, verilen anahtarın hash'inin bir kısmını ve anahtarın kendisini karşılaştırır.

Eşleşme bulursa değeri günceller. Bulamazsa, kovada boş bir slot arar ve anahtar/değeri oraya yazar. Kova doluysa, overflow mekanizmasını tetikler.

Okuma (val, ok := m["key"]): Süreç yazmaya çok benzer. Hash hesaplanır, kova bulunur, kova ve taşma kovaları (varsa) taranır. Anahtar bulunursa değeri ve true döner; taranacak yer kalmazsa sıfır değer ve false döner.

Büyüme ve Rehashing: Bir haritadaki eleman sayısı, kova sayısına göre belirli bir oranı (load factor, Go'da ~6.5) aştığında, harita verimsizleşir. Bu noktada Go, haritayı büyütme kararı alır.

Genellikle mevcut kova sayısının iki katı büyüklüğünde yeni bir kova dizisi heap'te oluşturulur.

Eski haritadaki tüm anahtar-değer çiftleri tek tek gezilir, anahtarları yeniden hash'lenir (rehashing) ve yeni, daha büyük kova dizisindeki doğru yerlerine yerleştirilir.

Bu, oldukça maliyetli bir operasyondur. Modern Go sürümleri, bu maliyeti programın çalışması sırasında ani duraksamalara yol açmayacak şekilde, kademeli (incrementally) olarak yapar. Yani tüm taşıma işlemi tek seferde değil, sonraki yazma/okuma işlemlerine yayılarak yapılır.


*/
