package Admin

import (
	"backend_jamijabal/encryption_decryption"
	"backend_jamijabal/entities"
	"backend_jamijabal/environment"
	"backend_jamijabal/models/Admin"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func getDataAdmin(c echo.Context) error {
	a := new(entities.Admin)
	if err := c.Bind(a); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	admin := entities.Admin{
		Username: a.Username,
		Password: a.Password,
	}

	admins := Admin.GetAdmins()
	adminSelected := entities.Admin{}
	for _, val := range admins {
		decryptPass := encryption_decryption.DoDecryption(val.Password)
		if admin.Password == decryptPass && admin.Username == val.Username {
			adminSelected.Name = val.Name
			adminSelected.Id = val.Id
			adminSelected.Username = val.Username
			adminSelected.Password = decryptPass
			adminSelected.Photo = val.Photo
			adminSelected.Status = val.Status
			break
		}
	}

	var (
		key string
		t   *jwt.Token
		s   string
	)

	key = environment.EnvVariable("jwtKey")
	t = jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"iss": "my-auth-server",
			"sub": "john",
			"foo": 2,
		},
	)
	s, _ = t.SignedString(key)

	return c.JSON(http.StatusOK, s)
}
