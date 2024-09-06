package main

// Vfox - Version manager with support for Java, Node.js, Flutter, .NET & more
// Homepage: https://vfox.lhan.me

import (
	"fmt"
	
	"os/exec"
)

func installVfox() {
	// Método 1: Descargar y extraer .tar.gz
	vfox_tar_url := "https://github.com/version-fox/vfox/archive/refs/tags/v0.5.4.tar.gz"
	vfox_cmd_tar := exec.Command("curl", "-L", vfox_tar_url, "-o", "package.tar.gz")
	err := vfox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vfox_zip_url := "https://github.com/version-fox/vfox/archive/refs/tags/v0.5.4.zip"
	vfox_cmd_zip := exec.Command("curl", "-L", vfox_zip_url, "-o", "package.zip")
	err = vfox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vfox_bin_url := "https://github.com/version-fox/vfox/archive/refs/tags/v0.5.4.bin"
	vfox_cmd_bin := exec.Command("curl", "-L", vfox_bin_url, "-o", "binary.bin")
	err = vfox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vfox_src_url := "https://github.com/version-fox/vfox/archive/refs/tags/v0.5.4.src.tar.gz"
	vfox_cmd_src := exec.Command("curl", "-L", vfox_src_url, "-o", "source.tar.gz")
	err = vfox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vfox_cmd_direct := exec.Command("./binary")
	err = vfox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
