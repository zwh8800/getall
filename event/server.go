package event

import (
	evbus "github.com/asaskevich/EventBus"
)

var Server = evbus.New()
