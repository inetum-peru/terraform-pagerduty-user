resource "pagerduty_user" "this" {
  name  = var.name
  email = var.email
}

resource "pagerduty_user_contact_method" "email" {
  user_id = pagerduty_user.this.id
  type    = "email_contact_method"
  address = var.email
  label   = "Work"
}

resource "pagerduty_user_contact_method" "phone" {
  user_id      = pagerduty_user.this.id
  type         = "phone_contact_method"
  label        = "Work"
  country_code = "+51"
  address      = var.mobile
}
