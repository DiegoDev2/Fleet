package main

// Zigmod - Package manager for the Zig programming language
// Homepage: https://nektro.github.io/zigmod/

import (
	"fmt"
	
	"os/exec"
)

func installZigmod() {
	// Método 1: Descargar y extraer .tar.gz
	zigmod_tar_url := "https://github.com/nektro/zigmod/archive/refs/tags/r90.tar.gz"
	zigmod_cmd_tar := exec.Command("curl", "-L", zigmod_tar_url, "-o", "package.tar.gz")
	err := zigmod_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zigmod_zip_url := "https://github.com/nektro/zigmod/archive/refs/tags/r90.zip"
	zigmod_cmd_zip := exec.Command("curl", "-L", zigmod_zip_url, "-o", "package.zip")
	err = zigmod_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zigmod_bin_url := "https://github.com/nektro/zigmod/archive/refs/tags/r90.bin"
	zigmod_cmd_bin := exec.Command("curl", "-L", zigmod_bin_url, "-o", "binary.bin")
	err = zigmod_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zigmod_src_url := "https://github.com/nektro/zigmod/archive/refs/tags/r90.src.tar.gz"
	zigmod_cmd_src := exec.Command("curl", "-L", zigmod_src_url, "-o", "source.tar.gz")
	err = zigmod_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zigmod_cmd_direct := exec.Command("./binary")
	err = zigmod_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: zig")
	exec.Command("latte", "install", "zig").Run()
}
