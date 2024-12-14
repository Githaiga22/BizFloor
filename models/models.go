package models

import (
    "time"
    "gorm.io/gorm"
)

// User represents either a client or business owner
type User struct {
    ID             uint           `gorm:"primaryKey" json:"id"`
    Name           string         `gorm:"type:varchar(255);not null" json:"name"`
    Phone          string         `gorm:"type:varchar(20);uniqueIndex;not null" json:"phone"`
    Email          string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
    PasswordHash   string         `gorm:"type:varchar(255);not null" json:"-"`  // "-" means don't show in JSON
    IsBusinessOwner bool          `gorm:"default:false" json:"is_business_owner"`
    CreatedAt      time.Time      `json:"created_at"`
    UpdatedAt      time.Time      `json:"updated_at"`
    DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
    
    // Relationships
    Businesses     []Business    `gorm:"foreignKey:OwnerID" json:"businesses,omitempty"`
    Bookings       []Booking     `gorm:"foreignKey:ClientID" json:"bookings,omitempty"`
    SavedPlaces    []SavedPlace  `gorm:"foreignKey:UserID" json:"saved_places,omitempty"`
}

// Business represents a service provider
type Business struct {
    ID           uint           `gorm:"primaryKey" json:"id"`
    OwnerID      uint           `gorm:"not null" json:"owner_id"`
    Name         string         `gorm:"not null" json:"name"`
    Type         string         `gorm:"not null" json:"type"`
    MpesaNumber  string         `gorm:"not null" json:"mpesa_number"`
    CreatedAt    time.Time      `json:"created_at"`
    UpdatedAt    time.Time      `json:"updated_at"`
    DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
    
    // Relationships
    Owner        User          `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
    Services     []Service     `json:"services,omitempty"`
    SavedByUsers []SavedPlace  `json:"saved_by_users,omitempty"`
    LoyalClients []LoyalClient `json:"loyal_clients,omitempty"`
}

// Service represents a specific service offered by a business
type Service struct {
    ID            uint           `gorm:"primaryKey" json:"id"`
    BusinessID    uint           `gorm:"not null" json:"business_id"`
    Name          string         `gorm:"not null" json:"name"`
    Price         float64        `gorm:"type:decimal(10,2);not null" json:"price"`
    DurationMins  int           `gorm:"not null" json:"duration_mins"`
    CreatedAt     time.Time      `json:"created_at"`
    UpdatedAt     time.Time      `json:"updated_at"`
    DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
    
    // Relationships
    Business      Business       `json:"business,omitempty"`
    Bookings      []Booking      `json:"bookings,omitempty"`
}

// Booking represents a service appointment
type Booking struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    ServiceID   uint           `gorm:"not null" json:"service_id"`
    ClientID    uint           `gorm:"not null" json:"client_id"`
    BusinessID  uint           `gorm:"not null" json:"business_id"`
    BookingTime time.Time      `gorm:"not null" json:"booking_time"`
    Status      string         `gorm:"not null;default:'pending'" json:"status"` // pending, confirmed, completed, cancelled
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
    
    // Relationships
    Service     Service        `json:"service,omitempty"`
    Client      User          `json:"client,omitempty"`
    Business    Business       `json:"business,omitempty"`
	PaymentID   *uint          `json:"payment_id"`
    Payment     *Payment        `json:"payment,omitempty"`
}

// Payment represents a payment transaction
type Payment struct {
    ID             uint           `gorm:"primaryKey" json:"id"`
    BookingID      uint           `gorm:"uniqueIndex;not null" json:"booking_id"`
    Amount         float64        `gorm:"type:decimal(10,2);not null" json:"amount"`
    Status         string         `gorm:"not null;default:'pending'" json:"status"` // pending, held, released, refunded
    PaymentTime    time.Time      `json:"payment_time"`
    MpesaReference string         `gorm:"type:varchar(255);uniqueIndex" json:"mpesa_reference"`
    CreatedAt      time.Time      `json:"created_at"`
    UpdatedAt      time.Time      `json:"updated_at"`
    DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
    
    // Relationships
    Booking        *Booking        `json:"booking,omitempty"`
}

// SavedPlace represents a user's saved business
type SavedPlace struct {
    ID         uint           `gorm:"primaryKey" json:"id"`
    UserID     uint           `gorm:"not null" json:"user_id"`
    BusinessID uint           `gorm:"not null" json:"business_id"`
    SavedAt    time.Time      `gorm:"not null" json:"saved_at"`
    CreatedAt  time.Time      `json:"created_at"`
    UpdatedAt  time.Time      `json:"updated_at"`
    DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
    
    // Relationships
    User       User           `json:"user,omitempty"`
    Business   Business       `json:"business,omitempty"`
}

// LoyalClient represents a business's tracked loyal customers
type LoyalClient struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    BusinessID  uint           `gorm:"not null" json:"business_id"`
    ClientID    uint           `gorm:"not null" json:"client_id"`
    VisitCount  int           `gorm:"default:0" json:"visit_count"`
    FirstVisit  time.Time      `json:"first_visit"`
    LastVisit   time.Time      `json:"last_visit"`
    LoyaltyTier string         `gorm:"default:'regular'" json:"loyalty_tier"` // regular, silver, gold
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
    
    // Relationships
    Business    Business       `json:"business,omitempty"`
    Client      User          `json:"client,omitempty"`
}