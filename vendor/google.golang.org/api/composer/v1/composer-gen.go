// Package composer provides access to the Cloud Composer API.
//
// See https://cloud.google.com/composer/
//
// Usage example:
//
//   import "google.golang.org/api/composer/v1"
//   ...
//   composerService, err := composer.New(oauthHttpClient)
package composer // import "google.golang.org/api/composer/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "composer:v1"
const apiName = "composer"
const apiVersion = "v1"
const basePath = "https://composer.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Locations = NewProjectsLocationsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Locations *ProjectsLocationsService
}

func NewProjectsLocationsService(s *Service) *ProjectsLocationsService {
	rs := &ProjectsLocationsService{s: s}
	rs.Environments = NewProjectsLocationsEnvironmentsService(s)
	rs.Operations = NewProjectsLocationsOperationsService(s)
	return rs
}

type ProjectsLocationsService struct {
	s *Service

	Environments *ProjectsLocationsEnvironmentsService

	Operations *ProjectsLocationsOperationsService
}

func NewProjectsLocationsEnvironmentsService(s *Service) *ProjectsLocationsEnvironmentsService {
	rs := &ProjectsLocationsEnvironmentsService{s: s}
	return rs
}

type ProjectsLocationsEnvironmentsService struct {
	s *Service
}

func NewProjectsLocationsOperationsService(s *Service) *ProjectsLocationsOperationsService {
	rs := &ProjectsLocationsOperationsService{s: s}
	return rs
}

