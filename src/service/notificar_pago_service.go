package service

import (
	"database/sql"
	"fmt"
	"servicio-rest-bbva/main/src/database"
	"servicio-rest-bbva/main/src/helper"
	"servicio-rest-bbva/main/src/model"
)

func NotificarPago(transaccion *model.Transaccion) model.BodyNotificarPagoResponse {
	if transaccion.NumeroReferenciaDeuda == "" {
		consultarDeuda := model.BodyNotificarPagoResponse{}
		consultarDeuda.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Codigo = helper.CodigoResputas["3013"].Codigo
		consultarDeuda.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Descripcion = helper.CodigoResputas["3013"].Descripcion
		return consultarDeuda
	}

	db, err := database.CreateConnection()
	if err != nil {
		fmt.Println(err)
		consultarDeuda := model.BodyNotificarPagoResponse{}
		consultarDeuda.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Codigo = helper.CodigoResputas["3002"].Codigo
		consultarDeuda.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Descripcion = helper.CodigoResputas["3002"].Descripcion
		return consultarDeuda
	}

	defer db.Close()

	var resultado string
	query := `exec Bancos.bbva_servicioRest_registroOperacion 1,@codigo,@numero,@monto`
	row := db.QueryRowContext(contx, query, sql.Named("codigo", transaccion.NumeroReferenciaDeuda), sql.Named("numero", transaccion.NumeroDocumento), sql.Named("monto", transaccion.ImporteDeudaPagada))
	err = row.Scan(&resultado)

	if err != nil {
		fmt.Println("1")
		consultarDeuda := model.BodyNotificarPagoResponse{}
		consultarDeuda.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Codigo = helper.CodigoResputas["3002"].Codigo
		consultarDeuda.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Descripcion = helper.CodigoResputas["3002"].Descripcion
		return consultarDeuda
	}

	if resultado == "2" {
		fmt.Println("2")
		consultarDeuda := model.BodyNotificarPagoResponse{}
		consultarDeuda.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Codigo = helper.CodigoResputas["0106"].Codigo
		consultarDeuda.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Descripcion = helper.CodigoResputas["0106"].Descripcion
		return consultarDeuda
	}

	if resultado == "3" {
		fmt.Println("3")
		consultarDeuda := model.BodyNotificarPagoResponse{}
		consultarDeuda.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Codigo = helper.CodigoResputas["0290"].Codigo
		consultarDeuda.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Descripcion = helper.CodigoResputas["0290"].Descripcion
		return consultarDeuda
	}

	if resultado == "0" {
		fmt.Println("4")
		consultarDeuda := model.BodyNotificarPagoResponse{}
		consultarDeuda.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Codigo = helper.CodigoResputas["0106"].Codigo
		consultarDeuda.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Descripcion = helper.CodigoResputas["0106"].Descripcion
		return consultarDeuda
	}

	consultarDeuda := model.BodyNotificarPagoResponse{}
	consultarDeuda.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Codigo = helper.CodigoResputas["0001"].Codigo
	consultarDeuda.NotificarPagoResponse.RecaudosRs.Detalle.Respuesta.Descripcion = helper.CodigoResputas["0001"].Descripcion
	return consultarDeuda
}
