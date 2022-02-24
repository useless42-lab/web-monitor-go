package models

import (
	"WebMonitor/tools"
	"time"
)

type DomainWhoisStruct struct {
	DefaultModel
	WebId                int64     `json:"web_id" gorm:"column:web_id"`
	DomainCreatedDate    time.Time `json:"domain_created_date" gorm:"column:domain_created_date"`
	DomainExpirationDate time.Time `json:"domain_expiration_date" gorm:"column:domain_expiration_date"`
	RegistrarName        string    `json:"registrar_name" gorm:"column:registrar_name"`
	RegistrantName       string    `json:"registrant_name" gorm:"column:registrant_name"`
	RegistrantEmail      string    `json:"registrant_email" gorm:"column:registrant_email"`
}

func AddDomainWhois(webId int64, domainCreatedDate time.Time, domainExpirationDate time.Time, registrarName string, registrantName string, registrantEmail string) {
	data := DomainWhoisStruct{
		DefaultModel:         DefaultModel{ID: tools.GenerateSnowflakeId()},
		WebId:                webId,
		DomainCreatedDate:    domainCreatedDate,
		DomainExpirationDate: domainExpirationDate,
		RegistrarName:        registrarName,
		RegistrantName:       registrantName,
		RegistrantEmail:      registrantEmail,
	}
	DB.Table("domain_whois").Create(&data)
}

type RDomainWhoisStruct struct {
	DomainCreatedDate    LocalTime `json:"domain_created_date" gorm:"column:domain_created_date"`
	DomainExpirationDate LocalTime `json:"domain_expiration_date" gorm:"column:domain_expiration_date"`
	RegistrarName        string    `json:"registrar_name" gorm:"column:registrar_name"`
	RegistrantName       string    `json:"registrant_name" gorm:"column:registrant_name"`
	UpdatedAt            LocalTime `json:"updated_at" gorm:"column:updated_at"`
}

func GetDomainWhois(webId int64) RDomainWhoisStruct {
	var result RDomainWhoisStruct
	sqlStr := `select * from domain_whois where web_id=@webId and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{
		"webId": webId,
	}).Scan(&result)
	return result
}

func CheckWhois(webId int64, domainCreatedDate time.Time, domainExpirationDate time.Time, registrarName string, registrantName string, registrantEmail string) {
	data := DomainWhoisStruct{
		DefaultModel:         DefaultModel{ID: tools.GenerateSnowflakeId()},
		WebId:                webId,
		DomainCreatedDate:    domainCreatedDate,
		DomainExpirationDate: domainExpirationDate,
		RegistrarName:        registrarName,
		RegistrantName:       registrantName,
		RegistrantEmail:      registrantEmail,
	}
	updateData := DomainWhoisStruct{
		WebId:                webId,
		DomainCreatedDate:    domainCreatedDate,
		DomainExpirationDate: domainExpirationDate,
		RegistrarName:        registrarName,
		RegistrantName:       registrantName,
		RegistrantEmail:      registrantEmail,
	}

	err := DB.Table("domain_whois").Where("web_id = ?", &webId).First(&updateData).Error
	if err != nil {
		DB.Table("domain_whois").Create(&data)
	} else {
		DB.Table("domain_whois").Where("web_id=?", webId).Updates(&updateData)
	}
}
