package handler

import (
	"app/internal/responses"
	"app/internal/services"

	"github.com/gofiber/fiber/v2"
)

type HealthCheckHandler struct {
	HealthCheckService services.HealthCheckService
}

func NewHealthCheckHandler(healthCheckService services.HealthCheckService) *HealthCheckHandler {
	return &HealthCheckHandler{
		HealthCheckService: healthCheckService,
	}
}

func (h *HealthCheckHandler) addServiceStatus(
	serviceList *[]responses.HealthCheck, name string, isUp bool, message *string,
) {
	status := "Up"

	if !isUp {
		status = "Down"
	}

	*serviceList = append(*serviceList, responses.HealthCheck{
		Name:    name,
		Status:  status,
		IsUp:    isUp,
		Message: message,
	})
}

// @Tags Health
// @Summary Health Check
// @Description Check the status of services and database connections
// @Accept json
// @Produce json
// @Success 200 {object} example.HealthCheckResponse
// @Failure 500 {object} example.HealthCheckResponseError
// @Router /health-check [get]
func (h *HealthCheckHandler) Check(c *fiber.Ctx) error {
	isHealthy := true
	var serviceList []responses.HealthCheck

	// Check the database connection
	if err := h.HealthCheckService.DBcheck(); err != nil {
		isHealthy = false
		errMsg := err.Error()
		h.addServiceStatus(&serviceList, "Postgre", false, &errMsg)
	} else {
		h.addServiceStatus(&serviceList, "Postgre", true, nil)
	}

	if err := h.HealthCheckService.MemoryHeapCheck(); err != nil {
		isHealthy = false
		errMsg := err.Error()
		h.addServiceStatus(&serviceList, "Memory", false, &errMsg)
	} else {
		h.addServiceStatus(&serviceList, "Memory", true, nil)
	}

	// Return the response based on health check result
	statusCode := fiber.StatusOK
	status := "success"

	if !isHealthy {
		statusCode = fiber.StatusInternalServerError
		status = "error"
	}

	return c.Status(statusCode).JSON(responses.HealthCheckResponse{
		Status:    status,
		Message:   "Health check completed",
		Code:      statusCode,
		IsHealthy: isHealthy,
		Result:    serviceList,
	})
}
