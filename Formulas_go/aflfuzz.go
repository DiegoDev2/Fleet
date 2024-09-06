package main

// AflFuzz - American fuzzy lop: Security-oriented fuzzer
// Homepage: https://github.com/google/AFL

import (
	"fmt"
	
	"os/exec"
)

func installAflFuzz() {
	// Método 1: Descargar y extraer .tar.gz
	aflfuzz_tar_url := "https://github.com/google/AFL/archive/refs/tags/v2.57b.tar.gz"
	aflfuzz_cmd_tar := exec.Command("curl", "-L", aflfuzz_tar_url, "-o", "package.tar.gz")
	err := aflfuzz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aflfuzz_zip_url := "https://github.com/google/AFL/archive/refs/tags/v2.57b.zip"
	aflfuzz_cmd_zip := exec.Command("curl", "-L", aflfuzz_zip_url, "-o", "package.zip")
	err = aflfuzz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aflfuzz_bin_url := "https://github.com/google/AFL/archive/refs/tags/v2.57b.bin"
	aflfuzz_cmd_bin := exec.Command("curl", "-L", aflfuzz_bin_url, "-o", "binary.bin")
	err = aflfuzz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aflfuzz_src_url := "https://github.com/google/AFL/archive/refs/tags/v2.57b.src.tar.gz"
	aflfuzz_cmd_src := exec.Command("curl", "-L", aflfuzz_src_url, "-o", "source.tar.gz")
	err = aflfuzz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aflfuzz_cmd_direct := exec.Command("./binary")
	err = aflfuzz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
