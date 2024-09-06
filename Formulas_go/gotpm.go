package main

// Gotpm - CLI for using TPM 2.0
// Homepage: https://github.com/google/go-tpm-tools

import (
	"fmt"
	
	"os/exec"
)

func installGotpm() {
	// Método 1: Descargar y extraer .tar.gz
	gotpm_tar_url := "https://github.com/google/go-tpm-tools/archive/refs/tags/v0.4.4.tar.gz"
	gotpm_cmd_tar := exec.Command("curl", "-L", gotpm_tar_url, "-o", "package.tar.gz")
	err := gotpm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gotpm_zip_url := "https://github.com/google/go-tpm-tools/archive/refs/tags/v0.4.4.zip"
	gotpm_cmd_zip := exec.Command("curl", "-L", gotpm_zip_url, "-o", "package.zip")
	err = gotpm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gotpm_bin_url := "https://github.com/google/go-tpm-tools/archive/refs/tags/v0.4.4.bin"
	gotpm_cmd_bin := exec.Command("curl", "-L", gotpm_bin_url, "-o", "binary.bin")
	err = gotpm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gotpm_src_url := "https://github.com/google/go-tpm-tools/archive/refs/tags/v0.4.4.src.tar.gz"
	gotpm_cmd_src := exec.Command("curl", "-L", gotpm_src_url, "-o", "source.tar.gz")
	err = gotpm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gotpm_cmd_direct := exec.Command("./binary")
	err = gotpm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
