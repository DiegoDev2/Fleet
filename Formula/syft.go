package main

// Syft - CLI for generating a Software Bill of Materials from container images
// Homepage: https://github.com/anchore/syft

import (
	"fmt"
	
	"os/exec"
)

func installSyft() {
	// Método 1: Descargar y extraer .tar.gz
	syft_tar_url := "https://github.com/anchore/syft/archive/refs/tags/v1.11.1.tar.gz"
	syft_cmd_tar := exec.Command("curl", "-L", syft_tar_url, "-o", "package.tar.gz")
	err := syft_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	syft_zip_url := "https://github.com/anchore/syft/archive/refs/tags/v1.11.1.zip"
	syft_cmd_zip := exec.Command("curl", "-L", syft_zip_url, "-o", "package.zip")
	err = syft_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	syft_bin_url := "https://github.com/anchore/syft/archive/refs/tags/v1.11.1.bin"
	syft_cmd_bin := exec.Command("curl", "-L", syft_bin_url, "-o", "binary.bin")
	err = syft_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	syft_src_url := "https://github.com/anchore/syft/archive/refs/tags/v1.11.1.src.tar.gz"
	syft_cmd_src := exec.Command("curl", "-L", syft_src_url, "-o", "source.tar.gz")
	err = syft_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	syft_cmd_direct := exec.Command("./binary")
	err = syft_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
