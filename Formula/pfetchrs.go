package main

// PfetchRs - Pretty system information tool written in Rust
// Homepage: https://github.com/Gobidev/pfetch-rs

import (
	"fmt"
	
	"os/exec"
)

func installPfetchRs() {
	// Método 1: Descargar y extraer .tar.gz
	pfetchrs_tar_url := "https://github.com/Gobidev/pfetch-rs/archive/refs/tags/v2.11.0.tar.gz"
	pfetchrs_cmd_tar := exec.Command("curl", "-L", pfetchrs_tar_url, "-o", "package.tar.gz")
	err := pfetchrs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pfetchrs_zip_url := "https://github.com/Gobidev/pfetch-rs/archive/refs/tags/v2.11.0.zip"
	pfetchrs_cmd_zip := exec.Command("curl", "-L", pfetchrs_zip_url, "-o", "package.zip")
	err = pfetchrs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pfetchrs_bin_url := "https://github.com/Gobidev/pfetch-rs/archive/refs/tags/v2.11.0.bin"
	pfetchrs_cmd_bin := exec.Command("curl", "-L", pfetchrs_bin_url, "-o", "binary.bin")
	err = pfetchrs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pfetchrs_src_url := "https://github.com/Gobidev/pfetch-rs/archive/refs/tags/v2.11.0.src.tar.gz"
	pfetchrs_cmd_src := exec.Command("curl", "-L", pfetchrs_src_url, "-o", "source.tar.gz")
	err = pfetchrs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pfetchrs_cmd_direct := exec.Command("./binary")
	err = pfetchrs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
