package cmd

import (
	"bufio"
	"fmt"
	"go/parser"
	"os"

	"github.com/Knetic/govaluate"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calc [expresion]",
	Short: "Calculator operations",
	Long:  "Operations calculator add(+), difference(-), divide(/) and others",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 1 {
			expression := args[0]
			result, err := evaluateExpression(expression)
			if err != nil {
				fmt.Println("Error evaluate expression  :", err)
				return
			}
			fmt.Printf("Resultado: %v\n", result)
		} else {
			startREPL()
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Modo interactivo - Escribe una expresion matematica o 'exit' para salir")

	for {
		fmt.Print(">>> ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		result, err := evaluateExpression(input)
		if err != nil {
			fmt.Println("Error evaluando la expresión:", err)
			continue
		}

		fmt.Printf("Resultado: %v\n", result)
	}

}

func evaluateExpression(expr string) (interface{}, error) {
	_, err := parser.ParseExpr(expr)
	if err != nil {
		return nil, fmt.Errorf("Expresion no valida")
	}

	/**
	es una función útil para crear y evaluar expresiones dinámicas en Go.
	La función NewEvaluableExpression en tu código
	es un envoltorio que facilita la creación de
	expresiones sin funciones personalizadas.
	**/
	evalExpr, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return nil, err
	}

	// Evaluar sin variables adicionales
	result, err := evalExpr.Evaluate(nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}
