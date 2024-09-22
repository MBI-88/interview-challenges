# The ECG Sequence

The ECG Sequence is a unique sequence of numbers that begins with the numbers 1 and 2. Each subsequent number in the sequence is the smallest number not already present that shares a common factor (other than 1) with its predecessor. This sequence diverges from a standard numerical sequence in that every number (after 1 and 2) occupies a different position than it would in a simple count.

## Sequence Generation

To generate the sequence, follow these steps:

1. Start with the initial sequence `[1, 2]`.
2. Find the smallest number not already in the sequence that shares a common factor with the last number in the sequence.
3. Add this number to the sequence.
4. Repeat steps 2 and 3.

### Example

- Starting sequence: `[1, 2]`
- `3` does not share a factor with `2`, but `4` does (factor of 2).
- New sequence: `[1, 2, 4]`
- Continuing this process, the sequence grows: `[1, 2, 4, 6, 3]`

Thus, the number `3` is at index `4` in the ECG Sequence.

## Function Implementation

Implement a function `ecg_seq_index(n)` that returns the index of the number `n` in the ECG Sequence.

### Examples

- `ecg_seq_index(3) ➞ 4`
- `ecg_seq_index(5) ➞ 9`
- `ecg_seq_index(7) ➞ 13`

## Notes

- ECG stands for electrocardiogram. The sequence visually resembles an ECG graph when plotted.
- An interesting property of this sequence is that every odd prime number `p` is preceded by `2p` and followed by `3p`.