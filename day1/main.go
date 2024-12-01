package main

import (
	"bufio"
	"log"
	"math"
	"net/http"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	arr1, arr2 := input(getSession())
	ans1 := part1(slices.Clone(arr1), slices.Clone(arr2))
	log.Printf("Answer for part 1 is %v", ans1)
	ans2 := part2(arr1, arr2)
	log.Printf("Answer for part 2 is %v", ans2)
}

func getSession() string{
	if err := godotenv.Load("../.env"); err != nil {
        log.Fatalln("No .env file found")
    }
	session, exists := os.LookupEnv("SESSION")
	if !exists {
		log.Fatalln("Env is absent")
	}
	return session
}

func part1(arr1 []int, arr2 []int) int {
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
	return ans;
}

func part2(arr1 []int, arr2 []int) int {
	mapper := make(map[int]int)
	for _, x := range(arr2) {
		mapper[x]++;
	}
	ans := 0;
	for _, x := range(arr1) {
		ans += x*mapper[x]
	}
	return ans;
}

func input(session string) ([]int, []int) {
	req, err := http.NewRequest("GET", "https://adventofcode.com/2024/day/1/input", nil)
	if err != nil {
		log.Fatalln(err)
	}
	cookie := http.Cookie{
		Name : "session",
		Value : session,
	}
	req.AddCookie(&cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Fetching input, status code: %v", resp.StatusCode)

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