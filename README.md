# Parcel Tracking

Open-source program that developed to simplify parcel tracking.
You can find you parcel from UPS, Ukrposhta, USPS, Nova Poshta, MeestExpress and much others just by calling
the single api endpoint:

You must use the ```POST``` method, and JSON array ```track_id```
```bash
curl -X POST /tracking
   -H 'Content-Type: application/json'
   -d '{"track_id":["******************", "******************"]}'
```
## Carriers
* [DHL](./pkg/determine-delivery/carriers/dhl)
* [FedEx](./pkg/determine-delivery/carriers/fedex)
* [Mess Express](./pkg/determine-delivery/carriers/me)
* [Nova Poshte](./pkg/determine-delivery/carriers/np)
* [Nova Poshte Sgopping](./pkg/determine-delivery/carriers/np-shopping)
* [Ukrposhta](./pkg/determine-delivery/carriers/ukrposhta)
* [UPS](./pkg/determine-delivery/carriers/ups)
* [USPS](./pkg/determine-delivery/carriers/usps)

## Config
[Config](./dependencies)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.

## License
[MIT](LICENSE.md)
