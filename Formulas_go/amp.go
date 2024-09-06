package main

// Amp - Text editor for your terminal
// Homepage: https://amp.rs

import (
	"fmt"
	
	"os/exec"
)

func installAmp() {
	// Método 1: Descargar y extraer .tar.gz
	amp_tar_url := "https://github.com/jmacdonald/amp/archive/refs/tags/0.7.0.tar.gz"
	amp_cmd_tar := exec.Command("curl", "-L", amp_tar_url, "-o", "package.tar.gz")
	err := amp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	amp_zip_url := "https://github.com/jmacdonald/amp/archive/refs/tags/0.7.0.zip"
	amp_cmd_zip := exec.Command("curl", "-L", amp_zip_url, "-o", "package.zip")
	err = amp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	amp_bin_url := "https://github.com/jmacdonald/amp/archive/refs/tags/0.7.0.bin"
	amp_cmd_bin := exec.Command("curl", "-L", amp_bin_url, "-o", "binary.bin")
	err = amp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	amp_src_url := "https://github.com/jmacdonald/amp/archive/refs/tags/0.7.0.src.tar.gz"
	amp_cmd_src := exec.Command("curl", "-L", amp_src_url, "-o", "source.tar.gz")
	err = amp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	amp_cmd_direct := exec.Command("./binary")
	err = amp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
}
