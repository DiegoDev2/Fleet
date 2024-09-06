package main

// Pedump - Dump Windows PE files using Ruby
// Homepage: https://pedump.me

import (
	"fmt"
	
	"os/exec"
)

func installPedump() {
	// Método 1: Descargar y extraer .tar.gz
	pedump_tar_url := "https://github.com/zed-0xff/pedump/archive/refs/tags/v0.6.10.tar.gz"
	pedump_cmd_tar := exec.Command("curl", "-L", pedump_tar_url, "-o", "package.tar.gz")
	err := pedump_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pedump_zip_url := "https://github.com/zed-0xff/pedump/archive/refs/tags/v0.6.10.zip"
	pedump_cmd_zip := exec.Command("curl", "-L", pedump_zip_url, "-o", "package.zip")
	err = pedump_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pedump_bin_url := "https://github.com/zed-0xff/pedump/archive/refs/tags/v0.6.10.bin"
	pedump_cmd_bin := exec.Command("curl", "-L", pedump_bin_url, "-o", "binary.bin")
	err = pedump_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pedump_src_url := "https://github.com/zed-0xff/pedump/archive/refs/tags/v0.6.10.src.tar.gz"
	pedump_cmd_src := exec.Command("curl", "-L", pedump_src_url, "-o", "source.tar.gz")
	err = pedump_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pedump_cmd_direct := exec.Command("./binary")
	err = pedump_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby")
exec.Command("latte", "install", "ruby")
}
