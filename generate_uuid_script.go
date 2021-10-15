package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

const inptuFileName = "input.sql"
const outputFileName = "output.sql"

func main() {
	generateUUID()
}

func generateUUID() {
	err := os.Remove(outputFileName)
	if err != nil {
		fmt.Println("Ocorreu um erro ao remover o arquivo:", err)
	}

	currentTimeBeforeExecution := time.Now()

	fmt.Println("Lendo arquivo " + inptuFileName)
	lines := readInputFile()
	fmt.Println("Quantidade de linhas lidas: " + strconv.Itoa(len(lines)))

	fmt.Println("Gerando uuid e escrevendo no " + outputFileName)
	for _, line := range lines {
		writeOutputFile(line)
	}
	fmt.Println("Arquivo gerado com sucesso " + outputFileName)
	currentTimeAfterExecution := time.Now()

	fmt.Println("Tempo de execução " + currentTimeAfterExecution.Sub(currentTimeBeforeExecution).String())
}

func readInputFile() []string {
	var lines []string
	file, err := os.Open(inptuFileName)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		lines = append(lines, line)

		if err == io.EOF {
			break
		}

	}

	file.Close()
	return lines
}

func writeOutputFile(line string) {
	arquivo, err := os.OpenFile(outputFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	myuuid, _ := uuid.NewV4()

	newLine := strings.Replace(line, "$uuid", quoteString(myuuid.String()), -1)
	arquivo.WriteString(newLine + "\n")

	arquivo.Close()
}

func quoteString(value string) string {
	return "'" + value + "'"
}
