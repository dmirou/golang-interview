# Task

It is New Year's Day and people are in line for the Wonderland rollercoaster ride. 
Each person wears a sticker indicating their initial position in the queue from  to n. 
Any person can bribe the person directly in front of them to swap positions, but they still 
wear their original sticker. One person can bribe at most two others.

Determine the minimum number of bribes that took place to get to a given queue order. 
Print the number of bribes, or, if anyone has bribed more than two people, print "Too chaotic".

## Function Description

minimumBribes has the following parameter(s):

int q[n]: the positions of the people after all bribes

## Returns

No value is returned. Print the minimum number of bribes necessary or "Too chaotic" 
if someone has bribed more than  people.

## Sample Input

t - number of test cases

STDIN       Function
-----       --------
2           t = 2
5           n = 5
2 1 5 3 4   q = [2, 1, 5, 3, 4]
5           n = 5
2 5 1 3 4   q = [2, 5, 1, 3, 4]


## Sample Output

3
Too chaotic
