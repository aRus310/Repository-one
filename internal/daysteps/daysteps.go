package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (steps int, duration time.Duration, err error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",") // делим строку на части
	if len(parts) != 2 {              // проверка наличия двух элементов
		return 0, 0, errors.New("неверный формат данных")
	}

	stepsStr := parts[0]    // первый элемент слайса (количество шагов)
	durationStr := parts[1] // второй элемент слайса (продолжительность шагов по времени)

	stepsInt, err := strconv.Atoi(stepsStr) // преобразование в тип int
	if err != nil || stepsInt <= 0 {
		return 0, 0, errors.New("ошибка преобразования шагов")
	}

	parsedDuration, err := time.ParseDuration(durationStr) // преобразование в тип time.Duration
	if err != nil {
		return 0, 0, errors.New("ошибка преобразования длительности")
	}

	return stepsInt, parsedDuration, nil
}
func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err.Error()) // вывод сообщения об ошибке
		return ""                // возвращаем пустую строку, если произошла ошибка
	}

	if steps <= 0 {
		return "" // возврат пустой строки, если если шагов нет или отрацательные
	}

	distanceMeters := float64(steps) * stepLength // вычислить дистанцию в метрах
	distanceKilometers := distanceMeters / mInKm  // перевести дистанцию в километры

	caloriesBurned, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration) // калорий на прогулке
	if err != nil {
		fmt.Println(err.Error()) // выводим сообщение об ошибке
		return ""                // возвращаем пустую строку, если возникла ошибка
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		steps, distanceKilometers, caloriesBurned,
	)

	return result
}
