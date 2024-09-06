package main

// Flatcc - FlatBuffers Compiler and Library in C for C
// Homepage: https://github.com/dvidelabs/flatcc

import (
	"fmt"
	
	"os/exec"
)

func installFlatcc() {
	// Método 1: Descargar y extraer .tar.gz
	flatcc_tar_url := "https://github.com/dvidelabs/flatcc/archive/refs/tags/v0.6.1.tar.gz"
	flatcc_cmd_tar := exec.Command("curl", "-L", flatcc_tar_url, "-o", "package.tar.gz")
	err := flatcc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flatcc_zip_url := "https://github.com/dvidelabs/flatcc/archive/refs/tags/v0.6.1.zip"
	flatcc_cmd_zip := exec.Command("curl", "-L", flatcc_zip_url, "-o", "package.zip")
	err = flatcc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flatcc_bin_url := "https://github.com/dvidelabs/flatcc/archive/refs/tags/v0.6.1.bin"
	flatcc_cmd_bin := exec.Command("curl", "-L", flatcc_bin_url, "-o", "binary.bin")
	err = flatcc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flatcc_src_url := "https://github.com/dvidelabs/flatcc/archive/refs/tags/v0.6.1.src.tar.gz"
	flatcc_cmd_src := exec.Command("curl", "-L", flatcc_src_url, "-o", "source.tar.gz")
	err = flatcc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flatcc_cmd_direct := exec.Command("./binary")
	err = flatcc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
