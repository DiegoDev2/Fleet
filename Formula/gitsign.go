package main

// Gitsign - Keyless Git signing using Sigstore
// Homepage: https://github.com/sigstore/gitsign

import (
	"fmt"
	
	"os/exec"
)

func installGitsign() {
	// Método 1: Descargar y extraer .tar.gz
	gitsign_tar_url := "https://github.com/sigstore/gitsign/archive/refs/tags/v0.10.2.tar.gz"
	gitsign_cmd_tar := exec.Command("curl", "-L", gitsign_tar_url, "-o", "package.tar.gz")
	err := gitsign_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitsign_zip_url := "https://github.com/sigstore/gitsign/archive/refs/tags/v0.10.2.zip"
	gitsign_cmd_zip := exec.Command("curl", "-L", gitsign_zip_url, "-o", "package.zip")
	err = gitsign_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitsign_bin_url := "https://github.com/sigstore/gitsign/archive/refs/tags/v0.10.2.bin"
	gitsign_cmd_bin := exec.Command("curl", "-L", gitsign_bin_url, "-o", "binary.bin")
	err = gitsign_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitsign_src_url := "https://github.com/sigstore/gitsign/archive/refs/tags/v0.10.2.src.tar.gz"
	gitsign_cmd_src := exec.Command("curl", "-L", gitsign_src_url, "-o", "source.tar.gz")
	err = gitsign_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitsign_cmd_direct := exec.Command("./binary")
	err = gitsign_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
