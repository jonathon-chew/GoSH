package sleep

import (
	"strconv"
	"strings"
	"time"
)

func checkFloat(argument string) (float64, error) {

	if !strings.Contains(argument, ".") {
		if rune(argument[len(argument)]) >= 30 && rune(argument[len(argument)]) <= 39 {
			argument += ".0" // If last thing is a number and doesn't have a period
		} else {
			argument = argument[:len(argument)-1] + ".0" + string(argument[len(argument)])
		}
	}

	sleepTime, ErrReadingArgumentAsNumber := strconv.ParseFloat(argument, 64)
	if ErrReadingArgumentAsNumber != nil {
		return 0.0, ErrReadingArgumentAsNumber
	}

	return sleepTime, nil
}

func Sleep(argumnet string) error {

	userTime, ErrGettingTime := checkFloat(argumnet)
	if ErrGettingTime != nil {
		return ErrGettingTime
	}

	switch argumnet[len(argumnet)] {
	default:
		time.Sleep(time.Duration(time.Duration(userTime).Seconds()))
	case 's':
		time.Sleep(time.Duration(time.Duration(userTime).Seconds()))
	case 'm':
		time.Sleep(time.Duration(time.Duration(userTime).Minutes()))
	case 'n':
		time.Sleep(time.Duration(time.Duration(userTime).Nanoseconds()))
	}
	return nil
}
