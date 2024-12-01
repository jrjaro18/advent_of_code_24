package main

import (
	"bufio"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"github.com/joho/godotenv"
)

func main() {
	Part2()
}

func getSession() string{
	if err := godotenv.Load("../.env"); err != nil {
        log.Print("No .env file found")
    }
	session, exists := os.LookupEnv("SESSION")
	if !exists {
		log.Fatalln("Env is absent")
	}
	return session
}

func Part1() {
 	arr1, arr2 := input()
	
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]	
	})

	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]	
	})
	ans := 0
	for i := 0; i < len(arr1); i++ {
		ans += int(math.Abs(float64(arr1[i] - arr2[i])));
	}
	log.Println(ans)
}

func Part2() {
	arr1, arr2 := input()
	map_of_2 := make(map[int]int)
	for _, x := range(arr2) {
		map_of_2[x]++;
	}
	ans := 0;
	for _, x := range(arr1) {
		ans += x*map_of_2[x]
	}
	log.Println(ans)
}

func input() ([]int, []int) {

	req, err := http.NewRequest("GET", "https://adventofcode.com/2024/day/1/input", nil)
	if err != nil {
		log.Fatalln(err)
	}
	cookie := http.Cookie{
		Name : "session",
		Value : getSession(),
	}
	req.AddCookie(&cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.StatusCode)

	x := resp.Body;
	scanner := bufio.NewScanner(x)

	var arr1 []int
	var arr2 []int

	for scanner.Scan() {
		resp := scanner.Text()
		arr := strings.Split(resp, "   ")
			
		item1, err := strconv.Atoi(arr[0]);
		if err != nil {
			log.Fatalln(err)
		}
		arr1 = append(arr1, item1)

		item2, err := strconv.Atoi(arr[1]);
		if err != nil {
			log.Fatalln(err)
		}
		arr2 = append(arr2, item2)
	}
	return arr1, arr2
}