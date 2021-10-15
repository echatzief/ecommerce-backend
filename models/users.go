package models

type User struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	ResetToken string `json:"resetToken"`
	ConfirmationToken string `json:"confirmationToken"`
	Confirmed bool `json:"confirmed"`
	Blocked bool `json:"blocked"`
	CellPhone string `json:"cellPhone"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	StripeAccountID string `json:"stripeAccountId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
