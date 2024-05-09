package config

import "os"

var WAKeyword string = os.Getenv("WAQRKEYWORD")

var WebhookSecret string = os.Getenv("WEBHOOKSECRET")

var WAAPIQRLogin string = os.Getenv("WAAPIQR")

var WAAPIMessage string = os.Getenv("WAAPITXTMSG")
