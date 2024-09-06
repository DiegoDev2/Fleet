package main

// Assh - Advanced SSH config - Regex, aliases, gateways, includes and dynamic hosts
// Homepage: https://manfred.life/assh

import (
	"fmt"
	
	"os/exec"
)

func installAssh() {
	// Método 1: Descargar y extraer .tar.gz
	assh_tar_url := "https://github.com/moul/assh/archive/refs/tags/v2.16.0.tar.gz"
	assh_cmd_tar := exec.Command("curl", "-L", assh_tar_url, "-o", "package.tar.gz")
	err := assh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	assh_zip_url := "https://github.com/moul/assh/archive/refs/tags/v2.16.0.zip"
	assh_cmd_zip := exec.Command("curl", "-L", assh_zip_url, "-o", "package.zip")
	err = assh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	assh_bin_url := "https://github.com/moul/assh/archive/refs/tags/v2.16.0.bin"
	assh_cmd_bin := exec.Command("curl", "-L", assh_bin_url, "-o", "binary.bin")
	err = assh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	assh_src_url := "https://github.com/moul/assh/archive/refs/tags/v2.16.0.src.tar.gz"
	assh_cmd_src := exec.Command("curl", "-L", assh_src_url, "-o", "source.tar.gz")
	err = assh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	assh_cmd_direct := exec.Command("./binary")
	err = assh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
