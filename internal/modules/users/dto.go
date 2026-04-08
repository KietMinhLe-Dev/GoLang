package users

import (
	"time"

	"github.com/google/uuid"
)

type UserRequest struct {
	Email    string    `json:"email" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Status   Status    `json:"status" binding:"required"`
	Ext      string    `json:"ext" binding:"required"`
	Role     Role      `json:"role" binding:"required"`
	TenantID uuid.UUID `json:"tenantId" binding:"required"`
}

type UserResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Status    Status     `json:"status"`
	Role      Role       `json:"role"`
	Ext       string     `json:"ext"`
	TenantID  string     `json:"tenantId"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
