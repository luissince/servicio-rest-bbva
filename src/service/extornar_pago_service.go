package service

import (
	"database/sql"
	"fmt"
	"servicio-rest-bbva/src/database"
	"servicio-rest-bbva/src/helper"
	"servicio-rest-bbva/src/model"
)

func ExtornarPago(recaudosRq *model.RecaudosRq) model.BodyExtornarPagoResponse {
	if recaudosRq.Detalle.Transaccion.NumeroReferenciaDeuda == "" {
		consultarDeuda := model.BodyExtornarPagoResponse{}
		consultarDeuda.ExtornarPagoResponse.RecaudosRs.Detalle.Respuesta.Codigo = helper.CodigoResputas["3013"].Codigo
		consultarDeuda.ExtornarPagoResponse.RecaudosRs.Detalle.Respuesta.Descripcion = helper.CodigoResputas["3013"].Descripcion
		return consultarDeuda
	}

	db, err := database.CreateConnection()
	if err != nil {
		fmt.Println(err)
		consultarDeuda := model.BodyExtornarPagoResponse{}
		consultarDeuda.ExtornarPagoResponse.RecaudosRs.Detalle.Respuesta.Codigo = helper.CodigoResputas["3002"].Codigo
		consultarDeuda.ExtornarPagoResponse.RecaudosRs.Detalle.Respuesta.Descripcion = helper.CodigoResputas["3002"].Descripcion
		return consultarDeuda
	}

	defer db.Close()

	var resultado string
	query := `exec Bancos.bbva_servicioRest_registroExtorno @codigo,@numero,@monto,@fecha,@hora,'123456','9785'`
	row := db.QueryRowContext(contx, query,
		sql.Named("codigo", recaudosRq.Detalle.Transaccion.NumeroReferenciaDeuda),
		sql.Named("numero", recaudosRq.Detalle.Transaccion.NumeroDocumento),
		sql.Named("monto", recaudosRq.Detalle.Transaccion.ImporteDeudaPagada),
		sql.Named("fecha", recaudosRq.Cabecera.Operacion.FechaOperacion),
		sql.Named("hora", recaudosRq.Cabecera.Operacion.HoraOperacion),
	)
	err = row.Scan(&resultado)

	if err != nil {
		fmt.Println("scan error")
		fmt.Println(err.Error())
		consultarDeuda := model.BodyExtornarPagoResponse{}
		consultarDeuda.ExtornarPagoResponse.RecaudosRs.Detalle.Respuesta.Codigo = helper.CodigoResputas["3002"].Codigo
		consultarDeuda.ExtornarPagoResponse.RecaudosRs.Detalle.Respuesta.Descripcion = helper.CodigoResputas["3002"].Descripcion
		return consultarDeuda
	}

	// Verificar si la clave "claveNoExistente" existe en el mapa
	if valor, ok := helper.CodigoResputas[resultado]; ok {
		consultarDeuda := model.BodyExtornarPagoResponse{}
		consultarDeuda.ExtornarPagoResponse.RecaudosRs.Detalle.Respuesta.Codigo = valor.Codigo
		consultarDeuda.ExtornarPagoResponse.RecaudosRs.Detalle.Respuesta.Descripcion = valor.Descripcion
		return consultarDeuda
	} else {
		fmt.Println("result error")
		fmt.Println(valor)
		consultarDeuda := model.BodyExtornarPagoResponse{}
		consultarDeuda.ExtornarPagoResponse.RecaudosRs.Detalle.Respuesta.Codigo = helper.CodigoResputas["3002"].Codigo
		consultarDeuda.ExtornarPagoResponse.RecaudosRs.Detalle.Respuesta.Descripcion = helper.CodigoResputas["3002"].Descripcion
		return consultarDeuda
	}
}
