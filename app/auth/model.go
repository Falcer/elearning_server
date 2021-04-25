package auth

// Register Model
type Register struct {
	Email        string `json:"email"`
	FullName     string `json:"fullname"`
	Password     string `json:"password"`
	ImageProfile string `json:"image_profile"`
}

// Login Model
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserWithoutPassword model
type UserWithoutPassword struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Email        string `json:"email"`
	FullName     string `json:"fullname"`
	ImageProfile string `json:"image_profile"`
}

// UserWithPassword model
type UserWithPassword struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Email        string `json:"email"`
	FullName     string `json:"fullname"`
	Password     string `json:"password"`
	ImageProfile string `json:"image_profile"`
}

// RoleInput model
type RoleInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

// RoleOutput model
type RoleOutput struct {
	ID          string  `json:"id" bson:"_id,omitempty"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

// UserWithRole model
type UserWithRole struct {
	ID           string       `json:"id" bson:"_id,omitempty"`
	Email        string       `json:"email"`
	FullName     string       `json:"fullname"`
	ImageProfile string       `json:"image_profile"`
	Role         []RoleOutput `json:"roles"`
}

// UserRoleInput model
type UserRoleInput struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}

// UserToken struct
type UserToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
