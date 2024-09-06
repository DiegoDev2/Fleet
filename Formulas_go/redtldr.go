package main

// RedTldr - Used to help red team staff quickly find the commands and key points
// Homepage: https://payloads.online/red-tldr/

import (
	"fmt"
	
	"os/exec"
)

func installRedTldr() {
	// Método 1: Descargar y extraer .tar.gz
	redtldr_tar_url := "https://github.com/Rvn0xsy/red-tldr/archive/refs/tags/v0.4.3.tar.gz"
	redtldr_cmd_tar := exec.Command("curl", "-L", redtldr_tar_url, "-o", "package.tar.gz")
	err := redtldr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	redtldr_zip_url := "https://github.com/Rvn0xsy/red-tldr/archive/refs/tags/v0.4.3.zip"
	redtldr_cmd_zip := exec.Command("curl", "-L", redtldr_zip_url, "-o", "package.zip")
	err = redtldr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	redtldr_bin_url := "https://github.com/Rvn0xsy/red-tldr/archive/refs/tags/v0.4.3.bin"
	redtldr_cmd_bin := exec.Command("curl", "-L", redtldr_bin_url, "-o", "binary.bin")
	err = redtldr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	redtldr_src_url := "https://github.com/Rvn0xsy/red-tldr/archive/refs/tags/v0.4.3.src.tar.gz"
	redtldr_cmd_src := exec.Command("curl", "-L", redtldr_src_url, "-o", "source.tar.gz")
	err = redtldr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	redtldr_cmd_direct := exec.Command("./binary")
	err = redtldr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
