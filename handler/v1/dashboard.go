package v1

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mazrean/todoList/domain/values"
	"github.com/mazrean/todoList/service"
)

type Dashboard struct {
	context          *Context
	session          *Session
	dashboardService service.Dashboard
}

func NewDashboard(
	context *Context,
	session *Session,
	dashboardService service.Dashboard,
) *Dashboard {
	return &Dashboard{
		context:          context,
		session:          session,
		dashboardService: dashboardService,
	}
}

type DashboardInfo struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"createdAt"`
}

type DashboardDetail struct {
	ID             uuid.UUID          `json:"id"`
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	CreatedAt      time.Time          `json:"createdAt"`
	TaskStatusList []TaskStatusDetail `json:"taskStatusList"`
}

func (d *Dashboard) PostDashboard(c *gin.Context) {
	var info DashboardInfo
	err := c.BindJSON(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	session := d.session.getSession(c)
	user, err := d.session.getUser(session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get user",
		})
		return
	}

	dashboard, err := d.dashboardService.CreateDashboard(
		c.Request.Context(),
		user,
		values.NewDashboardName(info.Name),
		values.NewDashboardDescription(info.Description),
	)
	if err != nil {
		log.Printf("failed to create dashboard: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create dashboard",
		})
		return
	}

	c.JSON(http.StatusCreated, DashboardInfo{
		ID:          uuid.UUID(dashboard.GetID()),
		Name:        string(dashboard.GetName()),
		Description: string(dashboard.GetDescription()),
		CreatedAt:   time.Time(dashboard.GetCreatedAt()),
	})
}

func (d *Dashboard) PatchDashboard(c *gin.Context) {
	strDashboardID := c.Param("dashboardID")
	dashboardID, err := uuid.Parse(strDashboardID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid dashboard id",
		})
		return
	}

	var info DashboardInfo
	err = c.BindJSON(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	dashboard, err := d.dashboardService.UpdateDashboard(
		c.Request.Context(),
		values.NewDashboardIDFromUUID(dashboardID),
		values.NewDashboardName(info.Name),
		values.NewDashboardDescription(info.Description),
	)
	if errors.Is(err, service.ErrNoDashboard) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no dashboard",
		})
		return
	}
	if err != nil {
		log.Printf("failed to create dashboard: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create dashboard",
		})
		return
	}

	c.JSON(http.StatusOK, DashboardInfo{
		ID:          uuid.UUID(dashboard.GetID()),
		Name:        string(dashboard.GetName()),
		Description: string(dashboard.GetDescription()),
		CreatedAt:   dashboard.GetCreatedAt(),
	})
}

func (d *Dashboard) DeleteDashboard(c *gin.Context) {
	strDashboardID := c.Param("dashboardID")
	dashboardID, err := uuid.Parse(strDashboardID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid dashboard id",
		})
		return
	}

	err = d.dashboardService.DeleteDashboard(
		c.Request.Context(),
		values.NewDashboardIDFromUUID(dashboardID),
	)
	if errors.Is(err, service.ErrNoDashboard) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no dashboard",
		})
		return
	}
	if err != nil {
		log.Printf("failed to delete dashboard: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete dashboard",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (d *Dashboard) GetMyDashboards(c *gin.Context) {
	session := d.session.getSession(c)
	user, err := d.session.getUser(session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get user",
		})
		return
	}

	dashboards, err := d.dashboardService.GetMyDashboards(
		c.Request.Context(),
		user,
	)
	if err != nil {
		log.Printf("failed to get my dashboards: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get my dashboards",
		})
		return
	}

	dashboardInfos := make([]*DashboardInfo, 0, len(dashboards))
	for _, dashboard := range dashboards {
		dashboardInfos = append(dashboardInfos, &DashboardInfo{
			ID:          uuid.UUID(dashboard.GetID()),
			Name:        string(dashboard.GetName()),
			Description: string(dashboard.GetDescription()),
			CreatedAt:   dashboard.GetCreatedAt(),
		})
	}

	c.JSON(http.StatusOK, dashboardInfos)
}

func (d *Dashboard) GetDashboardInfo(c *gin.Context) {
	strDashboardID := c.Param("dashboardID")
	dashboardID, err := uuid.Parse(strDashboardID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid dashboard id",
		})
		return
	}

	dashboard, err := d.dashboardService.GetDashboardInfo(
		c.Request.Context(),
		values.NewDashboardIDFromUUID(dashboardID),
	)
	if errors.Is(err, service.ErrNoDashboard) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no dashboard",
		})
		return
	}
	if err != nil {
		log.Printf("failed to get dashboard: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get dashboard",
		})
		return
	}

	taskStatusList := make([]TaskStatusDetail, 0, len(dashboard.TaskStatus))
	for _, taskStatus := range dashboard.TaskStatus {
		tasks := make([]TaskInfo, 0, len(taskStatus.Tasks))
		for _, task := range taskStatus.Tasks {
			tasks = append(tasks, TaskInfo{
				ID:          uuid.UUID(task.GetID()),
				Name:        string(task.GetName()),
				Description: string(task.GetDescription()),
				CreatedAt:   task.GetCreatedAt(),
			})
		}

		taskStatusList = append(taskStatusList, TaskStatusDetail{
			ID:    uuid.UUID(taskStatus.GetID()),
			Name:  string(taskStatus.GetName()),
			Tasks: tasks,
		})
	}

	c.JSON(http.StatusOK, DashboardDetail{
		ID:             uuid.UUID(dashboard.GetID()),
		Name:           string(dashboard.GetName()),
		Description:    string(dashboard.GetDescription()),
		CreatedAt:      dashboard.GetCreatedAt(),
		TaskStatusList: taskStatusList,
	})
}
