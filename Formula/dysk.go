package main

// Dysk - Linux utility to get information on filesystems, like df but better
// Homepage: https://dystroy.org/dysk/

import (
	"fmt"
	
	"os/exec"
)

func installDysk() {
	// Método 1: Descargar y extraer .tar.gz
	dysk_tar_url := "https://github.com/Canop/dysk/archive/refs/tags/v2.9.0.tar.gz"
	dysk_cmd_tar := exec.Command("curl", "-L", dysk_tar_url, "-o", "package.tar.gz")
	err := dysk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dysk_zip_url := "https://github.com/Canop/dysk/archive/refs/tags/v2.9.0.zip"
	dysk_cmd_zip := exec.Command("curl", "-L", dysk_zip_url, "-o", "package.zip")
	err = dysk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dysk_bin_url := "https://github.com/Canop/dysk/archive/refs/tags/v2.9.0.bin"
	dysk_cmd_bin := exec.Command("curl", "-L", dysk_bin_url, "-o", "binary.bin")
	err = dysk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dysk_src_url := "https://github.com/Canop/dysk/archive/refs/tags/v2.9.0.src.tar.gz"
	dysk_cmd_src := exec.Command("curl", "-L", dysk_src_url, "-o", "source.tar.gz")
	err = dysk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dysk_cmd_direct := exec.Command("./binary")
	err = dysk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
