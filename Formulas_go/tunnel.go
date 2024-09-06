package main

// Tunnel - Expose local servers to the internet securely
// Homepage: https://github.com/labstack/tunnel-client

import (
	"fmt"
	
	"os/exec"
)

func installTunnel() {
	// Método 1: Descargar y extraer .tar.gz
	tunnel_tar_url := "https://github.com/labstack/tunnel-client/archive/refs/tags/v0.5.15.tar.gz"
	tunnel_cmd_tar := exec.Command("curl", "-L", tunnel_tar_url, "-o", "package.tar.gz")
	err := tunnel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tunnel_zip_url := "https://github.com/labstack/tunnel-client/archive/refs/tags/v0.5.15.zip"
	tunnel_cmd_zip := exec.Command("curl", "-L", tunnel_zip_url, "-o", "package.zip")
	err = tunnel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tunnel_bin_url := "https://github.com/labstack/tunnel-client/archive/refs/tags/v0.5.15.bin"
	tunnel_cmd_bin := exec.Command("curl", "-L", tunnel_bin_url, "-o", "binary.bin")
	err = tunnel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tunnel_src_url := "https://github.com/labstack/tunnel-client/archive/refs/tags/v0.5.15.src.tar.gz"
	tunnel_cmd_src := exec.Command("curl", "-L", tunnel_src_url, "-o", "source.tar.gz")
	err = tunnel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tunnel_cmd_direct := exec.Command("./binary")
	err = tunnel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
