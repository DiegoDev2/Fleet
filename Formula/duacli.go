package main

// DuaCli - View disk space usage and delete unwanted data, fast
// Homepage: https://lib.rs/crates/dua-cli

import (
	"fmt"
	
	"os/exec"
)

func installDuaCli() {
	// Método 1: Descargar y extraer .tar.gz
	duacli_tar_url := "https://github.com/Byron/dua-cli/archive/refs/tags/v2.29.2.tar.gz"
	duacli_cmd_tar := exec.Command("curl", "-L", duacli_tar_url, "-o", "package.tar.gz")
	err := duacli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	duacli_zip_url := "https://github.com/Byron/dua-cli/archive/refs/tags/v2.29.2.zip"
	duacli_cmd_zip := exec.Command("curl", "-L", duacli_zip_url, "-o", "package.zip")
	err = duacli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	duacli_bin_url := "https://github.com/Byron/dua-cli/archive/refs/tags/v2.29.2.bin"
	duacli_cmd_bin := exec.Command("curl", "-L", duacli_bin_url, "-o", "binary.bin")
	err = duacli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	duacli_src_url := "https://github.com/Byron/dua-cli/archive/refs/tags/v2.29.2.src.tar.gz"
	duacli_cmd_src := exec.Command("curl", "-L", duacli_src_url, "-o", "source.tar.gz")
	err = duacli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	duacli_cmd_direct := exec.Command("./binary")
	err = duacli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
