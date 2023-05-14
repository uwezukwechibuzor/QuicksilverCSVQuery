# QuicksilverCSVQuery
An app in go that could query data from a quicksilver-node and output the Data in a CSV format.

## Get started

Install [Go](https://go.dev/dl/)

## Build and Install
```
$ make build
```

Run to see Available Commands

```
./app
```



<img width="1145" alt="Screenshot 2023-05-15 at 00 18 10" src="https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/assets/66339097/c378c0b7-c887-4022-8f76-3fba00835b2b">
<img width="772" alt="Screenshot 2023-05-15 at 00 19 11" src="https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/assets/66339097/4ea18db5-0c4f-493f-b432-53b0c795c3db">
<img width="772" alt="Screenshot 2023-05-15 at 00 19 39" src="https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/assets/66339097/3a136a7f-d46f-44af-ad93-798fe30a84d9">
<img width="772" alt="Screenshot 2023-05-15 at 00 20 02" src="https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/assets/66339097/e5704894-fa3f-4fc4-9921-01bdb3d4424e">

Commandline Samples


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

Samples of Data Output in CSV format


[Delegator Validator Mapping](https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/blob/master/delegator_validator_mapping.csv)
![Screenshot 2023-05-15 at 00 30 56](https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/assets/66339097/eb6aefb4-109c-4b97-8b42-1606287470da)

[Validator Delegator Mapping](https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/blob/master/validator_delegator_mapping.csv)
![Screenshot 2023-05-15 at 00 33 05](https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/assets/66339097/c441a1d6-49fc-4a1a-a01f-ecf4bf0efa5e)

[x/InterchainStaking Module](https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/blob/master/interchain-staking.csv)
![Screenshot 2023-05-15 at 00 32 35](https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/assets/66339097/2a458b68-65ea-4ae5-90a6-ec560376774a)

[Vesting Accounts Details(DelayedVestingAccount, PermanentLockedAccount and PeriodicVestingAccount)](https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/blob/master/vesting_accounts.csv)
![Screenshot 2023-05-15 at 00 33 37](https://github.com/uwezukwechibuzor/QuicksilverCSVQuery/assets/66339097/f34b9888-5b79-4a67-83a0-30d7ae5eb11b)
