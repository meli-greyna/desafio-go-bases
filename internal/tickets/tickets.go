package tickets

import (
	"errors"
	"strconv"
	"strings"
)

type Ticket struct {
	Id          int
	Name        string
	Email       string
	Destination string
	Time        string
	Price       int
}

// ejemplo 1
func GetTotalTickets(tickets []Ticket, destination string) int {
	count := 0
	for _, ticket := range tickets {
		if ticket.Destination == destination {
			count++
		}
	}

	return count
}

// ejemplo 2
func GetCountByPeriod(tickets []Ticket, time string) (int, error) {
	count := 0
	for _, ticket := range tickets {
		hora, err := strconv.ParseInt(strings.Split(ticket.Time, ":")[0], 10, 64)
		if err != nil {
			return 0, err
		}

		switch time {
		case "madrugada":
			if hora >= 0 && hora < 7 {
				count++
			}
		case "mañana":
			if hora >= 7 && hora < 13 {
				count++
			}
		case "tarde":
			if hora >= 13 && hora < 20 {
				count++
			}
		case "noche":
			if hora >= 20 && hora < 24 {
				count++
			}
		default:
			return 0, errors.New("time passed by argument is not a valid time")
		}
	}

	return count, nil
}

// ejemplo 3
func AverageDestination(tickets []Ticket, destination string) (float64, error) {
	total := GetTotalTickets(tickets, destination)
	if len(tickets) == 0 {
		return 0, errors.New("no tickets")
	}
	return float64(total) / float64(len(tickets)) * 100.0, nil
}
