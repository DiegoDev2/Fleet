package main

// StanfordParser - Statistical NLP parser
// Homepage: https://nlp.stanford.edu/software/lex-parser.shtml

import (
	"fmt"
	
	"os/exec"
)

func installStanfordParser() {
	// Método 1: Descargar y extraer .tar.gz
	stanfordparser_tar_url := "https://nlp.stanford.edu/software/stanford-parser-4.2.0.zip"
	stanfordparser_cmd_tar := exec.Command("curl", "-L", stanfordparser_tar_url, "-o", "package.tar.gz")
	err := stanfordparser_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stanfordparser_zip_url := "https://nlp.stanford.edu/software/stanford-parser-4.2.0.zip"
	stanfordparser_cmd_zip := exec.Command("curl", "-L", stanfordparser_zip_url, "-o", "package.zip")
	err = stanfordparser_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stanfordparser_bin_url := "https://nlp.stanford.edu/software/stanford-parser-4.2.0.zip"
	stanfordparser_cmd_bin := exec.Command("curl", "-L", stanfordparser_bin_url, "-o", "binary.bin")
	err = stanfordparser_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stanfordparser_src_url := "https://nlp.stanford.edu/software/stanford-parser-4.2.0.zip"
	stanfordparser_cmd_src := exec.Command("curl", "-L", stanfordparser_src_url, "-o", "source.tar.gz")
	err = stanfordparser_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stanfordparser_cmd_direct := exec.Command("./binary")
	err = stanfordparser_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
