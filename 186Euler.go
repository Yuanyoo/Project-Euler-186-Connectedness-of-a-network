package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type phoneNumber struct {
	leader int
	count  int
}

var allNumbers = []phoneNumber{}
var generatedNumbers = []int{}

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func searchForLeader(id int) int {
	current := id
	for true {
		leader := allNumbers[current].leader
		if leader == current {
			if allNumbers[id].leader != leader {
				allNumbers[id].leader = leader
			}
			return leader
		}
		current = leader
	}
	return 0
}

func mergeFriends(x int, y int) {
	leaderX := searchForLeader(x)
	leaderY := searchForLeader(y)

	if leaderX != leaderY {
		if allNumbers[leaderX].count >= allNumbers[leaderY].count {
			allNumbers[leaderX].count += allNumbers[leaderY].count
			allNumbers[leaderY].leader = leaderX
		} else {
			allNumbers[leaderY].count += allNumbers[leaderX].count
			allNumbers[leaderX].leader = leaderY
		}
	}
}

func generateNumber(j int) int {
	var sNumber int
	if j <= 55 {
		sNumber = (300007*j*j*j - 200003*j + 100003) % 1000000

	} else {
		sNumber = (generatedNumbers[j-56] + generatedNumbers[j-25]) % 1000000
	}
	generatedNumbers = append(generatedNumbers, sNumber)

	return sNumber

}
func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	firstLine := numbers(text)
	number := firstLine[0]
	p := firstLine[1]
	for i := 0; i <= 999999; i++ {
		ph := phoneNumber{
			i,
			1,
		}
		allNumbers = append(allNumbers, ph)
	}
	row := 1
	callsMade := 0
	for allNumbers[searchForLeader(number)].count < 1000000*p/100 {
		caller := generateNumber(2*row - 1)
		called := generateNumber(2 * row)
		if called != caller {

			callsMade++
			mergeFriends(caller, called)

		}
		row++
	}

	fmt.Println(callsMade)
}
