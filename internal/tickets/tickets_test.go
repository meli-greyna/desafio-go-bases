package tickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var nicholas = Ticket{
	Id:          214,
	Name:        "Nicholas Jiroutek",
	Email:       "njiroutek5x@wufoo.com",
	Destination: "Russia",
	Time:        "6:55",
	Price:       1436,
}

var wendel = Ticket{
	Id:          215,
	Name:        "Wendel Feldhuhn",
	Email:       "wfeldhuhn5y@digg.com",
	Destination: "Brazil",
	Time:        "13:41",
	Price:       1006,
}

var selig = Ticket{
	Id:          216,
	Name:        "Selig Demeter",
	Email:       "sdemeter5z@webnode.com",
	Destination: "Philippines",
	Time:        "16:40",
	Price:       1047,
}

var testTickets = []Ticket{nicholas, wendel, selig}

func TestGetTotalTickets(t *testing.T) {
	totalBrazil := GetTotalTickets(testTickets, "Brazil")
	totalChina := GetTotalTickets(testTickets, "China")

	assert.Equal(t, 1, totalBrazil)
	assert.Equal(t, 0, totalChina)
}

func TestGetCountByPeriod(t *testing.T) {
	totalMadrugada, errMadrugada := GetCountByPeriod(testTickets, "madrugada")
	totalNoche, errNoche := GetCountByPeriod(testTickets, "noche")
	_, err := GetCountByPeriod(testTickets, "fail")

	assert.NoError(t, errMadrugada)
	assert.Equal(t, 1, totalMadrugada)

	assert.NoError(t, errNoche)
	assert.Equal(t, 0, totalNoche)

	assert.Error(t, err)
}

func TestAverageDestination(t *testing.T) {
	percentBrazil, errBrazil := AverageDestination(testTickets, "Brazil")
	percentChina, errChina := AverageDestination(testTickets, "China")
	_, err := AverageDestination([]Ticket{}, "fail")

	assert.NoError(t, errBrazil)
	assert.Equal(t, 33.33333333333333, percentBrazil)

	assert.NoError(t, errChina)
	assert.Equal(t, 0.0, percentChina)

	assert.Error(t, err)
}
