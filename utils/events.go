package utils

import "vorta-go/models"

type VEvent struct {
	Topic   string
	Message string
	Profile *models.Profile
}
