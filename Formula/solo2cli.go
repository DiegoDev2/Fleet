package main

// Solo2Cli - CLI to update and use Solo 2 security keys
// Homepage: https://solokeys.com/

import (
	"fmt"
	
	"os/exec"
)

func installSolo2Cli() {
	// Método 1: Descargar y extraer .tar.gz
	solo2cli_tar_url := "https://github.com/solokeys/solo2-cli/archive/refs/tags/v0.2.2.tar.gz"
	solo2cli_cmd_tar := exec.Command("curl", "-L", solo2cli_tar_url, "-o", "package.tar.gz")
	err := solo2cli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	solo2cli_zip_url := "https://github.com/solokeys/solo2-cli/archive/refs/tags/v0.2.2.zip"
	solo2cli_cmd_zip := exec.Command("curl", "-L", solo2cli_zip_url, "-o", "package.zip")
	err = solo2cli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	solo2cli_bin_url := "https://github.com/solokeys/solo2-cli/archive/refs/tags/v0.2.2.bin"
	solo2cli_cmd_bin := exec.Command("curl", "-L", solo2cli_bin_url, "-o", "binary.bin")
	err = solo2cli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	solo2cli_src_url := "https://github.com/solokeys/solo2-cli/archive/refs/tags/v0.2.2.src.tar.gz"
	solo2cli_cmd_src := exec.Command("curl", "-L", solo2cli_src_url, "-o", "source.tar.gz")
	err = solo2cli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	solo2cli_cmd_direct := exec.Command("./binary")
	err = solo2cli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: pcsc-lite")
	exec.Command("latte", "install", "pcsc-lite").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}
