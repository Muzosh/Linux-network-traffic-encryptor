package gateway

import (
	"log"
	"time"

	"github.com/songgao/water"
)

const (
	// Message for keeping dynamic NAT translation
	keepAliveMessage = "Keep Alive"
)

func Listen() {
	// QKD IP
	// cert auth

	// Create TUNTAP
	_, err := water.New(water.Config{
		DeviceType: water.TUN,
		PlatformSpecificParams: water.PlatformSpecificParams{
			Name: "tun99",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Reference time for dynamic NAT translation
	_ = time.Now()

	for {

	}
}
