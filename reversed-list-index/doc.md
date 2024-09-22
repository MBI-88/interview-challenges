### Oddly or Evenly Positioned From Last
# Create a function that extracts the characters from a list on odd or even positions, depending on the specifier. The string "odd" for items on odd positions (... 5, 3, 1) and "even" for even positions (... 6, 4, 2) from the last item of that list or string.

## Examples
``` get_items_at([2, 4, 6, 8, 10], "odd") ➞ [2, 6, 10]
// 2, 6 & 10 occupy the 5th, 3rd and 1st positions from right.
// Odd positions, hence the parity, and from the opposite. ```

``` get_items_at(["E", "D", "A", "B", "I", "T"], "even") ➞ ["E", "A", "I"]
// E, A and I occupy the 6th, 4th and 2nd positions from right.
// Even positions, hence the parity, and from the opposite.
```

```` get_items_at([")", "(", "*", "&", "^", "%", "$", "#", "@", "!"], "even") ➞ [")", "*", ^", "$", "@"] ````

``` get_items_at(["A", "R", "B", "I", "T", "R", "A", "R", "I", "L", "Y"], "even") ➞ ["R", "I", "R", "R", "L"] ```

## Notes
Lists are zero-indexed, so, index+1 = position or position-1 = index.
Items in the list may contain duplicates. See example #4.
The last item in the list is always the first item to start a positional count.
The sequence of the items in the resulting list should be retained (i.e. example #1 - 6 should come after 2 and before 10, example #2 - "A" should come after "E" and before "I").