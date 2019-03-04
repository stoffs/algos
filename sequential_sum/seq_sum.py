#!/usr/bin/env python

def sequence_sum(begin_number, end_number, step):
    if begin_number > end_number:
        return 0

    r = xrange(begin_number, end_number+1, step)

    return sum(r)


print("final result {}".format(sequence_sum(1,5,1)))
