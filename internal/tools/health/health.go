package health

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

type Health struct {
}

func NewHealth() *Health {
	return &Health{}
}

func (h *Health) HealthHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}
