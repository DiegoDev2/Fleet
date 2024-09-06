package main

// Luaver - Manage and switch between versions of Lua, LuaJIT, and Luarocks
// Homepage: https://github.com/DhavalKapil/luaver

import (
	"fmt"
	
	"os/exec"
)

func installLuaver() {
	// Método 1: Descargar y extraer .tar.gz
	luaver_tar_url := "https://github.com/DhavalKapil/luaver/archive/refs/tags/v1.1.0.tar.gz"
	luaver_cmd_tar := exec.Command("curl", "-L", luaver_tar_url, "-o", "package.tar.gz")
	err := luaver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	luaver_zip_url := "https://github.com/DhavalKapil/luaver/archive/refs/tags/v1.1.0.zip"
	luaver_cmd_zip := exec.Command("curl", "-L", luaver_zip_url, "-o", "package.zip")
	err = luaver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	luaver_bin_url := "https://github.com/DhavalKapil/luaver/archive/refs/tags/v1.1.0.bin"
	luaver_cmd_bin := exec.Command("curl", "-L", luaver_bin_url, "-o", "binary.bin")
	err = luaver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	luaver_src_url := "https://github.com/DhavalKapil/luaver/archive/refs/tags/v1.1.0.src.tar.gz"
	luaver_cmd_src := exec.Command("curl", "-L", luaver_src_url, "-o", "source.tar.gz")
	err = luaver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	luaver_cmd_direct := exec.Command("./binary")
	err = luaver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: wget")
	exec.Command("latte", "install", "wget").Run()
}
