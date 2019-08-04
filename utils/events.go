package utils

import "vorta/models"

type VEvent struct {
	Topic   string
	Message string
	Profile *models.Profile
}
