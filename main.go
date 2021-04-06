package main

import (
	"fmt"
	"strconv"
)

func main() {
	// TODO: create a input over here
	//stringInput := "(1+(4+5+2)-3)+(6+8)"
	stringInput := "(123+2)"
	result := calculate(stringInput)
	fmt.Println(result)
}

// case 1: incoming input is a delimiter
// case 2: incoming input is not a delimiter
// case 2.a: TOS is an operation therefore iterate backwards until you find a delimiter then perform the operation
// case 2.b: TOS is a number therefore just insert
// case 2.c: TOS is a opening parenthesis just insert
// case 3: incomping input is a white space

// assumption is that the string numbers are already evaluated to a single element
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
					total = incomingValue + -presentValue
				}

				subExpressionStack[0] = strconv.Itoa(total)
				subExpressionStack = subExpressionStack[0:1]
			} else {
				continue
			}
		}
	}

	return total
}

// TODO: make sure that the numbers are cramped into one memory space
func calculate(expression string) int {
	calculatorStack := make([]string, 0, len(expression))
	var currentExpressionStart int
	var currentTotal int
	for index, s := range expression {
		incomingValue := string(s)
		if isWhiteSpace(incomingValue) {
			continue
		}

		if isOpeningParenthesis(incomingValue) {
			currentExpressionStart = index
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

			subExpressionValue := evaluate(calculatorStack[currentExpressionStart : index+1])
			currentTotal += subExpressionValue
			calculatorStack[currentExpressionStart] = strconv.Itoa(subExpressionValue)
			calculatorStack = calculatorStack[0:currentExpressionStart]
			currentExpressionStart = 0
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
