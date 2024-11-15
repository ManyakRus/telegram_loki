package telegram

import (
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/stopapp"
	"github.com/ManyakRus/starter/telegram_bot"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/ManyakRus/telegram_loki/internal/constants"
	"github.com/ManyakRus/telegram_loki/internal/load_json"
	"github.com/ManyakRus/telegram_loki/internal/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

// SendMessage - отправляет сообщения в Telegram, в общий чат, многим разрабочикам
// ChatName - в формате "@Имя1,@Имя2"
func SendMessage(Message1 types.Message) error {
	var err error

	//проверка отмены контекста
	err1 := contextmain.GetContext().Err()
	if err1 != nil {
		return err1
	}

	//
	ChatName := strings.TrimSpace(Message1.DeveloperName)
	//TextDate := Message1.Date.String()
	TextDate := Message1.Date.Format(constants.Layout)
	TextDate2 := micro.SubstringLeft(TextDate, 10)
	if strings.Contains(Message1.Text, TextDate2) == true {
		TextDate = ""
	} else {
		TextDate = "\n" + TextDate
	}
	if Message1.Date.IsZero() == true {
		TextDate = ""
	}
	DeveloperName := Message1.DeveloperName
	Text := Message1.Text

	TextServiceName := Message1.ServiceName
	if Message1.LokiURL != "" {
		TextServiceName = `<a href="` + Message1.LokiURL + `">` + Message1.ServiceName + "</a>"
	}

	//
	if Message1.IsSameLastText == true {
		DeveloperName = ""
	}

	Text = TextServiceName + " " + DeveloperName + " " + TextDate + "\n" + Text

	//отправка в общую группу
	if config.Settings.TELEGRAM_CHAT_NAME != "" {
		_, err = telegram_bot.SendMessage(config.Settings.TELEGRAM_CHAT_NAME, Text)
		if err != nil {
			log.Error("SendMessage() error: ", err)
			return err
		}
	}

	//отправка каждому программисту
	if ChatName == "" {
		return err
	}
	if ChatName == config.Settings.TELEGRAM_CHAT_NAME {
		return err
	}

	Text = Message1.Text
	if TextServiceName != "" {
		Text = TextServiceName + "\n" + Text
	}

	//
	MassChatNames := strings.Split(ChatName, ",")
	for _, ChatName1 := range MassChatNames {
		//проверка отмены контекста
		err1 := contextmain.GetContext().Err()
		if err1 != nil {
			return err1
		}
		ChatName1 = strings.TrimSpace(ChatName1)

		//найдём ИД по имени в телеграм
		ID, ok := types.MapTelegramUsers[ChatName1]
		if ok == false {
			continue
		}

		//отправка
		_, err = telegram_bot.SendMessageChatID(ID, Text)
		if err != nil {
			log.Error("SendMessage() error: ", err)
			continue
			//return err
		}
	}

	return err
}

// StartReadMessages - запускает чтение сообщений из Телеграм
func StartReadMessages() {
	stopapp.GetWaitGroup_Main().Add(1)
	go StartReadMessages_go()
}

// StartReadMessages_go - запускает чтение сообщений из Телеграм, внутри горутины
func StartReadMessages_go() {
	defer stopapp.GetWaitGroup_Main().Done()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := telegram_bot.Client.GetUpdatesChan(u)

	for update1 := range updates {
		if update1.Message == nil { // If we got a message
			continue
		}

		//
		ChatID := update1.Message.Chat.ID
		Name := update1.Message.Chat.UserName

		//
		AddMapTelegramUsers(Name, ChatID)

	}
}

// AddMapTelegramUsers - добавляет пользователя в map
func AddMapTelegramUsers(Name string, ChatID int64) {

	_, ok := types.MapTelegramUsers[Name]
	if ok == false {
		//добавим в карту
		types.MapTelegramUsers["@"+Name] = ChatID

		//
		load_json.SaveMapTelegramUsers()
	}
}
