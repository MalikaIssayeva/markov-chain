# markov chain

## General Criteria

- Your code MUST be written in accordance with gofumpt. If not, you will be graded 0 automatically.
- Your program MUST be able to compile successfully.
- Your program MUST not exit unexpectedly (any panics: nil-pointer dereference, index out of range etc.). If so, you will be get 0 during the defence.
- Only built-in packages are allowed. If not, you will get 0 grade.
- The project MUST be compiled by the following command in the project's root directory:

    
```bash
$ go build -o markovchain .
```
## Mandatory Part

### Baseline

By default your program must read from stdin the whole text and generate the text according to Markov Chain algorithm.

Outcomes:

- Program prints generated text according to the Markov Chain algorithm.

Notes:

- Suffix length is ALWAYS 1 word.
- Default prefix length is 2 words.
- Default starting prefix is the first N words of the text, where N is the length of the prefix.
- Default number of maximum words is 100.

Constraints:

- If any error print an error message indicating the reason.
- The code should stop generating code after it printed maximum number of words or encountered the very last word in the text.

Examples:

```bash
$ cat the_great_gatsby.txt | ./markovchain | cat -e
Chapter 1 In my younger and more stable, become for a job. He hadn't eat anything for a long, silent time. It was the sound of someone splashing after us over the confusion a long many-windowed room which overhung the terrace. Eluding Jordan's undergraduate who was well over sixty, and Maurice A. Flink and the great bursts of leaves growing on the air now. "How do you want? What do you like Europe?" she exclaimed surprisingly. "I just got here a minute. "Yes." He hesitated. "Was she killed?" "Yes." "I thought you didn't, if you'll pardon my--you see, I carry$
```
```bash
$ cat the_great_gatsby.txt | ./markovchain | wc -w
   100
```
```bash
$ ./markovchain
Error: no input text
```

### Number of words

Your program must be able to accept maximum number of words to be generated.

Outcomes:

- Program prints generated text according to the Markov Chain algorithm limited by the given maximum number of words.

Constraints:

- Given number can't be negative.
- Given number can't be more 10,000.
- If any error print an error message indicating the reason.

```bash
$ cat the_great_gatsby.txt | ./markovchain -w 10 | cat -e
Chapter 1 In my younger and more stable, become for$
```
### Prefix

Your program must be able to accept the starting prefix.

Outcomes:

- Program prints generated text according to the Markov Chain algorithm that starts with the given prefix.

Constraints:

- Given prefix must be present in the original text.
- If any error print an error message indicating the reason.

```bash
$ cat the_great_gatsby.txt | ./markovchain -w 10 -p "to play" | cat -e
to play for you in that vast obscurity beyond the$
```
### Prefix length

Your program must be able to accept the prefix length.

Outcomes:

- Program prints generated text according to the Markov Chain algorithm with the given prefix length.

Constraints:

- Given prefix length can't be negative.
- Given prefix length can't be greater than 5.
- If any error print an error message indicating the reason.

```bash
$ cat the_great_gatsby.txt | ./markovchain -w 10 -p "to something funny" -l 3
to something funny the last two days," remarked Wilson. "That's
```

### Usage

Your program must be able to print usage information.

Outcomes:

- Program prints usage text.

```bash
$ ./markovchain --help
Markov Chain text generator.

Usage:
  markovchain [-w <N>] [-p <S>] [-l <N>]
  markovchain --help

Options:
  --help  Show this screen.
  -w N    Number of maximum words
  -p S    Starting prefix
  -l N    Prefix length
```