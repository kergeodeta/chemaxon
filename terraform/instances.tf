resource "aws_key_pair" "chemaxon_hw" {
  key_name = "chemkey"
  public_key = file(abspath(var.ssh_public_key_path))
}

// EC2 instance for Application Lifecycle Management applications
resource "aws_instance" "chemaxon_alm" {
  ami = lookup(var.ami_ids, var.aws_region)
  instance_type = "t2.micro"
  key_name = aws_key_pair.chemaxon_hw.key_name

  provisioner "local-exec" {
    command = "echo 'chemaxon_alm ansible_host=${aws_instance.chemaxon_alm.public_ip} ansible_user=ubuntu docker_dev=/dev/xvdp data_dev=/dev/xvdq' >> ./ansible/inventory"
  }

  tags = {
    Name = "chemaxon_alm"
  }
}

// EC2 instance for run the application and the underlying database
resource "aws_instance" "chemaxon_app" {
  ami = lookup(var.ami_ids, var.aws_region)
  instance_type = "t2.micro"
  key_name = aws_key_pair.chemaxon_hw.key_name

  provisioner "local-exec" {
    command = "echo 'chemaxon_app ansible_host=${aws_instance.chemaxon_app.public_ip} ansible_user=ubuntu  docker_dev=/dev/xvdr data_dev=/dev/xvds' >> ./ansible/inventory"
  }

  tags = {
    Name = "chemaxon_app"
  }
}
