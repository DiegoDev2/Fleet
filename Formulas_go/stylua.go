package main

// Stylua - Opinionated Lua code formatter
// Homepage: https://github.com/JohnnyMorganz/StyLua

import (
	"fmt"
	
	"os/exec"
)

func installStylua() {
	// Método 1: Descargar y extraer .tar.gz
	stylua_tar_url := "https://github.com/JohnnyMorganz/StyLua/archive/refs/tags/v0.20.0.tar.gz"
	stylua_cmd_tar := exec.Command("curl", "-L", stylua_tar_url, "-o", "package.tar.gz")
	err := stylua_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stylua_zip_url := "https://github.com/JohnnyMorganz/StyLua/archive/refs/tags/v0.20.0.zip"
	stylua_cmd_zip := exec.Command("curl", "-L", stylua_zip_url, "-o", "package.zip")
	err = stylua_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stylua_bin_url := "https://github.com/JohnnyMorganz/StyLua/archive/refs/tags/v0.20.0.bin"
	stylua_cmd_bin := exec.Command("curl", "-L", stylua_bin_url, "-o", "binary.bin")
	err = stylua_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stylua_src_url := "https://github.com/JohnnyMorganz/StyLua/archive/refs/tags/v0.20.0.src.tar.gz"
	stylua_cmd_src := exec.Command("curl", "-L", stylua_src_url, "-o", "source.tar.gz")
	err = stylua_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stylua_cmd_direct := exec.Command("./binary")
	err = stylua_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
