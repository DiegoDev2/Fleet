package main

// Rdfind - Find duplicate files based on content (NOT file names)
// Homepage: https://rdfind.pauldreik.se/

import (
	"fmt"
	
	"os/exec"
)

func installRdfind() {
	// Método 1: Descargar y extraer .tar.gz
	rdfind_tar_url := "https://rdfind.pauldreik.se/rdfind-1.6.0.tar.gz"
	rdfind_cmd_tar := exec.Command("curl", "-L", rdfind_tar_url, "-o", "package.tar.gz")
	err := rdfind_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rdfind_zip_url := "https://rdfind.pauldreik.se/rdfind-1.6.0.zip"
	rdfind_cmd_zip := exec.Command("curl", "-L", rdfind_zip_url, "-o", "package.zip")
	err = rdfind_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rdfind_bin_url := "https://rdfind.pauldreik.se/rdfind-1.6.0.bin"
	rdfind_cmd_bin := exec.Command("curl", "-L", rdfind_bin_url, "-o", "binary.bin")
	err = rdfind_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rdfind_src_url := "https://rdfind.pauldreik.se/rdfind-1.6.0.src.tar.gz"
	rdfind_cmd_src := exec.Command("curl", "-L", rdfind_src_url, "-o", "source.tar.gz")
	err = rdfind_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rdfind_cmd_direct := exec.Command("./binary")
	err = rdfind_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: nettle")
	exec.Command("latte", "install", "nettle").Run()
}
