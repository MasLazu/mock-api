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

variable "ssh_private_key_path" {
  default = "~/.ssh/id_rsa"
}

variable "ssh_public_key_path" {
  default = "~/.ssh/id_rsa.pub"
}
