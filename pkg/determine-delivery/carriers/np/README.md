# Novaposhta
[Documentation](https://developers.novaposhta.ua/documentation)

## Tracking number patterns
14 digits start 59.., 20.., 1..

## Трекінг
Оновлений метод «getStatusDocuments» працює в моделі «TrackingDocument», цей метод дозволяє переглядати більш розширену інформацію щодо статусу відправлення.
При введеному номері телефону можна отримати наступні відомості: дані відправника або одержувача, номер телефону.
**Метод дозволяє переглядати одночасно до 100 відправлень.**
Доступність: Не вимагає використання API-ключа.
Актуальні статуси трекінгу

| Код  |  Статус                                                                                                                                                                     |
|------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  1   | Відправник самостійно створив цю накладну, але ще не надав до відправки                                                                                                     |
|  2   | Видалено                                                                                                                                                                    |
|  3   | Номер не знайдено                                                                                                                                                           |
|  4   | Відправлення у місті ХХXХ. (Статус для межобластных отправлений)                                                                                                            |
|  41  | Відправлення у місті ХХXХ. (Статус для услуг локал стандарт и локал экспресс - доставка в пределах города)                                                                  |
|  5   | Відправлення прямує до міста YYYY                                                                                                                                           |
|  6   | Відправлення у місті YYYY, орієнтовна доставка до ВІДДІЛЕННЯ-XXX dd-mm. Очікуйте додаткове повідомлення про прибуття                                                        |
|  7   | Прибув на відділення                                                                                                                                                        |
|  8   | Прибув на відділення (завантажено в Поштомат)                                                                                                                               |
|  9   | Відправлення отримано                                                                                                                                                       |
|  10  | Відправлення отримано %DateReceived%. Протягом доби ви одержите SMS-повідомлення про надходження грошового переказу та зможете отримати його в касі відділення «Нова пошта» |
|  11  | Відправлення отримано %DateReceived%. Грошовий переказ видано одержувачу.                                                                                                   |
|  12  | Нова Пошта комплектує ваше відправлення                                                                                                                                     |
|  101 | На шляху до одержувача                                                                                                                                                      |
|  102 | Відмова одержувача (створено заявку на повернення)                                                                                                                          |
|  103 | Відмова одержувача (отримувач відмовився від відправлення)                                                                                                                  |
|  104 | Змінено адресу                                                                                                                                                              |
|  105 | Припинено зберігання                                                                                                                                                        |
|  106 | Одержано і створено ЄН зворотньої доставки                                                                                                                                  |
|  111 | Невдала спроба доставки через відсутність Одержувача на адресі або зв'язку з ним                                                                                            |
|  112 | Дата доставки перенесена Одержувачем                                                                                                                                        |


## Request
```javascript
{
    "apiKey":"[ВАШ КЛЮЧ]",
        "modelName":"TrackingDocument",
        "calledMethod":"getStatusDocuments",
        "methodProperties":{
        "Documents":[
            {
                "DocumentNumber":"20400048799000",
                "Phone":"380600000000"
            },
            {
                "DocumentNumber":"20400048799001",
                "Phone":"380600000000"
            }
        ]
    }
}
```

## Response

[response](/fixtures)

| Параметр                             | Тип        | Опис                                                                                                                                                                               |
|--------------------------------------|------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| PossibilityCreateReturn              | bool       | Доступність замовлення послуги повернення вантажу, допустимі значення true/false: true - послуга доступна до замовлення; false - послуга не доступна до замовлення.                |
| PossibilityCreateRefusal             | bool       | Доступність замовлення послуги відмови від вантажу, допустимі значення true/false: true - послуга доступна до замовлення; false - послуга не доступна до замовлення.               |
| PossibilityChangeEW                  | bool       | Доступність замовлення послуги внесення змін в ЕН, допустимі значення true/false: true - послуга доступна до замовлення; false - послуга не доступна до замовлення.                |
| PossibilityCreateRedirecting         | bool       | Доступність замовлення послуги переадресування, допустимі значення true/false: true - послуга доступна до замовлення; false - послуга не доступна до замовлення.                   |
| Number                               | string[36] | Номер документу                                                                                                                                                                    |
| Redelivery                           | string[36] | Ідентифікатор зворотної доставки. 0 - відсутня ЗД, 1 - наявна ЗД                                                                                                                   |
| RedeliverySum                        | string[36] | Сума зворотної доставки                                                                                                                                                            |
| RedeliveryNum                        | string[36] | Ідентифікатор номеру ЕН воротної доставки (якщо таку ЕН створено)                                                                                                                  |
| RedeliveryPayer                      | string[36] | Ідентифікатор платника зворотної доставки (якщо є зворотна доставка)                                                                                                               |
| OwnerDocumentType                    | string[36] | Тип ЕН на підставі                                                                                                                                                                 |
| LastCreatedOnTheBasisDocumentType    | string[36] | Останні зміни типу документу                                                                                                                                                       |
| LastCreatedOnTheBasisPayerType       | string[36] | Останні зміни, тип платника                                                                                                                                                        |
| LastCreatedOnTheBasisDateTime        | string[36] | Останні зміни, дата створення                                                                                                                                                      |
| LastTransactionStatusGM              | string[36] | Останній статус транзакції грошового переказу                                                                                                                                      |
| LastTransactionDateTimeGM            | string[36] | Останній час та дата транзакції грошового переказу                                                                                                                                 |
| LastAmountTransferGM                 | string[36] | Поточне значення суми грошового переказу                                                                                                                                           |
| DateCreated                          | string[36] | Дата створення ЕН                                                                                                                                                                  |
| DocumentWeight                       | string[36] | Вага                                                                                                                                                                               |
| FactualWeight                        | string[36] | Фактична вага з ЕН                                                                                                                                                                 |
| VolumeWeight                         | string[36] | Об'ємна вага з ЕН                                                                                                                                                                  |
| CheckWeight                          | string[36] | Інформація після контрольного зважування                                                                                                                                           |
| CheckWeightMethod                    | string[36] | Тип зважування, згідно якого виконувалось контрольне зважування                                                                                                                    |
| DocumentCost                         | string[36] | Вартість доставки                                                                                                                                                                  |
| CalculatedWeight                     | string[36] | Розрахункова вага                                                                                                                                                                  |
| SumBeforeCheckWeight                 | string[36] | Інформація до контрольного зважування                                                                                                                                              |
| PayerType                            | string[36] | Ідентифікатор платника                                                                                                                                                             |
| RecipientFullName                    | string[36] | Дані ПІБ отримувача з накладної, за умови зазначення номеру телефону                                                                                                               |
| RecipientDateTime                    | string[36] | Дата, коли отримувач забрав вантаж                                                                                                                                                 |
| ScheduledDeliveryDate                | string[36] | Очікувана дата доставки                                                                                                                                                            |
| PaymentMethod                        | string[36] | Тип оплати                                                                                                                                                                         |
| CargoDescriptionString               | string[36] | Опис вантажу                                                                                                                                                                       |
| CargoType                            | string[36] | Тип вантажу                                                                                                                                                                        |
| CitySender                           | string[36] | Місто відправника                                                                                                                                                                  |
| CityRecipient                        | string[36] | Місто отримувача                                                                                                                                                                   |
| WarehouseRecipient                   | string[36] | Склад отримувача                                                                                                                                                                   |
| CounterpartyType                     | string[36] | Тип контрагенту                                                                                                                                                                    |
| AfterpaymentOnGoodsCost              | string[36] | Контроль оплати                                                                                                                                                                    |
| ServiceType                          | string[36] | Тип доставки                                                                                                                                                                       |
| UndeliveryReasonsSubtypeDescription  | string[36] | Опис причини нерозвезення                                                                                                                                                          |
| WarehouseRecipientNumber             | string[36] | Номер відділення отримувача                                                                                                                                                        |
| LastCreatedOnTheBasisNumber          | string[36] | Останні зміни, номер ЕН                                                                                                                                                            |
| PhoneRecipient                       | string[36] | Номер телефону отримувача                                                                                                                                                          |
| RecipientFullNameEW                  | string[36] | Дані ПІБ отримувача з накладної, за умови зазначення в запиті номеру телефону                                                                                                      |
| WarehouseRecipientInternetAddressRef | string[36] | Ідентифікатор (REF) складу отримувача                                                                                                                                              |
| MarketplacePartnerToken              | string[36] | Токен торгівельного майданчику                                                                                                                                                     |
| ClientBarcode                        | string[36] | Внутрішній номер замовлення                                                                                                                                                        |
| RecipientAddress                     | string[36] | Адреса отримувача                                                                                                                                                                  |
| CounterpartyRecipientDescription     | string[36] | Опис контрагента отримувача                                                                                                                                                        |
| CounterpartySenderType               | string[36] | Тип контрагента відправника                                                                                                                                                        |
| DateScan                             | string[36] | Дата сканування, що призвела до формування статусу                                                                                                                                 |
| PaymentStatus                        | string[36] | Статус для інтернет еквайрингу                                                                                                                                                     |
| PaymentStatusDate                    | string[36] | Дата оплати для інтернет еквайрингу                                                                                                                                                |
| AmountToPay                          | string[36] | Сума до сплати для інтернет еквайрингу                                                                                                                                             |
| AmountPaid                           | string[36] | Сплачено для інтернет еквайрингу                                                                                                                                                   |
| Status                               | string[36] | Статус                                                                                                                                                                             |
| StatusCode                           | string[36] | Ідентифікатор статусу                                                                                                                                                              |
| RefEW                                | string[36] | Ідентифікатор накладної для інтернет еквайрингу (використовується в робочих цілях)                                                                                                 |
| BackwardDeliverySubTypesActions      | string[36] | Інформація за нестандартними підтипами зворотної доставки                                                                                                                          |
| BackwardDeliverySubTypesServices     | string[36] | Інформація за нестандартними підтипами зворотної доставки                                                                                                                          |
| UndeliveryReasons                    | string[36] | Причина нерозвозу                                                                                                                                                                  |
| DatePayedKeeping                     | string[36] | Дата початку платного зберігання                                                                                                                                                   |
| InternationalDeliveryType            | string[36] | Тип міжнародної доставки                                                                                                                                                           |
| SeatsAmount                          | string[36] | Кількість місць                                                                                                                                                                    |
| CardMaskedNumber                     | string[36] | Замаскований номер платіжної карти                                                                                                                                                 |
| ExpressWaybillPaymentStatus          | string[36] | Статус оплати ЕН                                                                                                                                                                   |
| ExpressWaybillAmountToPay            | string[36] | Сума оплати по ЕН                                                                                                                                                                  |
| PhoneSender                          | string[36] | Телефон відправника                                                                                                                                                                |
| TrackingUpdateDate                   | string[36] | Оновлена дата відстеження                                                                                                                                                          |
| WarehouseSender                      | string[36] | Відділення відправника                                                                                                                                                             |
| DateReturnCargo                      | string[36] | Дата повернення вантажу                                                                                                                                                            |
| DateMoving                           | string[36] | Дата переміщення                                                                                                                                                                   |
| DateFirstDayStorage                  | string[36] | Дата початку зберігання                                                                                                                                                            |
| RefCityRecipient                     | string[36] | Ідентифікатор міста одержувача                                                                                                                                                     |
| RefCitySender                        | string[36] | Ідентифікатор міста відправника                                                                                                                                                    |
| RefSettlementRecipient               | string[36] | Ідентифікатор населеного пункту одержувача                                                                                                                                         |
| RefSettlementSender                  | string[36] | Ідентифікатор населеного пункту відправника                                                                                                                                        |
| SenderAddress                        | string[36] | Адреса відправника                                                                                                                                                                 |
| SenderFullNameEW                     | string[36] | Повне ім'я відправника                                                                                                                                                             |
| AnnouncedPrice                       | string[36] | Оголошена вартість                                                                                                                                                                 |
| AdditionalInformationEW              | string[36] | Додаткова інформація                                                                                                                                                               |
| ActualDeliveryDate                   | string[36] | Фактична дата доставки                                                                                                                                                             |
| PostomatV3CellReservationNumber      | string[36] | Номер бронювання комірки поштомату, за умови доставки до поштомату                                                                                                                 |
| OwnerDocumentNumber                  | string[36] | Номер ЕН на основі                                                                                                                                                                 |
| LastAmountReceivedCommissionGM       | string[36] | Сума комісії за грошовий переказ                                                                                                                                                   |
| DeliveryTimeframe                    | string[36] | Часовий інтервал, якщо було замовлено та тільки для ЕН до адреси                                                                                                                   |
| CreatedOnTheBasis                    | string[36] | Створено на основі                                                                                                                                                                 |
| UndeliveryReasonsDate                | string[36] | Дата причини нерозвезення                                                                                                                                                          |
| RecipientWarehouseTypeRef            | string[36] | Ідентифікатор типу відділення одержувача                                                                                                                                           |
| WarehouseRecipientRef                | string[36] | Ідентифікатор відділення одержувача                                                                                                                                                |
| CategoryOfWarehouse                  | string[36] | Категорія відділення                                                                                                                                                               |
| WarehouseRecipientAddress            | string[36] | Адреса відділення одержувача                                                                                                                                                       |
| WarehouseSenderInternetAddressRef    | string[36] | Інтернет адреса відділення відправника                                                                                                                                             |
| WarehouseSenderAddress               | string[36] | Адреса відділення відправника                                                                                                                                                      |
| CounterpartySenderType               | string[36] | Тип контрагенту відправника                                                                                                                                                        |
| AviaDelivery                         | string[36] | Авіа доставка                                                                                                                                                                      |
| BarcodeRedBox                        | string[36] | ШК пакування типу RedBox                                                                                                                                                           |
| CargoReturnRefusal                   | string[36] | Наявність послуги "Відмова від повернення", допустимі значення true - послугу замовлено, а false - послугу не замовлено                                                            |
| DaysStorageCargo                     | string[36] | День зберігання вантажу                                                                                                                                                            |
| Packaging                            | string[36] | Пакування                                                                                                                                                                          |
| PartialReturnGoods                   | string[36] | Часткове повернення                                                                                                                                                                |
| SecurePayment                        | string[36] | Ознака надійної покупки, допустимі значення true/false: true - послугу замовлено; false - послугу не замовлено.                                                                    |
| PossibilityChangeCash2Card           | bool       | Доступність можливості зміни виплати грошового переказу на карту, допустимі значення true/false: true - послуга доступна до замовлення; false - послуга не доступна до замовлення. |
| PossibilityChangeDeliveryIntervals   | bool       | Дуступність зміни інтервалу доставки, допустимі значення true/false: true - послуга доступна до замовлення; false - послуга не доступна до замовлення.                             |
| PossibilityTermExtensio              | bool       | Дуступність можливості подовження термінів зберігання, допустимі значення true/false: true - послуга доступна до замовлення; false - послуга не доступна до замовлення.            |
| StorageAmount                        | string[36] | Кількість днів зберігання вантажу                                                                                                                                                  |
| StoragePrice                         | string[36] | Вартість зберігання                                                                                                                                                                |
| FreeShipping                         | string[36] | Ознака безкоштовної доставки                                                                                                                                                       |
| LoyaltyCardRecipient                 | string[36] | Номер карти лояльності отримувача, за наявності                                                                                                                                    |


## Error codes

[Error codes](https://developers.novaposhta.ua/listerrorscodes?page=1&quality=3856)
