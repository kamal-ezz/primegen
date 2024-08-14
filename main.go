package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)


func isPrime(number int) bool {
	s := int(math.Sqrt(float64(number)))
	if number == 0 || number == 1 {
		return false
	}	

	for i := 2; i <= s; i++ {
		if number % i == 0 {
			return false
		}
	}
	return true
}


func firstPrime(start int, end int) (int, error) {
	for i := start; i <= end; i++ {	
		if(isPrime(i)){
			return i, nil
		}
	}

	return 0, errors.New("cannot find prime in the specific range")
}



func printPrimes(start int, end int){
	fileName := fmt.Sprintf("%d - %d.txt", start, end)
	f, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	j := 1
	p, err := firstPrime(start, end)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := start; i <= end; i++ {	
		if(isPrime(i)){
			_, err = f.WriteString(fmt.Sprintf("%d, %d, %d\n", j, i,i - p))

			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}

			j++
			p = i
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main(){
	if len(os.Args) < 3 {
		fmt.Println("Please provide the starting and ending point")
		return
	}

	i, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Please provide a valid number")
	}

	j, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Please provide a valid number")
	}

	start := time.Now()
	printPrimes(i,j)
	elapsed := time.Since(start)


	fmt.Println("Finished after", elapsed)
}