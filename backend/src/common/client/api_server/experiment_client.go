package api_server

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	apiclient "github.com/kubeflow/pipelines/backend/api/v1/go_http_client/experiment_client"
	params "github.com/kubeflow/pipelines/backend/api/v1/go_http_client/experiment_client/experiment_service"
	model "github.com/kubeflow/pipelines/backend/api/v1/go_http_client/experiment_model"
	"github.com/kubeflow/pipelines/backend/src/common/util"
	"golang.org/x/net/context"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
)

type ExperimentInterface interface {
	Create(params *params.ExperimentServiceCreateExperimentParams) (*model.V1Experiment, error)
	Get(params *params.ExperimentServiceGetExperimentParams) (*model.V1Experiment, error)
	List(params *params.ExperimentServiceListExperimentParams) ([]*model.V1Experiment, int, string, error)
	ListAll(params *params.ExperimentServiceListExperimentParams, maxResultSize int) ([]*model.V1Experiment, error)
	Archive(params *params.ExperimentServiceArchiveExperimentParams) error
	Unarchive(params *params.ExperimentServiceUnarchiveExperimentParams) error
}

type ExperimentClient struct {
	apiClient *apiclient.Experiment
}

func NewExperimentClient(clientConfig clientcmd.ClientConfig, debug bool) (
	*ExperimentClient, error) {

	runtime, err := NewHTTPRuntime(clientConfig, debug)
	if err != nil {
		return nil, err
	}

	apiClient := apiclient.New(runtime, strfmt.Default)

	// Creating upload client
	return &ExperimentClient{
		apiClient: apiClient,
	}, nil
}

func (c *ExperimentClient) Create(parameters *params.ExperimentServiceCreateExperimentParams) (*model.V1Experiment,
	error) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), apiServerDefaultTimeout)
	defer cancel()

	// Make service call
	parameters.Context = ctx
	response, err := c.apiClient.ExperimentService.ExperimentServiceCreateExperiment(parameters, PassThroughAuth)
	if err != nil {
		if defaultError, ok := err.(*params.ExperimentServiceCreateExperimentDefault); ok {
			err = CreateErrorFromAPIStatus(defaultError.Payload.Message, defaultError.Payload.Code)
		} else {
			err = CreateErrorCouldNotRecoverAPIStatus(err)
		}

		return nil, util.NewUserError(err,
			fmt.Sprintf("Failed to create experiment. Params: '%+v'. Body: '%+v'", parameters, parameters.Experiment),
			fmt.Sprintf("Failed to create experiment '%v'", parameters.Experiment.Name))
	}

	return response.Payload, nil
}

func (c *ExperimentClient) Get(parameters *params.ExperimentServiceGetExperimentParams) (*model.V1Experiment,
	error) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), apiServerDefaultTimeout)
	defer cancel()

	// Make service call
	parameters.Context = ctx
	response, err := c.apiClient.ExperimentService.ExperimentServiceGetExperiment(parameters, PassThroughAuth)
	if err != nil {
		if defaultError, ok := err.(*params.ExperimentServiceGetExperimentDefault); ok {
			err = CreateErrorFromAPIStatus(defaultError.Payload.Message, defaultError.Payload.Code)
		} else {
			err = CreateErrorCouldNotRecoverAPIStatus(err)
		}

		return nil, util.NewUserError(err,
			fmt.Sprintf("Failed to get experiment. Params: '%+v'", parameters),
			fmt.Sprintf("Failed to get experiment '%v'", parameters.ID))
	}

	return response.Payload, nil
}

func (c *ExperimentClient) List(parameters *params.ExperimentServiceListExperimentParams) (
	[]*model.V1Experiment, int, string, error) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), apiServerDefaultTimeout)
	defer cancel()

	// Make service call
	parameters.Context = ctx
	response, err := c.apiClient.ExperimentService.ExperimentServiceListExperiment(parameters, PassThroughAuth)
	if err != nil {
		if defaultError, ok := err.(*params.ExperimentServiceListExperimentDefault); ok {
			err = CreateErrorFromAPIStatus(defaultError.Payload.Message, defaultError.Payload.Code)
		} else {
			err = CreateErrorCouldNotRecoverAPIStatus(err)
		}

		return nil, 0, "", util.NewUserError(err,
			fmt.Sprintf("Failed to list experiments. Params: '%+v'", parameters),
			fmt.Sprintf("Failed to list experiments"))
	}

	return response.Payload.Experiments, int(response.Payload.TotalSize), response.Payload.NextPageToken, nil
}

