# USPS tracking info

[Here](https://www.usps.com/business/web-tools-apis/) you can find all information for developers presented by USPS

## How to work with the USPS API 
1. You need to be registered for the U.S.Postal Service's Web Tools Application Programming Interfaces (APIs). ***1**
   - As a result, you will have a UserID (Username) and Password.
2. You need to be registered as a USPS user for Mailer ID (MID).
3. To obtain Package Tracking API (API=TrackV2) access, users will need to contact the [USPS Web Tools Program Office](https://usps.force.com/emailus/s/web-tools-inquiry) to request access and supply additional information for customer verification. ***2**

Step-by-Step Instructions for All USPS Web Tools ([HTM](https://www.usps.com/business/web-tools-apis/general-api-developer-guide.htm) | [PDF](https://www.usps.com/business/web-tools-apis/general-api-developer-guide.pdf))

#### Error Handling

When an error condition exists, a specific XML return is generated. The following example shows the tags that are returned:

```
<?xml version="1.0"?> 
<Error>
    <Number>-2147217951</Number>
    <Source>EMI_Respond :EMI:clsEMI.ValidateParameters: 
        clsEMI.ProcessRequest;SOLServerIntl.EMI_Respond</Source>
    <Description>Missing value for To Phone number.</Description>
    <HelpFile></HelpFile>
    <HelpContext>1000440</HelpContext>
</Error>
```
#### Track testing
###### There is no capacity for testing in the USPS Web Tools infrastructure. Any account performing capacity/stress testing may be terminated.

This **track** test shows a multi-entry return that is arranged in reverse chronological order.
````
http://production.shippingapis.com/ShippingAPI.dll?API=TrackV2

REQUEST
//the "password" tag may be optional
&XML=<TrackRequest USERID="xxxxxxxx" PASSWORS="xxxxx"> 
    <TrackID ID="EJ958083578US"></TrackID>
</TrackRequest>

RESPONSE
<?xml version="1.0"?>
<TrackResponse>
    <TrackInfo ID="EJ958083578US">
        <TrackSummary>
Your item was delivered at 8:10 am on June 1 in Wilmington DE 19801.
</TrackSummary>
        <TrackDetail>
May 30 11:07 am NOTICE LEFT WILMINGTON DE 19801.
</TrackDetail>
        <TrackDetail>
May 30 10:08 am ARRIVAL AT UNIT WILMINGTON DE 19850.
</TrackDetail>
        <TrackDetail>
May 29 9:55 am ACCEPT OR PICKUP EDGEWATER NJ 07020.
</TrackDetail>
    </TrackInfo>
</TrackResponse>
````

###### *1 - всі секрети покладу в наш тімс чат.

###### *2 - вже робивсі запит і досі не відповідають(пінгалося 2 рази)