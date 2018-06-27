package projects

import (
	"github.com/machinebox/graphql"
	"log"
	"context"
	"github.com/enonic/projects/domain"
	"github.com/enonic/projects/queries"
	"github.com/enonic/projects/config"
)

type Client struct {
	config        *config.ClientConfig
	graphQLClient *graphql.Client
}

type ProjectResponse struct {
	Organizations []domain.Organization
}

func NewClient(config *config.ClientConfig) *Client {

	return &Client{
		config:        config,
		graphQLClient: graphql.NewClient(config.EndpointURL),
	}
}

func (c *Client) GetOrganizations() []domain.Organization {

	var respData ProjectResponse

	c.doQuery(queries.OrganizationsWithProjects, &respData)

	return respData.Organizations
}

func (c *Client) Query(respData interface{}) {

	c.doQuery(queries.OrganizationsWithProjects, respData)
}

func (c *Client) doQuery(query string, respData interface{}) {

	ctx := context.Background()

	req := graphql.NewRequest(query)

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("X-Auth-Token", c.config.AuthToken)

	if err := c.graphQLClient.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
}
