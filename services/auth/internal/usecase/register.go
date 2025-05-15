package usecase

import (
	"context"
	"github.com/nazarovnick/chat-platform/services/auth/internal/entity/user"
	pkgerrors "github.com/nazarovnick/chat-platform/services/auth/pkg/errors"
	"time"
)

// registerUseCase handles the logic for creating a new user account.
type registerUseCase struct {
	users       UserRepo
	hasher      user.PasswordHasher
	defaultRole user.Role
}

// registerUseCase handles the logic for creating a new user account.
func NewRegisterUseCase(users UserRepo, hasher user.PasswordHasher, defaultRole user.Role) RegisterUseCase {
	return &registerUseCase{users: users, hasher: hasher, defaultRole: defaultRole}
}

// Execute registers a new user account.
//
// Steps:
//  1. Create and validate the login value object.
//  2. Check if the login is already in use.
//  3. Create and validate the password value object.
//  4. Hash the password using the provided hasher.
//  5. Assign the role (either from input or use the default).
//  6. Create and store the user entity in the repository.
//
// Returns the new user's ID or an error if the registration fails.
func (uc *registerUseCase) Execute(ctx context.Context, in *RegisterInput) (_ *RegisterOutput, err error) {
	const op = "usecase.registerUseCase.Execute"
	defer func() {
		if err != nil {
			err = pkgerrors.Wrap(op, err)
		}
	}()

	// Step 1. Create and validate login
	login, err := user.NewLogin(in.Login)
	if err != nil {
		return nil, pkgerrors.WrapWith(ErrInvalidLogin, err)
	}
	if err := login.Validate(); err != nil {
		return nil, pkgerrors.WrapWith(ErrInvalidLogin, err)
	}

	// Step 2. Check if login is already taken
	_, err = uc.users.GetByLogin(ctx, login)
	if err == nil {
		return nil, pkgerrors.WrapWith(ErrLoginAlreadyExists, err)
	}

	// Step 3. Create and validate password
	password, err := user.NewPassword(in.Password)
	if err != nil {
		return nil, pkgerrors.WrapWith(ErrInvalidPassword, err)
	}
	if err := password.Validate(); err != nil {
		return nil, pkgerrors.WrapWith(ErrInvalidPassword, err)
	}

	// Step 4. Hash the password
	hashed, err := uc.hasher.Hash(password)
	if err != nil {
		return nil, pkgerrors.WrapWith(ErrHashingPassword, err)
	}

	// Step 5. Assign role
	var role user.Role
	if in.Role != "" {
		role, err = user.NewRole(in.Role)
		if err != nil {
			return nil, pkgerrors.WrapWith(ErrInvalidRole, err)
		}
	} else {
		role = uc.defaultRole
	}

	// Step 6. Create and store user
	id := user.NewUserID()
	u := user.NewUser(id, login, hashed, role, false, time.Now())
	if err := uc.users.Create(ctx, u); err != nil {
		return nil, pkgerrors.WrapWith(ErrUserCreationFailed, err)
	}

	return &RegisterOutput{UserID: id}, nil
}
