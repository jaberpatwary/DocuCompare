package validation

type CreateUser struct {
	Name      string `json:"name" validate:"required,max=100" example:"John Doe"`
	Email     string `json:"email" validate:"required,email,max=150" example:"john@example.com"`
	Phone     string `json:"phone" validate:"omitempty,max=20" example:"01712345678"`
	Password  string `json:"password" validate:"required,min=8,max=100" example:"password123"`
	AvatarURL string `json:"avatar_url,omitempty" validate:"omitempty,url" example:"https://example.com/avatar.jpg"`
	Status    string `json:"status,omitempty" validate:"omitempty,oneof=active inactive banned" example:"active"`
}

type UpdateUser struct {
	Name      string `json:"name,omitempty" validate:"omitempty,max=100" example:"John Doe"`
	Email     string `json:"email,omitempty" validate:"omitempty,email,max=150" example:"john@example.com"`
	Phone     string `json:"phone,omitempty" validate:"omitempty,max=20" example:"01712345678"`
	Password  string `json:"password,omitempty" validate:"omitempty,min=8,max=100" example:"password123"`
	AvatarURL string `json:"avatar_url,omitempty" validate:"omitempty,url" example:"https://example.com/avatar.jpg"`
	Status    string `json:"status,omitempty" validate:"omitempty,oneof=active inactive banned" example:"active"`
}

type UpdateUser2 struct {
	Name      string `json:"name,omitempty" validate:"omitempty,max=100" example:"Jaber Patwary"`
	Email     string `json:"email,omitempty" validate:"omitempty,email,max=150" example:"jaber@example.com"`
	Phone     string `json:"phone,omitempty" validate:"omitempty,max=20" example:"01712345678"`
	Password  string `json:"password_hash,omitempty" validate:"omitempty,min=8,max=100" example:"somehashedpassword"`
	AvatarURL string `json:"avatar_url,omitempty" validate:"omitempty,url" example:"https://example.com/avatar.jpg"`
	Status    string `json:"status,omitempty" validate:"omitempty,oneof=active inactive banned" example:"active"`
	CreatedAt string `json:"created_at,omitempty" validate:"omitempty" example:"2025-12-02T19:03:20Z"`
	UpdatedAt string `json:"updated_at,omitempty" validate:"omitempty" example:"2025-12-02T19:03:20Z"`
}

type UpdatePassOrVerify struct {
	Password      string `json:"password,omitempty" validate:"omitempty,min=8,max=20,password" example:"password1"`
	VerifiedEmail bool   `json:"verified_email" swaggerignore:"true" validate:"omitempty,boolean"`
}

type QueryUser struct {
	Page   int    `validate:"omitempty,number,max=50"`
	Limit  int    `validate:"omitempty,number,max=50"`
	Search string `validate:"omitempty,max=50"`
}
