package main

// Ccache - Object-file caching compiler wrapper
// Homepage: https://ccache.dev/

import (
	"fmt"
	
	"os/exec"
)

func installCcache() {
	// Método 1: Descargar y extraer .tar.gz
	ccache_tar_url := "https://github.com/ccache/ccache/releases/download/v4.10.2/ccache-4.10.2.tar.xz"
	ccache_cmd_tar := exec.Command("curl", "-L", ccache_tar_url, "-o", "package.tar.gz")
	err := ccache_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ccache_zip_url := "https://github.com/ccache/ccache/releases/download/v4.10.2/ccache-4.10.2.tar.xz"
	ccache_cmd_zip := exec.Command("curl", "-L", ccache_zip_url, "-o", "package.zip")
	err = ccache_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ccache_bin_url := "https://github.com/ccache/ccache/releases/download/v4.10.2/ccache-4.10.2.tar.xz"
	ccache_cmd_bin := exec.Command("curl", "-L", ccache_bin_url, "-o", "binary.bin")
	err = ccache_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ccache_src_url := "https://github.com/ccache/ccache/releases/download/v4.10.2/ccache-4.10.2.tar.xz"
	ccache_cmd_src := exec.Command("curl", "-L", ccache_src_url, "-o", "source.tar.gz")
	err = ccache_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ccache_cmd_direct := exec.Command("./binary")
	err = ccache_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoctor")
exec.Command("latte", "install", "asciidoctor")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: cpp-httplib")
exec.Command("latte", "install", "cpp-httplib")
	fmt.Println("Instalando dependencia: doctest")
exec.Command("latte", "install", "doctest")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: span-lite")
exec.Command("latte", "install", "span-lite")
	fmt.Println("Instalando dependencia: tl-expected")
exec.Command("latte", "install", "tl-expected")
	fmt.Println("Instalando dependencia: blake3")
exec.Command("latte", "install", "blake3")
	fmt.Println("Instalando dependencia: fmt")
exec.Command("latte", "install", "fmt")
	fmt.Println("Instalando dependencia: hiredis")
exec.Command("latte", "install", "hiredis")
	fmt.Println("Instalando dependencia: xxhash")
exec.Command("latte", "install", "xxhash")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
