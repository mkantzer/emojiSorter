output "url" {
  value       = aws_apprunner_service.this.service_url
  description = "The URL of the service"
}
