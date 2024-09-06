package main

// LuaAT53 - Powerful, lightweight programming language
// Homepage: https://www.lua.org/

import (
	"fmt"
	
	"os/exec"
)

func installLuaAT53() {
	// Método 1: Descargar y extraer .tar.gz
	luaat53_tar_url := "https://www.lua.org/ftp/lua-5.3.6.tar.gz"
	luaat53_cmd_tar := exec.Command("curl", "-L", luaat53_tar_url, "-o", "package.tar.gz")
	err := luaat53_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	luaat53_zip_url := "https://www.lua.org/ftp/lua-5.3.6.zip"
	luaat53_cmd_zip := exec.Command("curl", "-L", luaat53_zip_url, "-o", "package.zip")
	err = luaat53_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	luaat53_bin_url := "https://www.lua.org/ftp/lua-5.3.6.bin"
	luaat53_cmd_bin := exec.Command("curl", "-L", luaat53_bin_url, "-o", "binary.bin")
	err = luaat53_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	luaat53_src_url := "https://www.lua.org/ftp/lua-5.3.6.src.tar.gz"
	luaat53_cmd_src := exec.Command("curl", "-L", luaat53_src_url, "-o", "source.tar.gz")
	err = luaat53_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	luaat53_cmd_direct := exec.Command("./binary")
	err = luaat53_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
