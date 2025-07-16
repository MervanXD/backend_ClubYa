package solicitud

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/models/persona"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/awss3"
	"github.com/MervanXD/backend_ClubYa/internal/pkgs/utils"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type solicitudRepositoryDB struct{}

func NewSolicitudRepositoryDB() SolicitudRepository {
	return &solicitudRepositoryDB{}
}

func (r *solicitudRepositoryDB) ObtenerSolicitudesMembresia() ([]SolicitudDTO, error) {
	query := "call ingesoft.obtenerSolicitudesConTitulares()"
	rows, err := database.DB.Query(query)
	if err != nil {
		logs.Logger.Println("Error al obtener las solicitudes de membresia: ", err)
		return nil, err
	}
	defer rows.Close()
	var solicitudes []SolicitudDTO
	for rows.Next() {
		var solicitud SolicitudDTO
		if err := rows.Scan(&solicitud.Id, &solicitud.Fecha, &solicitud.Estado, &solicitud.IdTitular, &solicitud.Nombres, &solicitud.Apellidos); err != nil {
			logs.Logger.Println("Error al escanear solicitud de membresia: ", err)
			return nil, err
		}
		solicitudes = append(solicitudes, solicitud)
	}
	return solicitudes, nil
}

func (r *solicitudRepositoryDB) ActualizarEstadoSolicitud(id int, nuevoEstado string) error {
	estado := strings.Title(strings.ToLower(nuevoEstado))

	query := `CALL ActualizarEstadoSolicitud(?, ?)`
	result, err := database.DB.Exec(query, id, estado)
	if err != nil {
		logs.Logger.Println("Error al ejecutar el procedimiento:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logs.Logger.Println("Error al obtener rowsAffected:", err)
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}

func (r *solicitudRepositoryDB) ObtenerDatosSolicitudPorId(idSolicitud int) (*SolicitudMembresia, error) {
	query := "CALL ObtenerDatosSolicitudPorId(?)"
	rows := database.DB.QueryRow(query, idSolicitud)
	var solicitud SolicitudMembresia
	err := rows.Scan(&solicitud.Id, &solicitud.Fecha, &solicitud.Estado)
	if err != nil {
		logs.Logger.Println("Error al ejecutar el procedimiento:", err)
		return nil, err
	}
	return &solicitud, nil
}

func (r *solicitudRepositoryDB) ObtenerFamiliaresPorIdSolicitud(idSolicitud int) ([]persona.Familiar, error) {
	query := "CALL ObtenerFamiliaresPorSolicitud(?)"
	rows, err := database.DB.Query(query, idSolicitud)
	if err != nil {
		logs.Logger.Println("Error al obtener los datos de los familiares: ", err)
		return nil, err
	}
	defer rows.Close()
	var familiares []persona.Familiar
	for rows.Next() {
		var familia persona.Familiar
		if err := rows.Scan(&familia.Id, &familia.Nombre, &familia.Apellidos, &familia.TipoDocumento, &familia.NroDocumento, &familia.TipoFamiliar, &familia.FechaNacimiento); err != nil {
			logs.Logger.Println("Error al escanear al familiar: ", err)
			return nil, err
		}
		familiares = append(familiares, familia)
	}
	return familiares, nil
}

func (r *solicitudRepositoryDB) ObtenerDatosPersonaPorIdSolicitud(idSolicitud int) (*persona.Titular, error) {
	var urlDocumentoIdentidad, urlCartaRecomendacion1, urlCartaRecomendacion2 string
	query := "CALL ObtenerDatosPersonaPorIdSolicitud(?)"
	rows := database.DB.QueryRow(query, idSolicitud)
	var persona persona.Titular
	err := rows.Scan(&persona.Nombre, &persona.Apellidos, &persona.Sexo, &persona.TipoDocumento, &persona.NroDocumento, &persona.FechaNacimiento, &persona.TipoVia, &persona.Direccion, &persona.Distrito, &persona.Pais, &persona.Provincia, &persona.Telefono, &persona.Referencia, &persona.Ocupacion,
		&persona.IngresoPromedio, &persona.NombreEmpresa, &persona.DireccionEmpresa, &urlDocumentoIdentidad, &urlCartaRecomendacion1, &urlCartaRecomendacion2)
	if err != nil { //email lo estoy colocando en referencia , por ahora
		logs.Logger.Println("Error al ejecutar el procedimiento:", err)
		return nil, err
	}
	filnameDNI := awss3.ExtractFileNameFromURL(urlDocumentoIdentidad)
	filenameCarta1 := awss3.ExtractFileNameFromURL(urlCartaRecomendacion1)
	filenameCarta2 := awss3.ExtractFileNameFromURL(urlCartaRecomendacion2)
	persona.NombreDocumentoIdentidad = utils.NullString{NullString: sql.NullString{String: filnameDNI, Valid: true}}
	persona.NombreCartaRecomendacion1 = utils.NullString{NullString: sql.NullString{String: filenameCarta1, Valid: true}}
	persona.NombreCartaRecomendacion2 = utils.NullString{NullString: sql.NullString{String: filenameCarta2, Valid: true}}

	persona.UrlDocumentoIdentidad, err = awss3.GetPresignedURL(filnameDNI, 15*time.Minute)
	if err != nil {
		logs.Logger.Println("Error al obtener URL pre-firmada para el documento de identidad:", err)
		return nil, err
	}

	persona.UrlCartaRecomendacion1, err = awss3.GetPresignedURL(filenameCarta1, 15*time.Minute)
	if err != nil {
		logs.Logger.Println("Error al obtener URL pre-firmada para la carta de recomendación 1:", err)
		return nil, err
	}

	persona.UrlCartaRecomendacion2, err = awss3.GetPresignedURL(filenameCarta2, 15*time.Minute)
	if err != nil {
		logs.Logger.Println("Error al obtener URL pre-firmada para la carta de recomendación 2:", err)
		return nil, err
	}

	// Asignar los valores de las cartas y documento de identidad como nil
	persona.DocumentoIdentidad = nil
	persona.CartaRecomendacion1 = nil
	persona.CartaRecomendacion2 = nil

	return &persona, nil
}

func (r *solicitudRepositoryDB) ObtenerEstadoSolicitudPorID(idSolicitud int) (string, error) {
	query := "CALL ObtenerEstadoSolicitud(?)"
	rows, err := database.DB.Query(query, idSolicitud)
	if err != nil {
		logs.Logger.Println("Error al ejecutar el procedure ObtenerEstadoSolicitud: ", err)
		return "", err
	}
	defer rows.Close()

	var estado string
	if rows.Next() {
		if err := rows.Scan(&estado); err != nil {
			logs.Logger.Println("Error al escanear el resultado del estado: ", err)
			return "", err
		}
		return estado, nil
	}

	return "", fmt.Errorf("no se encontró el estado para la solicitud con ID %d", idSolicitud)
}
