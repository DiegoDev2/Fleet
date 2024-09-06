package main

// Cppzmq - Header-only C++ binding for libzmq
// Homepage: https://www.zeromq.org

import (
	"fmt"
	
	"os/exec"
)

func installCppzmq() {
	// Método 1: Descargar y extraer .tar.gz
	cppzmq_tar_url := "https://github.com/zeromq/cppzmq/archive/refs/tags/v4.10.0.tar.gz"
	cppzmq_cmd_tar := exec.Command("curl", "-L", cppzmq_tar_url, "-o", "package.tar.gz")
	err := cppzmq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cppzmq_zip_url := "https://github.com/zeromq/cppzmq/archive/refs/tags/v4.10.0.zip"
	cppzmq_cmd_zip := exec.Command("curl", "-L", cppzmq_zip_url, "-o", "package.zip")
	err = cppzmq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cppzmq_bin_url := "https://github.com/zeromq/cppzmq/archive/refs/tags/v4.10.0.bin"
	cppzmq_cmd_bin := exec.Command("curl", "-L", cppzmq_bin_url, "-o", "binary.bin")
	err = cppzmq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cppzmq_src_url := "https://github.com/zeromq/cppzmq/archive/refs/tags/v4.10.0.src.tar.gz"
	cppzmq_cmd_src := exec.Command("curl", "-L", cppzmq_src_url, "-o", "source.tar.gz")
	err = cppzmq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cppzmq_cmd_direct := exec.Command("./binary")
	err = cppzmq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: zeromq")
exec.Command("latte", "install", "zeromq")
}
