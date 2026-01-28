package main
import "fmt"
func main () {
	const dataTetap = "Data tetap" // const tidak bisa diubah value datanya 
	deklarasiawal:= "Deklarasi awal" // deklarasi awal bisa pakai :=
	var name string = "Hym4ht" // bisa di deklarasi tanpa menyebutkan tipe data jadi langsung var name = "Hym4ht"
	fmt.Println(name)
	name = "Hym4ht ke 2"
	fmt.Println(name)
	fmt.Println(deklarasiawal)
	var (
		firstName = "Azril"
		middleName = "Restu"
		lastName = "Mahinata"
	) // bisa di deklarasi multiple variable
	fmt.Println(firstName, middleName, lastName)
	fmt.Println(dataTetap) // const tidak bisa diubah value datanya  const tidak wajib dipakai atau digunakan
}
