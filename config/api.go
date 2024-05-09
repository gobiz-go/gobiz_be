package config

import "os"

var WAKeyword string = os.Getenv("WAKEYWORD")

var WebhookSecret string = os.Getenv("WASECRET")

var WAAPIToken string = os.Getenv("WAAPITOKEN")

var WAAPIQRLogin string = "https://api.wa.my.id/api/whatsauth/request"

var WAAPIMessage string = "https://api.wa.my.id/api/send/message/text"
