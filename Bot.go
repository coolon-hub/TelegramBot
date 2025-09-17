package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	randomNumber int
	hello        string
	creator      string
	help         string
)

func main() {
	text()

	bot, err := tgbotapi.NewBotAPI("8237516250:AAHkIyTXTTsktx11sywTgdl13qUXP_g1KhE")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Бот %s запущен и готов к работе!", bot.Self.UserName)

	// Настраиваем параметры получения обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Обрабатываем входящие сообщения
	for update := range updates {
		// ВЫВОДИМ ВСЕ СООБЩЕНИЯ В КОНСОЛЬ
		if update.Message != nil {
			printMessageToConsole(update.Message)
		}

		// Обрабатываем команды
		if update.Message != nil && update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, hello)
				msg.ReplyToMessageID = update.Message.MessageID
				if _, err := bot.Send(msg); err != nil {
					log.Println("Ошибка отправки:", err)
				}

			case "random":
				random()
				responseText := fmt.Sprintf("🎲 Ваше случайное число: %d", randomNumber)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
				msg.ReplyToMessageID = update.Message.MessageID
				if _, err := bot.Send(msg); err != nil {
					log.Println("Ошибка отправки:", err)
				}

			case "creator":
				responseText := fmt.Sprintf(creator)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
				msg.ReplyToMessageID = update.Message.MessageID
				if _, err := bot.Send(msg); err != nil {
					log.Println("Ошибка отправки:", err)
				}

			case "help":
				responseText := fmt.Sprintf(help)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
				msg.ReplyToMessageID = update.Message.MessageID
				if _, err := bot.Send(msg); err != nil {
					log.Println("Ошибка отправки:", err)
				}

			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "❌ Неизвестная команда. Используйте /help для списка команд.")
				bot.Send(msg)
			}
		}
	}
}

// Функция для вывода сообщений в консоль
func printMessageToConsole(message *tgbotapi.Message) {
	timestamp := time.Now().Format("15:04:05")

	fmt.Printf("\n=== 📨 Новое сообщение ===\n")
	fmt.Printf("Время: %s\n", timestamp)
	fmt.Printf("От: %s %s (@%s)\n",
		message.From.FirstName,
		message.From.LastName,
		message.From.UserName)
	fmt.Printf("ID чата: %d\n", message.Chat.ID)
	fmt.Printf("Текст: %s\n", message.Text)

	if message.IsCommand() {
		fmt.Printf("Команда: /%s\n", message.Command())
	}
	fmt.Printf("==========================\n")
}

// Функция для рандоматизации чисел
func random() {
	randomNumber = rand.Intn(100) + 1
	fmt.Printf("Сгенерировано число: %d\n", randomNumber)
}

// Функция для создание выводимого текста
func text() {
	hello = "Добро пожаловать в тестовую версию бота на языке go\n\n" +
		"Вот список доступных команд:\n" +
		"Предупреждаю, каждое ваше сообщение будет отсылаться создателю\n" +
		"/random - случайное число\n" +
		"/creator - узнать о разработчике"

	creator = "Разработчик данного бота:\n" +
		"Илья Лукьянов из 3ИС3."

	help = "Вот список доступных команд:\n" +
		"/random - случайное число\n" +
		"/creator - узнать о разработчике"
}
