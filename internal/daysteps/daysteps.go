package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/meetjoeblack13/champion-step-by-step/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) { // Функция для обработки входящих данных
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("input error")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, errors.New("steps parsing error")
	}
	if steps <= 0 {
		return 0, 0, errors.New("steps calculation error")
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, errors.New("duration parsing error")
	}
	if duration <= 0 {
		return 0, 0, errors.New("duration calculation error")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string { // Функция для возврата данных о тренировке
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distanceMeters := float64(steps) * stepLength
	distanceKilometers := distanceMeters / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distanceKilometers, calories)
}
