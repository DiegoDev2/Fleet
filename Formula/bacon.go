package main

// Bacon - Background rust code check
// Homepage: https://dystroy.org/bacon/

import (
	"fmt"
	
	"os/exec"
)

func installBacon() {
	// Método 1: Descargar y extraer .tar.gz
	bacon_tar_url := "https://github.com/Canop/bacon/archive/refs/tags/v2.20.0.tar.gz"
	bacon_cmd_tar := exec.Command("curl", "-L", bacon_tar_url, "-o", "package.tar.gz")
	err := bacon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bacon_zip_url := "https://github.com/Canop/bacon/archive/refs/tags/v2.20.0.zip"
	bacon_cmd_zip := exec.Command("curl", "-L", bacon_zip_url, "-o", "package.zip")
	err = bacon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bacon_bin_url := "https://github.com/Canop/bacon/archive/refs/tags/v2.20.0.bin"
	bacon_cmd_bin := exec.Command("curl", "-L", bacon_bin_url, "-o", "binary.bin")
	err = bacon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bacon_src_url := "https://github.com/Canop/bacon/archive/refs/tags/v2.20.0.src.tar.gz"
	bacon_cmd_src := exec.Command("curl", "-L", bacon_src_url, "-o", "source.tar.gz")
	err = bacon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bacon_cmd_direct := exec.Command("./binary")
	err = bacon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: rustup")
	exec.Command("latte", "install", "rustup").Run()
}
