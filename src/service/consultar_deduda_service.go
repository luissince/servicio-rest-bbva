package service

import (
	"context"
	"database/sql"
	"fmt"
	"servicio-rest-bbva/main/src/database"
	"servicio-rest-bbva/main/src/helper"
	"servicio-rest-bbva/main/src/model"
)

var contx = context.Background()

func ObtenerDeuda(codigo string) (model.ConsultarDeuda, string) {

	if codigo == "" {
		consultarDeudaError := model.ConsultarDeuda{}
		consultarDeudaError.Detalle.Respuesta.Codigo = helper.CodigoResputas["0101"].Codigo
		consultarDeudaError.Detalle.Respuesta.Descripcion = helper.CodigoResputas["0101"].Descripcion
		return consultarDeudaError, "ok"
	}

	db, err := database.CreateConnection()
	if err != nil {
		fmt.Println(err)
		consultarDeudaError := model.ConsultarDeuda{}
		consultarDeudaError.Detalle.Respuesta.Codigo = helper.CodigoResputas["3002"].Codigo
		consultarDeudaError.Detalle.Respuesta.Descripcion = helper.CodigoResputas["3002"].Descripcion
		return consultarDeudaError, "ok"
	}

	defer db.Close()

	queryCabecera := `exec bancos.bbva_servicioRest_consultaDeuda 1,@codigo`

	row := db.QueryRowContext(contx, queryCabecera, sql.Named("codigo", codigo))

	transaccion := &model.Transaccion{}

	err = row.Scan(
		&transaccion.NumeroReferenciaDeuda,
		&transaccion.NombreCliente,
		&transaccion.NumeroOperacionEmpresa,
		&transaccion.IndMasDeuda,
		&transaccion.CantidadDocsDeuda,
		&transaccion.DatosEmpresa,
	)

	if err != nil {
		fmt.Println(err)
		consultarDeudaError := model.ConsultarDeuda{}
		consultarDeudaError.Detalle.Respuesta.Codigo = helper.CodigoResputas["0101"].Codigo
		consultarDeudaError.Detalle.Respuesta.Descripcion = helper.CodigoResputas["0101"].Descripcion
		return consultarDeudaError, "ok"
	}

	queryDetalle := `exec bancos.bbva_servicioRest_consultaDeuda 2,@codigo`

	rows, err := db.QueryContext(contx, queryDetalle, sql.Named("codigo", codigo))

	if err != nil {
		fmt.Println(err)
		consultarDeudaError := model.ConsultarDeuda{}
		consultarDeudaError.Detalle.Respuesta.Codigo = helper.CodigoResputas["3009"].Codigo
		consultarDeudaError.Detalle.Respuesta.Descripcion = helper.CodigoResputas["3009"].Descripcion
		return consultarDeudaError, "ok"
	}

	defer rows.Close()

	documentos := []model.Documento{}

	for rows.Next() {
		documento := model.Documento{}

		err := rows.Scan(
			&documento.Numero,
			&documento.Descripcion,
			&documento.FechaEmision,
			&documento.FechaVencimiento,
			&documento.ImporteDeuda,
			&documento.ImporteDeudaMinima,
			&documento.IndicadorRestriccPago,
			&documento.CantidadSubconceptos,
		)

		if err != nil {
			fmt.Println(err)
		}

		documentos = append(documentos, documento)
	}

	if len(documentos) == 0 {
		consultarDeudaError := model.ConsultarDeuda{}
		consultarDeudaError.Detalle.Respuesta.Codigo = helper.CodigoResputas["3009"].Codigo
		consultarDeudaError.Detalle.Respuesta.Descripcion = helper.CodigoResputas["3009"].Descripcion
		return consultarDeudaError, "ok"
	}

	transaccion.ListaDocumentos = documentos

	consultarDeuda := model.ConsultarDeuda{}
	consultarDeuda.Detalle.Respuesta.Codigo = helper.CodigoResputas["0001"].Codigo
	consultarDeuda.Detalle.Respuesta.Descripcion = helper.CodigoResputas["0001"].Descripcion
	consultarDeuda.Detalle.Transaccion = transaccion

	return consultarDeuda, "ok"
}
