package main

// Macchina - System information fetcher, with an emphasis on performance and minimalism
// Homepage: https://github.com/Macchina-CLI/macchina

import (
	"fmt"
	
	"os/exec"
)

func installMacchina() {
	// Método 1: Descargar y extraer .tar.gz
	macchina_tar_url := "https://github.com/Macchina-CLI/macchina/archive/refs/tags/v6.1.8.tar.gz"
	macchina_cmd_tar := exec.Command("curl", "-L", macchina_tar_url, "-o", "package.tar.gz")
	err := macchina_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	macchina_zip_url := "https://github.com/Macchina-CLI/macchina/archive/refs/tags/v6.1.8.zip"
	macchina_cmd_zip := exec.Command("curl", "-L", macchina_zip_url, "-o", "package.zip")
	err = macchina_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	macchina_bin_url := "https://github.com/Macchina-CLI/macchina/archive/refs/tags/v6.1.8.bin"
	macchina_cmd_bin := exec.Command("curl", "-L", macchina_bin_url, "-o", "binary.bin")
	err = macchina_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	macchina_src_url := "https://github.com/Macchina-CLI/macchina/archive/refs/tags/v6.1.8.src.tar.gz"
	macchina_cmd_src := exec.Command("curl", "-L", macchina_src_url, "-o", "source.tar.gz")
	err = macchina_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	macchina_cmd_direct := exec.Command("./binary")
	err = macchina_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
