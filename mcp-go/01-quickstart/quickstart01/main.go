// Copyright 2023 igevin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"Calculator Demo",
		"0.1",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	s.AddTool(createCalculateTool(), calculateHandler)

	fmt.Println("Start serving...")
	err := server.ServeStdio(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Stopped serving")

}

func createCalculateTool() mcp.Tool {
	return mcp.NewTool(
		"calculate",
		mcp.WithDescription("Perform basic arithmetic operations"),
		mcp.WithString("operation",
			mcp.Required(),
			mcp.Description("The operation to perform (add, subtract, multiply, divide)"),
			mcp.Enum("add", "subtract", "multiply", "divide"),
		),
		mcp.WithNumber("x",
			mcp.Required(),
			mcp.Description("First Number"),
		),
		mcp.WithNumber("y",
			mcp.Required(),
			mcp.Description("Second Number"),
		),
	)
}

func calculateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	op := request.Params.Arguments["operation"].(string)
	x := request.Params.Arguments["x"].(float64)
	y := request.Params.Arguments["y"].(float64)
	var res float64
	switch op {
	case "add":
		res = x + y
	case "subtract":
		res = x - y
	case "multiply":
		res = x * y
	case "divide":
		res = x / y
	}
	return mcp.NewToolResultText(fmt.Sprintf("%.2f", res)), nil
}