func (c *ExperimentClient) Delete(parameters *params.ExperimentServiceDeleteExperimentParams) error {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), apiServerDefaultTimeout)
	defer cancel()

	// Make service call
	parameters.Context = ctx
	_, err := c.apiClient.ExperimentService.ExperimentServiceDeleteExperiment(parameters, PassThroughAuth)
	if err != nil {
		if defaultError, ok := err.(*params.ExperimentServiceDeleteExperimentDefault); ok {
			err = CreateErrorFromAPIStatus(defaultError.Payload.Message, defaultError.Payload.Code)
		} else {
			err = CreateErrorCouldNotRecoverAPIStatus(err)
		}

		return util.NewUserError(err,
			fmt.Sprintf("Failed to delete experiments. Params: '%+v'", parameters),
			fmt.Sprintf("Failed to delete experiment"))
	}

	return nil
}

func (c *ExperimentClient) ListAll(parameters *params.ExperimentServiceListExperimentParams, maxResultSize int) (
	[]*model.V1Experiment, error) {
	return listAllForExperiment(c, parameters, maxResultSize)
}

func listAllForExperiment(client ExperimentInterface, parameters *params.ExperimentServiceListExperimentParams,
	maxResultSize int) ([]*model.V1Experiment, error) {
	if maxResultSize < 0 {
		maxResultSize = 0
	}

	allResults := make([]*model.V1Experiment, 0)
	firstCall := true
	for (firstCall || (parameters.PageToken != nil && *parameters.PageToken != "")) &&
		(len(allResults) < maxResultSize) {
		results, _, pageToken, err := client.List(parameters)
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, results...)
		parameters.PageToken = util.StringPointer(pageToken)
		firstCall = false
	}
	if len(allResults) > maxResultSize {
		allResults = allResults[0:maxResultSize]
	}

	return allResults, nil
}

func (c *ExperimentClient) Archive(parameters *params.ExperimentServiceArchiveExperimentParams) error {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), apiServerDefaultTimeout)
	defer cancel()

	// Make service call
	parameters.Context = ctx
	_, err := c.apiClient.ExperimentService.ExperimentServiceArchiveExperiment(parameters, PassThroughAuth)

	if err != nil {
		if defaultError, ok := err.(*params.ExperimentServiceArchiveExperimentDefault); ok {
			err = CreateErrorFromAPIStatus(defaultError.Payload.Message, defaultError.Payload.Code)
		} else {
			err = CreateErrorCouldNotRecoverAPIStatus(err)
		}

		return util.NewUserError(err,
			fmt.Sprintf("Failed to archive experiments. Params: '%+v'", parameters),
			fmt.Sprintf("Failed to archive experiments"))
	}

	return nil
}

func (c *ExperimentClient) Unarchive(parameters *params.ExperimentServiceUnarchiveExperimentParams) error {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), apiServerDefaultTimeout)
	defer cancel()

	// Make service call
	parameters.Context = ctx
	_, err := c.apiClient.ExperimentService.ExperimentServiceUnarchiveExperiment(parameters, PassThroughAuth)

	if err != nil {
		if defaultError, ok := err.(*params.ExperimentServiceUnarchiveExperimentDefault); ok {
			err = CreateErrorFromAPIStatus(defaultError.Payload.Message, defaultError.Payload.Code)
		} else {
			err = CreateErrorCouldNotRecoverAPIStatus(err)
		}

		return util.NewUserError(err,
			fmt.Sprintf("Failed to unarchive experiments. Params: '%+v'", parameters),
			fmt.Sprintf("Failed to unarchive experiments"))
	}

	return nil
}
