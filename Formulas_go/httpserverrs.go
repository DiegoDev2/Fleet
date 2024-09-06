package main

// HttpServerRs - Simple and configurable command-line HTTP server
// Homepage: https://github.com/http-server-rs/http-server

import (
	"fmt"
	
	"os/exec"
)

func installHttpServerRs() {
	// Método 1: Descargar y extraer .tar.gz
	httpserverrs_tar_url := "https://github.com/http-server-rs/http-server/archive/refs/tags/v0.8.9.tar.gz"
	httpserverrs_cmd_tar := exec.Command("curl", "-L", httpserverrs_tar_url, "-o", "package.tar.gz")
	err := httpserverrs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	httpserverrs_zip_url := "https://github.com/http-server-rs/http-server/archive/refs/tags/v0.8.9.zip"
	httpserverrs_cmd_zip := exec.Command("curl", "-L", httpserverrs_zip_url, "-o", "package.zip")
	err = httpserverrs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	httpserverrs_bin_url := "https://github.com/http-server-rs/http-server/archive/refs/tags/v0.8.9.bin"
	httpserverrs_cmd_bin := exec.Command("curl", "-L", httpserverrs_bin_url, "-o", "binary.bin")
	err = httpserverrs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	httpserverrs_src_url := "https://github.com/http-server-rs/http-server/archive/refs/tags/v0.8.9.src.tar.gz"
	httpserverrs_cmd_src := exec.Command("curl", "-L", httpserverrs_src_url, "-o", "source.tar.gz")
	err = httpserverrs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	httpserverrs_cmd_direct := exec.Command("./binary")
	err = httpserverrs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
