package signup_utils

import (
	"appcraft_web/signup_models"
	"errors"
)

func ValidateSignup(signup signup_models.Signup) error {
	if signup.UserName == "" {
		return errors.New("username is required")
	}
	if signup.FirstName == "" {
		return errors.New("first name is required")
	}
	if signup.LastName == "" {
		return errors.New("last name is required")
	}
	if signup.Email == "" {
		return errors.New("email is required")
	}
	if signup.PhoneNo == 0 {
		return errors.New("phone number is required")
	}
	if signup.Password == "" {
		return errors.New("password is required")
	}
	if signup.ConfirmPassword == "" {
		return errors.New("confirm password is required")
	}
	if signup.Password != signup.ConfirmPassword {
		return errors.New("password and confirm password do not match")
	}
	return nil
}