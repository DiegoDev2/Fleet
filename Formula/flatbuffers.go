package main

// Flatbuffers - Serialization library for C++, supporting Java, C#, and Go
// Homepage: https://google.github.io/flatbuffers

import (
	"fmt"
	
	"os/exec"
)

func installFlatbuffers() {
	// Método 1: Descargar y extraer .tar.gz
	flatbuffers_tar_url := "https://github.com/google/flatbuffers/archive/refs/tags/v24.3.25.tar.gz"
	flatbuffers_cmd_tar := exec.Command("curl", "-L", flatbuffers_tar_url, "-o", "package.tar.gz")
	err := flatbuffers_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flatbuffers_zip_url := "https://github.com/google/flatbuffers/archive/refs/tags/v24.3.25.zip"
	flatbuffers_cmd_zip := exec.Command("curl", "-L", flatbuffers_zip_url, "-o", "package.zip")
	err = flatbuffers_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flatbuffers_bin_url := "https://github.com/google/flatbuffers/archive/refs/tags/v24.3.25.bin"
	flatbuffers_cmd_bin := exec.Command("curl", "-L", flatbuffers_bin_url, "-o", "binary.bin")
	err = flatbuffers_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flatbuffers_src_url := "https://github.com/google/flatbuffers/archive/refs/tags/v24.3.25.src.tar.gz"
	flatbuffers_cmd_src := exec.Command("curl", "-L", flatbuffers_src_url, "-o", "source.tar.gz")
	err = flatbuffers_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flatbuffers_cmd_direct := exec.Command("./binary")
	err = flatbuffers_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
