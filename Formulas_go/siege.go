package main

// Siege - HTTP regression testing and benchmarking utility
// Homepage: https://www.joedog.org/siege-home/

import (
	"fmt"
	
	"os/exec"
)

func installSiege() {
	// Método 1: Descargar y extraer .tar.gz
	siege_tar_url := "https://download.joedog.org/siege/siege-4.1.6.tar.gz"
	siege_cmd_tar := exec.Command("curl", "-L", siege_tar_url, "-o", "package.tar.gz")
	err := siege_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	siege_zip_url := "https://download.joedog.org/siege/siege-4.1.6.zip"
	siege_cmd_zip := exec.Command("curl", "-L", siege_zip_url, "-o", "package.zip")
	err = siege_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	siege_bin_url := "https://download.joedog.org/siege/siege-4.1.6.bin"
	siege_cmd_bin := exec.Command("curl", "-L", siege_bin_url, "-o", "binary.bin")
	err = siege_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	siege_src_url := "https://download.joedog.org/siege/siege-4.1.6.src.tar.gz"
	siege_cmd_src := exec.Command("curl", "-L", siege_src_url, "-o", "source.tar.gz")
	err = siege_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	siege_cmd_direct := exec.Command("./binary")
	err = siege_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
