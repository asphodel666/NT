package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Структура данных для хранения информации о пользователе
type User struct {
	Name  string
	Email string
}

// Обработчик для главной страницы
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Добро пожаловать на главную страницу!</h1>")
}

// Обработчик для страницы о пользователе
func userHandler(w http.ResponseWriter, r *http.Request) {
	// Создание структуры пользователя
	user := User{Name: "Иван", Email: "ivan@example.com"}

	// HTML-шаблон для отображения данных
	tmpl := `<html><body>
        <h1>Информация о пользователе</h1>
        <p>Имя: {{.Name}}</p>
        <p>Email: {{.Email}}</p>
    </body></html>`

	// Создание нового шаблона
	t, _ := template.New("user").Parse(tmpl)
	t.Execute(w, user)
}

// Обработчик для формы
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		fmt.Fprintf(w, "<h1>Спасибо за отправку формы!</h1><p>Имя: %s</p><p>Email: %s</p>", name, email)
	} else {
		// Отправка HTML-формы
		fmt.Fprintln(w, `
            <html>
            <body>
                <h1>Форма обратной связи</h1>
                <form method="post">
                    <label for="name">Имя:</label>
                    <input type="text" id="name" name="name" required><br><br>
                    <label for="email">Email:</label>
                    <input type="email" id="email" name="email" required><br><br>
                    <input type="submit" value="Отправить">
                </form>
            </body>
            </html>
        `)
	}
}

// Обработчик для статических файлов
func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "."+r.URL.Path)
}

func main() {
	// Регистрация обработчиков
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/form", formHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Запуск сервера
	fmt.Println("Сервер запущен на http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
