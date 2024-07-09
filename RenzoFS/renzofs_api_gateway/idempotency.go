/**
*	@author Elia Renzoni
*	@date 29/05/2024
*	@brief Module that ensure idempotent api gateway.
*
 */

package renzofsapigateway

import "net/http"

type Idempotency struct {
	requests map[]http.Request
}
