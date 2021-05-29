output "user" {
  description = "instance user"
  value       = pagerduty_user.this
}

output "user_id" {
  description = "instance user"
  value       = pagerduty_user.this.id
}
