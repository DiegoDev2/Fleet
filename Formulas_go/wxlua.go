package main

// Wxlua - Lua bindings for wxWidgets cross-platform GUI toolkit
// Homepage: https://github.com/pkulchenko/wxlua

import (
	"fmt"
	
	"os/exec"
)

func installWxlua() {
	// Método 1: Descargar y extraer .tar.gz
	wxlua_tar_url := "https://github.com/pkulchenko/wxlua/archive/refs/tags/v3.2.0.2.tar.gz"
	wxlua_cmd_tar := exec.Command("curl", "-L", wxlua_tar_url, "-o", "package.tar.gz")
	err := wxlua_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wxlua_zip_url := "https://github.com/pkulchenko/wxlua/archive/refs/tags/v3.2.0.2.zip"
	wxlua_cmd_zip := exec.Command("curl", "-L", wxlua_zip_url, "-o", "package.zip")
	err = wxlua_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wxlua_bin_url := "https://github.com/pkulchenko/wxlua/archive/refs/tags/v3.2.0.2.bin"
	wxlua_cmd_bin := exec.Command("curl", "-L", wxlua_bin_url, "-o", "binary.bin")
	err = wxlua_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wxlua_src_url := "https://github.com/pkulchenko/wxlua/archive/refs/tags/v3.2.0.2.src.tar.gz"
	wxlua_cmd_src := exec.Command("curl", "-L", wxlua_src_url, "-o", "source.tar.gz")
	err = wxlua_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wxlua_cmd_direct := exec.Command("./binary")
	err = wxlua_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
	fmt.Println("Instalando dependencia: wxwidgets")
exec.Command("latte", "install", "wxwidgets")
}
