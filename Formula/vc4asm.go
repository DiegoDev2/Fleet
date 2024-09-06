package main

// Vc4asm - Macro assembler for Broadcom VideoCore IV aka Raspberry Pi GPU
// Homepage: https://maazl.de/project/vc4asm/doc/index.html

import (
	"fmt"
	
	"os/exec"
)

func installVc4asm() {
	// Método 1: Descargar y extraer .tar.gz
	vc4asm_tar_url := "https://github.com/maazl/vc4asm/archive/refs/tags/V0.3.tar.gz"
	vc4asm_cmd_tar := exec.Command("curl", "-L", vc4asm_tar_url, "-o", "package.tar.gz")
	err := vc4asm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vc4asm_zip_url := "https://github.com/maazl/vc4asm/archive/refs/tags/V0.3.zip"
	vc4asm_cmd_zip := exec.Command("curl", "-L", vc4asm_zip_url, "-o", "package.zip")
	err = vc4asm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vc4asm_bin_url := "https://github.com/maazl/vc4asm/archive/refs/tags/V0.3.bin"
	vc4asm_cmd_bin := exec.Command("curl", "-L", vc4asm_bin_url, "-o", "binary.bin")
	err = vc4asm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vc4asm_src_url := "https://github.com/maazl/vc4asm/archive/refs/tags/V0.3.src.tar.gz"
	vc4asm_cmd_src := exec.Command("curl", "-L", vc4asm_src_url, "-o", "source.tar.gz")
	err = vc4asm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vc4asm_cmd_direct := exec.Command("./binary")
	err = vc4asm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
