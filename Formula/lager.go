package main

// Lager - C++ lib for value-oriented design using unidirectional data-flow architecture
// Homepage: https://sinusoid.es/lager/

import (
	"fmt"
	
	"os/exec"
)

func installLager() {
	// Método 1: Descargar y extraer .tar.gz
	lager_tar_url := "https://github.com/arximboldi/lager/archive/refs/tags/v0.1.1.tar.gz"
	lager_cmd_tar := exec.Command("curl", "-L", lager_tar_url, "-o", "package.tar.gz")
	err := lager_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lager_zip_url := "https://github.com/arximboldi/lager/archive/refs/tags/v0.1.1.zip"
	lager_cmd_zip := exec.Command("curl", "-L", lager_zip_url, "-o", "package.zip")
	err = lager_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lager_bin_url := "https://github.com/arximboldi/lager/archive/refs/tags/v0.1.1.bin"
	lager_cmd_bin := exec.Command("curl", "-L", lager_bin_url, "-o", "binary.bin")
	err = lager_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lager_src_url := "https://github.com/arximboldi/lager/archive/refs/tags/v0.1.1.src.tar.gz"
	lager_cmd_src := exec.Command("curl", "-L", lager_src_url, "-o", "source.tar.gz")
	err = lager_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lager_cmd_direct := exec.Command("./binary")
	err = lager_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: zug")
	exec.Command("latte", "install", "zug").Run()
}
