// Copyright 2018 Yaacov Zamir <kobi.zamir@gmail.com>
// and other contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/hokaccha/go-prettyjson"
	"gopkg.in/yaml.v2"

	"github.com/yaacov/tsl/pkg/parser"
	"github.com/yaacov/tsl/pkg/tsl"
)

func check(err error) {
	if err != nil {
		fmt.Printf("Err: %v\n", err)
	}
}

func main() {
	var err error
	var tree tsl.Node
	var s []byte

	// Setup the input
	inputPtr := flag.String("i", "", "the tsl string to parse (e.g. \"animal = 'kitty'\")")
	outputPtr := flag.String("o", "json", "output format [json/yaml/prettyjson]")
	flag.Parse()

	// Setup the ErrorListener
	errorListener := tsl.NewErrorListener()

	// Setup the input
	is := antlr.NewInputStream(*inputPtr)

	// Create the Lexer
	lexer := parser.NewTSLLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewTSLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	// Finally parse the expression (by walking the tree)
	var listener tsl.Listener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	err = errorListener.Err
	check(err)

	tree, err = listener.GetTree()
	check(err)

	// If listener has erros, we can not print the tree
	if err != nil {
		os.Exit(1)
	}

	switch *outputPtr {
	case "json":
		s, err = json.Marshal(tree)
	case "yaml":
		s, err = yaml.Marshal(tree)
	case "prettyjson":
		s, err = prettyjson.Marshal(tree)
	default:
		s, err = json.Marshal(tree)
	}

	check(err)
	fmt.Println(string(s))
}