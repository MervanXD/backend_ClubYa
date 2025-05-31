package models

type ApiResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

// funcion para generar una respuesta de exito
func Succes(message string, data interface{}) ApiResponse {
	return ApiResponse{
		Data:    data,
		Message: message,
		Status:  200,
	}
}

// funcion para generar una respuesta de no encontrado
func NotFound(message string) ApiResponse {
	return ApiResponse{
		Data:    nil,
		Message: message,
		Status:  404,
	}
}

// funcion para generar una respuesta de error
func Error(message string, data interface{}) ApiResponse {
	return ApiResponse{
		Data:    data,
		Message: message,
		Status:  500,
	}
}

// funcion para generar una respuesta de error
func BadRequest(message string, data interface{}) ApiResponse {
	return ApiResponse{
		Data:    data,
		Message: message,
		Status:  400,
	}
}
