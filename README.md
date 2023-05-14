# QuicksilverCSVQuery
An app in go that could query data from a quicksilver-node and output the Data in a CSV format.

## Get started

Install [Go](https://go.dev/dl/)

## Build and Install
```
$ make build
```

Run to Available Commands

```
./app
```



<img width="1145" alt="Screenshot 2023-05-15 at 00 18 10" src="https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/assets/66339097/c378c0b7-c887-4022-8f76-3fba00835b2b">
<img width="772" alt="Screenshot 2023-05-15 at 00 19 11" src="https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/assets/66339097/4ea18db5-0c4f-493f-b432-53b0c795c3db">
<img width="772" alt="Screenshot 2023-05-15 at 00 19 39" src="https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/assets/66339097/3a136a7f-d46f-44af-ad93-798fe30a84d9">
<img width="772" alt="Screenshot 2023-05-15 at 00 20 02" src="https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/assets/66339097/e5704894-fa3f-4fc4-9921-01bdb3d4424e">

Examples and data in csv format


```
 ./app  query-validator-delegator quickvaloper16pj5gljqnqs0ajxakccfjhu05yczp987nmxvcw --output validator_delegator_mapping.csv
```

```
./app  query-delegator-validator quick16pj5gljqnqs0ajxakccfjhu05yczp987dkgpk8 --output delegator_validator_mapping.csv
```
```
./app  query-interchain-staking  cosmoshub-4 --output interchain-staking.csv
```


```
./app  query-vesting-accounts --output vesting_accounts.csv 
```




