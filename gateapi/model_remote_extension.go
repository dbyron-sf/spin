/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RemoteExtension struct {
	Transport *RemoteExtensionTransport `json:"transport,omitempty"`
	Type_ string `json:"type,omitempty"`
	Config *interface{} `json:"config,omitempty"`
	Id string `json:"id,omitempty"`
}