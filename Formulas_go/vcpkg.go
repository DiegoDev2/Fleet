package main

// Vcpkg - C++ Library Manager
// Homepage: https://github.com/microsoft/vcpkg

import (
	"fmt"
	
	"os/exec"
)

func installVcpkg() {
	// Método 1: Descargar y extraer .tar.gz
	vcpkg_tar_url := "https://github.com/microsoft/vcpkg-tool/archive/refs/tags/2024-08-01.tar.gz"
	vcpkg_cmd_tar := exec.Command("curl", "-L", vcpkg_tar_url, "-o", "package.tar.gz")
	err := vcpkg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vcpkg_zip_url := "https://github.com/microsoft/vcpkg-tool/archive/refs/tags/2024-08-01.zip"
	vcpkg_cmd_zip := exec.Command("curl", "-L", vcpkg_zip_url, "-o", "package.zip")
	err = vcpkg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vcpkg_bin_url := "https://github.com/microsoft/vcpkg-tool/archive/refs/tags/2024-08-01.bin"
	vcpkg_cmd_bin := exec.Command("curl", "-L", vcpkg_bin_url, "-o", "binary.bin")
	err = vcpkg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vcpkg_src_url := "https://github.com/microsoft/vcpkg-tool/archive/refs/tags/2024-08-01.src.tar.gz"
	vcpkg_cmd_src := exec.Command("curl", "-L", vcpkg_src_url, "-o", "source.tar.gz")
	err = vcpkg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vcpkg_cmd_direct := exec.Command("./binary")
	err = vcpkg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: cmrc")
exec.Command("latte", "install", "cmrc")
	fmt.Println("Instalando dependencia: fmt")
exec.Command("latte", "install", "fmt")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}
