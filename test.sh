#!/bin/bash

echo "Running tests..."

name='andrews'

echo "Hello, $name! Welcome to the testing script."


fruits=('apple' 'banana' 'cherry')

echo "Here are some fruits: ${fruits[@]}" # we should use ${fruits[@]} to print all elements of the array. @ means all elements of the array.
 


 # associative array

 declare -A cars

 cars[benz]="white"
 cars[audi]="black"

 echo "${cars[@]}" # to print all values of the associative array, we can use ${cars[@]}. This will print "white black" because those are the values associated with the keys "benz" and "audi".


 echo ${cars[audi]} # to access the value of a specific key, we use ${cars[key]}. In this case, it will print "black" because the value associated with the key "audi" is "black".


 a=10
 b=20

if [ $a -lt $b ]; then
    echo "is less"

else
    echo "is greater"
    fi


# for loop
for i in {1..5}; do
    echo "Number: $i"
done  

# while loop
num=1

while [ $num -le 5 ]; do
    echo "current num: $num"
  ((num++))
done



