package main

// Dwarfs - Fast high compression read-only file system for Linux, Windows, and macOS
// Homepage: https://github.com/mhx/dwarfs

import (
	"fmt"
	
	"os/exec"
)

func installDwarfs() {
	// Método 1: Descargar y extraer .tar.gz
	dwarfs_tar_url := "https://github.com/mhx/dwarfs/releases/download/v0.10.1/dwarfs-0.10.1.tar.xz"
	dwarfs_cmd_tar := exec.Command("curl", "-L", dwarfs_tar_url, "-o", "package.tar.gz")
	err := dwarfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dwarfs_zip_url := "https://github.com/mhx/dwarfs/releases/download/v0.10.1/dwarfs-0.10.1.tar.xz"
	dwarfs_cmd_zip := exec.Command("curl", "-L", dwarfs_zip_url, "-o", "package.zip")
	err = dwarfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dwarfs_bin_url := "https://github.com/mhx/dwarfs/releases/download/v0.10.1/dwarfs-0.10.1.tar.xz"
	dwarfs_cmd_bin := exec.Command("curl", "-L", dwarfs_bin_url, "-o", "binary.bin")
	err = dwarfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dwarfs_src_url := "https://github.com/mhx/dwarfs/releases/download/v0.10.1/dwarfs-0.10.1.tar.xz"
	dwarfs_cmd_src := exec.Command("curl", "-L", dwarfs_src_url, "-o", "source.tar.gz")
	err = dwarfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dwarfs_cmd_direct := exec.Command("./binary")
	err = dwarfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: googletest")
	exec.Command("latte", "install", "googletest").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: brotli")
	exec.Command("latte", "install", "brotli").Run()
	fmt.Println("Instalando dependencia: double-conversion")
	exec.Command("latte", "install", "double-conversion").Run()
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: fmt")
	exec.Command("latte", "install", "fmt").Run()
	fmt.Println("Instalando dependencia: gflags")
	exec.Command("latte", "install", "gflags").Run()
	fmt.Println("Instalando dependencia: glog")
	exec.Command("latte", "install", "glog").Run()
	fmt.Println("Instalando dependencia: howard-hinnant-date")
	exec.Command("latte", "install", "howard-hinnant-date").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: nlohmann-json")
	exec.Command("latte", "install", "nlohmann-json").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: parallel-hashmap")
	exec.Command("latte", "install", "parallel-hashmap").Run()
	fmt.Println("Instalando dependencia: range-v3")
	exec.Command("latte", "install", "range-v3").Run()
	fmt.Println("Instalando dependencia: utf8cpp")
	exec.Command("latte", "install", "utf8cpp").Run()
	fmt.Println("Instalando dependencia: xxhash")
	exec.Command("latte", "install", "xxhash").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: jemalloc")
	exec.Command("latte", "install", "jemalloc").Run()
}
