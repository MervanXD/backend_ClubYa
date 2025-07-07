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
		return -1, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	query := "call ingesoft.InsertarTitular(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,@p_idTitular)"
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
	)
	if err != nil {
		logs.Logger.Println("Error al ejecutar InsertarTitular: ", err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				// Error 1062: Duplicate entry
				if strings.Contains(mysqlErr.Message, "nroDocumento") {
					return -1,fmt.Errorf("el nroDocumento ya está registrado")
				}

			}
		}
		tx.Rollback()
		return -1, err
	}

	var idTitular int
	err = tx.QueryRow("SELECT @p_idTitular").Scan(&idTitular)
	if err != nil {
		logs.Logger.Println("Error al obtener idTitular: ", err)
		tx.Rollback()
		return -1, err
	}

	if err := tx.Commit(); err != nil {
		return -1, err
	}

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
