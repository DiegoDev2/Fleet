package main

// Kcptun - Stable & Secure Tunnel based on KCP with N:M multiplexing and FEC
// Homepage: https://github.com/xtaci/kcptun

import (
	"fmt"
	
	"os/exec"
)

func installKcptun() {
	// Método 1: Descargar y extraer .tar.gz
	kcptun_tar_url := "https://github.com/xtaci/kcptun/archive/refs/tags/v20240831.tar.gz"
	kcptun_cmd_tar := exec.Command("curl", "-L", kcptun_tar_url, "-o", "package.tar.gz")
	err := kcptun_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kcptun_zip_url := "https://github.com/xtaci/kcptun/archive/refs/tags/v20240831.zip"
	kcptun_cmd_zip := exec.Command("curl", "-L", kcptun_zip_url, "-o", "package.zip")
	err = kcptun_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kcptun_bin_url := "https://github.com/xtaci/kcptun/archive/refs/tags/v20240831.bin"
	kcptun_cmd_bin := exec.Command("curl", "-L", kcptun_bin_url, "-o", "binary.bin")
	err = kcptun_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kcptun_src_url := "https://github.com/xtaci/kcptun/archive/refs/tags/v20240831.src.tar.gz"
	kcptun_cmd_src := exec.Command("curl", "-L", kcptun_src_url, "-o", "source.tar.gz")
	err = kcptun_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kcptun_cmd_direct := exec.Command("./binary")
	err = kcptun_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
