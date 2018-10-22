package mongo

import (
	"github.com/yoonsue/labchat/model/birthday"
	"github.com/yoonsue/labchat/model/menu"
	"github.com/yoonsue/labchat/model/phone"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MenuRepository struct definition
type MenuRepository struct {
	db         string
	session    *mgo.Session
	collection string
}

// NewMenuRepository does several services according to MongoDB
func NewMenuRepository(session *mgo.Session, collection string) (menu.Repository, error) {
	r := &MenuRepository{
		db:         "mongo",
		session:    session,
		collection: collection,
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

	c := sess.DB(r.db).C(r.collection)

	if err := c.EnsureIndex(index); err != nil {
		return nil, err
	}
	return r, nil
}

// Store saves menu model in MongoDB.
func (r *MenuRepository) Store(target *menu.Menu) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C(r.collection)

	_, err := c.Upsert(bson.M{"Restaurant": target.Restaurant}, bson.M{"$set": target})

	return err
}

// Find returns today's menus that match with the given restaurant.
func (r *MenuRepository) Find(rest menu.Restaurant) (*menu.Menu, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C(r.collection)

	menu := menu.Menu{}
	if err := c.Find(bson.M{"Restaurant": rest}).One(&menu); err != nil {
		if err == mgo.ErrNotFound {
			return nil, err
		}
		return nil, err
	}
	return &menu, nil
}

// Clean the menu repository.
func (r *MenuRepository) Clean() error {
	r.session.DB(r.db).C(r.collection).RemoveAll(nil)
	return nil
}

// PhoneRepository struct definition.
type PhoneRepository struct {
	db         string
	session    *mgo.Session
	collection string
}

// NewPhoneRepository return a new instance of MongoDB phone repository.
func NewPhoneRepository(session *mgo.Session, collection string) (phone.Repository, error) {
	r := &PhoneRepository{
		db:         "mongo",
		session:    session,
		collection: collection,
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

	c := sess.DB(r.db).C(collection)

	if err := c.EnsureIndex(index); err != nil {
		return nil, err
	}
	return r, nil
}

// Store saves phone model in MongoDB.
func (r *PhoneRepository) Store(target *phone.Phone) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C(r.collection)

	_, err := c.Upsert(bson.M{"Department": target.Department}, bson.M{"$set": target})

	return err
}

// Find returns phone model that match with the given deparment.
func (r *PhoneRepository) Find(dept phone.Department) ([]*phone.Phone, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C(r.collection)
	var mongoPhoneList []*phone.Phone
	// var resultPhoneList []*phone.Phone

	if err := c.Find(bson.M{"Department": bson.RegEx{".*" + dept.ToString() + ".*", ""}}).All(&mongoPhoneList); err != nil {
		if err == mgo.ErrNotFound {
			return nil, err
		}
		return nil, err
	}
	// for _, phone := range mongoPhoneList {
	// 	resultPhoneList = append(resultPhoneList, phone)
	// }
	return mongoPhoneList, nil
}

// Clean the phone repository.
func (r *PhoneRepository) Clean() error {
	r.session.DB(r.db).C(r.collection).RemoveAll(nil)
	return nil
}

// BirthdayRepository struct definition.
type BirthdayRepository struct {
	db         string
	session    *mgo.Session
	collection string
}

// NewBirthdayRepository return a new instance of MongoDB phone repository.
func NewBirthdayRepository(session *mgo.Session, collection string) (birthday.Repository, error) {
	r := &BirthdayRepository{
		db:         "mongo",
		session:    session,
		collection: collection,
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

	c := sess.DB(r.db).C(r.collection)

	if err := c.EnsureIndex(index); err != nil {
		return nil, err
	}
	return r, nil
}

// Store saves birthday model in MongoDB.
func (r *BirthdayRepository) Store(target *birthday.Birthday) error {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C(r.collection)

	_, err := c.Upsert(bson.M{"Name": target.Name}, bson.M{"$set": target})

	return err
}

// Find returns birthday model that match with the given name.
func (r *BirthdayRepository) Find(name string) (*birthday.Birthday, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C(r.collection)

	birthday := birthday.Birthday{}
	if err := c.Find(bson.M{"Name": name}).One(&birthday); err != nil {
		if err == mgo.ErrNotFound {
			return nil, err
		}
		return nil, err
	}
	return &birthday, nil
}

// FindAll returns birthday model that match with the given name.
func (r *BirthdayRepository) FindAll() ([]*birthday.Birthday, error) {
	sess := r.session.Copy()
	defer sess.Close()

	c := sess.DB(r.db).C(r.collection)

	var result []*birthday.Birthday
	if err := c.Find(nil).All(&result); err != nil {
		if err == mgo.ErrNotFound {
			return nil, err
		}
		return nil, err
	}
	return result, nil
}

// Clean the birthday repository.
func (r *BirthdayRepository) Clean() error {
	/////////////// HOW CAN I CHECK COLLECTION REMOVED???
	r.session.DB(r.db).C(r.collection).RemoveAll(nil)
	return nil
}
