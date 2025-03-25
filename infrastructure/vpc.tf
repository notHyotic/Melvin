resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}


# Public subnet
resource "aws_subnet" "subnet_a" {
  vpc_id                  = aws_vpc.main.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "us-east-1a"
  map_public_ip_on_launch = true # Public subnet can assign public IPs
}

# Private subnet
resource "aws_subnet" "subnet_b" {
  vpc_id                  = aws_vpc.main.id
  cidr_block              = "10.0.2.0/24"
  availability_zone       = "us-east-1b"
  map_public_ip_on_launch = false # Private subnet does not assign public IPs
}

# Security group for ECS tasks
resource "aws_security_group" "ecs_sg" {
  name        = "ecs_sg"
  description = "Allow inbound and outbound traffic for ECS tasks"
  vpc_id      = aws_vpc.main.id


  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}