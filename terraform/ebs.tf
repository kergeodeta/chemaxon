resource "aws_ebs_volume" "docker_alm" {
  availability_zone = aws_instance.chemaxon_alm.availability_zone
  size = 10
  type = "gp2"
}

resource "aws_ebs_volume" "data_alm" {
  availability_zone = aws_instance.chemaxon_alm.availability_zone
  size = 10
  type = "gp2"
}

resource "aws_ebs_volume" "docker_app" {
  availability_zone = aws_instance.chemaxon_app.availability_zone
  size = 5
  type = "gp2"
}

resource "aws_ebs_volume" "data_app" {
  availability_zone = aws_instance.chemaxon_app.availability_zone
  size = 5
  type = "gp2"
}

resource "aws_volume_attachment" "docker_alm_attach" {
  device_name = "/dev/xvdp"
  instance_id = aws_instance.chemaxon_alm.id
  volume_id = aws_ebs_volume.docker_alm.id
}

resource "aws_volume_attachment" "data_alm_attach" {
  device_name = "/dev/xvdq"
  instance_id = aws_instance.chemaxon_alm.id
  volume_id = aws_ebs_volume.data_alm.id
}

resource "aws_volume_attachment" "docker_app_attach" {
  device_name = "/dev/xvdr"
  instance_id = aws_instance.chemaxon_app.id
  volume_id = aws_ebs_volume.docker_app.id
}

resource "aws_volume_attachment" "data_app_attach" {
  device_name = "/dev/xvds"
  instance_id = aws_instance.chemaxon_app.id
  volume_id = aws_ebs_volume.data_app.id
}


