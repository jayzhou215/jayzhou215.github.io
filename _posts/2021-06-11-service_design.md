---
layout: post
title: design - some design ideas
tags: [architecture]
readtime: true
comments: true
---

## it's common many business process have a lot of similarities.

![](https://pt-starimg.didistatic.com/static/starimg/img/WZZtElqVEi1623416844350.png)

#### if we want to restore a calc after a few days, but many context variables may change.
We can use a data mirror table to record all the related variables used in the calc.

#### write an operation log in db
it can help to know who make the change. 

#### how to make sure the calc result saved correctly
We need idempotent db insert/update, which I talk a lot in [here](2021-06-10-idempotent_db_update.md)

#### retry
in idempotent db update, it may fail in some case, we need retry to make sure the submission of user can be executed finally. 

#### version/source field in DB
if there is an inconsistent data struct upgrade or interface upgrade, with version/source field, it can help to make decision to use the data or not.  

more space decrease not only the time complexity but also code complexity.


## db status vs hot status
which means the status saved in db do not must keep consist with the status returned, which may with some calc. 

#### the statement
include statement with outside and inside

#### there are a lot of details which depend on business requirements.

## reference
[system design primer](https://github.com/donnemartin/system-design-primer/blob/master/README.md)