# Meest
[Documentation](https://wiki.meest-group.com/uk/4-vidstezhennia-vidpravlen)

## Tracking number patterns
CV123456789US, MYCV123456789PL, MYCV123456789DE

## API URL
``https://apii.meest-group.com/T/1C_Query.php``

## Requesr
| Поле      | Опис                                                                                                                 |
|-----------|----------------------------------------------------------------------------------------------------------------------|
| login     | Ім’я користувача, присвоюєтьс я після внесення Контрагента в систему                                                 |
| function  | Назва функції, використовуємо для введення запису.                                                                   |
| where     | Умови пошуку: AgentUID - унікальний ідентифікатор агента (отримується при реєстрації) Parcel number - № відправлення |
| order     | Сортування записів.   Якщо не було внесено поля і умови сортування, результат - у довільному порядку.                |
| sign      |  = md5 ( login + password + function + where + order)                                                                |

## Response
[response](/fixtures)

## Response description
| Поле (Field)          | Опис (Description)                                                                                                                                  | Тип (Type) |
|-----------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|------------|
| ShipmentIDRef         | Ідентифікатор відправлення (в системі Росан) Shipment’s ID (in Rosan system)                                                                        | Text       |
| DocumentIDRef         | Ідентифікатор документа (в системі Росан), який генерує статус Document’s ID (in Rosan system), which generates this status                         | Text       |
| AgentsIdRef           | Унікальний ідентифікатор агента Unique Agent’s ID                                                                                                   | Text       |
| ShipmentNumberSender  | Номер відправлення, номер для трекінгу Shipment number, track number                                                                                | Text       |
| ShipmentNumberTransit | Transit number Транзитний номер                                                                                                                     | Text       |
| ShipmentClientID      | Ідентифікатор клієнта відправника Sender’s client ID                                                                                                | Text       |
| DateTimeAction        | Дата встановелння статусу Status datetime                                                                                                           | dateTime   |
| Country               | Країна встановлення статусу Country where this status was set                                                                                       | Text       |
| City                  | Місто встановлення статусу City where this status was set                                                                                           | Text       |
| ActionId              | Ідентифікатор статусу Status ID                                                                                                                     | Text       |
| StatusCode            | Код статусу Status Code                                                                                                                             | Text       |
| ActionMessages_UA     | Назва статусу українською мовою Status name in ukrainian                                                                                            | Text       |
| ActionMessages_RU     | Назва статусу російською мовою Status name in russian                                                                                               | Text       |
| ActionMessages_EN     | Назва статусу англійською мовою Status name in english                                                                                              | Text       |
| DetailMessages_UA     | Інформація про статус українською мовою: місто, країна, дата тощо Status details in ukrainian: city,country,date etc                                | Text       |
| DetailMessages_RU     | Інформація про статус російською мовою: місто, країна, дата тощо Status details in russian: city,country,date etc                                   | Text       |
| DetailMessages_EN     | Інформація про статус англійською мовою: місто, країна, дата тощо Status details in english: city,country,date etc                                  | Text       |
| DetailPlacesAction    | Кількість отриманих місць  / Кількість відправлених місць Quantity places received / Quantuty places sent                                           | Text       |
| CountryDel            | Країна, в яку буде доставлена посилка Country to which parcel is to be delivered                                                                    | Text       |
| Recipient_Country     | Параметр у розробці Parameter under construction                                                                                                    | Text       |
| Deliverydate          | Дата доставки посилки одержувачу (працює тільки коли StatusCode = 1622) Date of delivery of parcel to recipient (works only when StatusCode = 1622) | dateTime   |

## Errors Codes
| Код помилки ErrorsCode |            Опис            |        Description       |
|:----------------------:|:--------------------------:|:------------------------:|
|           000          | Успішне виконання          | ОК                       |
|           100          | Помилка з’єднання          | Connection Error         |
|           101          | Помилка авторизації        | Аuthorization Error      |
|           102          | Функція не знайдена        | Function is not found    |
|           103          | Документ не знайдено       | Document not found       |
|           104          | Каталог не знайдено        | Directory not found      |
|           105          | Не вдалося обробити запит  | Failed to parse request  |
|           106          | Внутрішня помилка 1С       | Internal error 1C        |
|           107          | Внутрішня помилка          | Internal error           |
|           108          | Помилка запиту             | Error request            |
|           109          | Помилка структури XML      | Error XML structure      |
|           110          | З'єднання розірвано        | Disconnected             |
|           111          | Запит обробляється         | The request is processed |
|           113          | Люба помилка при створенні | Any error                |

## Status Codes
| Код   | Назва UA                                                                | Назва EN                                                                      | Міжнародний поштовий код |
|-------|-------------------------------------------------------------------------|-------------------------------------------------------------------------------|--------------------------|
| 101   | ОФОРМЛЕНО ДЛЯ ВІДПРАВКИ                                                 | POSTING / COLLECTION                                                          | EMA                      |
| 202   | НАДХОДЖЕННЯ ДО МІЖНАРОДНОГО СОРТУВАЛЬНОГО ЦЕНТРУ                        | ARRIVAL AT OUTWARD OFFICE OF EXCHANGE                                         | EMB                      |
| 303   | ВІДПРАВКА З МІЖНАРОДНОГО СОРТУВАЛЬНОГО ЦЕНТРУ                           | DEPARTURE FROM OUTWARD OFFICE OF EXCHANGE                                     | EMC                      |
| 404   | НАДХОДЖЕННЯ ДО МІЖНАРОДНОГО СОРТУВАЛЬНОГО ЦЕНТРУ                        | ARRIVAL AT TRANSIT OFFICE OF EXCHANGE                                         | EMJ                      |
| 505   | НАДХОДЖЕННЯ ДО МІЖНАРОДНОГО СОРТУВАЛЬНОГО ЦЕНТРУ                        | ARRIVAL AT INWARD OFFICE OF EXCHANGE                                          | EMD                      |
| 606   | НАДХОДЖЕННЯ НА ПІДРОЗДІЛ ВІДПРАВНИКА                                    | ARRIVAL AT SENDER OFFICE                                                      |                          |
| 808   | ВІДПРАВКА З МІЖНАРОДНОГО СОРТУВАЛЬНОГО ЦЕНТРУ У НАПРЯМКУ МІСТА ДОСТАВКИ | DEPARTURE FROM INWARD OFFICE OF EXCHANGE IN THE DIRECTION OF DESTINATION CITY | EMF                      |
| 1 214 | НАДХОДЖЕННЯ НА ПІДРОЗДІЛ ДОРУЧЕННЯ                                      | ARRIVAL AT DELIVERY OFFICE                                                    | EMG                      |
| 1 315 | ВИДАНО КУР`ЄРУ                                                          | HANDED OVER TO THE COURIER                                                    |                          |
| 1 521 | ЗМІНА НАПРЯМКУ ПОСИЛКИ                                                  | CHANGE OF DIRECTION                                                           |                          |
| 1 622 | ДОРУЧЕНО                                                                | FINAL DELIVERY                                                                | EMI                      |
| 1 620 | ВІДМОВА                                                                 | DELIVERY REFUSED                                                              |  EMH                     |
| 1 621 | ПЕРЕАДРЕСАЦІЯ                                                           | RDELIVERY REDIRECTED                                                          |  EMH                     |
| 1 624 | НЕ КОРЕКТНІ ДАНІ ПРО ДОРУЧЕННЯ                                          | UNABLE TO VERIFY DELIVERY INFORMATION                                         |  EMH                     |
| 1 625 | ВІДСУТНІЙ КОНТАКТ З ОТРИМУВАЧЕМ                                         | RECEIPIENT NOT AVALIABLE                                                      |  EMH                     |
| 1 626 | ЗМІНА ДАТИ ДОРУЧЕННЯ                                                    | DELIVERY DATE HAS BEEN CHANGED                                                |  EMH                     |
| 1 627 | НЕ ВСТИГ ДОСТАВИТИ                                                      | NOT DELIVERED                                                                 |  EMH                     |
| 1 628 | НЕ КОРЕКТНІ ДАНІ ОТРИМУВАЧА                                             | UNABLE TO VERIFYRECEIPIENT                                                    |  EMH                     |
| 2 331 | ДОРУЧЕННЯ ПРИЗУПИНЕНО                                                   | DELIVERY STOPPED                                                              |                          |
| 2 432 | ЗНЯТО ОБМЕЖЕННЯ В ДОСТАВЦІ                                              | RESTRICTION REMOVED                                                           |                          |
| 1 825 | ПОВЕРНЕННЯ ВІДПРАВНИКУ                                                  | RETURNED TO SENDER                                                            |                          |
| 2 230 | МИТНЕ ОФОРМЛЕННЯ                                                        | HANDED OVER TO CUSTOM                                                         | EME                      |
| 2 533 | УТОЧНЕНО РЕКВІЗИТИ ОТРИМУВАЧА                                           | RECIPIENT'S ADDRESS VERIFICATION                                              |                          |
| 5 051 | ВІДПРАВЛЕННЯ НЕ ПРИБУЛО                                                 | PARCEL NOT ARRIVED                                                            |                          |
| 8 081 | ВІДПРАВКА З ПІДРОЗДІЛУ ВІДПРАВНИКА                                      | DEPARTURE FROM SENDER OFFICE                                                  |                          |


