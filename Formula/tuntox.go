package main

// Tuntox - Tunnel TCP connections over the Tox protocol
// Homepage: https://gdr.name/tuntox/

import (
	"fmt"
	
	"os/exec"
)

func installTuntox() {
	// Método 1: Descargar y extraer .tar.gz
	tuntox_tar_url := "https://github.com/gjedeer/tuntox/archive/refs/tags/0.0.10.1.tar.gz"
	tuntox_cmd_tar := exec.Command("curl", "-L", tuntox_tar_url, "-o", "package.tar.gz")
	err := tuntox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tuntox_zip_url := "https://github.com/gjedeer/tuntox/archive/refs/tags/0.0.10.1.zip"
	tuntox_cmd_zip := exec.Command("curl", "-L", tuntox_zip_url, "-o", "package.zip")
	err = tuntox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tuntox_bin_url := "https://github.com/gjedeer/tuntox/archive/refs/tags/0.0.10.1.bin"
	tuntox_cmd_bin := exec.Command("curl", "-L", tuntox_bin_url, "-o", "binary.bin")
	err = tuntox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tuntox_src_url := "https://github.com/gjedeer/tuntox/archive/refs/tags/0.0.10.1.src.tar.gz"
	tuntox_cmd_src := exec.Command("curl", "-L", tuntox_src_url, "-o", "source.tar.gz")
	err = tuntox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tuntox_cmd_direct := exec.Command("./binary")
	err = tuntox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cscope")
	exec.Command("latte", "install", "cscope").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: toxcore")
	exec.Command("latte", "install", "toxcore").Run()
}
