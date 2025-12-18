// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

// terraform-provider-f5xc provides Terraform resources for managing F5 Distributed Cloud services.
// For documentation, see https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs
//
// Documentation features OneOf property grouping for mutually exclusive arguments,
// improving clarity by grouping related properties with a single explanatory note.
//
// Version tagging and releases are automated via CI/CD using semantic versioning.
// Release workflow uses forked action with updated dependencies for improved cache performance.
// MCP server documentation is available at https://github.com/robinmordasiewicz/terraform-provider-f5xc/tree/main/mcp-server

package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/f5xc/terraform-provider-f5xc/internal/provider"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version string = "dev"

	// goreleaser can pass other information to the main package, such as the specific commit
	// https://goreleaser.com/cookbooks/using-main.version/
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/robinmordasiewicz/f5xc",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
