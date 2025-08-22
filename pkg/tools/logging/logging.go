// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logging

import (
	"context"

	"github.com/GoogleCloudPlatform/gke-mcp/pkg/config"
	"github.com/mark3labs/mcp-go/server"
)

// Install adds GCP logging related tools to an MCP server.
func Install(_ context.Context, s *server.MCPServer, c *config.Config) error {
	installQueryLogsTool(s, c)
	installGetLogSchemas(s)
	installNodeLogQueryTool(s, c)
	return nil
}
