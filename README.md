# prompter

Prompter is a small library that makes it easier to prompt user in the terminal for answers. It currently supports 5 types of questions:
1. Simple question - Just a question with an answer
2. Simple question with default answer - Like above, but if no response is provided, the default is returned.
3. Secret question - For passwords. Also, it's not implemented yet, but the placeholder is there
4. Selection question - A question with a list of possible answers. Only accepts values for 0<=ans<len(options) as answers.
5. Selection question with default answer - Like above, but if no response is provided, the default is returned.
