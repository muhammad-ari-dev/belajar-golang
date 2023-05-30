package main

// "golang/management"

// //pointer method
// type Gamer struct {
// 	Name  string
// 	Games []string
// }

// func (gamer *Gamer) Addgame(game string) {
// 	gamer.Games = append(gamer.Games, game)
// }

// func main() {
// 	gamer := Gamer{Name: "Ari"}
// 	gamer.Addgame("fifa")
// 	gamer.Addgame("konami")
// 	for _, game := range gamer.Games {
// 		fmt.Println(game)
// 	}

// }

// type Student struct {
// 	ID   int
// 	Name string
// 	GPA  float64
// }

// func (student *Student) graduation() {
// 	student.Name = student.Name + " S.Kom"

// }

// func main() {

// 	student := Student{1, "ari aja", 3.45}
// 	//pemanggilan funcion pointer
// 	graduation(&student)

// 	fmt.Println(student.Name)

// 	// pemanggilan funcion method
// 	student.graduation()

// 	fmt.Println(student.Name)

// }

// // pointer funcion
// func graduation(student *Student) {
// 	student.Name = student.Name + " S.Kom"
// 	fmt.Println(student.Name)

// }

// 	// pointer di funcion
// 	number := 5
// 	fmt.Println("alamat memory awal:", &number)
// 	fmt.Println("nilai awal :", number)

// 	change(&number, 100)

// 	fmt.Println("nilai akhir :", number)
// 	fmt.Println("alamat memory :", &number)

// }

// func change(old *int, new int) {
// 	*old = new
// 	fmt.Println("alamat memory :", old)
// 	fmt.Println("data dalam funcion :", *old)

// //pointer jika deklarasi make var type data pengambilan alamat di kasih bintang
// var nilaia int = 10
// var nilaib *int = &nilaia

// fmt.Println(nilaia)
// fmt.Println(*nilaib)

// nilaia = 20
// fmt.Println(nilaia)
// fmt.Println(*nilaib)

// pointer sederhana
// nilaiA := 5
// nilaiB := &nilaiA

// fmt.Println(nilaiA)
// fmt.Println(*nilaiB)

// *nilaiB = 10
// fmt.Println(*nilaiB)
// fmt.Println(nilaiA)

// student := management.Students{1, "muhammad", "Ferdiansyah", "bandar lampung", true}
// student2 := management.Students{1, "Mela Nur", "Hidayah", "Tulang bawang", true}

// // pemanggilan method
// fmt.Println(student.Display())

// fmt.Println(student2.Display())

// data_student := []management.Students{student, student2}

// ekskul := management.Ekskul{"Futsal", student, data_student, true}

// ekskul.Displayekskul()

// belajar struct
// type Students struct {
// 	ID            int
// 	Nama_Depan    string
// 	Nama_Belakang string
// 	Alamat        string
// 	IsActivate    bool
// }

// type Ekskul struct {
// 	Name         string
// 	Admin        Students
// 	Data_student []Students
// 	Isavailable  bool
// }

// // method
// func (student Students) display() string {
// 	return fmt.Sprintf("name : %s %s, alamat: %s", student.Nama_Depan, student.Nama_Belakang, student.Alamat)

// }
// func (coba Ekskul) displayekskul() {
// 	fmt.Printf("nama Ekskul: %s", coba.Name)
// 	fmt.Println(" ")
// 	fmt.Printf("admin: %s", coba.Admin.Nama_Depan)
// 	fmt.Println(" ")
// 	fmt.Printf("JUmlah Data : %d", len(coba.Data_student))
// 	fmt.Println(" ")
// 	fmt.Println("Nama Data :")

// 	for _, data := range coba.Data_student {
// 		fmt.Println(data.Nama_Depan)

// 	}
// }

// func displaystudent(student Students) string {
// return fmt.Sprintf("name : %s %s, alamat: %s", student.Nama_Depan, student.Nama_Belakang, student.Alamat)

// }

// //menampilkan funcion ekskul
// displayekskul(ekskul)

// // menampilkan funcion
// displaystudent1 := displaystudent(student)
// displaystudent2 := displaystudent(student2)

// fmt.Println(displaystudent1)
// fmt.Println(displaystudent2)

// func displayekskul(ekskul Ekskul) {
// 	fmt.Printf("nama Ekskul: %s", ekskul.Name)
// 	fmt.Println(" ")
// 	fmt.Printf("admin: %s", ekskul.Admin.Nama_Depan)
// 	fmt.Println(" ")
// 	fmt.Printf("JUmlah Data : %d", len(ekskul.Data_student))
// 	fmt.Println(" ")
// 	fmt.Println("Nama Data :")

// 	for _, data := range ekskul.Data_student {
// 		fmt.Println(data.Nama_Depan)

// 	}

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
