package persona

import (
	"fmt"
	"strings"

	"github.com/MervanXD/backend_ClubYa/database"
	"github.com/MervanXD/backend_ClubYa/logs"
	"github.com/go-sql-driver/mysql"
)

type FamiliarResquest struct {
	IdTitular  int        `json:"id_titular"`
	Familiares []Familiar `json:"familiares"`
}

type familiarRepositoryDB struct{}

func NewFamiliarRepositoryDB() FamiliarRepository {
	return &familiarRepositoryDB{}
}

func (r *familiarRepositoryDB) InsertarFamiliar(f Familiar, idTitular int) error {
	query := "call ingesoft.InsertarFamiliar(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err := database.DB.Exec(query, f.Nombre, f.Apellidos, f.Sexo.String(),
		f.TipoDocumento.String(), f.NroDocumento, f.FechaNacimiento, f.Telefono, f.Pais, f.Provincia, f.Distrito, f.TipoVia.String(),
		f.Direccion, f.Referencia, f.EsConyuge, idTitular, f.MismaDireccionPostulante, f.Ciudad, f.CodigoPostal, f.TipoFamiliar.String())
	if err != nil {
		logs.Logger.Println("Error al insertar Familiar: ", err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				// Error 1062: Duplicate entry
				if strings.Contains(mysqlErr.Message, "nroDocumento") {
					return fmt.Errorf("el nroDocumento ya está registrado")
				}

			}
		}
		return err
	}
	return nil
}

func (r *familiarRepositoryDB) RegistrarFamiliares(req FamiliarResquest) (int, error) {
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

    // Insertar cada familiar dentro de la transacción
    for _, familiar := range req.Familiares {
        query := "call ingesoft.InsertarFamiliar(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
        _, err = tx.Exec(query,
            familiar.Nombre,
            familiar.Apellidos,
            // ... resto de parámetros
        )
        if err != nil {
            logs.Logger.Println("Error al insertar familiar en transacción:", err)
            rollbackSafe()
            
            if mysqlErr, ok := err.(*mysql.MySQLError); ok {
                if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "nroDocumento") {
                    return -1, fmt.Errorf("el documento %s ya está registrado", familiar.NroDocumento)
                }
            }
            return -1, err
        }
    }

    // Insertar solicitud de membresía
    _, err = tx.Exec("call ingesoft.InsertarSolicitudMembresia(?, @s_id_solicitud)", req.IdTitular)
    if err != nil {
        logs.Logger.Println("Error al insertar solicitud membresía:", err)
        rollbackSafe()
        return -1, err
    }

    // Obtener id generado
    var idSolicitud int
    err = tx.QueryRow("SELECT @s_id_solicitud").Scan(&idSolicitud)
    if err != nil {
        logs.Logger.Println("Error al obtener idSolicitud:", err)
        rollbackSafe()
        return -1, err
    }

    // Commit final
    if err := tx.Commit(); err != nil {
        logs.Logger.Println("Error al hacer commit:", err)
        return -1, err
    }

    logs.Logger.Println("Familiares registrados exitosamente")
    return idSolicitud, nil
}

func (r *familiarRepositoryDB) ObtenerIdsFamiliaresPorTitular(idTitular int) ([]Familiar, error) {
	rows, err := database.DB.Query("CALL ObtenerFamiliaresPorTitular(?)", idTitular)
	if err != nil {
		logs.Logger.Println("Error al ejecutar procedimiento: ", err)
		return nil, err
	}
	defer rows.Close()

	var familiares []Familiar
	for rows.Next() {
		var fam Familiar
		if err := rows.Scan(&fam.Id,
			&fam.Nombre,
			&fam.Apellidos,
			&fam.Sexo,
			&fam.TipoDocumento,
			&fam.NroDocumento,
			&fam.FechaNacimiento,
			&fam.Telefono,
			&fam.Pais,
			&fam.Provincia,
			&fam.Distrito,
			&fam.Direccion,
			&fam.TipoVia,
			&fam.Referencia,
			&fam.Ciudad,
			&fam.CodigoPostal,
			&fam.MismaDireccionPostulante,
			&fam.EsConyuge,
			&fam.TipoFamiliar); err != nil {
			logs.Logger.Println("Error al escanear ID: ", err)
			return nil, err
		}
		familiares = append(familiares, fam)
	}

	return familiares, nil
}

func (r *familiarRepositoryDB) ObtenerFamiliarPorIDPersona(idPersona int) (*Familiar, error) {
	query := "call ingesoft.ObtenerFamiliarPorIdPersona(?)"
	rows := database.DB.QueryRow(query, idPersona)
	var f Familiar
	err := rows.Scan(&f.Id,
		&f.Nombre,
		&f.Apellidos,
		&f.Sexo,
		&f.TipoDocumento,
		&f.NroDocumento,
		&f.FechaNacimiento,
		&f.Telefono,
		&f.Pais,
		&f.Provincia,
		&f.Distrito,
		&f.Direccion,
		&f.TipoVia,
		&f.Referencia,
		&f.Ciudad,
		&f.CodigoPostal,
		&f.MismaDireccionPostulante,
		&f.EsConyuge,
		&f.TipoFamiliar)
	if err != nil {
		logs.Logger.Println("Error al obtener los datos de los familiares de la BD:", err)
		return nil, err
	}
	return &f, nil
}
