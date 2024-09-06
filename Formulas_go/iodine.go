package main

// Iodine - Tunnel IPv4 traffic through a DNS server
// Homepage: https://code.kryo.se/iodine

import (
	"fmt"
	
	"os/exec"
)

func installIodine() {
	// Método 1: Descargar y extraer .tar.gz
	iodine_tar_url := "https://github.com/yarrick/iodine/archive/refs/tags/v0.8.0.tar.gz"
	iodine_cmd_tar := exec.Command("curl", "-L", iodine_tar_url, "-o", "package.tar.gz")
	err := iodine_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iodine_zip_url := "https://github.com/yarrick/iodine/archive/refs/tags/v0.8.0.zip"
	iodine_cmd_zip := exec.Command("curl", "-L", iodine_zip_url, "-o", "package.zip")
	err = iodine_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iodine_bin_url := "https://github.com/yarrick/iodine/archive/refs/tags/v0.8.0.bin"
	iodine_cmd_bin := exec.Command("curl", "-L", iodine_bin_url, "-o", "binary.bin")
	err = iodine_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iodine_src_url := "https://github.com/yarrick/iodine/archive/refs/tags/v0.8.0.src.tar.gz"
	iodine_cmd_src := exec.Command("curl", "-L", iodine_src_url, "-o", "source.tar.gz")
	err = iodine_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iodine_cmd_direct := exec.Command("./binary")
	err = iodine_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
