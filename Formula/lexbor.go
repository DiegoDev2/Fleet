package main

// Lexbor - Fast embeddable web browser engine written in C with no dependencies
// Homepage: https://lexbor.com/

import (
	"fmt"
	
	"os/exec"
)

func installLexbor() {
	// Método 1: Descargar y extraer .tar.gz
	lexbor_tar_url := "https://github.com/lexbor/lexbor/archive/refs/tags/v2.3.0.tar.gz"
	lexbor_cmd_tar := exec.Command("curl", "-L", lexbor_tar_url, "-o", "package.tar.gz")
	err := lexbor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lexbor_zip_url := "https://github.com/lexbor/lexbor/archive/refs/tags/v2.3.0.zip"
	lexbor_cmd_zip := exec.Command("curl", "-L", lexbor_zip_url, "-o", "package.zip")
	err = lexbor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lexbor_bin_url := "https://github.com/lexbor/lexbor/archive/refs/tags/v2.3.0.bin"
	lexbor_cmd_bin := exec.Command("curl", "-L", lexbor_bin_url, "-o", "binary.bin")
	err = lexbor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lexbor_src_url := "https://github.com/lexbor/lexbor/archive/refs/tags/v2.3.0.src.tar.gz"
	lexbor_cmd_src := exec.Command("curl", "-L", lexbor_src_url, "-o", "source.tar.gz")
	err = lexbor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lexbor_cmd_direct := exec.Command("./binary")
	err = lexbor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
