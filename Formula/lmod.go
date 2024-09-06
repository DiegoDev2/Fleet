package main

// Lmod - Lua-based environment modules system to modify PATH variable
// Homepage: https://lmod.readthedocs.io

import (
	"fmt"
	
	"os/exec"
)

func installLmod() {
	// Método 1: Descargar y extraer .tar.gz
	lmod_tar_url := "https://github.com/TACC/Lmod/archive/refs/tags/8.7.49.tar.gz"
	lmod_cmd_tar := exec.Command("curl", "-L", lmod_tar_url, "-o", "package.tar.gz")
	err := lmod_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lmod_zip_url := "https://github.com/TACC/Lmod/archive/refs/tags/8.7.49.zip"
	lmod_cmd_zip := exec.Command("curl", "-L", lmod_zip_url, "-o", "package.zip")
	err = lmod_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lmod_bin_url := "https://github.com/TACC/Lmod/archive/refs/tags/8.7.49.bin"
	lmod_cmd_bin := exec.Command("curl", "-L", lmod_bin_url, "-o", "binary.bin")
	err = lmod_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lmod_src_url := "https://github.com/TACC/Lmod/archive/refs/tags/8.7.49.src.tar.gz"
	lmod_cmd_src := exec.Command("curl", "-L", lmod_src_url, "-o", "source.tar.gz")
	err = lmod_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lmod_cmd_direct := exec.Command("./binary")
	err = lmod_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: luarocks")
	exec.Command("latte", "install", "luarocks").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
}
