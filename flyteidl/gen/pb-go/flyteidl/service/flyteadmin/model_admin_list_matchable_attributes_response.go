/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Response for a request for all matching resource attributes.
type AdminListMatchableAttributesResponse struct {
	Configurations []AdminMatchableAttributesConfiguration `json:"configurations,omitempty"`
}
