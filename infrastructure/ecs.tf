resource "aws_ecs_cluster" "main" {
  name = "melvin-service-cluster"
}

resource "aws_ecs_service" "main" {
  name            = "melvin-service"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.ecs_task.arn
  desired_count   = 1
  launch_type     = "FARGATE" # Or "EC2" if you're using EC2 instances

  network_configuration {
    subnets          = [aws_subnet.subnet_a.id]
    security_groups  = [aws_security_group.ecs_sg.id]
    assign_public_ip = true
  }
}

resource "aws_ecs_task_definition" "ecs_task" {
  family                   = "melvin-task"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"] # This ensures it's Fargate compatible
  cpu                      = "256"
  memory                   = "512"
  execution_role_arn       = aws_iam_role.ecs_execution_role.arn

  container_definitions = jsonencode([{
    name   = "melvin-container"
    image  = "${aws_ecr_repository.melvin-ecr-repo.repository_url}:latest"
    memory = 512
    cpu    = 256

    logConfiguration = {
      logDriver = "awslogs"
      options = {
        "awslogs-group"         = aws_cloudwatch_log_group.ecs_log_group.name # Log group name
        "awslogs-region"        = "us-east-1"                                 # Your region
        "awslogs-stream-prefix" = "melvin-container"                          # Log stream prefix
      }
    }

    secrets = [
      {
        name      = "DISCORD_TOKEN"                     # Name of the environment variable
        valueFrom = aws_ssm_parameter.discord_token.arn # Reference to the SSM Parameter ARN
      }
    ]
  }])
}