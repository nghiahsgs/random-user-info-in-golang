package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var listName1 []string = []string{"An", "Uc", "Uat", "Dam", "Dao", "Dinh", "Doan", "An", "Banh", "Bach", "Cao", "Chau", "Chu", "Chu", "Chung", "Duu", "Diep", "Doan", "Giang", "Ha", "Han", "Kieu", "Kim", "Lam", "Luong", "Luu", "Lac", "Luc", "La", "Lieu", "Ma", "Mac", "Mach", "Mai", "Ngu", "Nghiem", "Phi", "Pho", "Phung", "Quach", "Quang", "Quyen", "To", "Ton", "Ta", "Tong", "Thai", "Sai", "Than", "Thach", "Thao", "Thuy", "Thi", "Tieu", "Truong", "Tram", "Trinh", "Trang", "Trieu", "Van", "Vinh", "Vuong", "Vuu", "Nguyen", "Tran", "Le", "Pham", "Huynh", "Hoang", "Phan", "Vu", "Vo", "Dang", "Bui", "Do", "Ho", "Ngo", "Duong", "Ly", "Nguyen", "Tran", "Le", "Pham", "Huynh", "Hoang", "Phan", "Vu", "Vo", "Dang", "Bui", "Do", "Ho", "Ngo", "Duong", "Ly", "Nguyen", "Nguyen", "Nguyen", "Nguyen", "Nguyen", "Nguyen", "Nguyen", "Nguyen", "Nguyen", "Nguyen", "Nguyen", "Nguyen", "Nguyen", "Nguyen", "Nguyen", "Nguyen", "Nguyen", "Tran", "Le", "Pham", "Huynh", "Hoang", "Phan", "Vu", "Vo", "Dang", "Bui", "Do", "Ho", "Ngo", "Duong", "Ly", "Nguyen", "Tran", "Le", "Pham", "Huynh", "Hoang", "Phan", "Vu", "Vo", "Dang", "Bui", "Do", "Ho", "Ngo", "Duong", "Ly"}
var listName2 []string = []string{"Bao An", "Binh An", "Сng An", "Duy An", "Khanh An", "Nam An", "Phuoc An", "Thanh An", "The An", "Thien An", "Truong An", "Viet An", "Xuan An", "Cong An", " An", "Gia An", "Hoang An", "Minh An", "Phu An", "Thanh An", "Thien An", "Thien An", "Vinh An", "Ngoc An", "Chi Anh", " Anh", "Duong Anh", "Gia Anh", "Hung Anh", "Huy Anh", "Minh Anh", "Quang Anh", "Quoc Anh", "The Anh", "Thieu Anh", "Thuan Anh", "Trung Anh", "Tuan Anh", "Tung Anh", "Tuong Anh", "Viet Anh", "Vu Anh", "Ho Bac", "Hoai Bac", "Gia Bach", "Cong Bang", " Bang", "Hai Bang", "Yen Bang", "Chi Bao", " Bao", "Duy Bao", "Gia Bao", "Huu Bao", "Nguyen Bao", "Quoc Bao", "Thieu Bao", "Tieu Bao", " Binh", "Gia Binh", "Hai Binh", "Hoa Binh", "Huu Binh", "Khanh Binh", "Kien Binh", "Kien Binh", "Phu Binh", "Quoc Binh", "Tan Binh", "Tat Binh", "Thai Binh", "The Binh", "Xuan Binh", "Yen Binh", "Quang Buu", "Thien Buu", "Khai Ca", "Gia Can", "Duy Can", "Gia Can", "Huu Canh", "Gia Canh", "Huu Canh", "Minh Canh", "Ngoc Canh", "Duc Cao", "Xuan Cao", "Bao Chan", "Bao Chau", "Huu Chau", "Phong Chau", "Thanh Chau", "Tuan Chau", "Tung Chau", "Dinh Chien", "Manh Chien", "Minh Chien", "Huu Chien", "Huy Chieu", "Truong Chinh", " Chinh", "Trong Chinh", "Trung Chinh", "Viet Chinh", "Tuan Chuong", "Minh Chuyen", "An Co", "Chi Cong", "Thanh Cong", "Xuan Cung", "Huu Cuong", "Manh Cuong", "Duy Cuong", "Viet Cuong", "Ba Cuong", " Cuong", "Duy Cuong", "Hung Cuong", "Huu Cuong", "Kien Cuong", "Manh Cuong", "Ngoc Cuong", "Phi Cuong", "Phuc Cuong", "Thinh Cuong", "Viet Cuong", "Ngoc Dai", "Quoc Сi", "Minh Dan", "The Dan", "Minh Сn", "Nguyen Сn", "Sy Сn", "Hai Сng", "Hong Сng", "Minh Danh", "Ngoc Danh", "Quang Danh", "Thanh Danh", "Hung Сo", "Thanh Сo", "Binh Сt", "Сng Сt", "Huu Сt", "Minh Сt", "Quang Сt", "Quang Сt", "Thanh Сt", "Сc Di", "Vinh Dieu", "Ngoc Ran", "Thanh Ran", "Thanh Doanh", "The Doanh", "Quang Dong", "Tu Rng", "Vien Rng", "Lam Rng", "Bach Du", "Thuy Du", "Hong Duc", "Anh ", "Gia ", "Kien ", "Minh ", "Quang ", "Tai ", "Thai ", "Thien ", "Thien ", "Tien ", "Trung ", "Tuan ", "Hoang Due", "Anh Dung", "Chi Dung", "Hoang Dung"}
var listName3 []string = []string{"Dieu Ai", "Kha Ai", "Ngoc Ai", "Hoai An", "Hue An", "Minh An", "Phuong An", "Thanh An", "Hai An", "Hue An", "Bao Anh", "Diep Anh", "Dieu Anh", "Hai Anh", "Hong Anh", "Huyen Anh", "Kieu Anh", "Kim Anh", "Lan Anh", "Mai Anh", "Minh Anh", "My Anh", "Ngoc Anh", "Nguyet Anh", "Nhu Anh", "Phuong Anh", "Que Anh", "Quynh Anh", "Thuc Anh", "Thuy Anh", "Thuy Anh", "Tram Anh", "Trang Anh", "Tu Anh", "Tuyet Anh", "Van Anh", "Yen Anh", "Kim Anh", "Ngoc Anh", "Nguyet Anh", "Nhat Anh", "Bang Bang", "Le Bang", "Tuyet Bang", "Nhu Bao", "Gia Bao", "Xuan Bao", "Ngoc Bich", "An Binh", "Thai Binh", "Son Ca", "Ngoc Cam", "Nguyet Cam", "Thi Cam", "Bao Chau", "Bich Chau", "Diem Chau", "Hai Chau", "Hoan Chau", "Hong Chau", "Linh Chau", "Loan Chau", "Ly Chau", "Mai Chau", "Minh Chau", "Tran Chau", "Diep Chi", "Diem Chi", "Hanh Chi", "Khanh Chi", "Kim Chi", "Lan Chi", "Le Chi", "Linh Chi", "Mai Chi", "Phuong Chi", "Que Chi", "Quynh Chi", "Bich Chieu", "Hoang Cuc", "Kim Cuong", "Trang Сi", "Tam Dan", "Thanh Dan", "Linh Сn", "Quynh Dao", "Anh Сo", "Bich Сo", "Hong Сo", "Ngoc Сo", "Thuc Сo", "Truc Сo", "An Di", "Thien Di", "Hong Diem", "Kieu Diem", "Phuong Diem", "Thuy Diem", "Bich Diep", "Hong Diep", "Ngoc Diep", "Huyen Dieu", "Tam Ran", "Thuc Ran", "Hanh Dung", "Kieu Dung", "Kim Dung", "My Dung", "Nghi Dung", "Ngoc Dung", "Phuong Dung", "Quynh Dung", "Thuy Dung", "Anh Duong", "Chieu Duong", "Thuy Duong", "Hai еong", "Bich Duyen", "Ky Duyen", "Linh Duyen", "Minh Duyen", "My Duyen", "Thu Duyen", "Ha Giang", "Hoai Giang", "Huong Giang", "Kieu Giang", "Linh Giang", "Phuong Giang", "Quynh Giang", "Thanh Giang", "Thien Giang", "Thu Giang", "Thuy Giang", "Hong Giang", "Tra Giang", "Khanh Giao", "Quynh Giao", "Bao Ha", "Bich Ha", "Hoang Ha", "Hong Ha", "Khanh Ha", "Lam Ha", "Linh Ha", "Mai Ha", "Minh Ha", "Ngan Ha", "Ngoc Ha", "Nguyet Ha", "Nhat Ha", "Quynh Ha", "Thai Ha", "Thanh Ha", "Thu Ha", "Thuy Ha", "Van Ha", "Viet Ha", "An Ha", "Mai Ha", "Ngoc Ha", "Nhat Ha", "Bich Hai", "Bao Han", "Gia Han", "Ngoc Han", "Tuyet Han", "Bich Hang", "Diem Hang", "Dieu Hang", "Minh Hang", "Thanh Hang", "Thu Hang", "Thuy Hang", "Bich Hanh", "Cam Hanh", "Diem Hanh", "Hieu Hanh", "Hong Hanh", "Kieu Hanh", "Minh Hanh", "My Hanh", "Phuong Hanh", "Thuy Hanh", "Bich Hao", "Thanh Hao", "Bich Hau", "Thu Hau", "Bich Hien", "Dieu Hien", "Mai Hien", "Minh Hien", "Ngoc Hien", "Phuong Hien", "Tam Hien", "Thanh Hien"}

func getTimeStamp() int {
	t := time.Now()
	kq, _ := strconv.Atoi(t.Format("20060102150405"))
	return kq
}

// list func get info
func randomName() string {
	rand.Seed(int64(getTimeStamp()))
	return listName1[rand.Intn(len(listName1))]
}
func randomFullName(typeChoice int) string {
	rand.Seed(int64(getTimeStamp()))
	if typeChoice == 0 {
		return listName2[rand.Intn(len(listName2))]
	} else {
		return listName3[rand.Intn(len(listName3))]
	}
}

func randomUsernameName(typeChoice int) string {
	rand.Seed(int64(getTimeStamp()))
	fullName := randomFullName(typeChoice)
	fullName = strings.Replace(fullName, " ", "", -1)
	return fullName + strconv.Itoa(rand.Intn(10000))
}
func randomPhone() string {
	rand.Seed(int64(getTimeStamp()))
	phone := "098"
	for i := 0; i < 7; i++ {
		ext := rand.Intn(10)
		phone += strconv.Itoa(ext)
	}
	return phone
}
func main() {
	fmt.Println(randomName())
	fmt.Println(randomFullName(0))
	fmt.Println(randomUsernameName(0))
	fmt.Println(randomPhone())

}
