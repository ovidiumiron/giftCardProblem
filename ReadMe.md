## Introduction

You have been given a gift card that is about to expire and you want to buy gifts for 2 friends. You want to spend the whole gift card, or if that’s not an option as close to the balance as possible. You have a list of sorted prices for a popular store that you know they both like to shop at. Your challenge is to find two distinct items in the list whose sum is minimally under (or equal to) the gift card balance.

The file contains two columns:

1. A unique identifier of the item. You can assume there are no duplicates.

2. The value of that item in cents. It is always a positive integer that represents the price in
cents (1000 = $10.00).

Write a program to find the best two items. It takes two inputs:

1. A filename with a list of sorted prices

2. The balance of your gift card

If no two items have a sum that is less than or equal to the balance on the gift card, print “Not possible”. You don’t have to return every possible pair that is under the balance, just one such pair.

Example:
$ cat prices.txt

Candy Bar, 500 Paperback Book, 700 Detergent, 1000 Headphones, 1400 Earmuffs, 2000 Bluetooth Stereo, 6000

$ find-pair prices.txt 2500 Candy Bar 500, Earmuffs 2000

$ find-pair prices.txt 2300 Paperback Book 700, Headphones 1400

$ find-pair prices.txt 10000 Earmuffs 2000, Bluetooth Stereo 6000

$ find-pair prices.txt 1100 Not possible
