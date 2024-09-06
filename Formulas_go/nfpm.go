package main

// Nfpm - Simple deb and rpm packager
// Homepage: https://nfpm.goreleaser.com/

import (
	"fmt"
	
	"os/exec"
)

func installNfpm() {
	// Método 1: Descargar y extraer .tar.gz
	nfpm_tar_url := "https://github.com/goreleaser/nfpm/archive/refs/tags/v2.39.0.tar.gz"
	nfpm_cmd_tar := exec.Command("curl", "-L", nfpm_tar_url, "-o", "package.tar.gz")
	err := nfpm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nfpm_zip_url := "https://github.com/goreleaser/nfpm/archive/refs/tags/v2.39.0.zip"
	nfpm_cmd_zip := exec.Command("curl", "-L", nfpm_zip_url, "-o", "package.zip")
	err = nfpm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nfpm_bin_url := "https://github.com/goreleaser/nfpm/archive/refs/tags/v2.39.0.bin"
	nfpm_cmd_bin := exec.Command("curl", "-L", nfpm_bin_url, "-o", "binary.bin")
	err = nfpm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nfpm_src_url := "https://github.com/goreleaser/nfpm/archive/refs/tags/v2.39.0.src.tar.gz"
	nfpm_cmd_src := exec.Command("curl", "-L", nfpm_src_url, "-o", "source.tar.gz")
	err = nfpm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nfpm_cmd_direct := exec.Command("./binary")
	err = nfpm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
