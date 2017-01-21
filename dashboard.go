package main

import (
	"github.com/kennygrant/sanitize"
	"github.com/night-codes/mgo-ai"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/night-codes/summer.v1"
	"gopkg.in/night-codes/types.v1"
	"time"
)

type (
	DashboardStruct struct {
		ID          uint64    `form:"id"  json:"id"  bson:"_id"`
		Name        string    `form:"name" json:"name" bson:"name" valid:"required"`
		Description string    `form:"description" json:"description" bson:"description"`
		Created     time.Time `form:"-" json:"created" bson:"created"`
		Updated     time.Time `form:"-" json:"updated" bson:"updated"`
		Deleted     bool      `form:"-" json:"deleted" bson:"deleted"`
	}
	DashboardModule struct {
		summer.Module
	}
)

var (
	dashboard = panel.AddModule(
		&summer.ModuleSettings{
			Name:   "dashboard",
			Title:  "Dashboard",
			Rights: summer.Rights{Groups: []string{"all"}},
		},
		&DashboardModule{},
	)
)

// Add new record
func (m *DashboardModule) Add(c *gin.Context) {
	var result DashboardStruct
	if !summer.PostBind(c, &result) {
		return
	}
	result.ID = ai.Next("dashboard")
	result.Created = time.Now()
	result.Updated = time.Now()
	result.Name = sanitize.HTML(result.Name)
	result.Description = sanitize.HTML(result.Description)

	if err := m.Collection.Insert(result); err != nil {
		c.String(400, "DB error")
		return
	}
	c.JSON(200, obj{"data": result})
}

// Edit record
func (m *DashboardModule) Edit(c *gin.Context) {
	id := types.Uint64(c.PostForm("id"))
	var result DashboardStruct
	var newValue DashboardStruct
	if !summer.PostBind(c, &newValue) {
		return
	}
	if err := m.Collection.FindId(id).One(&result); err == nil {
		result.Name = sanitize.HTML(newValue.Name)
		result.Description = sanitize.HTML(newValue.Description)
		result.Updated = time.Now()
		if err := m.Collection.UpdateId(newValue.ID, obj{"$set": result}); err != nil {
			c.String(400, "DB error")
			return
		}
	}
	c.JSON(200, obj{"data": result})
}

// Get record from DB
func (m *DashboardModule) Get(c *gin.Context) {
	id := types.Uint64(c.PostForm("id"))
	result := DashboardStruct{}
	if err := m.Collection.FindId(id).One(&result); err != nil {
		c.String(404, "Not found")
	}
	c.JSON(200, obj{"data": result})
}

// GetAll records
func (m *DashboardModule) GetAll(c *gin.Context) {
	results := []DashboardStruct{}
	request := obj{"deleted": false}

	// request to DB
	if err := m.Collection.Find(request).Sort("-_id").All(&results); err != nil {
		c.String(404, "Not found")
		return
	}

	c.JSON(200, obj{"data": results})
}

// Delete - remove record
func (m *DashboardModule) Delete(c *gin.Context) {
	id := types.Uint64(c.PostForm("id"))

	if err := m.Collection.UpdateId(id, obj{"$set": obj{"deleted": true}}); err != nil {
		c.String(404, "Not found")
		return
	}
	c.JSON(200, obj{"data": obj{"id": id}})
}
