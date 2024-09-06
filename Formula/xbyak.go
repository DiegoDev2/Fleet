package main

// Xbyak - C++ JIT assembler for x86 (IA32), x64 (AMD64, x86-64)
// Homepage: https://github.com/herumi/xbyak

import (
	"fmt"
	
	"os/exec"
)

func installXbyak() {
	// Método 1: Descargar y extraer .tar.gz
	xbyak_tar_url := "https://github.com/herumi/xbyak/archive/refs/tags/v7.07.1.tar.gz"
	xbyak_cmd_tar := exec.Command("curl", "-L", xbyak_tar_url, "-o", "package.tar.gz")
	err := xbyak_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xbyak_zip_url := "https://github.com/herumi/xbyak/archive/refs/tags/v7.07.1.zip"
	xbyak_cmd_zip := exec.Command("curl", "-L", xbyak_zip_url, "-o", "package.zip")
	err = xbyak_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xbyak_bin_url := "https://github.com/herumi/xbyak/archive/refs/tags/v7.07.1.bin"
	xbyak_cmd_bin := exec.Command("curl", "-L", xbyak_bin_url, "-o", "binary.bin")
	err = xbyak_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xbyak_src_url := "https://github.com/herumi/xbyak/archive/refs/tags/v7.07.1.src.tar.gz"
	xbyak_cmd_src := exec.Command("curl", "-L", xbyak_src_url, "-o", "source.tar.gz")
	err = xbyak_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xbyak_cmd_direct := exec.Command("./binary")
	err = xbyak_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
