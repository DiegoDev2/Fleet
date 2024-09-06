package main

// StanfordNer - Stanford NLP Group's implementation of a Named Entity Recognizer
// Homepage: https://nlp.stanford.edu/software/CRF-NER.shtml

import (
	"fmt"
	
	"os/exec"
)

func installStanfordNer() {
	// Método 1: Descargar y extraer .tar.gz
	stanfordner_tar_url := "https://nlp.stanford.edu/software/stanford-ner-4.2.0.zip"
	stanfordner_cmd_tar := exec.Command("curl", "-L", stanfordner_tar_url, "-o", "package.tar.gz")
	err := stanfordner_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stanfordner_zip_url := "https://nlp.stanford.edu/software/stanford-ner-4.2.0.zip"
	stanfordner_cmd_zip := exec.Command("curl", "-L", stanfordner_zip_url, "-o", "package.zip")
	err = stanfordner_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stanfordner_bin_url := "https://nlp.stanford.edu/software/stanford-ner-4.2.0.zip"
	stanfordner_cmd_bin := exec.Command("curl", "-L", stanfordner_bin_url, "-o", "binary.bin")
	err = stanfordner_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stanfordner_src_url := "https://nlp.stanford.edu/software/stanford-ner-4.2.0.zip"
	stanfordner_cmd_src := exec.Command("curl", "-L", stanfordner_src_url, "-o", "source.tar.gz")
	err = stanfordner_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stanfordner_cmd_direct := exec.Command("./binary")
	err = stanfordner_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
