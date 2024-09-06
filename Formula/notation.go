package main

// Notation - CLI tool to sign and verify OCI artifacts and container images
// Homepage: https://notaryproject.dev/

import (
	"fmt"
	
	"os/exec"
)

func installNotation() {
	// Método 1: Descargar y extraer .tar.gz
	notation_tar_url := "https://github.com/notaryproject/notation/archive/refs/tags/v1.2.0.tar.gz"
	notation_cmd_tar := exec.Command("curl", "-L", notation_tar_url, "-o", "package.tar.gz")
	err := notation_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	notation_zip_url := "https://github.com/notaryproject/notation/archive/refs/tags/v1.2.0.zip"
	notation_cmd_zip := exec.Command("curl", "-L", notation_zip_url, "-o", "package.zip")
	err = notation_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	notation_bin_url := "https://github.com/notaryproject/notation/archive/refs/tags/v1.2.0.bin"
	notation_cmd_bin := exec.Command("curl", "-L", notation_bin_url, "-o", "binary.bin")
	err = notation_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	notation_src_url := "https://github.com/notaryproject/notation/archive/refs/tags/v1.2.0.src.tar.gz"
	notation_cmd_src := exec.Command("curl", "-L", notation_src_url, "-o", "source.tar.gz")
	err = notation_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	notation_cmd_direct := exec.Command("./binary")
	err = notation_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
