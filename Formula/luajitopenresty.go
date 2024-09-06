package main

// LuajitOpenresty - OpenResty's Branch of LuaJIT 2
// Homepage: https://github.com/openresty/luajit2

import (
	"fmt"
	
	"os/exec"
)

func installLuajitOpenresty() {
	// Método 1: Descargar y extraer .tar.gz
	luajitopenresty_tar_url := "https://github.com/openresty/luajit2/archive/refs/tags/v2.1-20240815.tar.gz"
	luajitopenresty_cmd_tar := exec.Command("curl", "-L", luajitopenresty_tar_url, "-o", "package.tar.gz")
	err := luajitopenresty_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	luajitopenresty_zip_url := "https://github.com/openresty/luajit2/archive/refs/tags/v2.1-20240815.zip"
	luajitopenresty_cmd_zip := exec.Command("curl", "-L", luajitopenresty_zip_url, "-o", "package.zip")
	err = luajitopenresty_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	luajitopenresty_bin_url := "https://github.com/openresty/luajit2/archive/refs/tags/v2.1-20240815.bin"
	luajitopenresty_cmd_bin := exec.Command("curl", "-L", luajitopenresty_bin_url, "-o", "binary.bin")
	err = luajitopenresty_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	luajitopenresty_src_url := "https://github.com/openresty/luajit2/archive/refs/tags/v2.1-20240815.src.tar.gz"
	luajitopenresty_cmd_src := exec.Command("curl", "-L", luajitopenresty_src_url, "-o", "source.tar.gz")
	err = luajitopenresty_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	luajitopenresty_cmd_direct := exec.Command("./binary")
	err = luajitopenresty_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
