package main

// Luajit - Just-In-Time Compiler (JIT) for the Lua programming language
// Homepage: https://luajit.org/luajit.html

import (
	"fmt"
	
	"os/exec"
)

func installLuajit() {
	// Método 1: Descargar y extraer .tar.gz
	luajit_tar_url := "https://github.com/LuaJIT/LuaJIT/archive/87ae18af97fd4de790bb6c476b212e047689cc93.tar.gz"
	luajit_cmd_tar := exec.Command("curl", "-L", luajit_tar_url, "-o", "package.tar.gz")
	err := luajit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	luajit_zip_url := "https://github.com/LuaJIT/LuaJIT/archive/87ae18af97fd4de790bb6c476b212e047689cc93.zip"
	luajit_cmd_zip := exec.Command("curl", "-L", luajit_zip_url, "-o", "package.zip")
	err = luajit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	luajit_bin_url := "https://github.com/LuaJIT/LuaJIT/archive/87ae18af97fd4de790bb6c476b212e047689cc93.bin"
	luajit_cmd_bin := exec.Command("curl", "-L", luajit_bin_url, "-o", "binary.bin")
	err = luajit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	luajit_src_url := "https://github.com/LuaJIT/LuaJIT/archive/87ae18af97fd4de790bb6c476b212e047689cc93.src.tar.gz"
	luajit_cmd_src := exec.Command("curl", "-L", luajit_src_url, "-o", "source.tar.gz")
	err = luajit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	luajit_cmd_direct := exec.Command("./binary")
	err = luajit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
