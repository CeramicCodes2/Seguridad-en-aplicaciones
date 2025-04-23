package ports

import "exlora/internal/domain/entities"

type ForDataRequest interface {
	GetDashboardData() (entities.Dashboard, error)
	// GetDashboardData retrieves the dashboard data.

}
