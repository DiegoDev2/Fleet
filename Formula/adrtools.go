package main

// AdrTools - CLI tool for working with Architecture Decision Records
// Homepage: https://github.com/npryce/adr-tools

import (
	"fmt"
	
	"os/exec"
)

func installAdrTools() {
	// Método 1: Descargar y extraer .tar.gz
	adrtools_tar_url := "https://github.com/npryce/adr-tools/archive/refs/tags/3.0.0.tar.gz"
	adrtools_cmd_tar := exec.Command("curl", "-L", adrtools_tar_url, "-o", "package.tar.gz")
	err := adrtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	adrtools_zip_url := "https://github.com/npryce/adr-tools/archive/refs/tags/3.0.0.zip"
	adrtools_cmd_zip := exec.Command("curl", "-L", adrtools_zip_url, "-o", "package.zip")
	err = adrtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	adrtools_bin_url := "https://github.com/npryce/adr-tools/archive/refs/tags/3.0.0.bin"
	adrtools_cmd_bin := exec.Command("curl", "-L", adrtools_bin_url, "-o", "binary.bin")
	err = adrtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	adrtools_src_url := "https://github.com/npryce/adr-tools/archive/refs/tags/3.0.0.src.tar.gz"
	adrtools_cmd_src := exec.Command("curl", "-L", adrtools_src_url, "-o", "source.tar.gz")
	err = adrtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	adrtools_cmd_direct := exec.Command("./binary")
	err = adrtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
