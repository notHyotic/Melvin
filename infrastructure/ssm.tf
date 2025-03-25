resource "aws_ssm_parameter" "discord_token" {
  name        = "/melvin/discord-token"
  description = "discord bot token"
  type        = "SecureString"
  value       = "uninitialized"
  lifecycle {
    ignore_changes = [value] # Terraform will ignore changes to the value field
  }
}