package main

import (
	"fmt"
	"github.com/Shemistan/Lesson_6/internal/controller"
	"github.com/Shemistan/Lesson_6/internal/model"
	"github.com/Shemistan/Lesson_6/internal/service"
	"github.com/Shemistan/Lesson_6/internal/storage"
	"time"
)

func _init() *controller.UserController {
	db := storage.NewStorage()
	userService := service.NewUserService(db)
	userController := controller.NewUserController(userService)
	return userController
}

func cin[T any](msg string, value T) {
	fmt.Print(msg)
	_, err := fmt.Scan(value)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	userController := _init()

	var name, surname, login, password string
	var choice, userID uint32
	var isAuth bool

	fmt.Println("[1] Авторизация\n[2] Регистрация\n[0] Выход")
	cin("Выберите действие: ", &choice)

start:

	if choice == 2 {
		fmt.Println("\t\tРегистрация")
		cin("Введите имя: ", &name)
		cin("Введите фамилию: ", &surname)
		if login == "" || password == "" {
			cin("Введите логин: ", &login)
			cin("Введите пароль: ", &password)
		}
		id, err := userController.Add(name, surname, login, password)
		if err != nil {
			fmt.Println(err)
		}
		userID = id
		isAuth = true
	} else if choice == 1 {
		fmt.Println("\t\tАвторизация")
		cin("Логин: ", &login)
		cin("Пароль: ", &password)

		userId, err := userController.Auth(login, password)
		if err != nil {
			if userId == 0 {
				fmt.Println(err)
			}
			if userId == -1 {
				fmt.Println("Данного пользователя не существует\nБудет создан новый пользователь")
				choice = 2
				goto start
			}
		}
		isAuth = true
	} else if choice == 0 {
		return
	}

	if isAuth {
		var user *model.User
		var err error

		fmt.Println("Вывод пользователя с ID: ", userID)
		user, err = userController.Get(userID)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(&user)

		fmt.Println("\nВывод всех пользователей")
		users, _ := userController.GetAll()
		if err != nil {
			fmt.Println(err)
		}

		for _, user := range users {
			fmt.Println(&user)
		}

		time.Sleep(5 * time.Second)

		fmt.Println("\nОбновление данных пользователя")
		err = userController.Update(userID, &model.User{
			Name:         "Undefined",
			Surname:      "Undefined",
			Login:        "Undefined",
			HashPassword: 0,
			Status:       "Undefined",
			Role:         "Undefined",
		})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(userController.Get(userID))

		fmt.Println("Удаление пользователя с ID: ", userID)
		err = userController.Delete(userID)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("\nВывод статистики")
		fmt.Println(userController.GetStatistics())
	}
}
