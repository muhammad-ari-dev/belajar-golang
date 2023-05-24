package main

import (
	"fmt"
)

// belajar struct
type Students struct {
	ID            int
	Nama_Depan    string
	Nama_Belakang string
	Alamat        string
	IsActivate    bool
}

func main() {

	student := Students{1, "muhammad", "Ferdiansyah", "bandar lampung", true}
	student2 := Students{1, "Mela Nur", "Hidayah", "Tulang bawang", true}

	displaystudent1 := displaystudent(student)
	displaystudent2 := displaystudent(student2)

	fmt.Println(displaystudent1)
	fmt.Println(displaystudent2)

}

func displaystudent(student Students) string {
	return fmt.Sprintf("name : %s %s, alamat: %s", student.Nama_Depan, student.Nama_Belakang, student.Alamat)

}

// student := students{ID: 1, Nama_Depan: "Muhammad Ari", Nama_Belakang: "Ferdiansyah", Alamat: "Bandar Lampung", IsActivate: true}

// student := students{}
// student.ID = 1
// student.Nama_Depan = "Muhammad Ari"
// student.Nama_Belakang = "Ferdiansyah"
// student.Alamat = "Bandar Lampung"
// student.IsActivate = true

// student2 := students{}
// student2.ID = 2
// student2.Nama_Depan = "Mela Nur"
// student2.Nama_Belakang = "Nurhidayah"
// student2.Alamat = "Tulang Bawang"
// student2.IsActivate = false

// fmt.Println(student)
// // fmt.Println(student2)
// fmt.Println(student.Nama_Depan, student.Nama_Belakang)
// fmt.Println(student2.Nama_Depan)

// score := []int{10, 5, 8, 9, 7}
// total := sum(score)

// fmt.Println(total)

// 	result, err := calculate(10, 2, "+")
// 	if err != nil {
// 		fmt.Println("terjadi kesalahan")
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(result)

// var total int

// func sum(numbers []int) int {

// 	for _, value := range numbers {
// 		total = total + value

// 	}
// 	return total
// }

// func calculate(number, numberTwo int, operation string) (int, error) {

// 	var result int
// 	var errorResult error
// 	switch operation {
// 	case "+":
// 		result = number + numberTwo

// 	case "-":
// 		result = number - numberTwo

// 	case "*":
// 		result = number * numberTwo

// 	case "/":
// 		result = number / numberTwo
// 	default:
// 		errorResult = errors.New("Unknown Operation")
// 	}

// 	return result, errorResult
// }

//belajar funcion
// 	sentence := printmyresult("saya sedang")
// 	fmt.Println(sentence)

// 	hasil := tambahan(20, 30)
// 	fmt.Println(hasil)

// 	luas, keliling := luas(30, 20)
// 	fmt.Println(luas, keliling)
// }

// func printmyresult(sentence string) string {
// 	newsentence := sentence + "belajar golang"
// 	return newsentence

// }

// func tambahan(number, nummber2 int) int {

// 	return number + nummber2

// }

// func luas(panjang, lebar int) (int, int) {
// 	total := panjang * lebar
// 	keliling := 2 * (panjang + lebar)

// 	return total, keliling

//hitung rata-rata

// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

// var total int

// for _, score := range scores {
// 	total = total + score
// }

// banyak := len(scores)

// average := float64(total) / float64(banyak)

// fmt.Println(average)

// mencari nilai diatas 90

// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

// var highscores []int

// for _, highscore := range scores {
// 	if highscore >= 90 {
// 		highscores = append(highscores, highscore)
// 	}
// }

// fmt.Println(highscores)

//map didalam slice
// students := []map[string]string{
// 	{"name": "Ari", "score": "A"},
// 	{"name": "Mela", "score": "B"},
// 	{"name": "Riski", "score": "C"},
// }

// for _, student := range students {

// 	fmt.Println(student["name"], "-----", student["score"])

// }

// var myMap map[string]int
// myMap = map[string]int{}

// myMap["ruby"] = 10
// myMap["golang"] = 12
// myMap["javascript"] = 13

// value, isavailable := myMap["phyton"] //mencari nilai dalam map

// fmt.Println(value)
// fmt.Println(isavailable)

// var myMap map[string]int
// myMap = map[string]int{}

// myMap["ruby"] = 10
// myMap["golang"] = 12
// myMap["javascript"] = 13

// for key, value := range myMap {
// 	fmt.Println("key =", key, "value=", value)
// }

// fmt.Println("===============")

// delete(myMap, "ruby")

// for key, value := range myMap {
// 	fmt.Println("key =", key, "value=", value)
// }

// var gamingconsole []string

// gamingconsole = append(gamingconsole, "xborg")
// gamingconsole = append(gamingconsole, "nintendo")
// gamingconsole = append(gamingconsole, "ps3")

// fmt.Println(gamingconsole)

// for _, console := range gamingconsole {
// 	fmt.Println(console)
// }

// title := [...]string{"ruby", "java", "kotlin", "javascript", "golang", "lagi"}

// fmt.Println(title)

// length := len(title)

// fmt.Println(length)

// for index, ngambil := range title {
// 	fmt.Println("index=", index, "ngambil=", ngambil)
// }

// title := "Golang The best Language"

// for index, letter := range title {

// 	latterstring := string(letter)

// 	switch latterstring {
// 	case "a", "i", "u", "e", "o":
// 		fmt.Println("index=", index, "letter=", latterstring)

// 	}

// }

// if index%2 == 0 {
// 	fmt.Println("index=", index, "letter=", string(letter))
// }
