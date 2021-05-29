<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.13 |
| <a name="requirement_pagerduty"></a> [pagerduty](#requirement\_pagerduty) | >=1.9.6 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_pagerduty"></a> [pagerduty](#provider\_pagerduty) | >=1.9.6 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [pagerduty_user.this](https://registry.terraform.io/providers/PagerDuty/pagerduty/latest/docs/resources/user) | resource |
| [pagerduty_user_contact_method.email](https://registry.terraform.io/providers/PagerDuty/pagerduty/latest/docs/resources/user_contact_method) | resource |
| [pagerduty_user_contact_method.phone](https://registry.terraform.io/providers/PagerDuty/pagerduty/latest/docs/resources/user_contact_method) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_email"></a> [email](#input\_email) | email of user | `string` | n/a | yes |
| <a name="input_mobile"></a> [mobile](#input\_mobile) | nro mobile of user | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | name of user | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_user"></a> [user](#output\_user) | instance user |
| <a name="output_user_id"></a> [user\_id](#output\_user\_id) | instance user |
<!-- END_TF_DOCS -->