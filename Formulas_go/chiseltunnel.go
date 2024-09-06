package main

// ChiselTunnel - Fast TCP/UDP tunnel over HTTP
// Homepage: https://github.com/jpillora/chisel

import (
	"fmt"
	
	"os/exec"
)

func installChiselTunnel() {
	// Método 1: Descargar y extraer .tar.gz
	chiseltunnel_tar_url := "https://github.com/jpillora/chisel/archive/refs/tags/v1.10.0.tar.gz"
	chiseltunnel_cmd_tar := exec.Command("curl", "-L", chiseltunnel_tar_url, "-o", "package.tar.gz")
	err := chiseltunnel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chiseltunnel_zip_url := "https://github.com/jpillora/chisel/archive/refs/tags/v1.10.0.zip"
	chiseltunnel_cmd_zip := exec.Command("curl", "-L", chiseltunnel_zip_url, "-o", "package.zip")
	err = chiseltunnel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chiseltunnel_bin_url := "https://github.com/jpillora/chisel/archive/refs/tags/v1.10.0.bin"
	chiseltunnel_cmd_bin := exec.Command("curl", "-L", chiseltunnel_bin_url, "-o", "binary.bin")
	err = chiseltunnel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chiseltunnel_src_url := "https://github.com/jpillora/chisel/archive/refs/tags/v1.10.0.src.tar.gz"
	chiseltunnel_cmd_src := exec.Command("curl", "-L", chiseltunnel_src_url, "-o", "source.tar.gz")
	err = chiseltunnel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chiseltunnel_cmd_direct := exec.Command("./binary")
	err = chiseltunnel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
