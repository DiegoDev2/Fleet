package main

// ZLua - New cd command that helps you navigate faster by learning your habits
// Homepage: https://github.com/skywind3000/z.lua

import (
	"fmt"
	
	"os/exec"
)

func installZLua() {
	// Método 1: Descargar y extraer .tar.gz
	zlua_tar_url := "https://github.com/skywind3000/z.lua/archive/refs/tags/1.8.18.tar.gz"
	zlua_cmd_tar := exec.Command("curl", "-L", zlua_tar_url, "-o", "package.tar.gz")
	err := zlua_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zlua_zip_url := "https://github.com/skywind3000/z.lua/archive/refs/tags/1.8.18.zip"
	zlua_cmd_zip := exec.Command("curl", "-L", zlua_zip_url, "-o", "package.zip")
	err = zlua_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zlua_bin_url := "https://github.com/skywind3000/z.lua/archive/refs/tags/1.8.18.bin"
	zlua_cmd_bin := exec.Command("curl", "-L", zlua_bin_url, "-o", "binary.bin")
	err = zlua_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zlua_src_url := "https://github.com/skywind3000/z.lua/archive/refs/tags/1.8.18.src.tar.gz"
	zlua_cmd_src := exec.Command("curl", "-L", zlua_src_url, "-o", "source.tar.gz")
	err = zlua_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zlua_cmd_direct := exec.Command("./binary")
	err = zlua_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
}
