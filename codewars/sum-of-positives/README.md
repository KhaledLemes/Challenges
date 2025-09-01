# Sum of positives
You get an array of numbers<br>
Return the sum of all of the positives ones.<br>
Example [1, -4, 7, 12] => 1+7+12=20<br>
Note If there is nothing to sum, the sum is default to 0.

## Solving
For this challenge I have used a simple *for* loop that iterates within the slice if its lenght is greater than 0. So if the lenght is 0.<br>
<P>Inside the loop, only positive numbers are added to the sum.