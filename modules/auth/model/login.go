package model

type (
	// LoginRequest login
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// LoginResponse login response
	LoginResponse struct {
		Email        string `json:"email"`
		TokenID      string `json:"idToken"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    string `json:"expiresIn"`
		Registered   string `json:"registered"`
		DisplayName  string `json:"displayName"`
		LocalID      string `json:"localId"`
		Kind         string `json:"kind"`
	}
)
