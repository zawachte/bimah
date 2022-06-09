package providers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zawachte/bimah/internal/models"
	"github.com/zawachte/bimah/internal/services"
)

type Provider interface {
	// Returns a list of all fleets
	// (GET /fleet)
	GetFleet(c *gin.Context)
	// Create an fleet
	// (POST /fleet)
	PostFleet(c *gin.Context)
	// Delete fleet with fleet id
	// (DELETE /fleet/{id})
	DeleteFleetId(c *gin.Context, id int)
	// Get fleet with specific id
	// (GET /fleet/{id})
	GetFleetId(c *gin.Context, id int)
	// Returns a list of all paths
	// (GET /path)
	GetPath(c *gin.Context)
	// Create an path
	// (POST /path)
	PostPath(c *gin.Context)
	// Delete Path with path id
	// (DELETE /path/{id})
	DeletePathId(c *gin.Context, id int)
	// Get account with specific id
	// (GET /path/{id})
	GetPathId(c *gin.Context, id int)
	// Returns a list of all rules
	// (GET /rule)
	GetRule(c *gin.Context)
	// Create a rule
	// (POST /rule)
	PostRule(c *gin.Context)
	// Delete rule with specific id
	// (DELETE /rule/{id})
	DeleteRuleId(c *gin.Context, id int)
	// Get rule with specific id
	// (GET /rule/{id})
	GetRuleId(c *gin.Context, id int)
	// Returns a list of all tlsconfig
	// (GET /tlsconfig)
	GetTlsconfig(c *gin.Context)
	// Create an tlsconfig
	// (POST /tlsconfig)
	PostTlsconfig(c *gin.Context)
	// Delete tlsconfig with tlsconfig id
	// (DELETE /tlsconfig/{id})
	DeleteTlsconfigId(c *gin.Context, id int)
	// Get tlsconfig with specific id
	// (GET /tlsconfig/{id})
	GetTlsconfigId(c *gin.Context, id int)
}

type ProviderParams struct {
	DatabaseUrl string
}

type provider struct {
	fleetService     services.FleetService
	pathService      services.PathService
	ruleService      services.RuleService
	tlsconfigService services.TlsconfigService
}

func NewProvider(ctx context.Context, params ProviderParams) (*provider, error) {

	fleetService, err := services.NewFleetService(ctx, services.FleetServiceParams{
		DatabaseUrl: params.DatabaseUrl,
	})
	if err != nil {
		return nil, err
	}

	pathService, err := services.NewPathService(ctx, services.PathServiceParams{
		DatabaseUrl: params.DatabaseUrl,
	})
	if err != nil {
		return nil, err
	}

	ruleService, err := services.NewRuleService(ctx, services.RuleServiceParams{
		DatabaseUrl: params.DatabaseUrl,
	})
	if err != nil {
		return nil, err
	}

	tlsconfigService, err := services.NewTlsconfigService(ctx, services.TlsconfigServiceParams{
		DatabaseUrl: params.DatabaseUrl,
	})
	if err != nil {
		return nil, err
	}

	return &provider{
		fleetService:     fleetService,
		pathService:      pathService,
		ruleService:      ruleService,
		tlsconfigService: tlsconfigService,
	}, nil
}

func (p *provider) GetFleet(c *gin.Context) {
	fleets, err := p.fleetService.GetFleets(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, fleets)
}

func (p *provider) PostFleet(c *gin.Context) {
	fleet := models.Fleet{}
	err := json.NewDecoder(c.Request.Body).Decode(&fleet)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = p.fleetService.CreateFleet(c.Request.Context(), fleet)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (p *provider) DeleteFleetId(c *gin.Context, id int) {
	err := p.fleetService.DeleteFleetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (p *provider) GetFleetId(c *gin.Context, id int) {
	fleet, err := p.fleetService.GetFleetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, fleet)
}

func (p *provider) GetPath(c *gin.Context) {
	paths, err := p.pathService.GetPaths(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, paths)
}

func (p *provider) PostPath(c *gin.Context) {
	path := models.Path{}
	err := json.NewDecoder(c.Request.Body).Decode(&path)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = p.pathService.CreatePath(c.Request.Context(), path)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (p *provider) DeletePathId(c *gin.Context, id int) {
	err := p.pathService.DeletePathById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
func (p *provider) GetPathId(c *gin.Context, id int) {
	path, err := p.pathService.GetPathById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, path)
}

func (p *provider) GetRule(c *gin.Context) {
	rules, err := p.ruleService.GetRules(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, rules)
}

func (p *provider) PostRule(c *gin.Context) {
	rule := models.Rule{}
	err := json.NewDecoder(c.Request.Body).Decode(&rule)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = p.ruleService.CreateRule(c.Request.Context(), rule)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (p *provider) DeleteRuleId(c *gin.Context, id int) {
	err := p.ruleService.DeleteRuleById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
func (p *provider) GetRuleId(c *gin.Context, id int) {
	rule, err := p.ruleService.GetRuleById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, rule)
}

func (p *provider) GetTlsconfig(c *gin.Context) {
	tlsconfigs, err := p.tlsconfigService.GetTlsconfigs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, tlsconfigs)
}

func (p *provider) PostTlsconfig(c *gin.Context) {
	tlsconfig := models.Tlsconfig{}
	err := json.NewDecoder(c.Request.Body).Decode(&tlsconfig)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = p.tlsconfigService.CreateTlsconfig(c.Request.Context(), tlsconfig)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (p *provider) DeleteTlsconfigId(c *gin.Context, id int) {
	err := p.tlsconfigService.DeleteTlsconfigById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
func (p *provider) GetTlsconfigId(c *gin.Context, id int) {
	tlsconfig, err := p.tlsconfigService.GetTlsconfigById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, tlsconfig)
}
