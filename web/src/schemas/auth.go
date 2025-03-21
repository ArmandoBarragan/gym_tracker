package schemas

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"github.com/ArmandoBarragan/gym_tracker/conf"
	"github.com/ArmandoBarragan/gym_tracker/src/models"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

type CreateUser struct {
	Username             string
	Password             string
	PasswordConfirmation string
}

func (user *CreateUser) PasswordIsValid() bool {
	var re *regexp.Regexp = regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>_-]`)
	var reDigit *regexp.Regexp = regexp.MustCompile("[0-9]")
	var passwordContainsSpecialCharacter bool = re.MatchString(user.Password)
	var passwordContainsDigit bool = reDigit.MatchString(user.Password)
	return passwordContainsSpecialCharacter && passwordContainsDigit
}

func (user *CreateUser) Create() error {
	hashedPassword, err := hashPassword(user.Password)

	if err != nil {
		return err
	}

	newUser := &models.User{
		Username: user.Username,
		Password: hashedPassword,
	}

	result := conf.DB.Create(&newUser)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

type LogIn struct {
	Username string
	Password string
}
