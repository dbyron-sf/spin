/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Http struct {
	Url string `json:"url,omitempty"`
	QueryParams map[string]string `json:"queryParams,omitempty"`
	Config map[string]string `json:"config,omitempty"`
	Headers *Headers `json:"headers,omitempty"`
}
