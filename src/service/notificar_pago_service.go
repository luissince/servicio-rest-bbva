package service

import (
	"database/sql"
	"fmt"
	"servicio-rest-bbva/src/database"
	"servicio-rest-bbva/src/helper"
	"servicio-rest-bbva/src/model"
)

func NotificarPago(recaudosRq *model.RecaudosRq) model.BodyNotificarPagoResponse {
	//body.NotificarPago.RecaudosRq.Detalle.Transaccion
	if recaudosRq.Detalle.Transaccion.NumeroReferenciaDeuda == "" {
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
	query := `exec Bancos.bbva_servicioRest_registroOperacion @codigo,@numero,@monto,@fecha,@hora,'123456','9785'`
	row := db.QueryRowContext(contx, query,
		sql.Named("codigo", recaudosRq.Detalle.Transaccion.NumeroReferenciaDeuda),
		sql.Named("numero", recaudosRq.Detalle.Transaccion.NumeroDocumento),
		sql.Named("monto", recaudosRq.Detalle.Transaccion.ImporteDeudaPagada),
		sql.Named("fecha", recaudosRq.Cabecera.Operacion.FechaOperacion),
		sql.Named("hora", recaudosRq.Cabecera.Operacion.HoraOperacion),
	)
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
