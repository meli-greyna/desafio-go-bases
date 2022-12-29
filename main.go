package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func parseTicket(ticket []string) tickets.Ticket {
	id, _ := strconv.ParseInt(ticket[0], 10, 64)
	price, _ := strconv.ParseInt(ticket[5], 10, 64)
	return tickets.Ticket{
		Id:          int(id),
		Name:        ticket[1],
		Email:       ticket[2],
		Destination: ticket[3],
		Time:        ticket[4],
		Price:       int(price),
	}
}

func parseCSV(filename string) ([]tickets.Ticket, error) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	ticketData := csv.NewReader(strings.NewReader(string(raw)))
	tickets := []tickets.Ticket{}

	for {
		rawTicket, eof := ticketData.Read()
		if eof != nil {
			break
		}

		tickets = append(tickets, parseTicket(rawTicket))
	}

	return tickets, nil
}

func main() {
	data, err := parseCSV("tickets.csv")
	if err != nil {
		panic(err)
	}

	brazil := tickets.GetTotalTickets(data, "Brazil")
	mornings, err := tickets.GetCountByPeriod(data, "mañana")
	if err != nil {
		fmt.Println(err)
	}
	avg, err := tickets.AverageDestination(data, "Brazil")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Brazil flyers:", brazil)
	fmt.Println("Morning flyers:", mornings)
	fmt.Println("Brazil average:", avg)
}
