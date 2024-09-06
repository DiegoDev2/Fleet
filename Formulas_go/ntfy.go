package main

// Ntfy - Send push notifications to your phone or desktop via PUT/POST
// Homepage: https://ntfy.sh/

import (
	"fmt"
	
	"os/exec"
)

func installNtfy() {
	// Método 1: Descargar y extraer .tar.gz
	ntfy_tar_url := "https://github.com/binwiederhier/ntfy/archive/refs/tags/v2.11.0.tar.gz"
	ntfy_cmd_tar := exec.Command("curl", "-L", ntfy_tar_url, "-o", "package.tar.gz")
	err := ntfy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ntfy_zip_url := "https://github.com/binwiederhier/ntfy/archive/refs/tags/v2.11.0.zip"
	ntfy_cmd_zip := exec.Command("curl", "-L", ntfy_zip_url, "-o", "package.zip")
	err = ntfy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ntfy_bin_url := "https://github.com/binwiederhier/ntfy/archive/refs/tags/v2.11.0.bin"
	ntfy_cmd_bin := exec.Command("curl", "-L", ntfy_bin_url, "-o", "binary.bin")
	err = ntfy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ntfy_src_url := "https://github.com/binwiederhier/ntfy/archive/refs/tags/v2.11.0.src.tar.gz"
	ntfy_cmd_src := exec.Command("curl", "-L", ntfy_src_url, "-o", "source.tar.gz")
	err = ntfy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ntfy_cmd_direct := exec.Command("./binary")
	err = ntfy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
