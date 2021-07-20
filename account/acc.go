package account

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"genshin/bot_func"
)

type Users struct {
	Users []*User `json:"users"`
}
type User struct {
	ID           int    `json:"ID"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Uid          string `json:"uid"`
	Account_id   string `json:"account_id"`
	Cookie_token string `json:"cookie_token"`
	State        float64
}

var Accs Users

const my_info = 1
const change_info = 2
const gift = 3

func Instruction(text string) int {
	if text == "/my_info" || text == "/my_info"+bot_func.Bot_info.Username {
		return my_info
	} else if text == "/change_info" || text == "/change_info"+bot_func.Bot_info.Username {
		return change_info
	} else if text == "/gift" || text == "/gift"+bot_func.Bot_info.Username {
		return gift
	}
	return 0
}

func Save() {
	file, _ := json.MarshalIndent(Accs, "", "  ")
	_ = ioutil.WriteFile("./account/acc.json", file, 0600)
}

func (users *Users) Init() {
	jsonFile, err := os.Open("./account/acc.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, users)
	for i := 0; i < len(users.Users); i += 1 {
		fmt.Println(users.Users[i])
	}
}

func (users *Users)Gift(text string) {
	texts := strings.Split(text, "\n")
	for i := range texts {
		for j := range users.Users {
			users.Users[j].Gift(texts[i])
		}
	}
}

func (users *Users)Signin() {
	for i := range users.Users {
		users.Users[i].Signin()
	}
}

func (user *User) Update(text string) float64 {
	texts := strings.Split(text, "\n")
	if len(texts) < 3 {
		bot_func.TGBot.SendMessage(int64(user.ID), "Wrong input format")
		str := "請輸入你的 uid(game id), cookies(account_id & cookie_token)\n" +
			"example:\n" +
			"800800800\n108100100\njasdgjljknmwefibna"
		bot_func.TGBot.SendMessage(int64(user.ID), str)
		return user.State
	}
	user.Uid = texts[0]
	user.Account_id = texts[1]
	user.Cookie_token = texts[2]
	str := "Your information:\n" +
		"uid: " + user.Uid + "\n" +
		"accountd_id: " + user.Account_id + "\n" +
		"cookie_token: " + user.Cookie_token
	bot_func.TGBot.SendMessage(int64(user.ID), str)
	Save()
	return 1.0
}

func (user *User) Check() float64 {
	str := "uid: " + user.Uid + "\n" +
		"accountd_id: " + user.Account_id + "\n" +
		"cookie_token: " + user.Cookie_token
	bot_func.TGBot.SendMessage(int64(user.ID), str)
	return 1.0
}

func (user *User) Help(chatID int64) {
	str := "/my_info - 查詢自己的資訊\n" +
		"/change_info - 修改自己的資訊\n" +
		"/gift - 輸入兌換碼\n" +
		"/help - help"
	bot_func.TGBot.SendMessage(chatID, str)
}

func (user *User) Gift(text string) {
	uid := "uid=" + user.Uid
	code := uid + "&cdkey=" + text
	url := "https://hk4e-api-os.mihoyo.com/common/apicdkey/api/webExchangeCdkey?region=os_asia&lang=zh-tw&game_biz=hk4e_global&" + code
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.AddCookie(&http.Cookie{Name: "account_id", Value: user.Account_id})
	req.AddCookie(&http.Cookie{Name: "cookie_token", Value: user.Cookie_token})
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bot_func.TGBot.SendMessage(int64(user.ID), string(body))
}

func (user *User)Signin() {
	URL := "https://hk4e-api-os.mihoyo.com/event/sol/sign?lang=zh-tw"
	data := strings.NewReader(`{"act_id":"e202102251931481"}`)
	req, err := http.NewRequest("POST", URL, data)
	if err != nil {
		fmt.Println(err)
	}

	req.AddCookie(&http.Cookie{Name: "account_id", Value: user.Account_id})
	req.AddCookie(&http.Cookie{Name: "cookie_token", Value: user.Cookie_token})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func (users *Users) Acc_main(update tgbotapi.Update) {
	user := update.Message.From
	text := update.Message.Text
	chatID := update.Message.Chat.ID
	var acc *User
	for i := range users.Users {
		if Accs.Users[i].ID == user.ID {
			acc = Accs.Users[i]
			break
		}
	}
	if acc == nil {
		acc = &User{
			State: 0,
		}
	}
	if acc.State == 0 {
		acc.State = 2.0
		acc.ID = user.ID
		acc.FirstName = user.FirstName
		acc.LastName = user.LastName
		Accs.Users = append(Accs.Users, acc)
		str := "請輸入你的 uid(game id), cookies(account_id & cookie_token)\n" +
			"example:\n" +
			"800800800\n108100100\njasdgjljknmwefibna"
		bot_func.TGBot.SendMessage(int64(user.ID), str)
	} else if acc.State == 1.0 {
		if string(text[0]) == "/" {
			if Instruction(text) == my_info {
				acc.Check()
			} else if Instruction(text) == change_info {
				acc.State = 2.1
				str := "請輸入你的 uid(game id), cookies(account_id & cookie_token)\n" +
					"example:\n" +
					"800800800\n108100100\njasdgjljknmwefibna"
				bot_func.TGBot.SendMessage(int64(user.ID), str)
			} else if Instruction(text) == gift {
				acc.State = 3.0
				str := "請輸入兌換碼"
				bot_func.TGBot.SendMessage(int64(user.ID), str)
			} else {
				acc.Help(chatID)
			}
		} else {
			acc.Help(chatID)
		}
	} else if int(acc.State) == 2 {
		fmt.Println("state 2")
		acc.State = acc.Update(text)
		Save()
	} else if acc.State == 3.0 {
		fmt.Println("state 3")
		users.Gift(text)
		acc.State = 1.0
	}
}
