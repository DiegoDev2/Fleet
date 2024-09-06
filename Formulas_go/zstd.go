package main

// Zstd - Zstandard is a real-time compression algorithm
// Homepage: https://facebook.github.io/zstd/

import (
	"fmt"
	
	"os/exec"
)

func installZstd() {
	// Método 1: Descargar y extraer .tar.gz
	zstd_tar_url := "https://github.com/facebook/zstd/archive/refs/tags/v1.5.6.tar.gz"
	zstd_cmd_tar := exec.Command("curl", "-L", zstd_tar_url, "-o", "package.tar.gz")
	err := zstd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zstd_zip_url := "https://github.com/facebook/zstd/archive/refs/tags/v1.5.6.zip"
	zstd_cmd_zip := exec.Command("curl", "-L", zstd_zip_url, "-o", "package.zip")
	err = zstd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zstd_bin_url := "https://github.com/facebook/zstd/archive/refs/tags/v1.5.6.bin"
	zstd_cmd_bin := exec.Command("curl", "-L", zstd_bin_url, "-o", "binary.bin")
	err = zstd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zstd_src_url := "https://github.com/facebook/zstd/archive/refs/tags/v1.5.6.src.tar.gz"
	zstd_cmd_src := exec.Command("curl", "-L", zstd_src_url, "-o", "source.tar.gz")
	err = zstd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zstd_cmd_direct := exec.Command("./binary")
	err = zstd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
}
