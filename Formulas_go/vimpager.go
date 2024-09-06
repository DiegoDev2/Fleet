package main

// Vimpager - Use ViM as PAGER
// Homepage: https://github.com/rkitover/vimpager

import (
	"fmt"
	
	"os/exec"
)

func installVimpager() {
	// Método 1: Descargar y extraer .tar.gz
	vimpager_tar_url := "https://github.com/rkitover/vimpager/archive/refs/tags/2.06.tar.gz"
	vimpager_cmd_tar := exec.Command("curl", "-L", vimpager_tar_url, "-o", "package.tar.gz")
	err := vimpager_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vimpager_zip_url := "https://github.com/rkitover/vimpager/archive/refs/tags/2.06.zip"
	vimpager_cmd_zip := exec.Command("curl", "-L", vimpager_zip_url, "-o", "package.zip")
	err = vimpager_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vimpager_bin_url := "https://github.com/rkitover/vimpager/archive/refs/tags/2.06.bin"
	vimpager_cmd_bin := exec.Command("curl", "-L", vimpager_bin_url, "-o", "binary.bin")
	err = vimpager_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vimpager_src_url := "https://github.com/rkitover/vimpager/archive/refs/tags/2.06.src.tar.gz"
	vimpager_cmd_src := exec.Command("curl", "-L", vimpager_src_url, "-o", "source.tar.gz")
	err = vimpager_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vimpager_cmd_direct := exec.Command("./binary")
	err = vimpager_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
}
