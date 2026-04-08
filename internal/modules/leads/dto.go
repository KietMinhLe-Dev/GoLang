package leads

import (
	"time"

	"github.com/google/uuid"
)

type LeadsRequest struct {
	Name       string     `json:"name" binding:"required"`
	Phone      string     `json:"phone" binding:"required"`
	Address    string     `json:"address" binding:"required"`
	Source     string     `json:"source" binding:"required"`
	Stage      Stage      `json:"stage" binding:"required,oneof=NEW PROCESS CONTACTED BOOKED SHOWED CLOSED"`
	Notes      string     `json:"notes" binding:"required"`
	UserID     *uuid.UUID `json:"userId" binding:"required"`
	AssignedAt *time.Time `json:"assignedAt" binding:"required"`
	CampaignID *uuid.UUID `json:"campaignId" binding:"required"`
	TenantID   *uuid.UUID `json:"tenantId" binding:"required"`
	ShowAt     *time.Time `json:"showAt" binding:"required"`
	LastCallAt *time.Time `json:"lastCallAt" binding:"required"`
}

// H
type LeadsResponse struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Phone      string     `json:"phone"`
	Address    string     `json:"address"`
	Source     string     `json:"source"`
	Stage      Stage      `json:"stage,omitempty"`
	Notes      string     `json:"notes"`
	UserID     *uuid.UUID `json:"userId"`
	AssignedAt *time.Time `json:"assignedAt"`
	CampaignID *uuid.UUID `json:"campaignId"`
	TenantID   *uuid.UUID `json:"tenantId"`
	ShowAt     *time.Time `json:"showAt"`
	LastCallAt *time.Time `json:"lastCallAt"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}
