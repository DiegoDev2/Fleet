package main

// GoogleBenchmark - C++ microbenchmark support library
// Homepage: https://github.com/google/benchmark

import (
	"fmt"
	
	"os/exec"
)

func installGoogleBenchmark() {
	// Método 1: Descargar y extraer .tar.gz
	googlebenchmark_tar_url := "https://github.com/google/benchmark/archive/refs/tags/v1.9.0.tar.gz"
	googlebenchmark_cmd_tar := exec.Command("curl", "-L", googlebenchmark_tar_url, "-o", "package.tar.gz")
	err := googlebenchmark_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	googlebenchmark_zip_url := "https://github.com/google/benchmark/archive/refs/tags/v1.9.0.zip"
	googlebenchmark_cmd_zip := exec.Command("curl", "-L", googlebenchmark_zip_url, "-o", "package.zip")
	err = googlebenchmark_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	googlebenchmark_bin_url := "https://github.com/google/benchmark/archive/refs/tags/v1.9.0.bin"
	googlebenchmark_cmd_bin := exec.Command("curl", "-L", googlebenchmark_bin_url, "-o", "binary.bin")
	err = googlebenchmark_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	googlebenchmark_src_url := "https://github.com/google/benchmark/archive/refs/tags/v1.9.0.src.tar.gz"
	googlebenchmark_cmd_src := exec.Command("curl", "-L", googlebenchmark_src_url, "-o", "source.tar.gz")
	err = googlebenchmark_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	googlebenchmark_cmd_direct := exec.Command("./binary")
	err = googlebenchmark_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
