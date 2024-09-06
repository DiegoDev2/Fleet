package main

// IosSim - Command-line application launcher for the iOS Simulator
// Homepage: https://github.com/ios-control/ios-sim

import (
	"fmt"
	
	"os/exec"
)

func installIosSim() {
	// Método 1: Descargar y extraer .tar.gz
	iossim_tar_url := "https://github.com/ios-control/ios-sim/archive/refs/tags/9.0.0.tar.gz"
	iossim_cmd_tar := exec.Command("curl", "-L", iossim_tar_url, "-o", "package.tar.gz")
	err := iossim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iossim_zip_url := "https://github.com/ios-control/ios-sim/archive/refs/tags/9.0.0.zip"
	iossim_cmd_zip := exec.Command("curl", "-L", iossim_zip_url, "-o", "package.zip")
	err = iossim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iossim_bin_url := "https://github.com/ios-control/ios-sim/archive/refs/tags/9.0.0.bin"
	iossim_cmd_bin := exec.Command("curl", "-L", iossim_bin_url, "-o", "binary.bin")
	err = iossim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iossim_src_url := "https://github.com/ios-control/ios-sim/archive/refs/tags/9.0.0.src.tar.gz"
	iossim_cmd_src := exec.Command("curl", "-L", iossim_src_url, "-o", "source.tar.gz")
	err = iossim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iossim_cmd_direct := exec.Command("./binary")
	err = iossim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
