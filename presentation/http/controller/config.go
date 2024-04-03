package controller

import (
	"encoding/json"
	"net/http"
	configuse "trabalho-01-batalha-naval/application/use/config"
)

type ConfigController struct {
	listUseCase *configuse.ListConfigUseCase
}

func NewConfigController(listUseCase *configuse.ListConfigUseCase) *ConfigController {
	return &ConfigController{
		listUseCase,
	}
}

func (c *ConfigController) Index(w http.ResponseWriter, _ *http.Request) {
	list := c.listUseCase.Execute()

	w.Header().Set("Content-Type", JsonContentType)
	_ = json.NewEncoder(w).Encode(list)
}
