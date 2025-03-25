resource "aws_ecr_repository" "melvin-ecr-repo" {
  name                 = "melvin"
  image_tag_mutability = "MUTABLE" # You can use "IMMUTABLE" if you want to prevent overwriting images
}