---
layout: post
title: mysql - idempotent db update 
tags: [mysql]
readtime: true
comments: true
---

## how to do an idempotent db update
For example, Tom wants to save $100 into the bank

table user_account 

| id | name      | account_id     | amount | currency | 
|----|---------- |----------------|--------|----------|
| 1  | Tom       | 1001           | 0      | USD      |

#### 1. let's see a very simple update sql, it may cause problem if someone other just save $1000 to me...
```sql
UPDATE user_account SET amount = 100 WHERE account_id = 1001
```

#### 2. If I save the money from ATM, and click the submit button for twice... woo, my account turn to $200.
```sql
UPDATE user_account SET amount = amount + 100 WHERE account_id = 1001
```

#### 3. now quick click may not work if it is just myself. what if the account statement is like `+100, -100, +100, -100`. 
```sql
UPDATE user_account SET amount = amount + 100 WHERE account_id = 1001 and amount = 0
```

#### 4. let's create a new table account_record, and new a record_id whenever Tom open a save page.

account_record

| id | record_id | account_id     | amount | origin_amount| 
|----|---------- |----------------|--------|--------------|
| 1  | 1001      | 1001           | 100    | 0            |
| 1  | 1002      | 1001           | -50    | 0            |
| 1  | 1003      | 1001           | -50    | 0            |
| 1  | 1004      | 1001           | 150    | 0            |

user_account

| id | name      | account_id     | amount | currency | record_id | 
|----|---------- |----------------|--------|----------|-----------|
| 1  | Tom       | 1001           | 50     | USD      | 1002      |
| 2  | Jay       | 1002           | 200      | USD      | 1000      |

```sql
insert into account_record(`record_id`, `account_id`, `amount`, `origin_amount`) VALUES (1003, 1001, -50, 50)
UPDATE user_account SET amount = amount -50 WHERE account_id = 1001 and amount = 50 and record_id = 1002
```

#### 5. now if the money is transfer from Jay
```sql
begin transaction; 
insert into account_record(`record_id`, `account_id`, `amount`, `origin_amount`) VALUES (1004, 1002, -50, 200);
insert into account_record(`record_id`, `account_id`, `amount`, `origin_amount`) VALUES (1005, 1001, 50, 50);
UPDATE user_account SET amount = amount +50 WHERE account_id = 1001 and amount = 50 and record_id = 1002;
UPDATE user_account SET amount = amount -50 WHERE account_id = 1002 and amount = 200 and record_id = 1000;
commit;
```
