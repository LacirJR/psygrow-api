package router

import (
	"github.com/LacirJR/psygrow-api/src/internal/handler"
	"github.com/LacirJR/psygrow-api/src/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("/login", handler.Login)
			}

			// Protected routes
			protected := v1.Group("")
			protected.Use(middleware.AuthMiddleware())
			{

				users := v1.Group("/users")
				{
					users.POST("", handler.RegisterUser)
					users.GET("/:email", handler.GetUserByEmail)
				}

				// Anamnese routes
				anamnese := protected.Group("/anamnese")
				{
					templates := anamnese.Group("/templates")
					{
						templates.POST("", handler.CreateAnamneseTemplate)
						templates.GET("", handler.GetAnamneseTemplates)
						templates.GET("/:template_id", handler.GetAnamneseTemplate)
						templates.PUT("/:template_id", handler.UpdateAnamneseTemplate)
						templates.DELETE("/:template_id", handler.DeleteAnamneseTemplate)

						// Anamnese field routes
						fields := templates.Group("/:template_id/fields")
						{
							fields.POST("", handler.CreateAnamneseField)
							fields.GET("", handler.GetAnamneseFields)
							fields.PUT("/:field_id", handler.UpdateAnamneseField)
							fields.DELETE("/:field_id", handler.DeleteAnamneseField)

							// Anamnese field option routes
							options := fields.Group("/:field_id/options")
							{
								options.POST("", handler.CreateAnamneseFieldOption)
								options.POST("/bulk", handler.CreateAnamneseFieldOptionsBulk)
								options.GET("", handler.GetAnamneseFieldOptions)
							}
						}
					}

					// Patient anamnese routes
					patientAnamnese := anamnese.Group("/patients")
					{
						patientAnamnese.POST("", handler.CreatePatientAnamnese)
						patientAnamnese.GET("/:patient_id", handler.GetPatientAnamneses)
						patientAnamnese.GET("/:patient_id/details", handler.GetPatientAnamneseDetails)
					}
				}

				// Patient routes
				patients := protected.Group("/patients")
				{
					patients.POST("", handler.CreatePatient)
					patients.GET("", handler.GetPatients)
					patients.GET("/:patient_id", handler.GetPatient)
					patients.PUT("/:patient_id", handler.UpdatePatient)
					patients.DELETE("/:patient_id", handler.DeletePatient)
					patients.GET("/search", handler.SearchPatientsByName)
					patients.GET("/cost-center/:cost_center_id", handler.GetPatientsByCostCenter)

					// Patient family routes
					families := patients.Group("/:patient_id/families")
					{
						families.POST("", handler.CreatePatientFamily)
						families.GET("", handler.GetPatientFamilies)
						families.GET("/:id", handler.GetPatientFamily)
						families.PUT("/:id", handler.UpdatePatientFamily)
						families.DELETE("/:id", handler.DeletePatientFamily)
						families.GET("/relationship/:relationship", handler.GetPatientFamiliesByRelationship)
					}
				}

				// Appointment routes
				appointments := protected.Group("/appointments")
				{
					appointments.POST("", handler.CreateAppointment)
					appointments.GET("", handler.GetAppointments)
					appointments.GET("/:id", handler.GetAppointment)
					appointments.PUT("/:id", handler.UpdateAppointment)
					appointments.DELETE("/:id", handler.DeleteAppointment)
				}

				// Session routes
				sessions := protected.Group("/sessions")
				{
					sessions.GET("", handler.GetSessions)
					sessions.GET("/:session_id", handler.GetSession)

					// Evolution routes
					evolutions := sessions.Group("/:session_id/evolutions")
					{
						evolutions.POST("", handler.CreateEvolution)
						evolutions.GET("/:id", handler.GetEvolution)
					}

					// Get evolutions by patient
					protected.GET("/patients/:patient_id/evolutions", handler.GetEvolutionsByPatient)
				}

				// Lead routes
				leads := protected.Group("/leads")
				{
					leads.POST("", handler.CreateLead)
					leads.GET("", handler.GetLeads)
					leads.GET("/:id", handler.GetLead)
					leads.PUT("/:id", handler.UpdateLead)
					leads.DELETE("/:id", handler.DeleteLead)
					leads.POST("/:id/convert", handler.ConvertLeadToPatient)
				}

				// Financial routes
				financial := protected.Group("/financial")
				{
					// Cost center routes
					costCenters := financial.Group("/cost-centers")
					{
						costCenters.POST("", handler.CreateCostCenter)
						costCenters.GET("", handler.GetCostCenters)
						costCenters.GET("/:id", handler.GetCostCenter)
						costCenters.PUT("/:id", handler.UpdateCostCenter)
						costCenters.DELETE("/:id", handler.DeleteCostCenter)
					}

					// Payment routes
					payments := financial.Group("/payments")
					{
						payments.POST("", handler.CreatePayment)
						payments.GET("", handler.GetPayments)
						payments.GET("/:id", handler.GetPayment)
					}

					// Repasse routes
					repasses := financial.Group("/repasses")
					{
						repasses.POST("", handler.CreateRepasse)
						repasses.GET("", handler.GetRepasses)
						repasses.PUT("/:id/status", handler.UpdateRepasseStatus)
					}
				}
			}
		}
	}
}
