package main

// Luv - Bare libuv bindings for lua
// Homepage: https://github.com/luvit/luv

import (
	"fmt"
	
	"os/exec"
)

func installLuv() {
	// Método 1: Descargar y extraer .tar.gz
	luv_tar_url := "https://github.com/luvit/luv/archive/refs/tags/1.48.0-2.tar.gz"
	luv_cmd_tar := exec.Command("curl", "-L", luv_tar_url, "-o", "package.tar.gz")
	err := luv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	luv_zip_url := "https://github.com/luvit/luv/archive/refs/tags/1.48.0-2.zip"
	luv_cmd_zip := exec.Command("curl", "-L", luv_zip_url, "-o", "package.zip")
	err = luv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	luv_bin_url := "https://github.com/luvit/luv/archive/refs/tags/1.48.0-2.bin"
	luv_cmd_bin := exec.Command("curl", "-L", luv_bin_url, "-o", "binary.bin")
	err = luv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	luv_src_url := "https://github.com/luvit/luv/archive/refs/tags/1.48.0-2.src.tar.gz"
	luv_cmd_src := exec.Command("curl", "-L", luv_src_url, "-o", "source.tar.gz")
	err = luv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	luv_cmd_direct := exec.Command("./binary")
	err = luv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: luajit")
	exec.Command("latte", "install", "luajit").Run()
	fmt.Println("Instalando dependencia: libuv")
	exec.Command("latte", "install", "libuv").Run()
}
