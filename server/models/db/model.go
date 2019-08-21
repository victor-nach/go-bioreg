package db

import (
	"time"
)

// User - defines standard user struct
type User struct {
	ID 						int			`json:"id,omitempty"`
	Name 					string		`json:"name,omitempty"`
	Address 				string		`json:"address,omitempty"`
	Email 					string		`json:"email,omitempty"`
	FingerPrint 			string		`json:"fingerPrint,omitempty"`
	FingerPrintToSend 		string		`json:"fingerPrintToSend,omitempty"`
	FingerPosition 			string		`json:"fingerPosition,omitempty"`
	Gender 					string		`json:"gender,omitempty"`
	DateCreated 			string		`json:"dateCreated,omitempty"`
	DateMoved 				time.Time	`json:"dateMoved,omitempty" gorm:"default:CURRENT_TIMESTAMP" sql:"default:CURRENT_TIMESTAMP"`
}