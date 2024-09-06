package main

// Mitie - Library and tools for information extraction
// Homepage: https://github.com/mit-nlp/MITIE/

import (
	"fmt"
	
	"os/exec"
)

func installMitie() {
	// Método 1: Descargar y extraer .tar.gz
	mitie_tar_url := "https://github.com/mit-nlp/MITIE/archive/refs/tags/v0.7.tar.gz"
	mitie_cmd_tar := exec.Command("curl", "-L", mitie_tar_url, "-o", "package.tar.gz")
	err := mitie_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mitie_zip_url := "https://github.com/mit-nlp/MITIE/archive/refs/tags/v0.7.zip"
	mitie_cmd_zip := exec.Command("curl", "-L", mitie_zip_url, "-o", "package.zip")
	err = mitie_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mitie_bin_url := "https://github.com/mit-nlp/MITIE/archive/refs/tags/v0.7.bin"
	mitie_cmd_bin := exec.Command("curl", "-L", mitie_bin_url, "-o", "binary.bin")
	err = mitie_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mitie_src_url := "https://github.com/mit-nlp/MITIE/archive/refs/tags/v0.7.src.tar.gz"
	mitie_cmd_src := exec.Command("curl", "-L", mitie_src_url, "-o", "source.tar.gz")
	err = mitie_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mitie_cmd_direct := exec.Command("./binary")
	err = mitie_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
