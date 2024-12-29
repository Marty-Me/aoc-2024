package day_24

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	key     GateKey
	operand *Operand
	nodeA   *Node
	nodeB   *Node
	state   *bool
}

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	initialInputs := parseInitialInputs(scanner)
	logicGates := parseLogicGates(scanner)

	processedNodes := make(map[string]*Node, 0)
	topGates := make(map[string]*Node, 0)

	for _, gate := range logicGates {
		node := getOrCreateNode(gate.resultGate, &processedNodes, &initialInputs)
		nodeOperand := gate.operand
		node.operand = &nodeOperand

		a := getOrCreateNode(gate.inputA, &processedNodes, &initialInputs)
		b := getOrCreateNode(gate.inputB, &processedNodes, &initialInputs)

		node.nodeA = a
		node.nodeB = b

		if gate.resultGate[0] == 'z' {
			gateIndex := gate.resultGate[1:]
			topGates[gateIndex] = node
		}
	}

	topGateValues := make([]rune, len(topGates))

	for gateKey, topGate := range topGates {
		value := traverseGraph(topGate, &initialInputs)
		intKey, _ := strconv.Atoi(gateKey)

		bitRune := '0'
		if value {
			bitRune = '1'
		}
		topGateValues[len(topGates)-intKey-1] = bitRune
	}
	binaryString := string(topGateValues)
	result, _ := strconv.ParseInt(binaryString, 2, 0)
	return fmt.Sprint(result)
}

func traverseGraph(n *Node, knownGateValues *map[GateKey]*Node) bool {
	node, found := (*knownGateValues)[n.key]

	if found {
		return *node.state
	}

	a := traverseGraph(n.nodeA, knownGateValues)
	b := traverseGraph(n.nodeB, knownGateValues)

	var value bool

	if (*n.operand) == AND {
		value = a && b
	} else if (*n.operand) == OR {
		value = a || b
	} else {
		value = (a || b) && !(a && b)
	}

	n.state = &value
	(*knownGateValues)[n.key] = n

	return value
}

func getOrCreateNode(gateKey GateKey, nodes *map[string]*Node, initialInputs *map[GateKey]*Node) *Node {
	node, exists := (*nodes)[gateKey]
	initialInput, isInitialInput := (*initialInputs)[gateKey]

	if !exists {
		var state *bool

		if isInitialInput {
			state = initialInput.state
		}

		node = &Node{
			gateKey,
			nil,
			nil,
			nil,
			state,
		}
		(*nodes)[gateKey] = node
	}

	return node
}

type GateKey = string

type Operand int

const (
	AND = iota
	XOR
	OR
)

type LogicGate struct {
	resultGate GateKey
	inputA     GateKey
	inputB     GateKey
	operand    Operand
}

func parseInitialInputs(scanner *bufio.Scanner) map[GateKey]*Node {
	inputs := make(map[GateKey]*Node, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			return inputs
		}

		kv := strings.Split(line, ": ")
		key := kv[0]
		value := kv[1]
		gateState, _ := strconv.ParseBool(value)
		inputs[key] = &Node{
			key,
			nil,
			nil,
			nil,
			&gateState,
		}
	}
	return inputs
}

func parseLogicGates(scanner *bufio.Scanner) []LogicGate {
	logicGates := make([]LogicGate, 0)
	for scanner.Scan() {
		line := scanner.Text()
		eq := strings.Split(line, " -> ")
		logicGate := eq[0]
		output := eq[1]

		logicGateComponents := strings.Split(logicGate, " ")

		a := logicGateComponents[0]

		operand := logicGateComponents[1]
		parsedOperand := parseOperand(operand)

		b := logicGateComponents[2]

		logicGates = append(logicGates, LogicGate{
			output,
			a,
			b,
			parsedOperand,
		})
	}
	return logicGates
}

func parseOperand(operand string) Operand {
	if operand == "AND" {
		return AND
	} else if operand == "OR" {
		return OR
	} else {
		return XOR
	}
}
