package main

// Swift - High-performance system programming language
// Homepage: https://www.swift.org

import (
	"fmt"
	
	"os/exec"
)

func installSwift() {
	// Método 1: Descargar y extraer .tar.gz
	swift_tar_url := "https://github.com/apple/swift/archive/refs/tags/swift-5.10-RELEASE.tar.gz"
	swift_cmd_tar := exec.Command("curl", "-L", swift_tar_url, "-o", "package.tar.gz")
	err := swift_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swift_zip_url := "https://github.com/apple/swift/archive/refs/tags/swift-5.10-RELEASE.zip"
	swift_cmd_zip := exec.Command("curl", "-L", swift_zip_url, "-o", "package.zip")
	err = swift_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swift_bin_url := "https://github.com/apple/swift/archive/refs/tags/swift-5.10-RELEASE.bin"
	swift_cmd_bin := exec.Command("curl", "-L", swift_bin_url, "-o", "binary.bin")
	err = swift_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swift_src_url := "https://github.com/apple/swift/archive/refs/tags/swift-5.10-RELEASE.src.tar.gz"
	swift_cmd_src := exec.Command("curl", "-L", swift_src_url, "-o", "source.tar.gz")
	err = swift_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swift_cmd_direct := exec.Command("./binary")
	err = swift_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
}
