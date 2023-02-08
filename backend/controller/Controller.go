package controller

import (
	"net/http"

	"github.com/bxngearnx/sutlab/entity"
	"github.com/gin-gonic/gin"
)

// 3: GET Prefix เตรียมข้อมูลให้ combobox
func ListPrefix(c *gin.Context) {
	var Prefix []entity.PREFIX

	if err := entity.DB().Raw("SELECT * FROM prefixes").Scan(&Prefix).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ListPrefix_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Prefix})
}

// 4: GET Gender เตรียมข้อมูลให้ combobox
func ListGender(c *gin.Context) {
	var Gender []entity.GENDER

	if err := entity.DB().Raw("SELECT * FROM genders").Scan(&Gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ListGender_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Gender})
}

// 5: GET Province เตรียมข้อมูลให้ combobox
func ListProvince(c *gin.Context) {
	var Province []entity.PROVINCE

	if err := entity.DB().Raw("SELECT * FROM provinces").Scan(&Province).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ListProvince_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Province})
}
// 6: GET Degree เตรียมข้อมูลให้ combobox
func ListDegree(c *gin.Context) {
	var Degree []entity.DEGREE

	if err := entity.DB().Raw("SELECT * FROM degrees").Scan(&Degree).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ListDegree_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Degree})
}
// 7: GET Trimester เตรียมข้อมูลให้ combobox
func ListTrimester(c *gin.Context) {
	var Trimester []entity.TRIMESTER

	if err := entity.DB().Raw("SELECT * FROM trimesters").Scan(&Trimester).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ListTrimester_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Trimester})
}
// 8: GET Institute เตรียมข้อมูลให้ combobox
func ListInstitute(c *gin.Context) {
	var Institute []entity.INSTITUTE

	if err := entity.DB().Raw("SELECT * FROM institutes").Scan(&Institute).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ListInstitute_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Institute})
}
// 9: GET Branch เตรียมข้อมูลให้ combobox
func ListBranch(c *gin.Context) {
	var Branch []entity.BRANCH

	if err := entity.DB().Raw("SELECT * FROM branches").Scan(&Branch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ListBranch_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Branch})
}
///////////////////////////////// login //////////////////////////////////////////////////////////////////////

// GET /student/:id
func GetStudent(c *gin.Context) {
	var student entity.STUDENT
	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM students WHERE id = ?", id).Scan(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": student})
}
