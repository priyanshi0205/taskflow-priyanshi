package user

import "backedn_go/model"

func MapList(users []model.User) []userDropdownOptionResponse {
	response := make([]userDropdownOptionResponse, 0, len(users))
	for _, user := range users {
		response = append(response, mapSingle(user))
	}
	return response
}

func mapSingle(user model.User) userDropdownOptionResponse {
	return userDropdownOptionResponse{
		Name: user.Name,
		UUID: user.ID,
	}
}
