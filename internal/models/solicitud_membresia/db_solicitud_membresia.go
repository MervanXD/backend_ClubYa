package solicitud

import (
	"fmt"
	"strings"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/models/persona"
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
		if err := rows.Scan(&familia.Id, &familia.Nombre, &familia.Apellidos, &familia.TipoDocumento, &familia.NroDocumento, &familia.TipoFamiliar, &familia.FechaNacimiento, &familia.DocumentoIdentidad); err != nil {
			logs.Logger.Println("Error al escanear al familiar: ", err)
			return nil, err
		}
		familiares = append(familiares, familia)
	}
	return familiares, nil
}

func (r *solicitudRepositoryDB) ObtenerDatosPersonaPorIdSolicitud(idSolicitud int) (*persona.Titular, error) {
	query := "CALL ObtenerDatosPersonaPorIdSolicitud(?)"
	rows := database.DB.QueryRow(query, idSolicitud)
	var persona persona.Titular
	err := rows.Scan(&persona.Nombre, &persona.Apellidos, &persona.Sexo, &persona.TipoDocumento, &persona.NroDocumento, &persona.FechaNacimiento, &persona.TipoVia, &persona.Direccion, &persona.Distrito, &persona.Pais, &persona.Provincia, &persona.Telefono, &persona.Referencia, &persona.Ocupacion, &persona.IngresoPromedio, &persona.NombreEmpresa, &persona.DireccionEmpresa,
		&persona.DocumentoIdentidad, &persona.CartaRecomendacion1, &persona.CartaRecomendacion2)
	if err != nil { //email lo estoy colocando en referencia , por ahora
		logs.Logger.Println("Error al ejecutar el procedimiento:", err)
		return nil, err
	}
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
