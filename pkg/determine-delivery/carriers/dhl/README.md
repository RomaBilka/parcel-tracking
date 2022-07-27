# DHL
[Documentation](https://developer.dhl.com/api-reference/shipment-tracking#get-started-section/overview)

## Tracking number patterns
[Patterns](https://www.dhl.com/pf-en/home/tracking/id-labels.html)

|         Type        |                          Value                          |            Option           |            Service            |              Response             |
|:-------------------:|:-------------------------------------------------------:|:---------------------------:|:-----------------------------:|:---------------------------------:|
| Tracking            | 1Z12345E0205271688 (Signature Availability)             | Activity (All)              | 2nd Day Air                   | Delivered                         |
| Tracking            | 1Z12345E6605272234                                      | None (Last) World Wide      | Express                       | Delivered                         |
| Shipping            | 1Z12345E0305271640 (Second Package: 1Z12345E0393657226) | None (Last)                 | Ground                        | Delivered                         |
| Tracking            | 1Z12345E1305277940                                      | None (Last)                 | Next Day Air Saver            | ORIGIN SCAN                       |
| Tracking            | 1Z12345E6205277936                                      | Activity (All) Next Day Air | Saver                         | 2nd Delivery attempt              |
| Tracking            | 1Z12345E020527079                                       | None (Last)                 |                               | Invalid Tracking Number           |
| Tracking            | 1Z12345E1505270452                                      | None (Last)                 |                               | No Tracking Information Available |
| Tracking            | 990728071                                               | Activity (All)              | UPS Freight LTL               | In Transit                        |
| Tracking            | 3251026119                                              | Activity (All)              |                               | Delivered Origin CFS              |
| MI Tracking Number  | 9102084383041101186729                                  | None (Last)                 |                               |                                   |
| MI Reference Number | cgish000116630                                          | None (Last)                 |                               |                                   |
| Tracking            | 1Z648616E192760718                                      | Activity                    | UPS Worldwide Express Freight | Order Process by UPS              |
| Tracking            | 5548789114                                              | Activity                    | UPS Express Freight           | Response for UPS Air Freight      |
| Tracking            | ER751105042015062                                       | Activity                    | UPS Ocean                     | Response for UPS Ocean Freight    |
| Tracking            | 1ZWX0692YP40636269                                      | Activity                    | UPS SUREPOST                  | Response for UPS SUREPOST         |

## Scope
This API covers services provided by DPDHL under these brand names:

Post & Parcel Germany (incl. mail / letter tracking)
DHL Global Forwarding (incl. DHL Same Day)
DHL Freight
DHL Express
DHL Supply Chain
DHL eCommerce Solutions
Asia-Pacific
US, Canada
EU (Belgium, Luxemburg, Netherlands, Poland, Portugal, Spain, United Kingdom)
It does not cover:

Services that require a login (eg. B2B systems). These are excluded as a backend.
The API provides users with tracking information on:

Shipment location
Shipment delivery time*
Shipment travel history*
The Proof of Delivery*
Shipment timestamp
Origin and destination information
Shipment number of pieces*
Shipment piece level events*
Shipment dimensions*
Shipment weight*
Not available for all DPDHL services.

## Authentication
Every call to the API requires a subscription key. This key needs to be either passed through a query string parameter or specified in the request header (DHL-API-Key).

## API URL
``https://api-eu.dhl.com/track/shipments``

## Rate limits
Rate limits protect the DHL infrastructure from suspicious requests that exceed defined thresholds.

When you first request access to the Shipment Tracking - Unified API, you will get the initial service level which allows 250 calls per day with a maximum of 1 call per second.

Additional rate limits are available and they are granted according to your specific use case. If you would like to request for additional limits, please proceed with the following steps:

* Create an app as described under Get Access section.
* Click My Apps on the portal website.
* Click on the App you created
* Scroll down to the APIs list and click on the "Request Upgrade" button.

## Request
| Name                                 | Description                                                                                                                                                                                                                                                                  |
|--------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| trackingNumber * string  (query)     | The tracking number of the shipment for which to return the information.                                                                                                                                                                                                     |
| service string  (query)              | Hint which service (provider) should be used to resolve the tracking number. Available   values : express, parcel-de, ecommerce, dgf, parcel-uk, post-de, sameday, freight, parcel-nl, parcel-pl, dsc--                                                                      |
| requesterCountryCode string  (query) | Optional ISO 3166-1 alpha-2 country code represents country of the consumer of the API response. It optimizes the return of the API response.                                                                                                                                |
| originCountryCode string  (query)    | Optional ISO 3166-1 alpha-2 country code of the shipment origin to further qualify the shipment tracking number (trackingNumber) parameter of the request. This parameter is necessary to search for the shipment in dsc service.                                            |
| recipientPostalCode string  (query)  | Postal code of the destination address to further qualify the shipment tracking number (trackingNumber) parameter of the request or parcel-nl and parcel-de services to display full set of data in the response.                                                            |
| language string  (query)             | ISO 639-1 2-character language code for the response. This parameter serves as an indication of the client preferences ONLY. Language availability depends on the service used. The actual response language is indicated by the Content-Language header. Default value : en |
| offset number  (query)               | Pagination parameter. Offset from the start of the result set at which to retrieve the remainder of the results (if any). Default value : 0                                                                                                                                  |
| limit number  (query)                | Pagination parameter. Maximal number of results to retireve. Default value : 5                                                                                                                                                                                               |

## Response

[response](/fixtures)

