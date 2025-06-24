package main

import (
	"bufio" // สำหรับการอ่านข้อมูลแบบมีบัฟเฟอร์ (buffered input)
	"fmt"   // สำหรับการพิมพ์ข้อความ (format)
	"os"    // สำหรับการเข้าถึง Standard Input/Output
	"strconv"
)

func calculateBMI(weight, height float64) float64 {
	bmi := weight / (height * height)
	return bmi
}

func checkBMI(bmi float64) {
	if bmi < 18.5 {
		fmt.Println("You are Underweight")
	} else if bmi < 25.0 {
		fmt.Println("You are Normalweight")
	} else if bmi < 30.0 {
		fmt.Println("You are Overweight")
	} else {
		fmt.Println("You are Obese")
	}
}

func main() {
	var weight, height string
	// 1. สร้างตัวอ่าน (Scanner) จาก os.Stdin
	//    bufio.NewScanner(os.Stdin) จะสร้าง Scanner ที่พร้อมจะอ่านข้อมูลจากคีย์บอร์ด
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("How much do you weight (kg)?") // พิมพ์ข้อความแจ้งผู้ใช้ให้กรอกข้อมูล

	// 2. อ่านข้อมูลหนึ่งบรรทัดจาก Standard Input
	//    scanner.Scan() จะอ่านบรรทัดถัดไปจาก input และคืนค่าเป็น true หากอ่านสำเร็จ
	//    หากไม่มีข้อมูลให้ประมวลผลอีกหรือเกิดข้อผิดพลาด จะคืนค่า false
	if scanner.Scan() {
		// 3. ดึงข้อความที่อ่านได้
		//    scanner.Text() จะคืนค่า string ของบรรทัดที่อ่านได้ล่าสุด
		weight = scanner.Text()
		fmt.Printf("Your weight is %s kg\n", weight) // พิมพ์ข้อความทักทาย
	} else {
		// 4. จัดการข้อผิดพลาด (ถ้ามี)
		//    scanner.Err() จะคืนค่า error ที่เกิดขึ้นระหว่างการสแกน
		if err := scanner.Err(); err != nil {
			fmt.Println("Error :", err)
		}
	}

	weightf, err := strconv.ParseFloat(weight, 64)
	//ถ้าไม่สามารถแปลงค่าเป็น Float ได้จะส่ง Error กลับไป
	if err != nil {
		fmt.Println("That's strange. Why do you weight aren't number?")
		fmt.Println("The system is suspicious and shutdown")
		os.Exit(1)
	}
	if weightf <= 0 {
		fmt.Println("That's strange. Why do you weight less than 0?")
		fmt.Println("The system is suspicious and shutdown")
		os.Exit(1)
	}

	fmt.Println("How much do you height (m)?")
	if scanner.Scan() {
		height = scanner.Text()
		fmt.Printf("You height is %s m\n", height)
	} else {
		if err := scanner.Err(); err != nil {
			fmt.Println("Error :", err)
		}
	}

	heightf, err := strconv.ParseFloat(height, 64)
	if err != nil {
		fmt.Println("That's strange. Why do you height aren't number?")
		fmt.Println("The system is suspicious and shutdown")
		os.Exit(1)
	}
	if heightf <= 0 {
		fmt.Println("That's strange. Why do you height less than 0?")
		fmt.Println("The system is suspicious and shutdown")
		os.Exit(1)
	}

	bmi := calculateBMI(weightf, heightf)
	fmt.Printf("Your BMI is %f\n", bmi)
	checkBMI(bmi)
}
