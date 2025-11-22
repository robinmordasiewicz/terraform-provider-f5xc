terraform {
  required_providers {
    f5xc = {
      source  = "f5xc/f5xc"
      version = "~> 0.1"
    }
  }
}

provider "f5xc" {
  # API token can be set via F5XC_API_TOKEN environment variable
  # api_token = "your-api-token-here"

  # API URL defaults to https://console.ves.volterra.io/api
  # Can be set via F5XC_API_URL environment variable
  # api_url = "https://console.ves.volterra.io/api"
}
