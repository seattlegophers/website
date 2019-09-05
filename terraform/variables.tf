variable "location" {
  default = "westus"
}

variable "projectName" {
  default = "seattleGo"
}

variable "prefix" {
  default = "Dev"
}
variable "sku" {
    default = {
        westus = "16.04-LTS"
        eastus = "18.04-LTS"
    }
}
