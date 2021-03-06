package models

import "github.com/go-openapi/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Complete complete

swagger:model Complete
*/
type Complete struct {

	/* Time when task was completed. Always in UTC.
	 */
	CompletedAt strfmt.DateTime `json:"completed_at,omitempty"`

	/* Error message, if status=error. Only used by the /error endpoint.
	 */
	Error string `json:"error,omitempty"`

	/* Machine readable reason failure, if status=error. Only used by the /error endpoint.
	 */
	Reason string `json:"reason,omitempty"`
}

// Validate validates this complete
func (m *Complete) Validate(formats strfmt.Registry) error {
	return nil
}
