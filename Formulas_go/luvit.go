package main

// Luvit - Asynchronous I/O for Lua
// Homepage: https://luvit.io

import (
	"fmt"
	
	"os/exec"
)

func installLuvit() {
	// Método 1: Descargar y extraer .tar.gz
	luvit_tar_url := "https://github.com/luvit/luvit/archive/refs/tags/2.18.1.tar.gz"
	luvit_cmd_tar := exec.Command("curl", "-L", luvit_tar_url, "-o", "package.tar.gz")
	err := luvit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	luvit_zip_url := "https://github.com/luvit/luvit/archive/refs/tags/2.18.1.zip"
	luvit_cmd_zip := exec.Command("curl", "-L", luvit_zip_url, "-o", "package.zip")
	err = luvit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	luvit_bin_url := "https://github.com/luvit/luvit/archive/refs/tags/2.18.1.bin"
	luvit_cmd_bin := exec.Command("curl", "-L", luvit_bin_url, "-o", "binary.bin")
	err = luvit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	luvit_src_url := "https://github.com/luvit/luvit/archive/refs/tags/2.18.1.src.tar.gz"
	luvit_cmd_src := exec.Command("curl", "-L", luvit_src_url, "-o", "source.tar.gz")
	err = luvit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	luvit_cmd_direct := exec.Command("./binary")
	err = luvit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libuv")
exec.Command("latte", "install", "libuv")
	fmt.Println("Instalando dependencia: luajit")
exec.Command("latte", "install", "luajit")
	fmt.Println("Instalando dependencia: luv")
exec.Command("latte", "install", "luv")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
}
