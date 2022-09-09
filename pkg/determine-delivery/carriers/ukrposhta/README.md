# Ukrposhta
[Documentation](https://dev.ukrposhta.ua/uploads/Status-tracking-API-05042022.pdf)

## Tracking number patterns
13 (1234567890123) digits or starts with 2 letters 9 numbers and ends with 2 letters (AA123456789AA)

## Getting all statuses by list of barcodes not found in system
Brief description. If the query contains the barcodes that are not registered in the
system, it returns the data only on the registered shipments, followed by list of
barcodes that were not found in system.
URI:/statuses/with-not-found

