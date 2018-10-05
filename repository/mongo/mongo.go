package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/yoonsue/labchat/model/birthday"
	"github.com/yoonsue/labchat/model/menu"
	"github.com/yoonsue/labchat/model/phone"
)

// MenuRepository struct definition
type MenuRepository struct {
	db      string
	session *mgo.Session
}

// NewMenuRepository does several services according to MongoDB
func NewMenuRepository(session *mgo.Session) (menu.Repository, error) {
	r := &MenuRepository{
		db:      "mongo",
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

// Store saves menu model in MongoDB.
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

// PhoneRepository struct definition.
type PhoneRepository struct {
	db      string
	session *mgo.Session
}

// NewPhoneRepository return a new instance of MongoDB phone repository.
func NewPhoneRepository(session *mgo.Session) (phone.Repository, error) {
	r := &PhoneRepository{
		db:      "mongo",
		session: session,
	}

	index := mgo.Index{
		Key:        []string{"Department"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("phone")

	if err := c.EnsureIndex(index); err != nil {
		return nil, err
	}
	return r, nil
}

// Store saves phone model in MongoDB.
func (r *PhoneRepository) Store(target *phone.Phone) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("phone")

	_, err := c.Upsert(bson.M{"Department": target.Department}, bson.M{"$set": target})

	return err
}

// Find returns phone model that match with the given deparment.
func (r *PhoneRepository) Find(dept phone.Department) (*phone.Phone, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("phone")

	phone := phone.Phone{}
	if err := c.Find(bson.M{"Department": dept}).One(&phone); err != nil {
		if err == mgo.ErrNotFound {
			return nil, err
		}
		return nil, err
	}
	return &phone, nil
}

// BirthdayRepository struct definition.
type BirthdayRepository struct {
	db      string
	session *mgo.Session
}

// NewBirthdayRepository return a new instance of MongoDB phone repository.
func NewBirthdayRepository(session *mgo.Session) (birthday.Repository, error) {
	r := &BirthdayRepository{
		db:      "mongo",
		session: session,
	}

	index := mgo.Index{
		Key:        []string{"Name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("birthday")

	if err := c.EnsureIndex(index); err != nil {
		return nil, err
	}
	return r, nil
}

// Store saves birthday model in MongoDB.
func (r *BirthdayRepository) Store(target *birthday.Birthday) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("birthday")

	_, err := c.Upsert(bson.M{"Name": target.Name}, bson.M{"$set": target})

	return err
}

// Find returns birthday model that match with the given name.
func (r *BirthdayRepository) Find(name string) (*birthday.Birthday, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C("birthday")

	birthday := birthday.Birthday{}
	if err := c.Find(bson.M{"Name": name}).One(&birthday); err != nil {
		if err == mgo.ErrNotFound {
			return nil, err
		}
		return nil, err
	}
	return &birthday, nil
}
