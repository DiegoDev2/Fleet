package main

// Shell2http - Executing shell commands via HTTP server
// Homepage: https://github.com/msoap/shell2http

import (
	"fmt"
	
	"os/exec"
)

func installShell2http() {
	// Método 1: Descargar y extraer .tar.gz
	shell2http_tar_url := "https://github.com/msoap/shell2http/archive/refs/tags/v1.17.0.tar.gz"
	shell2http_cmd_tar := exec.Command("curl", "-L", shell2http_tar_url, "-o", "package.tar.gz")
	err := shell2http_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shell2http_zip_url := "https://github.com/msoap/shell2http/archive/refs/tags/v1.17.0.zip"
	shell2http_cmd_zip := exec.Command("curl", "-L", shell2http_zip_url, "-o", "package.zip")
	err = shell2http_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shell2http_bin_url := "https://github.com/msoap/shell2http/archive/refs/tags/v1.17.0.bin"
	shell2http_cmd_bin := exec.Command("curl", "-L", shell2http_bin_url, "-o", "binary.bin")
	err = shell2http_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shell2http_src_url := "https://github.com/msoap/shell2http/archive/refs/tags/v1.17.0.src.tar.gz"
	shell2http_cmd_src := exec.Command("curl", "-L", shell2http_src_url, "-o", "source.tar.gz")
	err = shell2http_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shell2http_cmd_direct := exec.Command("./binary")
	err = shell2http_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
