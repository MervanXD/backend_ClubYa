package cuenta

type CuentaGmailDTO struct {
	Username      *string `json:"username"`
	Email         string  `json:"email"`
	EmailVerified bool    `json:"email_verified"`
}
