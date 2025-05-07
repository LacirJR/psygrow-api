package migration

import (
	"github.com/LacirJR/psygrow-api/src/internal/config"
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"log"
)

func Migrate() {
	db := config.DB

	err := db.AutoMigrate(
		&model.User{},
		&model.AnamneseTemplate{},
		&model.AnamneseField{},
		&model.PatientAnamnese{},
		&model.PatientAnamneseField{},
		&model.Appointment{},
		&model.Session{},
		&model.Evolution{},
		&model.CostCenter{},
		&model.Payment{},
		&model.PaymentAppointment{},
		&model.Repasse{},
		&model.Lead{},
		&model.Patient{},
		&model.PatientFamily{},
	)

	if err != nil {
		log.Fatalf("Erro ao rodar migrations: %v", err)
	}

	log.Println("Migrations aplicadas com sucesso.")

}
