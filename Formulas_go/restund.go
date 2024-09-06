package main

// Restund - Modular STUN/TURN server
// Homepage: https://web.archive.org/web/20200427184619/www.creytiv.com/restund.html

import (
	"fmt"
	
	"os/exec"
)

func installRestund() {
	// Método 1: Descargar y extraer .tar.gz
	restund_tar_url := "https://sources.openwrt.org/restund-0.4.12.tar.gz"
	restund_cmd_tar := exec.Command("curl", "-L", restund_tar_url, "-o", "package.tar.gz")
	err := restund_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	restund_zip_url := "https://sources.openwrt.org/restund-0.4.12.zip"
	restund_cmd_zip := exec.Command("curl", "-L", restund_zip_url, "-o", "package.zip")
	err = restund_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	restund_bin_url := "https://sources.openwrt.org/restund-0.4.12.bin"
	restund_cmd_bin := exec.Command("curl", "-L", restund_bin_url, "-o", "binary.bin")
	err = restund_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	restund_src_url := "https://sources.openwrt.org/restund-0.4.12.src.tar.gz"
	restund_cmd_src := exec.Command("curl", "-L", restund_src_url, "-o", "source.tar.gz")
	err = restund_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	restund_cmd_direct := exec.Command("./binary")
	err = restund_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libre")
exec.Command("latte", "install", "libre")
}
