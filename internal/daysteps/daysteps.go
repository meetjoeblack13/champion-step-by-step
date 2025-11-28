package daysteps

import (
	"time"
	"strings"
	"fmt"
	"strconv"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
return 0, 0, fmt.Errorf("Ошибка ввода!")
	}
steps, err := strconv.Atoi(parts[0])
if err != nil || steps < 0 {
return 0, 0, fmt.Errorf("Ошибка подсчета количества шагов")
}
walkDuration, err := time.ParseDuration(parts[1])
if err != nil || walkDuration < 0 {
return 0, 0, fmt.Errorf("Ошибка подсчета продолжительности прогулки")
}
return steps, walkDuration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, walkDuration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	_ = walkDuration // zaglushka
	if steps == 0 {
		return ""
	}
	distanceMeters := float64(steps) * stepLength
distanceKilometers := distanceMeters / float64(mInKm)
calories := 0.0 //zaglushka
return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", 
steps, distanceKilometers, calories)
}
