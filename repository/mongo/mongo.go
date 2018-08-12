package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/yoonsue/labchat/model/menu"
)

// MenuRepository struct definition
type MenuRepository struct {
	db      string
	session *mgo.Session
}

// NewMenuRepository does several services according to MongoDB
func NewMenuRepository(db string, session *mgo.Session) (menu.Repository, error) {
	r := &MenuRepository{
		db:      db,
		session: session,
	}

	index := mgo.Index{
		Key:        []string{"Restaurant"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("menu")

	if err := c.EnsureIndex(index); err != nil {
		return nil, err
	}
	return r, nil
}

// Store saves menu model in memory.
func (r *MenuRepository) Store(target *menu.Menu) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("menu")

	_, err := c.Upsert(bson.M{"Restaurant": target.Restaurant}, bson.M{"$set": target})

	return err
}

// Find returns today's menus that match with the given restaurant.
func (r *MenuRepository) Find(rest menu.Restaurant) (*menu.Menu, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("menu")

	menu := menu.Menu{}
	if err := c.Find(bson.M{"Restaurant": rest}).One(&menu); err != nil {
		if err == mgo.ErrNotFound {
			return nil, err
		}
		return nil, err
	}
	return &menu, nil
}
