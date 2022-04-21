package repository

import (
	"github.com/rahul.chaube/CurdeDemo/model"
	"gopkg.in/mgo.v2"
)

const (
	collectionMovie = "Movies"
)

type MovieRepository struct {
	session  *mgo.Session
	database string
}

func NewMovieRepository(session *mgo.Session, databaseName string) *MovieRepository {
	return &MovieRepository{
		session:  session,
		database: databaseName,
	}

}

func (repo *MovieRepository) AddMovie(m model.Movie) (err error) {
	session := repo.session.Copy()
	defer session.Close()
	err = session.DB(repo.database).C(collectionMovie).Insert(m)
	if err != nil {
		return err
	}
	return
}
