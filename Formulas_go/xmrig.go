package main

// Xmrig - Monero (XMR) CPU miner
// Homepage: https://github.com/xmrig/xmrig

import (
	"fmt"
	
	"os/exec"
)

func installXmrig() {
	// Método 1: Descargar y extraer .tar.gz
	xmrig_tar_url := "https://github.com/xmrig/xmrig/archive/refs/tags/v6.22.0.tar.gz"
	xmrig_cmd_tar := exec.Command("curl", "-L", xmrig_tar_url, "-o", "package.tar.gz")
	err := xmrig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xmrig_zip_url := "https://github.com/xmrig/xmrig/archive/refs/tags/v6.22.0.zip"
	xmrig_cmd_zip := exec.Command("curl", "-L", xmrig_zip_url, "-o", "package.zip")
	err = xmrig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xmrig_bin_url := "https://github.com/xmrig/xmrig/archive/refs/tags/v6.22.0.bin"
	xmrig_cmd_bin := exec.Command("curl", "-L", xmrig_bin_url, "-o", "binary.bin")
	err = xmrig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xmrig_src_url := "https://github.com/xmrig/xmrig/archive/refs/tags/v6.22.0.src.tar.gz"
	xmrig_cmd_src := exec.Command("curl", "-L", xmrig_src_url, "-o", "source.tar.gz")
	err = xmrig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xmrig_cmd_direct := exec.Command("./binary")
	err = xmrig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: hwloc")
exec.Command("latte", "install", "hwloc")
	fmt.Println("Instalando dependencia: libuv")
exec.Command("latte", "install", "libuv")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
