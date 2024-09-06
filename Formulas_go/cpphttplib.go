package main

// CppHttplib - C++ header-only HTTP/HTTPS server and client library
// Homepage: https://github.com/yhirose/cpp-httplib

import (
	"fmt"
	
	"os/exec"
)

func installCppHttplib() {
	// Método 1: Descargar y extraer .tar.gz
	cpphttplib_tar_url := "https://github.com/yhirose/cpp-httplib/archive/refs/tags/v0.17.1.tar.gz"
	cpphttplib_cmd_tar := exec.Command("curl", "-L", cpphttplib_tar_url, "-o", "package.tar.gz")
	err := cpphttplib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpphttplib_zip_url := "https://github.com/yhirose/cpp-httplib/archive/refs/tags/v0.17.1.zip"
	cpphttplib_cmd_zip := exec.Command("curl", "-L", cpphttplib_zip_url, "-o", "package.zip")
	err = cpphttplib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpphttplib_bin_url := "https://github.com/yhirose/cpp-httplib/archive/refs/tags/v0.17.1.bin"
	cpphttplib_cmd_bin := exec.Command("curl", "-L", cpphttplib_bin_url, "-o", "binary.bin")
	err = cpphttplib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpphttplib_src_url := "https://github.com/yhirose/cpp-httplib/archive/refs/tags/v0.17.1.src.tar.gz"
	cpphttplib_cmd_src := exec.Command("curl", "-L", cpphttplib_src_url, "-o", "source.tar.gz")
	err = cpphttplib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpphttplib_cmd_direct := exec.Command("./binary")
	err = cpphttplib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}
