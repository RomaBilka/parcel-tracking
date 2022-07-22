Rate Limits

FedEx assigns one API key to each project which is used for all APIs in the project.The rate limit for each API key(and project) is 750 transactions per 10 seconds. Review the FedEx API Throttling Guide for more information, best practices and upgrade options.

API Certification

Select APIs must be certified by FedEx before you can move to production. The level of  certification can vary by API. See the certification requirements for the APIs you selected below. Note that Shipping Label certification can take up to 5 days.


## API Authorization

Once you have secured the API credentials on FedEx Developer portal, use this endpoint to get an access token to use as credentials with each API transaction.

Following are the required input information associated with this request:

* grant_type – Type of customer. (Valid values: client_credentials, csp_credentials)
* client_id – Refers to the Project API Key.
* client_secret – Refers to the Project API Secret Key.
For FedEx®Internal or Compatible customers, send the below additional inputs:

* child_id – Customer Key returned through Credential Registration API request.
* child_secret – Customer password returned through Credential Registration API request
The result of this request should return below:

* access_token – The encrypted OAuth token that needs to be used in the API transaction.
* token_type – Type of token. In this case, it is bearer authentication.
* expires_in – Token expiration time in milliseconds. One hour is the standard Token expiration time.
* Scope – Scope of authorization provided to the consumer.