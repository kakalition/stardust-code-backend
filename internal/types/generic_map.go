package types

import (
	"stardustcode/backend/internal/projects/parcus/models"
	"time"
)

type GenericMap map[interface{}]interface{}

func (m *GenericMap) GetUser() *models.NetworkUser {
	id, _ := (*m)["id"].(string)
	email, _ := (*m)["email"].(string)
	displayName, _ := (*m)["displayName"].(string)
	lastSignedIn, _ := (*m)["lastSignedIn"].(time.Time)

	return &models.NetworkUser{
		Id:           &id,
		Email:        &email,
		DisplayName:  &displayName,
		LastSignedIn: &lastSignedIn,
	}
}
