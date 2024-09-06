package main

// Wordnet - Lexical database for the English language
// Homepage: https://wordnet.princeton.edu/

import (
	"fmt"
	
	"os/exec"
)

func installWordnet() {
	// Método 1: Descargar y extraer .tar.gz
	wordnet_tar_url := "https://wordnetcode.princeton.edu/3.0/WordNet-3.0.tar.bz2"
	wordnet_cmd_tar := exec.Command("curl", "-L", wordnet_tar_url, "-o", "package.tar.gz")
	err := wordnet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wordnet_zip_url := "https://wordnetcode.princeton.edu/3.0/WordNet-3.0.tar.bz2"
	wordnet_cmd_zip := exec.Command("curl", "-L", wordnet_zip_url, "-o", "package.zip")
	err = wordnet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wordnet_bin_url := "https://wordnetcode.princeton.edu/3.0/WordNet-3.0.tar.bz2"
	wordnet_cmd_bin := exec.Command("curl", "-L", wordnet_bin_url, "-o", "binary.bin")
	err = wordnet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wordnet_src_url := "https://wordnetcode.princeton.edu/3.0/WordNet-3.0.tar.bz2"
	wordnet_cmd_src := exec.Command("curl", "-L", wordnet_src_url, "-o", "source.tar.gz")
	err = wordnet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wordnet_cmd_direct := exec.Command("./binary")
	err = wordnet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: tcl-tk")
	exec.Command("latte", "install", "tcl-tk").Run()
}
