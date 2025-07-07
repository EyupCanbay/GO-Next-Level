//package main

import "fmt"

/*


Örnek 1: "Private" Değişkenler ve Nesne Simülasyonu
Closure'lar, nesne yönelimli dillerdeki "private" alanlara sahip nesneleri taklit etmek için kullanılabilir. Dışarıdan doğrudan erişilemeyen, sadece closure'ın döndürdüğü metodlar aracılığıyla değiştirilebilen bir durum yaratabilirsiniz.

Go
*/
// Bu fonksiyon, Deposit ve Balance "metodları" olan bir banka hesabı döndürür.
func NewBankAccount(initialBalance int) (func(int), func() int) {
	balance := initialBalance // Bu 'balance' değişkeni "private" gibidir.

	deposit := func(amount int) {
		balance += amount
	}

	getBalance := func() int {
		return balance
	}
	return deposit, getBalance
}

func main() {
	myDeposit, myBalance := NewBankAccount(100)

	fmt.Println("Initial Balance:", myBalance()) // 100
	myDeposit(50)
	fmt.Println("New Balance:", myBalance()) // 150
	// balance değişkenine doğrudan erişemezsiniz!
}

//BU KOD ÇOK GÜZEL BİR ÖRNEK



Örnek 2: Fonksiyon Dekoratörleri
Dekoratör, mevcut bir fonksiyona dokunmadan ona yeni yetenekler (logging, zaman ölçümü, yetkilendirme vb.) ekleyen bir sarmalayıcı (wrapper) fonksiyondur. Closure'lar, bu deseni Go'da uygulamak için mükemmeldir.

Go

// Bu dekoratör, bir fonksiyona loglama yeteneği ekler.
func LoggingDecorator(fn func(string)) func(string) {
	return func(s string) { // Bu bir closure
		fmt.Println("Fonksiyon çağrılmadan önce...")
		fn(s) // Orijinal fonksiyonu çağır
		fmt.Println("Fonksiyon çağrıldıktan sonra...")
	}
}

func SayHello(name string) {
	fmt.Println("Merhaba,", name)
}

func main() {
	decoratedHello := LoggingDecorator(SayHello)
	decoratedHello("Alice")
}
Örnek 3: Goroutine'lerde Döngü Değişkeni Tuzağı (ÇOK ÖNEMLİ!)
Bu, Go'da closure'lar ile ilgili en yaygın ve en kritik hatadır. Bir for döngüsü içinde goroutine'ler başlatırken, döngü değişkenine doğrudan referans vermek genellikle yanlış çalışır.

YANLIŞ KOD:

Go

for i := 0; i < 5; i++ {
go func() {
// Bu goroutine çalıştığında, döngü çoktan bitmiş olacak
// ve 'i' değişkeninin son değeri (5) ne ise onu görecek.
fmt.Println(i)
}()
}
time.Sleep(time.Second) // Çıktı genellikle "5 5 5 5 5" olur.
Neden? Tüm goroutine'ler aynı i değişkenine referans verir (onu yakalar). Goroutine'ler çalışmaya başladığında döngü çoktan bitmiş ve i son değerini almıştır.

DOĞRU ÇÖZÜM 1 (Argüman Olarak Geçmek):

Go

for i := 0; i < 5; i++ {
go func(val int) { // 'i'yi argüman olarak al
fmt.Println(val)
}(i) // Her döngüde 'i'nin o anki DEĞERİNİ fonksiyona yolla
}
DOĞRU ÇÖZÜM 2 (Gölgeleme - Shadowing):

Go

for i := 0; i < 5; i++ {
i := i // Döngünün içinde yeni bir 'i' değişkeni yarat (gölgele)
go func() {
// Closure, artık her döngü için özel olan bu YENİ 'i'yi yakalar.
fmt.Println(i)
}()
}
3. & 5. Arka Plandaki Veri Yapısı ve Mantığı
Go'da bir fonksiyon değeri (function value) aslında bir pointer'dır. Ancak closure'lar için bu yapı biraz daha karmaşıktır.

Normal Fonksiyon: SayHello gibi normal bir fonksiyonun değeri, sadece o fonksiyonun bellekteki makine kodunu gösteren tek bir pointer'dır.

Closure: Bir closure'ın değeri ise iki parçalı bir veri yapısına işaret eden bir pointer'dır:

Fonksiyon Koduna İşaretçi (Function Pointer): Closure'ın çalıştırılacak olan asıl makine kodunun adresini tutar.

Ortam İşaretçisi (Environment Pointer): Closure'ın yakaladığı dış değişkenleri (sum gibi) barındıran, heap üzerinde özel olarak ayrılmış bir bellek bloğuna işaret eder.

Diyagram (pos := adder() için):

pos değişkeni, içinde iki pointer tutan bir yapıdır. Biri koda, diğeri ise heap'teki sum değerine gider. neg := adder() çağrıldığında, heap'te yeni bir sum alanı daha oluşturulur ve neg'in ortam işaretçisi orayı gösterir.

4. İşlemci ve Bellek Seviyesinde Ne Oluyor?
Kaçış Analizi (Escape Analysis): adder() fonksiyonu derlenirken, Go derleyicisi "kaçış analizi" yapar. sum değişkeninin, adder fonksiyonunun kapsamı dışına "kaçacak" olan (döndürülen fonksiyon tarafından referans alındığı için) bir değişken olduğunu anlar.

Heap'e Yükseltme (Promotion to Heap): Normalde sum gibi yerel değişkenler stack'te oluşturulur ve fonksiyon bitince yok olur. Ancak kaçış analizi sonucunda derleyici, sum değişkenini stack yerine doğrudan heap'te oluşturur. Çünkü bu değişkenin, onu oluşturan adder fonksiyonundan daha uzun yaşaması gerekmektedir.

Closure Çağrısı (pos(i)):

pos(i) çağrıldığında, işlemci closure'ın yapısındaki ilk pointer'ı takip ederek fonksiyonun makine koduna atlar.

Bu kod, gizli bir argüman olarak ortam pointer'ını (environment pointer) alacak şekilde tasarlanmıştır.

Kod içinde sum değişkenine erişim gerektiğinde, bu ortam pointer'ı takip edilir, heap'teki sum'ın adresi bulunur ve oradaki değer okunur/yazılır.

Çöp Toplama (Garbage Collection): Heap üzerinde oluşturulan sum değişkeni, normal bir nesne gibi davranır. Ne zaman ki ona referans veren son closure da (pos değişkeni) kapsam dışı kalır ve yok olursa, Go'nun çöp toplayıcısı (Garbage Collector) heap'teki bu sum alanını temizler. Durumun kalıcılığı bu şekilde sağlanır.


