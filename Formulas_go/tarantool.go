package main

// Tarantool - In-memory database and Lua application server
// Homepage: https://tarantool.org/

import (
	"fmt"
	
	"os/exec"
)

func installTarantool() {
	// Método 1: Descargar y extraer .tar.gz
	tarantool_tar_url := "https://download.tarantool.org/tarantool/src/tarantool-3.2.0.tar.gz"
	tarantool_cmd_tar := exec.Command("curl", "-L", tarantool_tar_url, "-o", "package.tar.gz")
	err := tarantool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tarantool_zip_url := "https://download.tarantool.org/tarantool/src/tarantool-3.2.0.zip"
	tarantool_cmd_zip := exec.Command("curl", "-L", tarantool_zip_url, "-o", "package.zip")
	err = tarantool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tarantool_bin_url := "https://download.tarantool.org/tarantool/src/tarantool-3.2.0.bin"
	tarantool_cmd_bin := exec.Command("curl", "-L", tarantool_bin_url, "-o", "binary.bin")
	err = tarantool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tarantool_src_url := "https://download.tarantool.org/tarantool/src/tarantool-3.2.0.src.tar.gz"
	tarantool_cmd_src := exec.Command("curl", "-L", tarantool_src_url, "-o", "source.tar.gz")
	err = tarantool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tarantool_cmd_direct := exec.Command("./binary")
	err = tarantool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: curl")
exec.Command("latte", "install", "curl")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: libunwind")
exec.Command("latte", "install", "libunwind")
}
