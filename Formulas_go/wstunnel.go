package main

// Wstunnel - Tunnel all your traffic over Websocket or HTTP2
// Homepage: https://github.com/erebe/wstunnel

import (
	"fmt"
	
	"os/exec"
)

func installWstunnel() {
	// Método 1: Descargar y extraer .tar.gz
	wstunnel_tar_url := "https://github.com/erebe/wstunnel/archive/refs/tags/v10.1.1.tar.gz"
	wstunnel_cmd_tar := exec.Command("curl", "-L", wstunnel_tar_url, "-o", "package.tar.gz")
	err := wstunnel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wstunnel_zip_url := "https://github.com/erebe/wstunnel/archive/refs/tags/v10.1.1.zip"
	wstunnel_cmd_zip := exec.Command("curl", "-L", wstunnel_zip_url, "-o", "package.zip")
	err = wstunnel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wstunnel_bin_url := "https://github.com/erebe/wstunnel/archive/refs/tags/v10.1.1.bin"
	wstunnel_cmd_bin := exec.Command("curl", "-L", wstunnel_bin_url, "-o", "binary.bin")
	err = wstunnel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wstunnel_src_url := "https://github.com/erebe/wstunnel/archive/refs/tags/v10.1.1.src.tar.gz"
	wstunnel_cmd_src := exec.Command("curl", "-L", wstunnel_src_url, "-o", "source.tar.gz")
	err = wstunnel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wstunnel_cmd_direct := exec.Command("./binary")
	err = wstunnel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
