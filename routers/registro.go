package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DuocGauss/YapGo/bd"
	"github.com/DuocGauss/YapGo/models"
)

func Registro(ctx context.Context) models.RespApi {
	var t models.Usuario
	var r models.RespApi
	r.Status = 400

	fmt.Println("Entre a Registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}
	if len(t.Email) == 0 {
		r.Message = "Debe especificar el email"
		fmt.Println(r.Message)
		return r
	}
	if len(t.Password) < 6 {
		r.Message = "La contraseña debe tener al menos 6 caracteres"
		fmt.Println(r.Message)
		return r
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		r.Message = "Ya existe un usuario con este email"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		r.Message = "Hubo un error al intentar realizar un registro de usuario " + err.Error()
		fmt.Println(r.Message)
		return r
	}
	if !status {
		r.Message = "No se ha logrado inserta el registro de usuario"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Registro ok"
	fmt.Println(r.Message)
	return r
}
