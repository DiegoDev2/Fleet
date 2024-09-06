package main

// Zchunk - Compressed file format for efficient deltas
// Homepage: https://github.com/zchunk/zchunk

import (
	"fmt"
	
	"os/exec"
)

func installZchunk() {
	// Método 1: Descargar y extraer .tar.gz
	zchunk_tar_url := "https://github.com/zchunk/zchunk/archive/refs/tags/1.5.1.tar.gz"
	zchunk_cmd_tar := exec.Command("curl", "-L", zchunk_tar_url, "-o", "package.tar.gz")
	err := zchunk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zchunk_zip_url := "https://github.com/zchunk/zchunk/archive/refs/tags/1.5.1.zip"
	zchunk_cmd_zip := exec.Command("curl", "-L", zchunk_zip_url, "-o", "package.zip")
	err = zchunk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zchunk_bin_url := "https://github.com/zchunk/zchunk/archive/refs/tags/1.5.1.bin"
	zchunk_cmd_bin := exec.Command("curl", "-L", zchunk_bin_url, "-o", "binary.bin")
	err = zchunk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zchunk_src_url := "https://github.com/zchunk/zchunk/archive/refs/tags/1.5.1.src.tar.gz"
	zchunk_cmd_src := exec.Command("curl", "-L", zchunk_src_url, "-o", "source.tar.gz")
	err = zchunk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zchunk_cmd_direct := exec.Command("./binary")
	err = zchunk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: argp-standalone")
exec.Command("latte", "install", "argp-standalone")
}
