package management

import "fmt"

type Students struct {
	ID            int
	Nama_Depan    string
	Nama_Belakang string
	Alamat        string
	IsActivate    bool
}

type Ekskul struct {
	Name         string
	Admin        Students
	Data_student []Students
	Isavailable  bool
}

// method
func (student Students) Display() string {
	return fmt.Sprintf("name : %s %s, alamat: %s", student.Nama_Depan, student.Nama_Belakang, student.Alamat)

}
func (coba Ekskul) Displayekskul() {
	fmt.Printf("nama Ekskul: %s", coba.Name)
	fmt.Println(" ")
	fmt.Printf("admin: %s", coba.Admin.Nama_Depan)
	fmt.Println(" ")
	fmt.Printf("JUmlah Data : %d", len(coba.Data_student))
	fmt.Println(" ")
	fmt.Println("Nama Data :")

	for _, data := range coba.Data_student {
		fmt.Println(data.Nama_Depan)

	}
}
