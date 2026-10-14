// ---------------------------------------------------------------
//
//	Подсчет индекса массы тела (ИМТ)
//
// ---------------------------------------------------------------

package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const Version = "0.1.2"

// Интерпретация результатов для взрослых
func HelpResult() {
	fmt.Println()
	fmt.Println("Менее 16: Выраженный дефицит массы тела.")
	fmt.Println("16 – 18,5: Недостаточная масса тела.")
	fmt.Println("18,5 – 25: Норма (оптимальный вес).")
	fmt.Println("25 – 30: Избыточная масса тела (предожирение).")
	fmt.Println("30 – 35: Ожирение первой степени.")
	fmt.Println("35 – 40: Ожирение второй степени.")
	fmt.Println("40 и более: Ожирение третьей степени.")
}

// Чтение и проверка числового значения
func ReadFloat(scanner *bufio.Scanner, prompt string) (float64, error) {
	fmt.Print(prompt)

	if !scanner.Scan() {
		return 0, fmt.Errorf("не удалось прочитать ввод")
	}

	input := strings.TrimSpace(scanner.Text())

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("ожидается ввод числового значения")
	}

	return value, nil
}

func main() {
	showVersion := flag.Bool("version", false, "Показать версию программы")
	flag.Parse()

	if *showVersion {
		fmt.Println("IMT Calculator", Version)
		return
	}

	fmt.Println("IMT Calculator v" + Version)
	fmt.Println("")

	scanner := bufio.NewScanner(os.Stdin)

	userHeight, err := ReadFloat(scanner, "Введите рост в см.: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if userHeight <= 0 {
		fmt.Println("Ошибка: рост должен быть больше нуля.")
		return
	}

	userWeight, err := ReadFloat(scanner, "Введите вес в кг.: ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if userWeight <= 0 {
		fmt.Println("Ошибка: вес должен быть больше нуля.")
		return
	}

	imt := userWeight / math.Pow(userHeight/100, 2)

	fmt.Printf("ИМТ = %.2f\n", imt)

	HelpResult()
}
