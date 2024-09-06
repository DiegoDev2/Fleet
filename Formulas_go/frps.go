package main

// Frps - Server app of fast reverse proxy to expose a local server to the internet
// Homepage: https://github.com/fatedier/frp

import (
	"fmt"
	
	"os/exec"
)

func installFrps() {
	// Método 1: Descargar y extraer .tar.gz
	frps_tar_url := "https://github.com/fatedier/frp/archive/refs/tags/v0.60.0.tar.gz"
	frps_cmd_tar := exec.Command("curl", "-L", frps_tar_url, "-o", "package.tar.gz")
	err := frps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	frps_zip_url := "https://github.com/fatedier/frp/archive/refs/tags/v0.60.0.zip"
	frps_cmd_zip := exec.Command("curl", "-L", frps_zip_url, "-o", "package.zip")
	err = frps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	frps_bin_url := "https://github.com/fatedier/frp/archive/refs/tags/v0.60.0.bin"
	frps_cmd_bin := exec.Command("curl", "-L", frps_bin_url, "-o", "binary.bin")
	err = frps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	frps_src_url := "https://github.com/fatedier/frp/archive/refs/tags/v0.60.0.src.tar.gz"
	frps_cmd_src := exec.Command("curl", "-L", frps_src_url, "-o", "source.tar.gz")
	err = frps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	frps_cmd_direct := exec.Command("./binary")
	err = frps_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
