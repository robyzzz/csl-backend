package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/robyzzz/csl-backend/model"
	"github.com/solovev/steam_go"
)

type ErrorResponse struct {
	Code     int
	ErrorMsg string
}

// ErrorResponse "constructor"
func NewAPIError(c int, m string) ErrorResponse {
	return ErrorResponse{Code: c, ErrorMsg: m}
}

// Error response in JSON format
func APIErrorRespond(w http.ResponseWriter, res ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	log.Printf(res.ErrorMsg)
	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}

// Converts PlayerSummaries to SteamUser
func PlayerSummariesToSteamUser(user *steam_go.PlayerSummaries) model.SteamUser {
	return model.SteamUser{
		ID:             0,
		SteamID:        user.SteamId,
		PersonaName:    user.PersonaName,
		LastLogOff:     user.LastLogOff,
		ProfileUrl:     user.ProfileUrl,
		Avatar:         user.Avatar,
		AvatarMedium:   user.AvatarMedium,
		AvatarFull:     user.AvatarFull,
		RealName:       user.RealName,
		PrimaryClanID:  user.PrimaryClanId,
		TimeCreated:    user.TimeCreated,
		LocCountryCode: user.LocCountryCode,
		GameID:         user.GameId,
	}
}
