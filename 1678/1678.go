package leetcode1678

/**
You own a Goal Parser that can interpret a string command. The command consists of an alphabet of "G", "()" and/or "(al)" in some order. The Goal Parser will interpret "G" as the string "G", "()" as the string "o", and "(al)" as the string "al". The interpreted strings are then concatenated in the original order.

Given the string command, return the Goal Parser's interpretation of command.

Example 1:

Input: command = "G()(al)"
Output: "Goal"
Explanation: The Goal Parser interprets the command as follows:
G -> G
() -> o
(al) -> al
The final concatenated result is "Goal".

Example 2:

Input: command = "G()()()()(al)"
Output: "Gooooal"

Example 3:

Input: command = "(al)G(al)()()G"
Output: "alGalooG"

*/
import "strings"

func interpret1(command string) string {
	command = strings.ReplaceAll(command, "(al)", "al")
	return strings.ReplaceAll(command, "()", "o")
}

func interpret2(command string) string {
	var sb strings.Builder
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			sb.WriteByte('G')
		} else if command[i] == '(' && command[i+1] == ')' {
			sb.WriteByte('o')
			i++
		} else {
			sb.WriteString("al")
			i += 3
		}
	}
	return sb.String()
}
