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

func scanFloat(str string) (float64, error) {
	var text string
	var num float64
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(str)
	scanner.Scan()
	text = scanner.Text()
	if text == "" {
		return 0, fmt.Errorf("this is should not be null")
	}

	num, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, fmt.Errorf("this is not number?")
	}
	if num <= 0 {
		return 0, fmt.Errorf("it's should be > 0")
	}
	return num, nil
}

func main() {

	weight, err := scanFloat("How much do you weight (kg)?")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	height, err := scanFloat("How much do you height (m)?")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bmi := calculateBMI(weight, height)
	fmt.Printf("Your BMI is %.2f\n", bmi)
	checkBMI(bmi)
}
