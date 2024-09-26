package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// создаем список кафе
var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

// через метод Get получаем количество требуемых ресторанов
func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("count missing"))
		return
	}

	// переводим в число
	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("wrong count value"))
		return
	}

	// извлекаем город
	city := req.URL.Query().Get("city")

	// ищем в нашей мапе список кафе по ключу
	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("wrong city value"))
		return
	}

	// если кол-во превышает, то возвращаем наибольшее кол-во имеющихся ресторанов
	if count > len(cafe) {
		count = len(cafe)
	}

	// соединяем ответ через запятую
	answer := strings.Join(cafe[:count], ",")

	// возвращаем статус
	w.WriteHeader(http.StatusOK)

	// возвращаем ответ
	_, _ = w.Write([]byte(answer))
}

func main() {
	http.HandleFunc("/", mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}

}
