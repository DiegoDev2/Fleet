package main

// MinizipNg - Zip file manipulation library with minizip 1.x compatibility layer
// Homepage: https://github.com/zlib-ng/minizip-ng

import (
	"fmt"
	
	"os/exec"
)

func installMinizipNg() {
	// Método 1: Descargar y extraer .tar.gz
	minizipng_tar_url := "https://github.com/zlib-ng/minizip-ng/archive/refs/tags/4.0.7.tar.gz"
	minizipng_cmd_tar := exec.Command("curl", "-L", minizipng_tar_url, "-o", "package.tar.gz")
	err := minizipng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minizipng_zip_url := "https://github.com/zlib-ng/minizip-ng/archive/refs/tags/4.0.7.zip"
	minizipng_cmd_zip := exec.Command("curl", "-L", minizipng_zip_url, "-o", "package.zip")
	err = minizipng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minizipng_bin_url := "https://github.com/zlib-ng/minizip-ng/archive/refs/tags/4.0.7.bin"
	minizipng_cmd_bin := exec.Command("curl", "-L", minizipng_bin_url, "-o", "binary.bin")
	err = minizipng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minizipng_src_url := "https://github.com/zlib-ng/minizip-ng/archive/refs/tags/4.0.7.src.tar.gz"
	minizipng_cmd_src := exec.Command("curl", "-L", minizipng_src_url, "-o", "source.tar.gz")
	err = minizipng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minizipng_cmd_direct := exec.Command("./binary")
	err = minizipng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
