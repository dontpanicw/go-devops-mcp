package port

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Health interface {
	HealthHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
}
