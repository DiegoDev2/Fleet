package main

// StanfordCorenlp - Java suite of core NLP tools
// Homepage: https://stanfordnlp.github.io/CoreNLP/

import (
	"fmt"
	
	"os/exec"
)

func installStanfordCorenlp() {
	// Método 1: Descargar y extraer .tar.gz
	stanfordcorenlp_tar_url := "https://nlp.stanford.edu/software/stanford-corenlp-4.4.0.zip"
	stanfordcorenlp_cmd_tar := exec.Command("curl", "-L", stanfordcorenlp_tar_url, "-o", "package.tar.gz")
	err := stanfordcorenlp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stanfordcorenlp_zip_url := "https://nlp.stanford.edu/software/stanford-corenlp-4.4.0.zip"
	stanfordcorenlp_cmd_zip := exec.Command("curl", "-L", stanfordcorenlp_zip_url, "-o", "package.zip")
	err = stanfordcorenlp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stanfordcorenlp_bin_url := "https://nlp.stanford.edu/software/stanford-corenlp-4.4.0.zip"
	stanfordcorenlp_cmd_bin := exec.Command("curl", "-L", stanfordcorenlp_bin_url, "-o", "binary.bin")
	err = stanfordcorenlp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stanfordcorenlp_src_url := "https://nlp.stanford.edu/software/stanford-corenlp-4.4.0.zip"
	stanfordcorenlp_cmd_src := exec.Command("curl", "-L", stanfordcorenlp_src_url, "-o", "source.tar.gz")
	err = stanfordcorenlp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stanfordcorenlp_cmd_direct := exec.Command("./binary")
	err = stanfordcorenlp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
