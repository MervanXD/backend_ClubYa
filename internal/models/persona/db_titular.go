package persona

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/go-sql-driver/mysql"
)

type titularRepositoryDB struct{}

func NewTitularRepositoryDB() TitularRepository {
	return &titularRepositoryDB{}
}

func (r *titularRepositoryDB) InsertarTitular(p Titular, idCuenta int) (int, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		logs.Logger.Println("Error al iniciar transacción:", err)
		return -1, err
	}

	// Función para rollback seguro
	rollbackSafe := func() {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			logs.Logger.Printf("Error en rollback: %v", rollbackErr)
		}
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Logger.Printf("PANIC recuperado en InsertarTitular: %v", r)
			rollbackSafe()
			// NO re-panic para mantener el servidor estable
		}
	}()

	query := "call ingesoft.InsertarTitular(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,@p_idTitular)"
	_, err = tx.Exec(query,
		p.Nombre,
		p.Apellidos,
		p.Sexo.String(),
		p.TipoDocumento.String(),
		p.NroDocumento,
		p.FechaNacimiento,
		p.Telefono,
		p.Pais,
		p.Provincia,
		p.Distrito,
		p.TipoVia.String(),
		p.Direccion,
		p.Referencia,
		p.Ocupacion,
		p.NombreEmpresa,
		p.DireccionEmpresa,
		p.IngresoPromedio,
		p.EsPostulante,
		p.Ciudad,
		p.CodigoPostal,
		idCuenta,
		p.DocumentoIdentidad,
		p.CartaRecomendacion1,
		p.CartaRecomendacion2,
	)
	if err != nil {
		logs.Logger.Println("Error al ejecutar InsertarTitular:", err)
		rollbackSafe()

		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "nroDocumento") {
				return -1, fmt.Errorf("el número de documento ya está registrado")
			}
		}
		return -1, err
	}

	var idTitular int
	err = tx.QueryRow("SELECT @p_idTitular").Scan(&idTitular)
	if err != nil {
		logs.Logger.Println("Error al obtener idTitular:", err)
		rollbackSafe()
		return -1, err
	}

	// Commit final
	if err := tx.Commit(); err != nil {
		logs.Logger.Println("Error al hacer commit:", err)
		return -1, err
	}

	logs.Logger.Printf("Titular insertado exitosamente con ID: %d", idTitular)
	return idTitular, nil
}

// ObtenerTitularPorID obtiene la información de un t por su ID
// Returns:
//   - *Titular: Puntero a la estructura Titular con los datos
//   - error: Error si ocurre alguno
func (r *titularRepositoryDB) ObtenerTitularPorID(ctx context.Context, idPersona int) (*Titular, error) {
	query := "call ingesoft.ObtenerTitularPorID(?)"

	if idPersona <= 0 {
		return nil, errors.New("ID de persona inválido")

	}

	var t Titular
	row := database.DB.QueryRowContext(ctx, query, idPersona)
	err := row.Scan(&t.Id,
		&t.Nombre,
		&t.Apellidos,
		&t.Sexo,
		&t.TipoDocumento,
		&t.NroDocumento,
		&t.FechaNacimiento,
		&t.Telefono,
		&t.Pais,
		&t.Provincia,
		&t.Distrito,
		&t.Direccion,
		&t.TipoVia,
		&t.Referencia,
		&t.Ciudad,
		&t.CodigoPostal,
		&t.IngresoPromedio,
		&t.Ocupacion,
		&t.NombreEmpresa,
		&t.DireccionEmpresa)

	if err != nil {
		logs.Logger.Println("Error al escanear ID: ", err)
		return nil, err
	}

	return &t, nil
}

func (r *titularRepositoryDB) BuscarTitulares(nroDocumento string, nombre string) ([]TitularResumen, error) {
	sesionesQuery := "CALL BuscarTitulares(?,?)"
	rows, err := database.DB.Query(sesionesQuery, nroDocumento, nombre)
	if err != nil {
		logs.Logger.Println("Error al obtener los titulares:", err)
		return nil, err
	}
	defer rows.Close()
	var titulares []TitularResumen
	for rows.Next() {
		var titular TitularResumen
		err := rows.Scan(&titular.IdPersona, &titular.NumeroDocumento, &titular.NombreCompleto)
		if err != nil {
			logs.Logger.Println("Error al escanear el titular:", err)
			return nil, err
		}
		titulares = append(titulares, titular)
	}
	return titulares, nil
}
