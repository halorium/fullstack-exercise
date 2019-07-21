package main

type Person struct {
	ID           int      `json:"id"       gorm:"primary_key"`
	Name         string   `json:"name"     gorm:"name"`
	IMGURL       string   `json:"imgUrl"   gorm:"img_url"`
	Location     string   `json:"location" gorm:"location"`
	PeopleColors []Colors `json:"-"        gorm:"foreignkey:PeopleID"`
	Colors       []string `json:"colors"   gorm:"-"`
}

func (Person) TableName() string {
	return "people"
}

type Colors struct {
	ID       int    `json:"id"       gorm:"primary_key"`
	PeopleID int    `json:"peopleId" gorm:"people_id"`
	Color    string `json:"color"    gorm:"color"`
}

func (Colors) TableName() string {
	return "people_colors"
}
