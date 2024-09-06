package main

// SpanLite - C++20-like span for C++98, C++11 and later in a single-file header-only library
// Homepage: https://github.com/martinmoene/span-lite

import (
	"fmt"
	
	"os/exec"
)

func installSpanLite() {
	// Método 1: Descargar y extraer .tar.gz
	spanlite_tar_url := "https://github.com/martinmoene/span-lite/archive/refs/tags/v0.11.0.tar.gz"
	spanlite_cmd_tar := exec.Command("curl", "-L", spanlite_tar_url, "-o", "package.tar.gz")
	err := spanlite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spanlite_zip_url := "https://github.com/martinmoene/span-lite/archive/refs/tags/v0.11.0.zip"
	spanlite_cmd_zip := exec.Command("curl", "-L", spanlite_zip_url, "-o", "package.zip")
	err = spanlite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spanlite_bin_url := "https://github.com/martinmoene/span-lite/archive/refs/tags/v0.11.0.bin"
	spanlite_cmd_bin := exec.Command("curl", "-L", spanlite_bin_url, "-o", "binary.bin")
	err = spanlite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spanlite_src_url := "https://github.com/martinmoene/span-lite/archive/refs/tags/v0.11.0.src.tar.gz"
	spanlite_cmd_src := exec.Command("curl", "-L", spanlite_src_url, "-o", "source.tar.gz")
	err = spanlite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spanlite_cmd_direct := exec.Command("./binary")
	err = spanlite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
