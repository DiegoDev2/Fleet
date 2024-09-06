package main

// PlzCli - Copilot for your terminal
// Homepage: https://github.com/m1guelpf/plz-cli

import (
	"fmt"
	
	"os/exec"
)

func installPlzCli() {
	// Método 1: Descargar y extraer .tar.gz
	plzcli_tar_url := "https://github.com/m1guelpf/plz-cli/archive/refs/tags/v0.1.9.tar.gz"
	plzcli_cmd_tar := exec.Command("curl", "-L", plzcli_tar_url, "-o", "package.tar.gz")
	err := plzcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	plzcli_zip_url := "https://github.com/m1guelpf/plz-cli/archive/refs/tags/v0.1.9.zip"
	plzcli_cmd_zip := exec.Command("curl", "-L", plzcli_zip_url, "-o", "package.zip")
	err = plzcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	plzcli_bin_url := "https://github.com/m1guelpf/plz-cli/archive/refs/tags/v0.1.9.bin"
	plzcli_cmd_bin := exec.Command("curl", "-L", plzcli_bin_url, "-o", "binary.bin")
	err = plzcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	plzcli_src_url := "https://github.com/m1guelpf/plz-cli/archive/refs/tags/v0.1.9.src.tar.gz"
	plzcli_cmd_src := exec.Command("curl", "-L", plzcli_src_url, "-o", "source.tar.gz")
	err = plzcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	plzcli_cmd_direct := exec.Command("./binary")
	err = plzcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
