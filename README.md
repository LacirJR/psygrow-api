# PsyGrow API - Model Improvements

## Docker Setup

### Prerequisites
- Docker and Docker Compose installed on your machine

### Running with Docker Compose
1. Clone the repository
2. Navigate to the project directory
3. Run the application using Docker Compose:
   ```bash
   docker-compose up -d
   ```
4. The API will be available at http://localhost:8080
5. The PostgreSQL database will be available at localhost:5432

### Environment Variables
The following environment variables are used by the application and are set in the docker-compose.yml file:
- `APP_PORT`: The port on which the API server runs
- `DB_HOST`: PostgreSQL database host
- `DB_USER`: PostgreSQL database user
- `DB_PASSWORD`: PostgreSQL database password
- `DB_NAME`: PostgreSQL database name
- `DB_PORT`: PostgreSQL database port
- `JWT_SECRET`: Secret key for JWT token generation/validation

## Changes Implemented

### 1. Enumerations for Fixed Values
Created constants for fixed values to avoid using strings directly:
- `RepasseModel`: `RepasseModelClinicPays`, `RepasseModelProfessionalPays`
- `RepasseType`: `RepasseTypePercent`, `RepasseTypeFixed`
- `AppointmentStatus`: `AppointmentStatusScheduled`, `AppointmentStatusDone`, etc.
- `PaymentMethod`: `PaymentMethodPix`, `PaymentMethodCash`, etc.
- `RepasseStatus`: `RepasseStatusPending`, `RepasseStatusPaid`, etc.

### 2. Improved Data Validation
Added validation using the go-playground/validator library:
- String length validations for fields like `Name`, `ServiceTitle`
- Validation for `RepasseValue` based on `RepasseType` (percentages between 0-100%)
- Required field validations
- Custom validation functions for complex validations

### 3. Database Indices
Added indices for frequently queried fields:
- `UserID`, `CostCenterID`, `PatientID`, `Status`, etc.

### 4. Explicit Entity Relationships
Improved entity relationships with explicit references:
```go
CostCenterID uuid.UUID  `gorm:"type:uuid;not null;index" validate:"required"`
CostCenter   CostCenter `gorm:"foreignKey:CostCenterID"`
```

### 5. Monetary Values
Replaced `float64` with `int64` for monetary values:
- Values are stored as cents (e.g., $10.50 = 1050)
- Percentages are stored as basis points (e.g., 10.50% = 1050)

### 6. Descriptive Boolean Field Names
Made boolean field names more descriptive:
- `Active` → `IsActive`
- `ClinicReceives` → `DoesClinicReceive`

### 7. Improved Documentation
Added detailed documentation comments to structs and fields.

### 8. Transaction Support Framework
Created a transaction helper framework to support transactions for related operations.

## Next Steps

To fully implement these changes, the following additional steps are needed:

### 1. Update DTOs
Update the DTOs to match the model changes:
- Change `float64` to `int64` for monetary values
- Update field names to match the model changes
- Add validation tags to DTOs

### 2. Update Handlers
Update the handlers to use the new model features:
- Use constants instead of string literals
- Validate models before saving
- Use the transaction helper for operations involving multiple tables

### 3. Update Repositories
Update the repositories to support transactions:
- Add transaction-aware methods to repository interfaces
- Implement these methods in repository implementations

### 4. Update Migrations
Update database migrations to reflect the model changes:
- Add indices for frequently queried fields
- Change column types for monetary values
- Rename boolean fields

### 5. Testing
Test the changes to ensure they work as expected:
- Unit tests for validation logic
- Integration tests for database operations
- End-to-end tests for API endpoints
