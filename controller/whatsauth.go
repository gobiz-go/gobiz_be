package controller

import (
	"gocroot/config"
	"gocroot/helper"
	"gocroot/model"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func WhatsAuthReceiver(c *fiber.Ctx) error {
	var h model.Header
	err := c.ReqHeaderParser(&h)
	if err != nil {
		return err
	}
	var resp model.Response
	if h.Secret == config.WebhookSecret {
		var msg model.IteungMessage
		err = c.BodyParser(&msg)
		if err != nil {
			return err
		}
		if IsLoginRequest(msg, config.WAKeyword) { //untuk whatsauth request login
			resp = HandlerQRLogin(msg, config.WAKeyword)
		} else { //untuk membalas pesan masuk
			resp = HandlerIncomingMessage(msg)
		}
	} else {
		resp.Response = "Secret Salah"
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func IsLoginRequest(msg model.IteungMessage, keyword string) bool {
	return strings.Contains(msg.Message, keyword) && msg.From_link
}

func GetUUID(msg model.IteungMessage, keyword string) string {
	return strings.Replace(msg.Message, keyword, "", 1)
}

func HandlerQRLogin(msg model.IteungMessage, WAKeyword string) (resp model.Response) {
	dt := &model.WhatsauthRequest{
		Uuid:        GetUUID(msg, WAKeyword),
		Phonenumber: msg.Phone_number,
		Delay:       msg.From_link_delay,
	}
	resp, _ = helper.PostStructWithToken[model.Response]("Token", WAAPIToken(), dt, config.WAAPIQRLogin)
	return
}

func HandlerIncomingMessage(msg model.IteungMessage) (resp model.Response) {
	dt := &model.TextMessage{
		To:       msg.Chat_number,
		IsGroup:  false,
		Messages: "Hai hai hai kak " + msg.Alias_name,
	}
	if msg.Chat_server == "g.us" { //jika pesan datang dari group maka balas ke group
		dt.IsGroup = true
	}
	if (msg.Phone_number != "628112000279") && (msg.Phone_number != "6283131895000") { //ignore pesan datang dari iteung
		resp, _ = helper.PostStructWithToken[model.Response]("Token", WAAPIToken(), dt, config.WAAPIMessage)
	}
	return
}

func WAAPIToken() string {
	token := "112"
	return token
}
