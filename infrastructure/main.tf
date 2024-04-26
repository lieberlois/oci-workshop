# Azure Provider source and version being used
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.101.0"
    }
  }
}

provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "rg" {
  name = "ociworkshop"
  location = "West Europe"
}

resource "azurerm_container_registry" "acr" {
  sku = "Basic"
  name = "ociworkshopacr"
  resource_group_name = azurerm_resource_group.rg.name
  location = azurerm_resource_group.rg.location

  admin_enabled = true
}

output "acr_host" {
  value = azurerm_container_registry.acr.login_server
}

output "admin_username" {
  value = azurerm_container_registry.acr.admin_username
}

output "admin_password" {
  value = azurerm_container_registry.acr.admin_password
  sensitive = true
}