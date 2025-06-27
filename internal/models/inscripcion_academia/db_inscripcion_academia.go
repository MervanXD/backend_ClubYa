package inscripcionacademia

import (
	"fmt"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/internal/models/sesiones"
	servicios "github.com/MervanXD/backend_ClubYa/internal/services"
	"github.com/MervanXD/backend_ClubYa/logs"
)

type inscripcionAcademiaRepositoryDB struct{}

func NewInscripcionAcademiaRepositoryDB() InscripcionAcademiaRepository {
	return &inscripcionAcademiaRepositoryDB{}
}

func (r *inscripcionAcademiaRepositoryDB) RegistrarInscripcionAcademia(idPersonaint int, idGrupo int, idTarifa int, uniforme int, costo_total float64, idTitular int, metodoPago string) error {
	query := "CALL RegistrarInscripcionAcademia(?, ?, ?,?,?,?,?)"
	_, err := database.DB.Exec(query, idPersonaint, idGrupo, idTarifa, uniforme, costo_total, idTitular, metodoPago)
	if err != nil {
		logs.Logger.Println("Error al registrar la inscripción al evento: ", err)
		return err
	}
	return nil
}

func (r *inscripcionAcademiaRepositoryDB) ObtenerFamiliaresInscritosAcademia(idSocio int) ([]InscritoAcademiaDTO, error) {
	inscritosQuery := "CALL ListarInscripcionesFamiliares(?)"
	inscritosRows, err := database.DB.Query(inscritosQuery, idSocio)
	if err != nil {
		logs.Logger.Println("Error al obtener las personas inscritas:", err)
		return nil, err
	}
	defer inscritosRows.Close()
	var inscritos []InscritoAcademiaDTO
	for inscritosRows.Next() {
		var inscrito InscritoAcademiaDTO
		if err := inscritosRows.Scan(&inscrito.IdInscripcion, &inscrito.IdPersona, &inscrito.NombrePersona, &inscrito.ApellidoPersona,
			&inscrito.TipoSocio, &inscrito.NombreGrupo, &inscrito.IdGrupo, &inscrito.EdadMinima, &inscrito.EdadMaxima, &inscrito.NombreAcademia, &inscrito.FechaInicio, &inscrito.FechaFin, &inscrito.EstadoInscripcion, &inscrito.FechaInscripcion, &inscrito.Monto); err != nil {
			logs.Logger.Println("Error al escanear a la persona inscrita:", err)
			return nil, err
		}
		//leemos sus sesiones
		repo := sesiones.NewSesionRepositoryDB()
		sesiones, err := repo.ObtenerSesionesGrupo(inscrito.IdGrupo)
		if err != nil {
			logs.Logger.Println("Error al obtener las sesiones del grupo :", err)
			return nil, err
		}
		inscrito.Sesiones = sesiones
		inscritos = append(inscritos, inscrito)
	}
	return inscritos, nil
}

func (r *inscripcionAcademiaRepositoryDB) ListarInscritosPorIdAcademia(idAcademia int) ([]InscritosAcademiaRequest, error) {
	query := "CALL ListarInscritosPorIdAcademia(?)"
	rows, err := database.DB.Query(query, idAcademia)
	if err != nil {
		logs.Logger.Println("Error al obtener los inscritos de la academia:", err)
		return nil, err
	}
	defer rows.Close()

	var inscritos []InscritosAcademiaRequest
	for rows.Next() {
		var inscrito InscritosAcademiaRequest
		if err := rows.Scan(
			&inscrito.IdPersona,
			&inscrito.NombrePersona,
			&inscrito.ApellidoPersona,
			&inscrito.TipoSocio,
			&inscrito.NombreGrupo,
			&inscrito.IdGrupo,
			&inscrito.EdadPersona,
			&inscrito.FechaInscripcion,
			&inscrito.Monto,
			&inscrito.EstadoInscripcion,
			&inscrito.AnulaccionInscripcion.IdInscripcion,
			&inscrito.AnulaccionInscripcion.IdAnulacion,
			&inscrito.AnulaccionInscripcion.Fecha,
			&inscrito.AnulaccionInscripcion.Motivo,
			&inscrito.AnulaccionInscripcion.Devolucion,
		); err != nil {
			logs.Logger.Println("Error al escanear el inscrito:", err)
			return nil, err
		}
		inscritos = append(inscritos, inscrito)
	}

	return inscritos, nil
}

func (r *inscripcionAcademiaRepositoryDB) AnularInscripcionAcademia(idInscripcion int, idPersona int, motivo string) error {
	query := "call ingesoft.AnularInscripcionAcademia(?, ?, ?)"
	_, err := database.DB.Exec(query, idInscripcion, idPersona, motivo)
	if err != nil {
		logs.Logger.Println("Error al anular la inscripcion a la academia ", err)
		return err
	}
	return nil
}

func (r *inscripcionAcademiaRepositoryDB) AceptarAnulacionInscripcionAcademia(idAnulacion int) error {
	query := "call ingesoft.AceptarAnulacionInscripcionAcademia(?,  @p_correo, @p_monto)"
	_, err := database.DB.Exec(query, idAnulacion)
	if err != nil {
		logs.Logger.Println("Error al aceptar la devolución de la anulación de academia: ", err)
		return err
	}
	// Recupera el valor del parámetro de salida
	var correo string
	var monto float64
	row := database.DB.QueryRow("SELECT @p_correo, @p_monto")
	if err := row.Scan(&correo, &monto); err != nil {
		logs.Logger.Println("Error al obtener el correo de salida: ", err)
		return err
	}
	//Mandamos un correo de confirmacion al socio

	if correo != "" {
		err = servicios.EnviarCorreo([]string{correo},
			"Confirmación de devolución de anulación de reserva",
			fmt.Sprintf("Su solicitud de devolución ha sido aceptada. El monto a devolver es: %.2f.", monto),
		)
		if err != nil {
			logs.Logger.Println("Error al enviar correo de confirmación de devolución: ", err)
			return err
		}
	}

	return nil
}
