package model

// RepasseModel defines the repasse model constants
const (
	RepasseModelClinicPays       = "clinic_pays"
	RepasseModelProfessionalPays = "professional_pays"
)

// RepasseType defines the repasse type constants
const (
	RepasseTypePercent = "percent"
	RepasseTypeFixed   = "fixed"
)

// AppointmentStatus defines the appointment status constants
const (
	AppointmentStatusScheduled = "scheduled"
	AppointmentStatusDone      = "done"
	AppointmentStatusCanceled  = "canceled"
	AppointmentStatusNoShow    = "no_show"
)

// PaymentMethod defines the payment method constants
const (
	PaymentMethodPix   = "pix"
	PaymentMethodCash  = "cash"
	PaymentMethodCard  = "card"
	PaymentMethodOther = "other"
)

// RepasseStatus defines the repasse status constants
const (
	RepasseStatusPending       = "pending"
	RepasseStatusPaid          = "paid"
	RepasseStatusInformational = "informational"
)

// LeadStatus defines the lead status constants
const (
	LeadStatusNew        = "new"
	LeadStatusInAnalysis = "in_analysis"
	LeadStatusConverted  = "converted"
	LeadStatusLost       = "lost"
)

// PatientFamilyRelationship defines the relationship constants
const (
	RelationshipFather      = "father"
	RelationshipMother      = "mother"
	RelationshipSpouse      = "spouse"
	RelationshipChild       = "child"
	RelationshipResponsible = "guardian" // Ou "responsible_party", mas "guardian" é comum para responsável legal.
	RelationshipGrandparent = "grandparent"
	RelationshipSibling     = "sibling"
	RelationshipOther       = "other"
)
