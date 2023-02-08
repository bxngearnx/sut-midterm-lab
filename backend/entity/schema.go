package entity

import (
	"time"

	"gorm.io/gorm"
)

type PREFIX struct {
	gorm.Model
	Prefix_Name string

	Postponement []POSTPONEMENT `gorm:"foreignKey:PrefixID"`
	Admin        []ADMIN        `gorm:"foreignKey:PrefixID"`
}
type GENDER struct {
	gorm.Model
	Gender_Type string
	Admin       []ADMIN `gorm:"foreignKey:GenderID"`
}
type PROVINCE struct {
	gorm.Model
	Province_Name string
	Admin         []ADMIN `gorm:"foreignKey:ProvinceID"`
}
type TRIMESTER struct {
	gorm.Model
	Trimester_Name string

	Postponement []POSTPONEMENT `gorm:"foreignKey:TrimesterID"`
}
type DEGREE struct {
	gorm.Model
	Degree_Name string

	Course       []COURSE       `gorm:"foreignKey:DegreeID"`
	Student      []STUDENT      `gorm:"foreignKey:DegreeID"`
	Postponement []POSTPONEMENT `gorm:"foreignKey:DegreeID"`
}
type INSTITUTE struct {
	gorm.Model
	Institute_Name string

	Branch       []BRANCH       `gorm:"foreignKey:InstituteID"`
	Course       []COURSE       `gorm:"foreignKey:InstituteID"`
	Postponement []POSTPONEMENT `gorm:"foreignKey:InstituteID"`
	Student      []STUDENT      `gorm:"foreignKey:InstituteID"`
}
type BRANCH struct {
	gorm.Model
	Branch_Name    string
	Branch_Teacher string
	Branch_Info    string

	PrefixID    *uint
	InstituteID *uint
	AdminID     *uint

	Prefix    PREFIX
	Institute INSTITUTE
	Admin     ADMIN

	Course       []COURSE       `gorm:"foreignKey:BranchID"`
	Postponement []POSTPONEMENT `gorm:"foreignKey:BranchID"`
	Student      []STUDENT      `gorm:"foreignKey:BranchID"`
}
type COURSE struct {
	gorm.Model
	Course_Name    string
	Course_Teacher string
	Course_Credit  uint
	Course_Detail  string
	Course_Year    uint

	DegreeID    *uint
	PrefixID    *uint
	InstituteID *uint
	BranchID    *uint
	AdminID     *uint

	Degree    DEGREE
	Prefix    PREFIX
	Institute INSTITUTE
	Branch    BRANCH
	Admin     ADMIN

	Student []STUDENT `gorm:"foreignKey:CourseID"`
}
type STUDENT struct {
	gorm.Model
	Student_Year_Of_Entry time.Time
	Student_Number        string `gorm:"uniqueIndex"`
	Student_Name          string
	Student_Birthday      time.Time
	Student_Tel           string
	Student_Identity_Card string `gorm:"uniqueIndex"`
	Student_Nationality   string
	Student_Religion      string
	Student_Address       string
	Student_Fathers_Name  string
	Student_Mothers_Name  string

	GenderID    *uint
	DegreeID    *uint
	PrefixID    *uint
	InstituteID *uint
	ProvinceID  *uint
	BranchID    *uint
	CourseID    *uint
	AdminID     *uint

	Gender    GENDER
	Degree    DEGREE
	Prefix    PREFIX
	Institute INSTITUTE
	Province  PROVINCE
	Branch    BRANCH
	Course    COURSE
	Admin     ADMIN

	Postponement []POSTPONEMENT `gorm:"foreignKey:StudentID"`
}
type ADMIN struct {
	gorm.Model
	Admin_Name     string
	Admin_Email    string `gorm:"uniqueIndex"`
	Admin_Password string `valid:"minstringlength(8)"`
	Admin_Tel      string `valid:"matches(^\\d{10}$)"`
	Admin_Address  string `gorm:"uniqueIndex" valid:"minstringlength(20)"`

	PrefixID   *uint
	GenderID   *uint
	ProvinceID *uint

	Prefix   PREFIX
	Gender   GENDER
	Province PROVINCE

	Student []STUDENT `gorm:"foreignKey:AdminID"`
	Branch  []BRANCH  `gorm:"foreignKey:AdminID"`
	Course  []COURSE  `gorm:"foreignKey:AdminID"`
}

type POSTPONEMENT struct {
	gorm.Model
	Postponement_Student_Number string `gorm:"uniqueIndex" valid:"required,matches(^[BMD]\\d{7}$)"`
	Postponement_Student_Name   string `valid:"required~name cannot be blank"`
	Postponement_AcademicYear   string
	Postponement_Gpax           string
	Postponement_Credit         string
	Postponement_Date           time.Time
	Postponement_Reasons        string `valid:"required, minstringlength(10)~กรุณากรอกเหตุผลอย่างน้อย10ตัวอักษร"`
	PrefixID                    *uint
	DegreeID                    *uint
	TrimesterID                 *uint
	InstituteID                 *uint
	BranchID                    *uint
	StudentID                   *uint

	Prefix    PREFIX
	Degree    DEGREE
	Trimester TRIMESTER
	Institute INSTITUTE
	Branch    BRANCH
	Student   STUDENT
}
