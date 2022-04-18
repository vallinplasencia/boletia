package models

const (
	// KeyUserContext clave en el contexto de gin para obtener el usuario logueado en la app
	KeyUserContext string = "user_in"
)

// AuthUser datos del usuario logueado y q se obtiene a traves del contexto de gin con esta clave KeyUserContext
type AuthUser struct {
	Username string
	Roles    []RoleType
}

// IsAnonymous retorna true si el usuario no esta logueado
func (u *AuthUser) IsAnonymous() bool {
	for _, v := range u.Roles {
		if v == RoleAnonymous {
			return true
		}
	}
	return false
}

// IsAdmin retorna true si el usuario es administrator
func (u *AuthUser) IsAdmin() bool {
	for _, v := range u.Roles {
		if v == RoleAdmin {
			return true
		}
	}
	return false
}

// === types === //

// RoleType users role
type RoleType string

const (
	RoleAnonymous RoleType = "anonymous"
	RoleUser      RoleType = "user"
	RoleAdmin     RoleType = "admin"
)

// ToRoleFromString retorna un tipo de rol segun s.
func ToRoleFromString(s string) RoleType {
	values := map[string]RoleType{
		"anonymous": RoleAnonymous,
		"user":      RoleUser,
		"admin":     RoleAdmin,
	}
	r, ok := values[s]
	if ok {
		return r
	}
	return RoleAnonymous
}
