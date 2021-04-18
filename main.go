package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: change the algorihtm

func main() {
	// TODO: create a input over here
	stringInput := "(5-(1+(5)))"
	result := calculate(stringInput)
	fmt.Println(result)
}

func evaluate(expressionStack []string) int {
	subExpressionStack := make([]string, 0, len(expressionStack))
	var total int
	var TOS int
	for _, s := range expressionStack {
		if !isClosingParenthesis(s) && !isOpeningParenthesis(s) {
			subExpressionStack = append(subExpressionStack, s)
			TOS = len(subExpressionStack) - 1
			if isNumber(s) && TOS > 1 {
				previousString := subExpressionStack[TOS-1]
				if isAddition(previousString) {
					incomingValue, _ := strconv.Atoi(s)
					presentValue, _ := strconv.Atoi(subExpressionStack[TOS-2])
					// pop it
					total = incomingValue + presentValue
				}
				if isSubtraction(previousString) {
					incomingValue, _ := strconv.Atoi(s)
					presentValue, _ := strconv.Atoi(subExpressionStack[TOS-2])
					// pop it
					total = -incomingValue + presentValue
				}

				subExpressionStack[0] = strconv.Itoa(total)
				subExpressionStack = subExpressionStack[0:1]
			} else {
				continue
			}
		}
	}

	if len(subExpressionStack) == 1 {
		value, _ := strconv.Atoi(subExpressionStack[0])
		return value
	}

	return total
}

func calculate(expression string) int {
	expression = "(" + expression + ")"

	calculatorStack := make([]string, 0, len(expression))
	var currentExpressionStart int
	var currentTotal int
	for index, s := range expression {
		incomingValue := string(s)
		if isWhiteSpace(incomingValue) {
			continue
		}

		if isOpeningParenthesis(incomingValue) {
			currentExpressionStart = len(calculatorStack)
		}

		if isSubtraction(incomingValue) {
			TOS := len(calculatorStack) - 1
			TOSValue := calculatorStack[TOS]
			isTOSNumber := isNumber(TOSValue)
			if !isTOSNumber {
				calculatorStack = append(calculatorStack, "0", incomingValue)
				continue
			}
		}

		if !isClosingParenthesis(incomingValue) {
			if isNumber(incomingValue) {
				TOSCalculatorStack := len(calculatorStack) - 1
				TOSIsNumber := isNumber(calculatorStack[TOSCalculatorStack])

				if TOSIsNumber {
					TOSValue, _ := strconv.Atoi(calculatorStack[TOSCalculatorStack])
					TOSValue *= 10
					incomingValueInInt, _ := strconv.Atoi(incomingValue)
					calculatorStack[TOSCalculatorStack] = strconv.Itoa(TOSValue + incomingValueInInt)
				} else {
					calculatorStack = append(calculatorStack, incomingValue)
				}
			} else {
				calculatorStack = append(calculatorStack, incomingValue)
			}
			continue
		}

		if isClosingParenthesis(incomingValue) {
			calculatorStack = append(calculatorStack, incomingValue)

			for i := 0; i <= len(calculatorStack)-1; i++ {
				if isOpeningParenthesis(calculatorStack[i]) {
					currentExpressionStart = i
					continue
				}
			}

			subExpressionValue := evaluate(calculatorStack[currentExpressionStart:])
			currentTotal = subExpressionValue
			calculatorStack[currentExpressionStart] = strconv.Itoa(subExpressionValue)
			calculatorStack = calculatorStack[0 : currentExpressionStart+1]

			//currentExpressionStart = 0
			if index == len(expression)-1 && len(calculatorStack) > 1 {
				currentTotal = calculate(strings.Join(calculatorStack, ""))
			}

			continue
		}

	}

	return currentTotal
}

func isAddition(s string) bool {
	return s == "+"
}

func isSubtraction(s string) bool {
	return s == "-"
}

func isClosingParenthesis(s string) bool {
	return s == ")"
}

func isOpeningParenthesis(s string) bool {
	return s == "("
}

func isNumber(s string) bool {
	return !isClosingParenthesis(s) && !isOpeningParenthesis(s) && !isSubtraction(s) && !isAddition(s)
}

func isWhiteSpace(s string) bool {
	return s == " "
}
