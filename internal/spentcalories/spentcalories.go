package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("Ошибка ввода!")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return 0, "", 0, fmt.Errorf("Ошибка подсчета количества шагов")
	}
	activity := strings.Title(strings.ToLower(parts[1]))
	if activity != "Ходьба" && activity != "Бег" {
		return 0, "", 0, fmt.Errorf("Ошибка: неизвестный тип тренировки")
	}
	duration, err := time.ParseDuration(parts[2])
	if err != nil || duration <= 0 {
		return 0, "", 0, fmt.Errorf("Ошибка подсчета продолжительности прогулки")
	}
	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength
	distanceKilometers := distanceMeters / mInKm
	return distanceKilometers
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0.0
	}
	distanceKilometers := distance(steps, height)
	return distanceKilometers / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// считаем общие значения
	distanceKm := distance(steps, height) // одинаково для ходьбы и бега
	speedKmH := meanSpeed(steps, height, duration)
	var calories float64
	switch activity {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, duration.Hours(), distanceKm, speedKmH, calories)
	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0.0, fmt.Errorf("Ошибка: некорректные входные параметры")
	}
	speed := meanSpeed(steps, height, duration)
	runningCalories := (weight * speed * duration.Minutes()) / minInH
	return runningCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0.0, fmt.Errorf("Ошибка: некорректные входные параметры")
	}
	speed := meanSpeed(steps, height, duration)
	walkingCalories := (weight * speed * duration.Minutes()) / minInH
	walkingCaloriesWithCoef := walkingCalories * walkingCaloriesCoefficient
	return walkingCaloriesWithCoef, nil
}
