package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_acvp "github.com/input-api/mcp-server/tools/acvp"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_acvp.CreateDelete_acvp_v1Tool(cfg),
	}
}
