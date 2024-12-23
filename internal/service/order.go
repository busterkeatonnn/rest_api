package service

import (
	"tspo_final/internal/dto"
	"tspo_final/internal/models"
	"tspo_final/internal/repository"
)

func CreateOrder(order *models.Order, repository repository.OrderRepository) dto.InfoResponse {
	operationResult := repository.Save(order)

	if operationResult.Error != nil {
		return dto.InfoResponse{IsResponseSuccess: false, TextResponse: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Order)

	return dto.InfoResponse{IsResponseSuccess: true, PayloadResponse: data}
}

func FindAllOrders(repository repository.OrderRepository) dto.InfoResponse {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dto.InfoResponse{IsResponseSuccess: false, TextResponse: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*models.Orders)

	return dto.InfoResponse{IsResponseSuccess: true, PayloadResponse: datas}
}

func FindOneOrderById(id uint, repository repository.OrderRepository) dto.InfoResponse {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dto.InfoResponse{IsResponseSuccess: false, TextResponse: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Order)

	return dto.InfoResponse{IsResponseSuccess: true, PayloadResponse: data}
}

func UpdateOrderById(id uint, order *models.Order, repository repository.OrderRepository) dto.InfoResponse {
	existingOrderResponse := FindOneOrderById(id, repository)

	if !existingOrderResponse.IsResponseSuccess {
		return existingOrderResponse
	}

	existingOrder := existingOrderResponse.PayloadResponse.(*models.Order)

	existingOrder.DeliveryType = order.DeliveryType
	existingOrder.DeliveryTime = order.DeliveryTime
	existingOrder.OrderTime = order.OrderTime
	existingOrder.TotalPrice = order.TotalPrice
	existingOrder.Canceled = order.Canceled

	operationResult := repository.Save(existingOrder)

	if operationResult.Error != nil {
		return dto.InfoResponse{IsResponseSuccess: false, TextResponse: operationResult.Error.Error()}
	}

	return dto.InfoResponse{IsResponseSuccess: true, PayloadResponse: operationResult.Result}
}

func DeleteOneOrderById(id uint, repository repository.OrderRepository) dto.InfoResponse {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dto.InfoResponse{IsResponseSuccess: false, TextResponse: operationResult.Error.Error()}
	}

	return dto.InfoResponse{IsResponseSuccess: true}
}
