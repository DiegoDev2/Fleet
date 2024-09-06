package main

// Hotbuild - Cross platform hot compilation tool for go
// Homepage: https://hotbuild.ffactory.org

import (
	"fmt"
	
	"os/exec"
)

func installHotbuild() {
	// Método 1: Descargar y extraer .tar.gz
	hotbuild_tar_url := "https://github.com/wandercn/hotbuild/archive/refs/tags/v1.0.8.tar.gz"
	hotbuild_cmd_tar := exec.Command("curl", "-L", hotbuild_tar_url, "-o", "package.tar.gz")
	err := hotbuild_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hotbuild_zip_url := "https://github.com/wandercn/hotbuild/archive/refs/tags/v1.0.8.zip"
	hotbuild_cmd_zip := exec.Command("curl", "-L", hotbuild_zip_url, "-o", "package.zip")
	err = hotbuild_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hotbuild_bin_url := "https://github.com/wandercn/hotbuild/archive/refs/tags/v1.0.8.bin"
	hotbuild_cmd_bin := exec.Command("curl", "-L", hotbuild_bin_url, "-o", "binary.bin")
	err = hotbuild_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hotbuild_src_url := "https://github.com/wandercn/hotbuild/archive/refs/tags/v1.0.8.src.tar.gz"
	hotbuild_cmd_src := exec.Command("curl", "-L", hotbuild_src_url, "-o", "source.tar.gz")
	err = hotbuild_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hotbuild_cmd_direct := exec.Command("./binary")
	err = hotbuild_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
