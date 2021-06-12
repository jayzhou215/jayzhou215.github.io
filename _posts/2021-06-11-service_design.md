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

#### how to make sure the calc result saved correctly
We need idempotent db insert/update, which I talk a lot in [here](2021-06-10-idempotent_db_update.md)

#### there are a lot of details which depend on business requirements.

## reference
[system design primer](https://github.com/donnemartin/system-design-primer/blob/master/README.md)