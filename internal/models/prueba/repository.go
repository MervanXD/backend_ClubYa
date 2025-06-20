package prueba

import (
	"github.com/MervanXD/backend_ClubYa/database"
)

type PruebaRepository interface {
	ListarPruebas() ([]Prueba, error)
	InsertarPrueba(req Prueba) (int, error)
	ModificarPrueba(req Prueba) error
	EliminarPrueba(id int) error
}

type pruebaRepositoryDB struct{}

func NewPruebaRepositoryDB() PruebaRepository {
	return &pruebaRepositoryDB{}
}

func (r *pruebaRepositoryDB) ListarPruebas() ([]Prueba, error) {
	query := "SELECT id_prueba, titulo, descripcion, url_imagen FROM ingesoft.Prueba"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pruebas []Prueba
	for rows.Next() {
		var p Prueba
		if err := rows.Scan(&p.IdPrueba, &p.Titulo, &p.Descripcion, &p.UrlImagen); err != nil {
			return nil, err
		}
		pruebas = append(pruebas, p)
	}
	return pruebas, nil
}

func (r *pruebaRepositoryDB) InsertarPrueba(req Prueba) (int, error) {
	query := "INSERT INTO ingesoft.Prueba (titulo, descripcion, url_imagen) VALUES (?, ?, ?)"
	res, err := database.DB.Exec(query, req.Titulo, req.Descripcion, req.UrlImagen)
	if err != nil {
		return -1, err
	}
	id, _ := res.LastInsertId()
	return int(id), nil
}

func (r *pruebaRepositoryDB) ModificarPrueba(req Prueba) error {
	query := "UPDATE ingesoft.Prueba SET titulo = ?, descripcion = ?, url_imagen = ? WHERE id_prueba = ?"
	_, err := database.DB.Exec(query, req.Titulo, req.Descripcion, req.UrlImagen, req.IdPrueba)
	return err
}

func (r *pruebaRepositoryDB) EliminarPrueba(id int) error {
	query := "DELETE FROM ingesoft.Prueba WHERE id_prueba = ?"
	_, err := database.DB.Exec(query, id)
	return err
}
