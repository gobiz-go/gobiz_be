package config

import "os"

var WAKeyword string = os.Getenv("WAQRKEYWORD")

var WebhookSecret string = os.Getenv("WEBHOOKSECRET")

var WAAPIQRLogin string = "https://api.wa.my.id/api/whatsauth/request"

var WAAPIMessage string = "https://api.wa.my.id/api/send/message/text"
