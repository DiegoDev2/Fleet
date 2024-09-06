package main

// Confd - Manage local application configuration files using templates
// Homepage: https://github.com/kelseyhightower/confd

import (
	"fmt"
	
	"os/exec"
)

func installConfd() {
	// Método 1: Descargar y extraer .tar.gz
	confd_tar_url := "https://github.com/kelseyhightower/confd/archive/refs/tags/v0.16.0.tar.gz"
	confd_cmd_tar := exec.Command("curl", "-L", confd_tar_url, "-o", "package.tar.gz")
	err := confd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	confd_zip_url := "https://github.com/kelseyhightower/confd/archive/refs/tags/v0.16.0.zip"
	confd_cmd_zip := exec.Command("curl", "-L", confd_zip_url, "-o", "package.zip")
	err = confd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	confd_bin_url := "https://github.com/kelseyhightower/confd/archive/refs/tags/v0.16.0.bin"
	confd_cmd_bin := exec.Command("curl", "-L", confd_bin_url, "-o", "binary.bin")
	err = confd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	confd_src_url := "https://github.com/kelseyhightower/confd/archive/refs/tags/v0.16.0.src.tar.gz"
	confd_cmd_src := exec.Command("curl", "-L", confd_src_url, "-o", "source.tar.gz")
	err = confd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	confd_cmd_direct := exec.Command("./binary")
	err = confd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
