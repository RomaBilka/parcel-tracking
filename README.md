# Parcel Tracking

Open-source program that developed to simplify parcel tracking.
You can find you parcel from UPS, Ukrposhta, USPS, Nova Poshta, MeestExpress, Amazaon and much others just by calling
the single api endpoint:

```bash
/tracking?track_id=[Your parcel number]
```
## Carriers
* [DHL](./pkg/determine-delivery/carriers/dhl)
* [FedEx](./pkg/determine-delivery/carriers/fedex)
* [Mess Express](./pkg/determine-delivery/carriers/me)
* [Nova Poshte](./pkg/determine-delivery/carriers/np)
* [Nova Poshte Sgopping](./pkg/determine-delivery/carriers/np-shopping)
* [UPS](./pkg/determine-delivery/carriers/ups)
* [USPS](./pkg/determine-delivery/carriers/usps)

## Config
[Config](./dependencies)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.

## License
[MIT](LICENSE.md)
