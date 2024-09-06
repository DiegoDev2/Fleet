package main

// Cpuid - CPU feature identification for Go
// Homepage: https://github.com/klauspost/cpuid

import (
	"fmt"
	
	"os/exec"
)

func installCpuid() {
	// Método 1: Descargar y extraer .tar.gz
	cpuid_tar_url := "https://github.com/klauspost/cpuid/archive/refs/tags/v2.2.8.tar.gz"
	cpuid_cmd_tar := exec.Command("curl", "-L", cpuid_tar_url, "-o", "package.tar.gz")
	err := cpuid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpuid_zip_url := "https://github.com/klauspost/cpuid/archive/refs/tags/v2.2.8.zip"
	cpuid_cmd_zip := exec.Command("curl", "-L", cpuid_zip_url, "-o", "package.zip")
	err = cpuid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpuid_bin_url := "https://github.com/klauspost/cpuid/archive/refs/tags/v2.2.8.bin"
	cpuid_cmd_bin := exec.Command("curl", "-L", cpuid_bin_url, "-o", "binary.bin")
	err = cpuid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpuid_src_url := "https://github.com/klauspost/cpuid/archive/refs/tags/v2.2.8.src.tar.gz"
	cpuid_cmd_src := exec.Command("curl", "-L", cpuid_src_url, "-o", "source.tar.gz")
	err = cpuid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpuid_cmd_direct := exec.Command("./binary")
	err = cpuid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
