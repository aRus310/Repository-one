package spentcalories

import (
	"errors"
	"fmt"
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

func parseTraining(data string) (steps int, activityType string, duration time.Duration, err error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",") // разделим строку на слайс
	if len(parts) != 3 {              // проверка наличая трёх элементов в слайсе
		return 0, "", 0, errors.New("некорректный формат ввода данных")
	}

	stepsStr := parts[0]    // перый элемент слайса количество шагов
	activityType = parts[1] // второй элемент вид активности
	durationStr := parts[2] // третий элемент продолжительность активности

	stepsInt, err := strconv.Atoi(stepsStr) // преобразование количество шагов в тип int
	if err != nil {
		return 0, "", 0, errors.New("ошибка преобразования шагов")
	}

	parsedDuration, err := time.ParseDuration(durationStr) // продолжительность активности в тип time.Duration
	if err != nil {
		return 0, "", 0, errors.New("ошибка преобразования времени")
	}

	return stepsInt, activityType, parsedDuration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient       // длина шага
	totalDistanceMeters := float64(steps) * stepLength // пройденная дистанция в метрах
	return totalDistanceMeters / mInKm                 // дистанция в км
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 { // проверка наличая времени активности
		return 0
	}

	distanceKilometers := distance(steps, height) // вызов функции distance

	hours := duration.Hours() // переводим продолжительность в часы

	speed := distanceKilometers / hours // вычисляем среднюю скорость
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	parsedSteps, activityType, parsedDuration, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	var caloriesBurned float64
	var avgSpeed float64
	switch activityType { // проверяем вид тренировки
	case "Бег":
		caloriesBurned, _ = RunningSpentCalories(parsedSteps, weight, height, parsedDuration) // сожгли калорий при беге
		avgSpeed = meanSpeed(parsedSteps, height, parsedDuration)                             // средняя сорость при беге
	case "Ходьба":
		caloriesBurned, _ = WalkingSpentCalories(parsedSteps, weight, height, parsedDuration) // сожгли калорий при ходьбе
		avgSpeed = meanSpeed(parsedSteps, height, parsedDuration)                             // средняя сорость при ходьбе
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	infoString := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		activityType,
		parsedDuration.Hours(),
		distance(parsedSteps, height),
		avgSpeed,
		caloriesBurned,
	)
	return infoString, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные данные бег-калорий")
	}

	avgSpeed := meanSpeed(steps, height, duration) // средняя скорость бега

	minutes := duration.Minutes() // переводим продолжительность в минуты

	spentCalories := weight * avgSpeed * minutes / minInH // количество потраченных калорий при беге

	return spentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные данные ходьба-калорий")
	}

	avgSpeed := meanSpeed(steps, height, duration) // // средняя скорость ходьбы

	minutes := duration.Minutes() // переводим продолжительность в минуты

	spentCalories := weight * avgSpeed * minutes * walkingCaloriesCoefficient / minInH // количество потраченных калорий при ходьбе

	return spentCalories, nil
}
