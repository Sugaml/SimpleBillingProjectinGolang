package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var count int = 1
var discount int = 0
var total int = 0
var asum int = 0
var servicecharge int = 0
var grand int = 0
var tax int = 0
var netTotal int = 0

type Coffee struct {
	id       int
	itemName string
	price    int
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

var iid int
var qun int
var status int

func controlunit() {
	f, err := os.Open("coffeemenu.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)

	}
	var coffee []Coffee
	for _, record := range records {
		coffee = append(coffee, Coffee{
			id:       atoi(record[0]),
			itemName: record[1],
			price:    atoi(record[2]),
		})
	}
	file, err := os.Create("store.csv")
	if err != nil {
		panic(err)
	}
	for true {

		fmt.Println("Enter Item ID : ")
		fmt.Scanln(&iid)
		fmt.Println("Enter orderd quantity : ")
		fmt.Scanln(&qun)
		if iid >= 1 && iid <= 72 {
			file.WriteString(coffee[iid].itemName)
			file.WriteString(",")
			file.WriteString(strconv.Itoa(coffee[iid].price))
			file.WriteString(",")
			file.WriteString(strconv.Itoa(qun))
			file.WriteString("\n")
		}
		fmt.Println("If item is finished Enter '0' ")
		fmt.Scanln(&status)
		if status == 0 {
			file.Close()
			break
		}
	}
}
func calculation() {
	f, err := os.Open("store.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)

	}
	var coffee []Coffee
	for _, record := range records {
		coffee = append(coffee, Coffee{
			id:       atoi(record[0]),
			itemName: record[1],
			price:    atoi(record[2]),
		})
	}
}

func calculationAndSimpleFormat() {
	f, err := os.Open("store.csv")
	if err != nil {
		panic(err)
	}
	file, er := os.Create("DemoBill.txt")
	if er != nil {
		panic(er)
	}
	var sum int = 0
	currentTime := time.Now()
	scanner := bufio.NewScanner(f)
	file.WriteString("----------------------------------------------------------")
	file.WriteString("           Demo Cafe           \n")
	file.WriteString("           Kathmandu           ")
	file.WriteString("\n")
	file.WriteString("-----------------------------------------------\n")
	file.WriteString("           Invoice     \t\t\tDate   ")
	file.WriteString(currentTime.Format("01-02-2006"))
	file.WriteString("\n-----------------------------------------------\n")
	file.WriteString("Invoic Number:\n")
	file.WriteString("101")
	file.WriteString("\n")
	file.WriteString("-----------------------------------------------------\n")
	file.WriteString("S.N.\t Particular \t Price \tQunatity\tTotal\n")
	file.WriteString("-----------------------------------------------------\n")

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		price, _ := strconv.Atoi(items[1])
		quntity, _ := strconv.Atoi(items[2])
		total = (price * quntity)
		sum += total
		file.WriteString(strconv.Itoa(count))
		file.WriteString("  ")
		file.WriteString(items[0])
		file.WriteString("\t")
		file.WriteString(items[1])
		file.WriteString("\t")
		file.WriteString(items[2])
		file.WriteString("  ")
		file.WriteString(strconv.Itoa(total))
		file.WriteString("\n")
		count++
	}
label:
	fmt.Println("\nWhat Discount % Available For This Bill? : : ")
	fmt.Scanln(&discount)
	if discount >= 0 && discount <= 30 {
		discount = (sum * discount) / 100
	} else {
		fmt.Println("Enter valid Discount Percentage between 0 to 30 percentage")
		goto label
	}
	asum = (sum - discount)
	servicecharge = (asum * 10) / 100 //service charg
	grand = asum + servicecharge      //afterservice charge Total
	tax = ((grand) * 13) / 100        //For tax
	netTotal = asum + tax + servicecharge
	file.WriteString("-----------------------------------------------------\n")
	file.WriteString("\t\t\tTotal    ")
	file.WriteString(strconv.Itoa(sum))
	file.WriteString("\n")
	file.WriteString("\t\t\tDiscount\t  ")
	file.WriteString(strconv.Itoa(discount))
	file.WriteString("\n")
	file.WriteString("\t\tService Charge\t")
	file.WriteString(strconv.Itoa(servicecharge))
	file.WriteString("\n")
	file.WriteString("\t\tGrand Total   ")
	file.WriteString(strconv.Itoa(grand))
	file.WriteString("\n")
	file.WriteString("\t\tTAX\t    ")
	file.WriteString(strconv.Itoa(tax))
	file.WriteString("\n")
	file.WriteString("\t\tNet Amount\t")
	file.WriteString(strconv.Itoa(netTotal))
	file.WriteString("\n-----------------------------------------------------\n")
	file.WriteString("\t\t\tCashier\n")
	file.WriteString("\t\t  Name of Cashier")
	fmt.Println("Bill created successfully")
	//
	fmt.Println("Bill is ready in 'DemoBill.txt' File")
	fmt.Println("Thank you")

}

func main() {

	controlunit()
	calculationAndSimpleFormat()
}
