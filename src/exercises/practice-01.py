#! /usr/bin/env python
# -*- coding: utf-8 -*-
# __author__ = ""
# Date: 2019-07-09

count = 0
for i in [1, 2, 3, 4]:
    for j in [1, 2, 3, 4]:
        for k in [1, 2, 3, 4]:
            if i != j and i != k and j != k:
                count += 1
                print(i, j, k)
print(count)