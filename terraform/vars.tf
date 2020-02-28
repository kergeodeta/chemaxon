variable "aws_access_key" {
  type = string
  description = ""
}

variable "aws_secret_key" {
  type = string
  description = ""
}

variable "aws_region" {
  type = string
  description = ""
  default = "eu-central-1"
}

variable "ami_ids" {
  type = map(string)
  default = {
    eu-central-1 = "ami-0718a1ae90971ce4d"
    eu-north-1 = "ami-09c1e351d7adfb5d8"
    eu-west-1 = "ami-07042e91d04b1c30d"
    eu-west-2 = "ami-0904845532eecd472"
    eu-west-3 = "ami-0c367ebddcf279dc6"
  }
}

variable "ssh_public_key_path" {
  type = string
  description = ""
  default = "./ssh/chemkey.pub"
}