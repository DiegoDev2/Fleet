package main

// Ldc - Portable D programming language compiler
// Homepage: https://wiki.dlang.org/LDC

import (
	"fmt"
	
	"os/exec"
)

func installLdc() {
	// Método 1: Descargar y extraer .tar.gz
	ldc_tar_url := "https://github.com/ldc-developers/ldc/releases/download/v1.39.0/ldc-1.39.0-src.tar.gz"
	ldc_cmd_tar := exec.Command("curl", "-L", ldc_tar_url, "-o", "package.tar.gz")
	err := ldc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ldc_zip_url := "https://github.com/ldc-developers/ldc/releases/download/v1.39.0/ldc-1.39.0-src.zip"
	ldc_cmd_zip := exec.Command("curl", "-L", ldc_zip_url, "-o", "package.zip")
	err = ldc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ldc_bin_url := "https://github.com/ldc-developers/ldc/releases/download/v1.39.0/ldc-1.39.0-src.bin"
	ldc_cmd_bin := exec.Command("curl", "-L", ldc_bin_url, "-o", "binary.bin")
	err = ldc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ldc_src_url := "https://github.com/ldc-developers/ldc/releases/download/v1.39.0/ldc-1.39.0-src.src.tar.gz"
	ldc_cmd_src := exec.Command("curl", "-L", ldc_src_url, "-o", "source.tar.gz")
	err = ldc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ldc_cmd_direct := exec.Command("./binary")
	err = ldc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libconfig")
exec.Command("latte", "install", "libconfig")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
