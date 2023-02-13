// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package model

import (
	"errors"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type lang_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var LangBinding = lang_EntityInfo{
	Entity: objectbox.Entity{
		Id: 1,
	},
	Uid: 2005915493830902783,
}

// Lang_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Lang_ = struct {
	ID            *objectbox.PropertyUint64
	Code          *objectbox.PropertyString
	Name          *objectbox.PropertyString
	AncestorCodes *objectbox.PropertyStringVector
	Desc          *objectbox.PropertyString
}{
	ID: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &LangBinding.Entity,
		},
	},
	Code: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &LangBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &LangBinding.Entity,
		},
	},
	AncestorCodes: &objectbox.PropertyStringVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &LangBinding.Entity,
		},
	},
	Desc: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &LangBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (lang_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (lang_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Lang", 1, 2005915493830902783)
	model.Property("ID", 6, 1, 3286233471437341963)
	model.PropertyFlags(1)
	model.Property("Code", 9, 2, 7268864952445343482)
	model.PropertyFlags(2080)
	model.PropertyIndex(1, 3620402759221610712)
	model.Property("Name", 9, 3, 5762888112877023092)
	model.Property("AncestorCodes", 30, 4, 5054523526498907470)
	model.Property("Desc", 9, 5, 7067698906285608907)
	model.EntityLastPropertyId(5, 7067698906285608907)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (lang_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*Lang).ID, nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (lang_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Lang).ID = id
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (lang_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (lang_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*Lang)
	var offsetCode = fbutils.CreateStringOffset(fbb, obj.Code)
	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetAncestorCodes = fbutils.CreateStringVectorOffset(fbb, obj.AncestorCodes)
	var offsetDesc = fbutils.CreateStringOffset(fbb, obj.Desc)

	// build the FlatBuffers object
	fbb.StartObject(5)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetCode)
	fbutils.SetUOffsetTSlot(fbb, 2, offsetName)
	fbutils.SetUOffsetTSlot(fbb, 4, offsetDesc)
	fbutils.SetUOffsetTSlot(fbb, 3, offsetAncestorCodes)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (lang_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'Lang' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propID = table.GetUint64Slot(4, 0)

	return &Lang{
		ID:            propID,
		Code:          fbutils.GetStringSlot(table, 6),
		Name:          fbutils.GetStringSlot(table, 8),
		Desc:          fbutils.GetStringSlot(table, 12),
		AncestorCodes: fbutils.GetStringVectorSlot(table, 10),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (lang_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Lang, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (lang_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*Lang), nil)
	}
	return append(slice.([]*Lang), object.(*Lang))
}

// Box provides CRUD access to Lang objects
type LangBox struct {
	*objectbox.Box
}

// BoxForLang opens a box of Lang objects
func BoxForLang(ob *objectbox.ObjectBox) *LangBox {
	return &LangBox{
		Box: ob.InternalBox(1),
	}
}

// Put synchronously inserts/updates a single object.
// In case the ID is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Lang.ID property on the passed object will be assigned the new ID as well.
func (box *LangBox) Put(object *Lang) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the ID is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Lang.ID property on the passed object will be assigned the new ID as well.
func (box *LangBox) Insert(object *Lang) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *LangBox) Update(object *Lang) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *LangBox) PutAsync(object *Lang) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case IDs are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Lang.ID property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Lang.ID assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *LangBox) PutMany(objects []*Lang) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *LangBox) Get(id uint64) (*Lang, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Lang), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *LangBox) GetMany(ids ...uint64) ([]*Lang, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Lang), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *LangBox) GetManyExisting(ids ...uint64) ([]*Lang, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Lang), nil
}

// GetAll reads all stored objects
func (box *LangBox) GetAll() ([]*Lang, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Lang), nil
}

// Remove deletes a single object
func (box *LangBox) Remove(object *Lang) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *LangBox) RemoveMany(objects ...*Lang) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.ID
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Lang_ struct to create conditions.
// Keep the *LangQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *LangBox) Query(conditions ...objectbox.Condition) *LangQuery {
	return &LangQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Lang_ struct to create conditions.
// Keep the *LangQuery if you intend to execute the query multiple times.
func (box *LangBox) QueryOrError(conditions ...objectbox.Condition) (*LangQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &LangQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See LangAsyncBox for more information.
func (box *LangBox) Async() *LangAsyncBox {
	return &LangAsyncBox{AsyncBox: box.Box.Async()}
}

// LangAsyncBox provides asynchronous operations on Lang objects.
//
// Asynchronous operations are executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "execute & forget:" you gain faster put/remove operations as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
// In situations with (extremely) high async load, an async method may be throttled (~1ms) or delayed up to 1 second.
// In the unlikely event that the object could still not be enqueued (full queue), an error will be returned.
//
// Note that async methods do not give you hard durability guarantees like the synchronous Box provides.
// There is a small time window in which the data may not have been committed durably yet.
type LangAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForLang creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use LangBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForLang(ob *objectbox.ObjectBox, timeoutMs uint64) *LangAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 1, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 1: %s" + err.Error())
	}
	return &LangAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the ID property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *LangAsyncBox) Put(object *Lang) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The ID property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *LangAsyncBox) Insert(object *Lang) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *LangAsyncBox) Update(object *Lang) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *LangAsyncBox) Remove(object *Lang) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all Lang which ID is either 42 or 47:
//
//	box.Query(Lang_.ID.In(42, 47)).Find()
type LangQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *LangQuery) Find() ([]*Lang, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Lang), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *LangQuery) Offset(offset uint64) *LangQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *LangQuery) Limit(limit uint64) *LangQuery {
	query.Query.Limit(limit)
	return query
}
