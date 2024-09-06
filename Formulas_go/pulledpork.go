package main

// Pulledpork - Snort rule management
// Homepage: https://github.com/shirkdog/pulledpork

import (
	"fmt"
	
	"os/exec"
)

func installPulledpork() {
	// Método 1: Descargar y extraer .tar.gz
	pulledpork_tar_url := "https://github.com/shirkdog/pulledpork/archive/refs/tags/v0.7.4.tar.gz"
	pulledpork_cmd_tar := exec.Command("curl", "-L", pulledpork_tar_url, "-o", "package.tar.gz")
	err := pulledpork_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pulledpork_zip_url := "https://github.com/shirkdog/pulledpork/archive/refs/tags/v0.7.4.zip"
	pulledpork_cmd_zip := exec.Command("curl", "-L", pulledpork_zip_url, "-o", "package.zip")
	err = pulledpork_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pulledpork_bin_url := "https://github.com/shirkdog/pulledpork/archive/refs/tags/v0.7.4.bin"
	pulledpork_cmd_bin := exec.Command("curl", "-L", pulledpork_bin_url, "-o", "binary.bin")
	err = pulledpork_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pulledpork_src_url := "https://github.com/shirkdog/pulledpork/archive/refs/tags/v0.7.4.src.tar.gz"
	pulledpork_cmd_src := exec.Command("curl", "-L", pulledpork_src_url, "-o", "source.tar.gz")
	err = pulledpork_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pulledpork_cmd_direct := exec.Command("./binary")
	err = pulledpork_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
