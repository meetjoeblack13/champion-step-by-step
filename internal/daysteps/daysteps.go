package daysteps

import (
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
		return 0, 0, fmt.Errorf("Ошибка ввода!")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return 0, 0, fmt.Errorf("Ошибка подсчета количества шагов")
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil || duration <= 0 {
		return 0, 0, fmt.Errorf("Ошибка подсчета продолжительности прогулки")
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
	distanceKilometers := distanceMeters / float64(mInKm)
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distanceKilometers, calories)
}
