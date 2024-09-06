package main

// Asmfmt - Go Assembler Formatter
// Homepage: https://github.com/klauspost/asmfmt

import (
	"fmt"
	
	"os/exec"
)

func installAsmfmt() {
	// Método 1: Descargar y extraer .tar.gz
	asmfmt_tar_url := "https://github.com/klauspost/asmfmt/archive/refs/tags/v1.3.2.tar.gz"
	asmfmt_cmd_tar := exec.Command("curl", "-L", asmfmt_tar_url, "-o", "package.tar.gz")
	err := asmfmt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asmfmt_zip_url := "https://github.com/klauspost/asmfmt/archive/refs/tags/v1.3.2.zip"
	asmfmt_cmd_zip := exec.Command("curl", "-L", asmfmt_zip_url, "-o", "package.zip")
	err = asmfmt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asmfmt_bin_url := "https://github.com/klauspost/asmfmt/archive/refs/tags/v1.3.2.bin"
	asmfmt_cmd_bin := exec.Command("curl", "-L", asmfmt_bin_url, "-o", "binary.bin")
	err = asmfmt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asmfmt_src_url := "https://github.com/klauspost/asmfmt/archive/refs/tags/v1.3.2.src.tar.gz"
	asmfmt_cmd_src := exec.Command("curl", "-L", asmfmt_src_url, "-o", "source.tar.gz")
	err = asmfmt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asmfmt_cmd_direct := exec.Command("./binary")
	err = asmfmt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
