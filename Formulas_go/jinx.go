package main

// Jinx - Embeddable scripting language for real-time applications
// Homepage: https://github.com/JamesBoer/Jinx

import (
	"fmt"
	
	"os/exec"
)

func installJinx() {
	// Método 1: Descargar y extraer .tar.gz
	jinx_tar_url := "https://github.com/JamesBoer/Jinx/archive/refs/tags/v1.3.10.tar.gz"
	jinx_cmd_tar := exec.Command("curl", "-L", jinx_tar_url, "-o", "package.tar.gz")
	err := jinx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jinx_zip_url := "https://github.com/JamesBoer/Jinx/archive/refs/tags/v1.3.10.zip"
	jinx_cmd_zip := exec.Command("curl", "-L", jinx_zip_url, "-o", "package.zip")
	err = jinx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jinx_bin_url := "https://github.com/JamesBoer/Jinx/archive/refs/tags/v1.3.10.bin"
	jinx_cmd_bin := exec.Command("curl", "-L", jinx_bin_url, "-o", "binary.bin")
	err = jinx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jinx_src_url := "https://github.com/JamesBoer/Jinx/archive/refs/tags/v1.3.10.src.tar.gz"
	jinx_cmd_src := exec.Command("curl", "-L", jinx_src_url, "-o", "source.tar.gz")
	err = jinx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jinx_cmd_direct := exec.Command("./binary")
	err = jinx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
