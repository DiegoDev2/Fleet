package main

// Blackbox - Safely store secrets in Git/Mercurial/Subversion
// Homepage: https://github.com/StackExchange/blackbox

import (
	"fmt"
	
	"os/exec"
)

func installBlackbox() {
	// Método 1: Descargar y extraer .tar.gz
	blackbox_tar_url := "https://github.com/StackExchange/blackbox/archive/refs/tags/v1.20220610.tar.gz"
	blackbox_cmd_tar := exec.Command("curl", "-L", blackbox_tar_url, "-o", "package.tar.gz")
	err := blackbox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	blackbox_zip_url := "https://github.com/StackExchange/blackbox/archive/refs/tags/v1.20220610.zip"
	blackbox_cmd_zip := exec.Command("curl", "-L", blackbox_zip_url, "-o", "package.zip")
	err = blackbox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	blackbox_bin_url := "https://github.com/StackExchange/blackbox/archive/refs/tags/v1.20220610.bin"
	blackbox_cmd_bin := exec.Command("curl", "-L", blackbox_bin_url, "-o", "binary.bin")
	err = blackbox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	blackbox_src_url := "https://github.com/StackExchange/blackbox/archive/refs/tags/v1.20220610.src.tar.gz"
	blackbox_cmd_src := exec.Command("curl", "-L", blackbox_src_url, "-o", "source.tar.gz")
	err = blackbox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	blackbox_cmd_direct := exec.Command("./binary")
	err = blackbox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnupg")
exec.Command("latte", "install", "gnupg")
}
