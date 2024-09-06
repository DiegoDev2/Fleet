package main

// Melange - Build APKs from source code
// Homepage: https://github.com/chainguard-dev/melange

import (
	"fmt"
	
	"os/exec"
)

func installMelange() {
	// Método 1: Descargar y extraer .tar.gz
	melange_tar_url := "https://github.com/chainguard-dev/melange/archive/refs/tags/v0.11.6.tar.gz"
	melange_cmd_tar := exec.Command("curl", "-L", melange_tar_url, "-o", "package.tar.gz")
	err := melange_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	melange_zip_url := "https://github.com/chainguard-dev/melange/archive/refs/tags/v0.11.6.zip"
	melange_cmd_zip := exec.Command("curl", "-L", melange_zip_url, "-o", "package.zip")
	err = melange_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	melange_bin_url := "https://github.com/chainguard-dev/melange/archive/refs/tags/v0.11.6.bin"
	melange_cmd_bin := exec.Command("curl", "-L", melange_bin_url, "-o", "binary.bin")
	err = melange_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	melange_src_url := "https://github.com/chainguard-dev/melange/archive/refs/tags/v0.11.6.src.tar.gz"
	melange_cmd_src := exec.Command("curl", "-L", melange_src_url, "-o", "source.tar.gz")
	err = melange_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	melange_cmd_direct := exec.Command("./binary")
	err = melange_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
