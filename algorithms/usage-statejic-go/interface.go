// package main

// import "fmt"

// /*
// ------------INTERFACE---------------
// benzertipteki objelerin davranış yapılarını tanımlıyourz

// 2 neden için kullanılır
// 1. kod kapsülleme : fonksiyonları kullanıcıyı ilgilendirmediği durumlarda methodları
// kapsüllemek izole etmek için kullanır
// -esnek kullanıma izin vermektedir
// */

// type Kisi struct {
// 	id     int
// 	adi    string
// 	soyadi string
// }

// type Pstgresql struct {
// 	kisi Kisi
// 	// bir method oluşturcaz bu methodu çağırdığımızda
// 	//oluşturulan kişi sql veritabanına kaydedilcek
// }

// func (p *Pstgresql) Kaydet() {
// 	fmt.Println(p.kisi.adi, p.kisi.soyadi, "ait veriler kaydedildi")
// }

// type Oracle struct {
// 	kisi Kisi
// }

// func (o *Oracle) Kaydet() {
// 	fmt.Println(o.kisi.adi, o.kisi.soyadi, "oracle veri tabınına kaydedildi")
// }

// type Kaydedici interface {
// 	Kaydet()
// }

// func vtKaydet(k Kaydedici) {
// 	k.Kaydet()
// }

// func main() {
// 	k := Kisi{
// 		id:     1,
// 		adi:    "yoshi",
// 		soyadi: "takamatsu",
// 	}

// 	// p := Pstgresql{
// 	// 	kisi: k,
// 	// }
//  //--------------------------------------------
// 	// var p Kaydedici = &Pstgresql{
// 	// 	kisi: k,
// 	// }

// 	// var o Kaydedici = &Oracle{
// 	// 	kisi: k,
// 	// }

// 	// p.Kaydet()
// 	// o.Kaydet()
// //---------------------------------------------

// 	m := Oracle{k}
// 	p := Pstgresql{k}

// 	vtKaydet(&m)
// 	vtKaydet(&p)

// }

// ------------------------------------------
// ---------empty-interface------------------

// package main

// import (
// 	"fmt"
// )

// func main(){
// 	var hersey interface{} = 1.5
// 	birString, ok := hersey.(string)// belirttiğim herşeyin tipi mi gibi düşünebiliriz
// 	if !ok { //type assertion
// 		fmt.Println("tip hatası")
// 	}

// 	fmt.Println("merhaba", birString)

// }

package main

import (
	"fmt"
)

type kisi struct {
	adi    string
	soyadi string
}

func veriTipiGoster(a interface{}) {

	switch v := a.(type) {
	case string:
		fmt.Println("string tipidir",v)
	case int:
		fmt.Println("int tipidir",v)
	case float64:
		fmt.Println("float64 tipidir",v)
	case kisi:
		fmt.Println("kisi tipidir",v)
	default:
		fmt.Println("bu verinin tipi yoktur")
	}
}

func main() {
	veriTipiGoster(10)
	veriTipiGoster("ad")
	veriTipiGoster(kisi{adi: "s", soyadi: "d"})


}
