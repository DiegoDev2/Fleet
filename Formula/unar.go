package main

// Unar - Command-line unarchiving tools supporting multiple formats
// Homepage: https://theunarchiver.com/command-line

import (
	"fmt"
	
	"os/exec"
)

func installUnar() {
	// Método 1: Descargar y extraer .tar.gz
	unar_tar_url := "https://github.com/MacPaw/XADMaster/archive/refs/tags/v1.10.8.tar.gz"
	unar_cmd_tar := exec.Command("curl", "-L", unar_tar_url, "-o", "package.tar.gz")
	err := unar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unar_zip_url := "https://github.com/MacPaw/XADMaster/archive/refs/tags/v1.10.8.zip"
	unar_cmd_zip := exec.Command("curl", "-L", unar_zip_url, "-o", "package.zip")
	err = unar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unar_bin_url := "https://github.com/MacPaw/XADMaster/archive/refs/tags/v1.10.8.bin"
	unar_cmd_bin := exec.Command("curl", "-L", unar_bin_url, "-o", "binary.bin")
	err = unar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unar_src_url := "https://github.com/MacPaw/XADMaster/archive/refs/tags/v1.10.8.src.tar.gz"
	unar_cmd_src := exec.Command("curl", "-L", unar_src_url, "-o", "source.tar.gz")
	err = unar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unar_cmd_direct := exec.Command("./binary")
	err = unar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnustep-base")
	exec.Command("latte", "install", "gnustep-base").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: libobjc2")
	exec.Command("latte", "install", "libobjc2").Run()
	fmt.Println("Instalando dependencia: wavpack")
	exec.Command("latte", "install", "wavpack").Run()
}
