package service

import (
	"tspo_final/internal/dto"
	"tspo_final/internal/models"
	"tspo_final/internal/repository"
)

func CreateUser(user *models.User, repository repository.UserRepository) dto.InfoResponse {
	operationResult := repository.Save(user)

	if operationResult.Error != nil {
		return dto.InfoResponse{IsResponseSuccess: false, TextResponse: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.User)

	return dto.InfoResponse{IsResponseSuccess: true, PayloadResponse: data}
}

func FindAllUsers(repository repository.UserRepository) dto.InfoResponse {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dto.InfoResponse{IsResponseSuccess: false, TextResponse: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Users)

	return dto.InfoResponse{IsResponseSuccess: true, PayloadResponse: datas}
}

func FindOneUserById(id uint, repository repository.UserRepository) dto.InfoResponse {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dto.InfoResponse{IsResponseSuccess: false, TextResponse: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.User)

	return dto.InfoResponse{IsResponseSuccess: true, PayloadResponse: data}
}

func UpdateUserById(id uint, user *models.User, repository repository.UserRepository) dto.InfoResponse {
	existingUserResponse := FindOneUserById(id, repository)

	if !existingUserResponse.IsResponseSuccess {
		return existingUserResponse
	}

	existingUser := existingUserResponse.PayloadResponse.(*models.User)

	existingUser.Name = user.Name

	operationResult := repository.Save(existingUser)

	if operationResult.Error != nil {
		return dto.InfoResponse{IsResponseSuccess: false, TextResponse: operationResult.Error.Error()}
	}

	return dto.InfoResponse{IsResponseSuccess: true, PayloadResponse: operationResult.Result}
}

func DeleteOneUserById(id uint, repository repository.UserRepository) dto.InfoResponse {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dto.InfoResponse{IsResponseSuccess: false, TextResponse: operationResult.Error.Error()}
	}

	return dto.InfoResponse{IsResponseSuccess: true}
}