type ProjectsLocationsOperationsService struct {
	s *Service
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// Environment: An environment for running orchestration tasks.
type Environment struct {
	// Config: Configuration parameters for this environment.
	Config *EnvironmentConfig `json:"config,omitempty"`

	// CreateTime: Output only.
	// The time at which this environment was created.
	CreateTime string `json:"createTime,omitempty"`

	// Labels: Optional. User-defined labels for this environment.
	// The labels map can contain no more than 64 entries. Entries of the
	// labels
	// map are UTF8 strings that comply with the following restrictions:
	//
	// * Keys must conform to regexp: \p{Ll}\p{Lo}{0,62}
	// * Values must conform to regexp:  [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// * Both keys and values are additionally constrained to be <= 128
	// bytes in
	// size.
	Labels map[string]string `json:"labels,omitempty"`

	// Name: The resource name of the environment, in the
	// form:
	// "projects/{projectId}/locations/{locationId}/environments/{envir
	// onmentId}"
	Name string `json:"name,omitempty"`

	// State: The current state of the environment.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - The state of the environment is unknown.
	//   "CREATING" - The environment is in the process of being created.
	//   "RUNNING" - The environment is currently running and healthy. It is
	// ready for use.
	//   "UPDATING" - The environment is being updated. It remains usable
	// but cannot receive
	// additional update requests or be deleted at this time.
	//   "DELETING" - The environment is undergoing deletion. It cannot be
	// used.
	//   "ERROR" - The environment has encountered an error and cannot be
	// used.
	State string `json:"state,omitempty"`

	// UpdateTime: Output only.
	// The time at which this environment was last modified.
	UpdateTime string `json:"updateTime,omitempty"`

	// Uuid: Output only.
	// The UUID (Universally Unique IDentifier) associated with this
	// environment.
	// This value is generated when the environment is created.
	Uuid string `json:"uuid,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Config") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Config") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Environment) MarshalJSON() ([]byte, error) {
	type NoMethod Environment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// EnvironmentConfig: Configuration information for an environment.
type EnvironmentConfig struct {
	// AirflowUri: The URI of the Apache Airflow Web UI hosted within this
	// environment (see
	// [Airflow web
	// interface](/composer/docs/how-to/accessing/airflow-web-interface)).
	AirflowUri string `json:"airflowUri,omitempty"`

	// DagGcsPrefix: Output only.
	// The Cloud Storage prefix of the DAGs for this environment. Although
	// Cloud
	// Storage objects reside in a flat namespace, a hierarchical file
	// tree
	// can be simulated using "/"-delimited object name prefixes. DAG
	// objects for
	// this environment reside in a simulated directory with the given
	// prefix.
	DagGcsPrefix string `json:"dagGcsPrefix,omitempty"`

	// GkeCluster: Output only.
	// The Kubernetes Engine cluster used to run this environment.
	GkeCluster string `json:"gkeCluster,omitempty"`

	// NodeConfig: The configuration used for the Kubernetes Engine cluster.
	NodeConfig *NodeConfig `json:"nodeConfig,omitempty"`

	// NodeCount: The number of nodes in the Kubernetes Engine cluster that
	// will be
	// used to run this environment.
	NodeCount int64 `json:"nodeCount,omitempty"`

	// SoftwareConfig: The configuration settings for software inside the
	// environment.
	SoftwareConfig *SoftwareConfig `json:"softwareConfig,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AirflowUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AirflowUri") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *EnvironmentConfig) MarshalJSON() ([]byte, error) {
	type NoMethod EnvironmentConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListEnvironmentsResponse: The environments in a project and location.
type ListEnvironmentsResponse struct {
	// Environments: The list of environments returned by a
	// ListEnvironmentsRequest.
	Environments []*Environment `json:"environments,omitempty"`

	// NextPageToken: The page token used to query for the next page if one
	// exists.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Environments") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Environments") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListEnvironmentsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListEnvironmentsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListOperationsResponse: The response message for
// Operations.ListOperations.
type ListOperationsResponse struct {
	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: A list of operations that matches the specified filter in
	// the request.
	Operations []*Operation `json:"operations,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListOperationsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListOperationsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NodeConfig: The configuration information for the Kubernetes Engine
// nodes running
// the Apache Airflow software.
type NodeConfig struct {
	// DiskSizeGb: Optional. The disk size in GB used for node VMs. Minimum
	// size is 20GB.
	// If unspecified, defaults to 100GB. Cannot be updated.
	DiskSizeGb int64 `json:"diskSizeGb,omitempty"`

	// Location: Optional. The Compute Engine
	// [zone](/compute/docs/regions-zones) in which
	// to deploy the VMs used to run the Apache Airflow software, specified
	// as a
	// [relative resource
	// name](/apis/design/resource_names#relative_resource_name).
	// For example: "projects/{projectId}/zones/{zoneId}".
	//
	// This `location` must belong to the enclosing environment's project
	// and
	// location. If both this field and `nodeConfig.machineType` are
	// specified,
	// `nodeConfig.machineType` must belong to this `location`; if both
	// are
	// unspecified, the service will pick a zone in the Compute Engine
	// region
	// corresponding to the Cloud Composer location, and propagate that
	// choice to
	// both fields. If only one field (`location` or
	// `nodeConfig.machineType`) is
	// specified, the location information from the specified field will
	// be
	// propagated to the unspecified field.
	Location string `json:"location,omitempty"`

	// MachineType: Optional. The Compute Engine
	// [machine type](/compute/docs/machine-types) used for cluster
	// instances,
	// specified as a
	// [relative resource
	// name](/apis/design/resource_names#relative_resource_name).
	// For
	// example:
	// "projects/{projectId}/zones/{zoneId}/machineTypes/{machineTyp
	// eId}".
	//
	// The `machineType` must belong to the enclosing environment's project
	// and
	// location. If both this field and `nodeConfig.location` are
	// specified,
	// this `machineType` must belong to the `nodeConfig.location`; if both
	// are
	// unspecified, the service will pick a zone in the Compute Engine
	// region
	// corresponding to the Cloud Composer location, and propagate that
	// choice to
	// both fields. If exactly one of this field and `nodeConfig.location`
	// is
	// specified, the location information from the specified field will
	// be
	// propagated to the unspecified field.
	//
	// If this field is unspecified, the `machineTypeId` defaults
	// to "n1-standard-1".
	MachineType string `json:"machineType,omitempty"`

	// Network: Optional. The Compute Engine network to be used for
	// machine
	// communications, specified as a
	// [relative resource
	// name](/apis/design/resource_names#relative_resource_name).
	// For example:
	// "projects/{projectId}/global/networks/{networkId}".
	//
	// [Shared VPC](/vpc/docs/shared-vpc) is not currently supported.
	// The
	// network must belong to the environment's project. If unspecified,
	// the
	// "default" network ID in the environment's project is used.  If
	// a
	// [Custom Subnet Network]((/vpc/docs/vpc#vpc_networks_and_subnets)
	// is provided, `nodeConfig.subnetwork` must also be provided.
	Network string `json:"network,omitempty"`

	// OauthScopes: Optional. The set of Google API scopes to be made
	// available on all
	// node VMs. If `oauth_scopes` is empty, defaults
	// to
	// ["https://www.googleapis.com/auth/cloud-platform"]. Cannot be
	// updated.
	OauthScopes []string `json:"oauthScopes,omitempty"`

	// ServiceAccount: Optional. The Google Cloud Platform Service Account
	// to be used by the node
	// VMs. If a service account is not specified, the "default" Compute
	// Engine
	// service account is used. Cannot be updated.
	ServiceAccount string `json:"serviceAccount,omitempty"`

	// Subnetwork: Optional. The Compute Engine subnetwork to be used for
	// machine
	// communications, specified as a
	// [relative resource
	// name](/apis/design/resource_names#relative_resource_name).
	// For
	// example:
	// "projects/{projectId}/regions/{regionId}/subnetworks/{subnetw
	// orkId}"
	//
	// If a subnetwork is provided, `nodeConfig.network` must also be
	// provided,
	// and the subnetwork must belong to the enclosing environment's project
	// and
	// location.
	Subnetwork string `json:"subnetwork,omitempty"`

	// Tags: Optional. The list of instance tags applied to all node VMs.
	// Tags are used
	// to identify valid sources or targets for network firewalls. Each tag
	// within
	// the list must comply with
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
	// Cannot be updated.
	Tags []string `json:"tags,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DiskSizeGb") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DiskSizeGb") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NodeConfig) MarshalJSON() ([]byte, error) {
	type NoMethod NodeConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Operation: This resource represents a long-running operation that is
// the result of a
// network API call.
type Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If `true`, the operation is completed, and either `error` or
	// `response` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such as create
	// time.
	// Some services might not provide such metadata.  Any method that
	// returns a
	// long-running operation should document the metadata type, if any.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`, the response
	// is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For
	// other
	// methods, the response should have the type `XxxResponse`, where
	// `Xxx`
	// is the original method name.  For example, if the original method
	// name
	// is `TakeSnapshot()`, the inferred response type
	// is
	// `TakeSnapshotResponse`.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type NoMethod Operation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OperationMetadata: Metadata describing an operation.
type OperationMetadata struct {
	// CreateTime: Output only.
	// The time the operation was submitted to the server.
	CreateTime string `json:"createTime,omitempty"`

	// EndTime: Output only.
	// The time when the operation terminated, regardless of its
	// success.
	// This field is unset if the operation is still ongoing.
	EndTime string `json:"endTime,omitempty"`

	// OperationType: Output only.
	// The type of operation being performed.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Unused.
	//   "CREATE" - A resource creation operation.
	//   "DELETE" - A resource deletion operation.
	//   "UPDATE" - A resource update operation.
	OperationType string `json:"operationType,omitempty"`

	// Resource: Output only.
	// The resource being operated on, as a [relative resource
	// name](
	// /apis/design/resource_names#relative_resource_name).
	Resource string `json:"resource,omitempty"`

	// ResourceUuid: Output only.
	// The UUID of the resource being operated on.
	ResourceUuid string `json:"resourceUuid,omitempty"`

	// State: Output only.
	// The current operation state.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - Unused.
	//   "PENDING" - The operation has been created but is not yet started.
	//   "RUNNING" - The operation is underway.
	//   "SUCCEEDED" - The operation completed successfully.
	//   "SUCCESSFUL"
	//   "FAILED" - The operation is no longer running but did not succeed.
	State string `json:"state,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OperationMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod OperationMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SoftwareConfig: Specifies the selection and configuration of software
// inside the environment.
type SoftwareConfig struct {
	// AirflowConfigOverrides: Optional. Apache Airflow configuration
	// properties to override.
	//
	// Property keys contain the section and property names, separated by a
	// hyphen,
	// for example "core-dags_are_paused_at_creation". Section names must
	// not
	// contain hyphens ("-"), opening square brackets ("["),  or closing
	// square
	// brackets ("]"). The property name must not be empty and must not
	// contain
	// an equals sign ("=") or semicolon (";"). Section and property names
	// must
	// not contain a period ("."). Apache Airflow configuration property
	// names
	// must be written in
	// [snake_case](https://en.wikipedia.org/wiki/Snake_case).
	// Property values can contain any character, and can be written in
	// any
	// lower/upper case format.
	//
	// Certain Apache Airflow configuration property values
	// are
	// [blacklisted](/composer/docs/how-to/managing/setting-airflow-confi
	// gurations#airflow_configuration_blacklists),
	// and cannot be overridden.
	AirflowConfigOverrides map[string]string `json:"airflowConfigOverrides,omitempty"`

	// EnvVariables: Optional. Additional environment variables to provide
	// to the Apache Airflow
	// scheduler, worker, and webserver processes.
	//
	// Environment variable names must match the regular
	// expression
	// `a-zA-Z_*`. They cannot specify Apache Airflow
	// software configuration overrides (they cannot match the regular
	// expression
	// `AIRFLOW__[A-Z0-9_]+__[A-Z0-9_]+`), and they cannot match any of
	// the
	// following reserved names:
	//
	// * `AIRFLOW_HOME`
	// * `C_FORCE_ROOT`
	// * `CONTAINER_NAME`
	// * `DAGS_FOLDER`
	// * `GCP_PROJECT`
	// * `GCS_BUCKET`
	// * `GKE_CLUSTER_NAME`
	// * `SQL_DATABASE`
	// * `SQL_INSTANCE`
	// * `SQL_PASSWORD`
	// * `SQL_PROJECT`
	// * `SQL_REGION`
	// * `SQL_USER`
	EnvVariables map[string]string `json:"envVariables,omitempty"`

	// ImageVersion: Output only.
	// The version of the software running in the environment.
	// This encapsulates both the version of Cloud Composer functionality
	// and the
	// version of Apache Airflow. It must match the regular
	// expression
	// `composer-[0-9]+\.[0-9]+(\.[0-9]+)?-airflow-[0-9]+\.[0-9]+(
	// \.[0-9]+.*)?`.
	//
	// The Cloud Composer portion of the version is a
	// [semantic version](https://semver.org). The portion of the image
	// version
	// following _airflow-_ is an official Apache Airflow
	// repository
	// [release
	// name](https://github.com/apache/incubator-airflow/releases).
	//
	// See also [Release Notes](/composer/docs/release-notes).
	ImageVersion string `json:"imageVersion,omitempty"`

	// PypiPackages: Optional. Custom Python Package Index (PyPI) packages
	// to be installed in
	// the environment.
	//
	// Keys refer to the lowercase package name such as "numpy"
	// and values are the lowercase extras and version specifier such
	// as
	// "==1.12.0", "[devel,gcp_api]", or "[devel]>=1.8.2, <1.9.2". To
	// specify a
	// package without pinning it to a version specifier, use the empty
	// string as
	// the value.
	PypiPackages map[string]string `json:"pypiPackages,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AirflowConfigOverrides") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AirflowConfigOverrides")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SoftwareConfig) MarshalJSON() ([]byte, error) {
	type NoMethod SoftwareConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "composer.projects.locations.environments.create":

type ProjectsLocationsEnvironmentsCreateCall struct {
	s           *Service
	parent      string
	environment *Environment
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Create: Create a new environment.
func (r *ProjectsLocationsEnvironmentsService) Create(parent string, environment *Environment) *ProjectsLocationsEnvironmentsCreateCall {
	c := &ProjectsLocationsEnvironmentsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.environment = environment
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsEnvironmentsCreateCall) Fields(s ...googleapi.Field) *ProjectsLocationsEnvironmentsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsEnvironmentsCreateCall) Context(ctx context.Context) *ProjectsLocationsEnvironmentsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsEnvironmentsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsEnvironmentsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.environment)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/environments")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "composer.projects.locations.environments.create" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsLocationsEnvironmentsCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a new environment.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/environments",
	//   "httpMethod": "POST",
	//   "id": "composer.projects.locations.environments.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "The parent must be of the form \"projects/{projectId}/locations/{locationId}\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/environments",
	//   "request": {
	//     "$ref": "Environment"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "composer.projects.locations.environments.delete":

type ProjectsLocationsEnvironmentsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Delete an environment.
func (r *ProjectsLocationsEnvironmentsService) Delete(name string) *ProjectsLocationsEnvironmentsDeleteCall {
	c := &ProjectsLocationsEnvironmentsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsEnvironmentsDeleteCall) Fields(s ...googleapi.Field) *ProjectsLocationsEnvironmentsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsEnvironmentsDeleteCall) Context(ctx context.Context) *ProjectsLocationsEnvironmentsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsEnvironmentsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsEnvironmentsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "composer.projects.locations.environments.delete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsLocationsEnvironmentsDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Delete an environment.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/environments/{environmentsId}",
	//   "httpMethod": "DELETE",
	//   "id": "composer.projects.locations.environments.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The environment to delete, in the form:\n\"projects/{projectId}/locations/{locationId}/environments/{environmentId}\"",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/environments/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "composer.projects.locations.environments.get":

type ProjectsLocationsEnvironmentsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Get an existing environment.
func (r *ProjectsLocationsEnvironmentsService) Get(name string) *ProjectsLocationsEnvironmentsGetCall {
	c := &ProjectsLocationsEnvironmentsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsEnvironmentsGetCall) Fields(s ...googleapi.Field) *ProjectsLocationsEnvironmentsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsEnvironmentsGetCall) IfNoneMatch(entityTag string) *ProjectsLocationsEnvironmentsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsEnvironmentsGetCall) Context(ctx context.Context) *ProjectsLocationsEnvironmentsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsEnvironmentsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsEnvironmentsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "composer.projects.locations.environments.get" call.
// Exactly one of *Environment or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Environment.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsLocationsEnvironmentsGetCall) Do(opts ...googleapi.CallOption) (*Environment, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Environment{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get an existing environment.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/environments/{environmentsId}",
	//   "httpMethod": "GET",
	//   "id": "composer.projects.locations.environments.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the environment to get, in the form:\n\"projects/{projectId}/locations/{locationId}/environments/{environmentId}\"",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/environments/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Environment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "composer.projects.locations.environments.list":

type ProjectsLocationsEnvironmentsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List environments.
func (r *ProjectsLocationsEnvironmentsService) List(parent string) *ProjectsLocationsEnvironmentsListCall {
	c := &ProjectsLocationsEnvironmentsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of environments to return.
func (c *ProjectsLocationsEnvironmentsListCall) PageSize(pageSize int64) *ProjectsLocationsEnvironmentsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous List request, if any.
func (c *ProjectsLocationsEnvironmentsListCall) PageToken(pageToken string) *ProjectsLocationsEnvironmentsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsEnvironmentsListCall) Fields(s ...googleapi.Field) *ProjectsLocationsEnvironmentsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsEnvironmentsListCall) IfNoneMatch(entityTag string) *ProjectsLocationsEnvironmentsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsEnvironmentsListCall) Context(ctx context.Context) *ProjectsLocationsEnvironmentsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsEnvironmentsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsEnvironmentsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/environments")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "composer.projects.locations.environments.list" call.
// Exactly one of *ListEnvironmentsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListEnvironmentsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsEnvironmentsListCall) Do(opts ...googleapi.CallOption) (*ListEnvironmentsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListEnvironmentsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List environments.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/environments",
	//   "httpMethod": "GET",
	//   "id": "composer.projects.locations.environments.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of environments to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The next_page_token value returned from a previous List request, if any.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "List environments in the given project and location, in the form:\n\"projects/{projectId}/locations/{locationId}\"",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/environments",
	//   "response": {
	//     "$ref": "ListEnvironmentsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsLocationsEnvironmentsListCall) Pages(ctx context.Context, f func(*ListEnvironmentsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "composer.projects.locations.environments.patch":

type ProjectsLocationsEnvironmentsPatchCall struct {
	s           *Service
	name        string
	environment *Environment
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Patch: Update an environment.
func (r *ProjectsLocationsEnvironmentsService) Patch(name string, environment *Environment) *ProjectsLocationsEnvironmentsPatchCall {
	c := &ProjectsLocationsEnvironmentsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.environment = environment
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required. A
// comma-separated list of paths, relative to `Environment`, of
// fields to update.
// For example, to set the version of scikit-learn to install in
// the
// environment to 0.19.0 and to remove an existing installation
// of
// numpy, the `updateMask` parameter would include the following
// two
// `paths` values: "config.softwareConfig.pypiPackages.scikit-learn"
// and
// "config.softwareConfig.pypiPackages.numpy". The included
// patch
// environment would specify the scikit-learn version as follows:
//
//     {
//       "config":{
//         "softwareConfig":{
//           "pypiPackages":{
//             "scikit-learn":"==0.19.0"
//           }
//         }
//       }
//     }
//
// Note that in the above example, any existing PyPI packages
// other than scikit-learn and numpy will be unaffected.
//
// Only one update type may be included in a single request's
// `updateMask`.
// For example, one cannot update both the PyPI packages and
// labels in the same request. However, it is possible to update
// multiple
// members of a map field simultaneously in the same request. For
// example,
// to set the labels "label1" and "label2" while clearing "label3"
// (assuming
// it already exists), one can
// provide the paths "labels.label1", "labels.label2", and
// "labels.label3"
// and populate the patch environment as follows:
//
//     {
//       "labels":{
//         "label1":"new-label1-value"
//         "label2":"new-label2-value"
//       }
//     }
//
// Note that in the above example, any existing labels that are
// not
// included in the `updateMask` will be unaffected.
//
// It is also possible to replace an entire map field by providing
// the
// map field's path in the `updateMask`. The new value of the field
// will
// be that which is provided in the patch environment. For example,
// to
// delete all pre-existing user-specified PyPI packages and
// install botocore at version 1.7.14, the `updateMask` would
// contain
// the path "config.softwareConfig.pypiPackages", and
// the patch environment would be the following:
//
//     {
//       "config":{
//         "softwareConfig":{
//           "pypiPackages":{
//             "botocore":"==1.7.14"
//           }
//         }
//       }
//     }
//
// **Note:** Only the following fields can be updated:
//
//  <table>
//  <tbody>
//  <tr>
//  <td><strong>Mask</strong></td>
//  <td><strong>Purpose</strong></td>
//  </tr>
//  <tr>
//  <td>config.softwareConfig.pypiPackages
//  </td>
//  <td>Replace all custom custom PyPI packages. If a replacement
//  package map is not included in `environment`, all custom
//  PyPI packages are cleared. It is an error to provide both this mask
// and a
//  mask specifying an individual package.</td>
//  </tr>
//  <tr>
//  <td>config.softwareConfig.pypiPackages.<var>packagename</var></td>
//  <td>Update the custom PyPI package <var>packagename</var>,
//  preserving other packages. To delete the package, include it in
//  `updateMask`, and omit the mapping for it in
//  `environment.config.softwareConfig.pypiPackages`. It is an error
//  to provide both a mask of this form and the
//  "config.softwareConfig.pypiPackages" mask.</td>
//  </tr>
//  <tr>
//  <td>labels</td>
//  <td>Replace all environment labels. If a replacement labels map is
// not
//  included in `environment`, all labels are cleared. It is an error
// to
//  provide both this mask and a mask specifying one or more individual
//  labels.</td>
//  </tr>
//  <tr>
//  <td>labels.<var>labelName</var></td>
//  <td>Set the label named <var>labelName</var>, while preserving
// other
//  labels. To delete the label, include it in `updateMask` and omit
// its
//  mapping in `environment.labels`. It is an error to provide both a
//  mask of this form and the "labels" mask.</td>
//  </tr>
//  <tr>
//  <td>config.nodeCount</td>
//  <td>Horizontally scale the number of nodes in the environment. An
// integer
//  greater than or equal to 3 must be provided in the
// `config.nodeCount` field.
//  </td>
//  </tr>
//  <tr>
//  <td>config.softwareConfig.airflowConfigOverrides</td>
//  <td>Replace all Apache Airflow config overrides. If a replacement
// config
//  overrides map is not included in `environment`, all config
// overrides
//  are cleared.
//  It is an error to provide both this mask and a mask specifying one
// or
//  more individual config overrides.</td>
//  </tr>
//  <tr>
//  <td>config.softwareConfig.properties.<var>section</var>-<var>name
//  </var></td>
//  <td>Override the Apache Airflow property <var>name</var> in the
// section
//  named <var>section</var>, preserving other properties. To delete
// the
//  property override, include it in `updateMask` and omit its mapping
//  in `environment.config.softwareConfig.properties`.
//  It is an error to provide both a mask of this form and the
//  "config.softwareConfig.properties" mask.</td>
//  </tr>
//  <tr>
//  <td>config.softwareConfig.envVariables</td>
//  <td>Replace all environment variables. If a replacement environment
//  variable map is not included in `environment`, all custom
// environment
//  variables  are cleared.
//  It is an error to provide both this mask and a mask specifying one
// or
//  more individual environment variables.</td>
//  </tr>
//  </tbody>
//  </table>
func (c *ProjectsLocationsEnvironmentsPatchCall) UpdateMask(updateMask string) *ProjectsLocationsEnvironmentsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsEnvironmentsPatchCall) Fields(s ...googleapi.Field) *ProjectsLocationsEnvironmentsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsEnvironmentsPatchCall) Context(ctx context.Context) *ProjectsLocationsEnvironmentsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsEnvironmentsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsEnvironmentsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.environment)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "composer.projects.locations.environments.patch" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsLocationsEnvironmentsPatchCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update an environment.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/environments/{environmentsId}",
	//   "httpMethod": "PATCH",
	//   "id": "composer.projects.locations.environments.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The relative resource name of the environment to update, in the form:\n\"projects/{projectId}/locations/{locationId}/environments/{environmentId}\"",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/environments/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required. A comma-separated list of paths, relative to `Environment`, of\nfields to update.\nFor example, to set the version of scikit-learn to install in the\nenvironment to 0.19.0 and to remove an existing installation of\nnumpy, the `updateMask` parameter would include the following two\n`paths` values: \"config.softwareConfig.pypiPackages.scikit-learn\" and\n\"config.softwareConfig.pypiPackages.numpy\". The included patch\nenvironment would specify the scikit-learn version as follows:\n\n    {\n      \"config\":{\n        \"softwareConfig\":{\n          \"pypiPackages\":{\n            \"scikit-learn\":\"==0.19.0\"\n          }\n        }\n      }\n    }\n\nNote that in the above example, any existing PyPI packages\nother than scikit-learn and numpy will be unaffected.\n\nOnly one update type may be included in a single request's `updateMask`.\nFor example, one cannot update both the PyPI packages and\nlabels in the same request. However, it is possible to update multiple\nmembers of a map field simultaneously in the same request. For example,\nto set the labels \"label1\" and \"label2\" while clearing \"label3\" (assuming\nit already exists), one can\nprovide the paths \"labels.label1\", \"labels.label2\", and \"labels.label3\"\nand populate the patch environment as follows:\n\n    {\n      \"labels\":{\n        \"label1\":\"new-label1-value\"\n        \"label2\":\"new-label2-value\"\n      }\n    }\n\nNote that in the above example, any existing labels that are not\nincluded in the `updateMask` will be unaffected.\n\nIt is also possible to replace an entire map field by providing the\nmap field's path in the `updateMask`. The new value of the field will\nbe that which is provided in the patch environment. For example, to\ndelete all pre-existing user-specified PyPI packages and\ninstall botocore at version 1.7.14, the `updateMask` would contain\nthe path \"config.softwareConfig.pypiPackages\", and\nthe patch environment would be the following:\n\n    {\n      \"config\":{\n        \"softwareConfig\":{\n          \"pypiPackages\":{\n            \"botocore\":\"==1.7.14\"\n          }\n        }\n      }\n    }\n\n**Note:** Only the following fields can be updated:\n\n \u003ctable\u003e\n \u003ctbody\u003e\n \u003ctr\u003e\n \u003ctd\u003e\u003cstrong\u003eMask\u003c/strong\u003e\u003c/td\u003e\n \u003ctd\u003e\u003cstrong\u003ePurpose\u003c/strong\u003e\u003c/td\u003e\n \u003c/tr\u003e\n \u003ctr\u003e\n \u003ctd\u003econfig.softwareConfig.pypiPackages\n \u003c/td\u003e\n \u003ctd\u003eReplace all custom custom PyPI packages. If a replacement\n package map is not included in `environment`, all custom\n PyPI packages are cleared. It is an error to provide both this mask and a\n mask specifying an individual package.\u003c/td\u003e\n \u003c/tr\u003e\n \u003ctr\u003e\n \u003ctd\u003econfig.softwareConfig.pypiPackages.\u003cvar\u003epackagename\u003c/var\u003e\u003c/td\u003e\n \u003ctd\u003eUpdate the custom PyPI package \u003cvar\u003epackagename\u003c/var\u003e,\n preserving other packages. To delete the package, include it in\n `updateMask`, and omit the mapping for it in\n `environment.config.softwareConfig.pypiPackages`. It is an error\n to provide both a mask of this form and the\n \"config.softwareConfig.pypiPackages\" mask.\u003c/td\u003e\n \u003c/tr\u003e\n \u003ctr\u003e\n \u003ctd\u003elabels\u003c/td\u003e\n \u003ctd\u003eReplace all environment labels. If a replacement labels map is not\n included in `environment`, all labels are cleared. It is an error to\n provide both this mask and a mask specifying one or more individual\n labels.\u003c/td\u003e\n \u003c/tr\u003e\n \u003ctr\u003e\n \u003ctd\u003elabels.\u003cvar\u003elabelName\u003c/var\u003e\u003c/td\u003e\n \u003ctd\u003eSet the label named \u003cvar\u003elabelName\u003c/var\u003e, while preserving other\n labels. To delete the label, include it in `updateMask` and omit its\n mapping in `environment.labels`. It is an error to provide both a\n mask of this form and the \"labels\" mask.\u003c/td\u003e\n \u003c/tr\u003e\n \u003ctr\u003e\n \u003ctd\u003econfig.nodeCount\u003c/td\u003e\n \u003ctd\u003eHorizontally scale the number of nodes in the environment. An integer\n greater than or equal to 3 must be provided in the `config.nodeCount` field.\n \u003c/td\u003e\n \u003c/tr\u003e\n \u003ctr\u003e\n \u003ctd\u003econfig.softwareConfig.airflowConfigOverrides\u003c/td\u003e\n \u003ctd\u003eReplace all Apache Airflow config overrides. If a replacement config\n overrides map is not included in `environment`, all config overrides\n are cleared.\n It is an error to provide both this mask and a mask specifying one or\n more individual config overrides.\u003c/td\u003e\n \u003c/tr\u003e\n \u003ctr\u003e\n \u003ctd\u003econfig.softwareConfig.properties.\u003cvar\u003esection\u003c/var\u003e-\u003cvar\u003ename\n \u003c/var\u003e\u003c/td\u003e\n \u003ctd\u003eOverride the Apache Airflow property \u003cvar\u003ename\u003c/var\u003e in the section\n named \u003cvar\u003esection\u003c/var\u003e, preserving other properties. To delete the\n property override, include it in `updateMask` and omit its mapping\n in `environment.config.softwareConfig.properties`.\n It is an error to provide both a mask of this form and the\n \"config.softwareConfig.properties\" mask.\u003c/td\u003e\n \u003c/tr\u003e\n \u003ctr\u003e\n \u003ctd\u003econfig.softwareConfig.envVariables\u003c/td\u003e\n \u003ctd\u003eReplace all environment variables. If a replacement environment\n variable map is not included in `environment`, all custom environment\n variables  are cleared.\n It is an error to provide both this mask and a mask specifying one or\n more individual environment variables.\u003c/td\u003e\n \u003c/tr\u003e\n \u003c/tbody\u003e\n \u003c/table\u003e",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "request": {
	//     "$ref": "Environment"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "composer.projects.locations.operations.delete":

type ProjectsLocationsOperationsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a long-running operation. This method indicates that
// the client is
// no longer interested in the operation result. It does not cancel
// the
// operation. If the server doesn't support this method, it
// returns
// `google.rpc.Code.UNIMPLEMENTED`.
func (r *ProjectsLocationsOperationsService) Delete(name string) *ProjectsLocationsOperationsDeleteCall {
	c := &ProjectsLocationsOperationsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsOperationsDeleteCall) Fields(s ...googleapi.Field) *ProjectsLocationsOperationsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsOperationsDeleteCall) Context(ctx context.Context) *ProjectsLocationsOperationsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsOperationsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsOperationsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "composer.projects.locations.operations.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsOperationsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a long-running operation. This method indicates that the client is\nno longer interested in the operation result. It does not cancel the\noperation. If the server doesn't support this method, it returns\n`google.rpc.Code.UNIMPLEMENTED`.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/operations/{operationsId}",
	//   "httpMethod": "DELETE",
	//   "id": "composer.projects.locations.operations.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be deleted.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "composer.projects.locations.operations.get":

type ProjectsLocationsOperationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the latest state of a long-running operation.  Clients can
// use this
// method to poll the operation result at intervals as recommended by
// the API
// service.
func (r *ProjectsLocationsOperationsService) Get(name string) *ProjectsLocationsOperationsGetCall {
	c := &ProjectsLocationsOperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsOperationsGetCall) Fields(s ...googleapi.Field) *ProjectsLocationsOperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsOperationsGetCall) IfNoneMatch(entityTag string) *ProjectsLocationsOperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsOperationsGetCall) Context(ctx context.Context) *ProjectsLocationsOperationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsOperationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsOperationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "composer.projects.locations.operations.get" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsLocationsOperationsGetCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the latest state of a long-running operation.  Clients can use this\nmethod to poll the operation result at intervals as recommended by the API\nservice.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "composer.projects.locations.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "composer.projects.locations.operations.list":

type ProjectsLocationsOperationsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists operations that match the specified filter in the
// request. If the
// server doesn't support this method, it returns
// `UNIMPLEMENTED`.
//
// NOTE: the `name` binding allows API services to override the
// binding
// to use different resource name schemes, such as `users/*/operations`.
// To
// override the binding, API services can add a binding such
// as
// "/v1/{name=users/*}/operations" to their service configuration.
// For backwards compatibility, the default name includes the
// operations
// collection id, however overriding users must ensure the name
// binding
// is the parent resource, without the operations collection id.
func (r *ProjectsLocationsOperationsService) List(name string) *ProjectsLocationsOperationsListCall {
	c := &ProjectsLocationsOperationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": The standard list
// filter.
func (c *ProjectsLocationsOperationsListCall) Filter(filter string) *ProjectsLocationsOperationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The standard list
// page size.
func (c *ProjectsLocationsOperationsListCall) PageSize(pageSize int64) *ProjectsLocationsOperationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The standard list
// page token.
func (c *ProjectsLocationsOperationsListCall) PageToken(pageToken string) *ProjectsLocationsOperationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsOperationsListCall) Fields(s ...googleapi.Field) *ProjectsLocationsOperationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsOperationsListCall) IfNoneMatch(entityTag string) *ProjectsLocationsOperationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsOperationsListCall) Context(ctx context.Context) *ProjectsLocationsOperationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsOperationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsOperationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}/operations")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "composer.projects.locations.operations.list" call.
// Exactly one of *ListOperationsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListOperationsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsOperationsListCall) Do(opts ...googleapi.CallOption) (*ListOperationsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListOperationsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists operations that match the specified filter in the request. If the\nserver doesn't support this method, it returns `UNIMPLEMENTED`.\n\nNOTE: the `name` binding allows API services to override the binding\nto use different resource name schemes, such as `users/*/operations`. To\noverride the binding, API services can add a binding such as\n`\"/v1/{name=users/*}/operations\"` to their service configuration.\nFor backwards compatibility, the default name includes the operations\ncollection id, however overriding users must ensure the name binding\nis the parent resource, without the operations collection id.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/operations",
	//   "httpMethod": "GET",
	//   "id": "composer.projects.locations.operations.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The standard list filter.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name of the operation's parent resource.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The standard list page size.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The standard list page token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}/operations",
	//   "response": {
	//     "$ref": "ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsLocationsOperationsListCall) Pages(ctx context.Context, f func(*ListOperationsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}