variable "do_token" {
  type        = string
  description = "Digital Ocean API token"
  sensitive   = true
}

variable "region" {
  default = "sgp1"
}

variable "image" {
  default = "ubuntu-22-04-x64"
}

variable "size" {
  default = "s-1vcpu-1gb"
}
