package main

// BitwardenCli - Secure and free password manager for all of your devices
// Homepage: https://bitwarden.com/

import (
	"fmt"
	
	"os/exec"
)

func installBitwardenCli() {
	// Método 1: Descargar y extraer .tar.gz
	bitwardencli_tar_url := "https://github.com/bitwarden/clients/archive/refs/tags/cli-v2024.8.2.tar.gz"
	bitwardencli_cmd_tar := exec.Command("curl", "-L", bitwardencli_tar_url, "-o", "package.tar.gz")
	err := bitwardencli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bitwardencli_zip_url := "https://github.com/bitwarden/clients/archive/refs/tags/cli-v2024.8.2.zip"
	bitwardencli_cmd_zip := exec.Command("curl", "-L", bitwardencli_zip_url, "-o", "package.zip")
	err = bitwardencli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bitwardencli_bin_url := "https://github.com/bitwarden/clients/archive/refs/tags/cli-v2024.8.2.bin"
	bitwardencli_cmd_bin := exec.Command("curl", "-L", bitwardencli_bin_url, "-o", "binary.bin")
	err = bitwardencli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bitwardencli_src_url := "https://github.com/bitwarden/clients/archive/refs/tags/cli-v2024.8.2.src.tar.gz"
	bitwardencli_cmd_src := exec.Command("curl", "-L", bitwardencli_src_url, "-o", "source.tar.gz")
	err = bitwardencli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bitwardencli_cmd_direct := exec.Command("./binary")
	err = bitwardencli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
