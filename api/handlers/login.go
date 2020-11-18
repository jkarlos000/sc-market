package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/hashicorp/go-hclog"
	"github.com/jkarlos000/sc-market/api/conf"
	"github.com/jkarlos000/sc-market/client/cmd/client"
	proto "github.com/jkarlos000/sc-market/client/internal/infrastructure/delivery/grpc/proto"
	"google.golang.org/grpc"
	"time"
)

type login struct {
	l hclog.Logger
}
type Login struct {
	Email	string `json:"email" validate:"required,min=4,max=100"`
	Password string `json:"password" validate:"required,min=4,max=100"`
}

type usuario struct {
	ID int
	Email string
	Names string
	Lastname string
	Type string
	Address string
	Telephone string
}

func NewLogin() *login {
	l := hclog.Default()
	return &login{l: l}
}

func (lh *login) SetupRoutes(router fiber.Router)  {
	router.Post("/clientes")
	router.Post("/usuarios")
	router.Post("/proveedores")
}

func (lh *login) loginClient(c *fiber.Ctx)  {
	var login Login
	if err := c.BodyParser(&login); err != nil {
		c.SendStatus(fiber.StatusUnauthorized)
		c.JSON(fiber.Map{"status":"failed", "message":"Datos incorrectos" })
		return
	}
	cc, errDial := grpc.Dial(conf.Config("CLIENT_ADDR")+":"+conf.Config("CLIENT_PORT"), grpc.WithInsecure())
	if errDial != nil {
		lh.l.Error("Login: No se pudo conectar al servicio de Cliente", "error", errDial)
	}
	defer cc.Close()
	clientGrpc := proto.NewClientServiceClient(cc)
	connection := client.NewClientCli(clientGrpc)
	_, user, errLogin := connection.Login(login.Email, login.Password)
	if errLogin != nil {
		lh.l.Error("Login: Datos incorrectos", "error", errLogin)
	}
	usr := &usuario{
		ID:        int(user.ID),
		Email:     user.Email,
		Names:      user.Names,
		Lastname:  user.LastNames,
		Type:      "cliente",
		Address:   user.Address,
		Telephone: user.Telephone,
	}
	t, err := lh.generateToken(usr)
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		c.JSON(fiber.Map{"status":"failed", "message":"Algo salió mal..." })
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Success login", "token": t})
}

func (lh *login) generateToken(user *usuario) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["nombre"] = user.Names
	claims["apellido"] = user.Lastname
	claims["tipo"] = user.Type
	claims["email"] = user.Email
	claims["direccion"] = user.Address
	claims["telefono"] = user.Telephone
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(conf.Config("SECRET")))
}


