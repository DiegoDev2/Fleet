package main

// LuaAT51 - Powerful, lightweight programming language (v5.1.5)
// Homepage: https://www.lua.org/

import (
	"fmt"
	
	"os/exec"
)

func installLuaAT51() {
	// Método 1: Descargar y extraer .tar.gz
	luaat51_tar_url := "https://www.lua.org/ftp/lua-5.1.5.tar.gz"
	luaat51_cmd_tar := exec.Command("curl", "-L", luaat51_tar_url, "-o", "package.tar.gz")
	err := luaat51_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	luaat51_zip_url := "https://www.lua.org/ftp/lua-5.1.5.zip"
	luaat51_cmd_zip := exec.Command("curl", "-L", luaat51_zip_url, "-o", "package.zip")
	err = luaat51_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	luaat51_bin_url := "https://www.lua.org/ftp/lua-5.1.5.bin"
	luaat51_cmd_bin := exec.Command("curl", "-L", luaat51_bin_url, "-o", "binary.bin")
	err = luaat51_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	luaat51_src_url := "https://www.lua.org/ftp/lua-5.1.5.src.tar.gz"
	luaat51_cmd_src := exec.Command("curl", "-L", luaat51_src_url, "-o", "source.tar.gz")
	err = luaat51_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	luaat51_cmd_direct := exec.Command("./binary")
	err = luaat51_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
