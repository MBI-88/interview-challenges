# Longest Alternating Substring (expert level)

Given a string of digits, return the longest substring with alternating odd/even or even/odd digits. 
If two or more substrings have the same length, return the substring that occurs first.

### Examples
- `longest_substring("225424272163254474441338664823")` ➞ "272163254"
  - Substrings: 254, 272163254, 474, 41, 38, 23

- `longest_substring("594127169973391692147228678476")` ➞ "16921472"
  - Substrings: 94127, 169, 16921472, 678, 476

- `longest_substring("721449827599186159274227324466")` ➞ "7214"
  - Substrings: 7214, 498, 27, 18, 61, 9274, 27, 32
  - 7214 and 9274 have the same length, but 7214 occurs first.

### Notes
- The minimum alternating substring size is 2, and there will always be at least one alternating substring.
