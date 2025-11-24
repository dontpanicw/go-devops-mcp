package main

import (
	"go-devops-mcp/internal/mcp"
	"go-devops-mcp/internal/tools/health"
)

func main() {

	HealthTool := health.NewHealth()
	mcp.StartServer(HealthTool)
}
