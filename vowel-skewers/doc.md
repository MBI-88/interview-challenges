# Vowel Skewers Documentation

## Overview

An authentic vowel skewer is a skewer with a delicious and juicy mix of consonants and vowels. However, the way they are made must be just right:

- Skewers must begin and end with a consonant.
- Skewers must alternate between consonants and vowels.
- There must be an even spacing between each letter on the skewer, so that there is a consistent flavour throughout.

## Functionality

Create a function which returns whether a given vowel skewer is authentic.

### Examples

- `is_authentic_skewer("B--A--N--A--N--A--S")` returns `True`
- `is_authentic_skewer("A--X--E")` returns `False` (Should start and end with a consonant.)
- `is_authentic_skewer("C-L-A-P")` returns `False` (Should alternate between consonants and vowels.)
- `is_authentic_skewer("M--A---T-E-S")` returns `False` (Should have consistent spacing between letters.)

### Notes

- All tests will be given in uppercase.
- Strings without any actual skewer "-" or letters should return False.