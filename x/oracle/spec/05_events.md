<!--
order: 5
-->

# Events

The oracle module emits the following events:

## EndBlocker

| Type                 | Attribute Key | Attribute Value |
| -------------------- | ------------- | --------------- |
| exchange_rate_update | denom         | {denom}         |
| exchange_rate_update | exchange_rate | {exchangeRate}  |

## Handlers

### MsgDelegateFeedConsent

| Type          | Attribute Key | Attribute Value    |
| ------------- | ------------- | ------------------ |
| feed_delegate | operator      | {validatorAddress} |
| feed_delegate | feeder        | {feederAddress}    |
| message       | module        | oracle             |
| message       | action        | delegatefeeder     |
| message       | sender        | {senderAddress}    |

### MsgAggregateExchangeRatePrevote

| Type              | Attribute Key | Attribute Value              |
| ----------------- | ------------- | ---------------------------- |
| aggregate_prevote | voter         | {validatorAddress}           |
| aggregate_prevote | feeder        | {feederAddress}              |
| message           | module        | oracle                       |
| message           | action        | aggregateexchangerateprevote |
| message           | sender        | {senderAddress}              |

### MsgAggregateExchangeRateVote

| Type           | Attribute Key  | Attribute Value           |
| -------------- | -------------- | ------------------------- |
| aggregate_vote | voter          | {validatorAddress}        |
| aggregate_vote | exchange_rates | {exchangeRates}           |
| aggregate_vote | feeder         | {feederAddress}           |
| message        | module         | oracle                    |
| message        | action         | aggregateexchangeratevote |
| message        | sender         | {senderAddress}           |
