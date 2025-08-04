package user

import "ToDo/internal/domain/user"

// ToUser преобразует CreateUserDTO в доменную сущность user.User.
func ToUser(dto CreateUserDTO) user.User {
	return user.NewUser(dto.Name, dto.Email, dto.Password)
}

// ToUserDTO преобразует доменную сущность user.User в DTO для передачи клиенту.
func ToUserDTO(u user.User) DTO {
	return DTO{
		ID:    u.ID().String(),
		Name:  u.Name(),
		Email: u.Email(),
	}
}
