package main

// Rizin - UNIX-like reverse engineering framework and command-line toolset
// Homepage: https://rizin.re

import (
	"fmt"
	
	"os/exec"
)

func installRizin() {
	// Método 1: Descargar y extraer .tar.gz
	rizin_tar_url := "https://github.com/rizinorg/rizin/releases/download/v0.7.3/rizin-src-v0.7.3.tar.xz"
	rizin_cmd_tar := exec.Command("curl", "-L", rizin_tar_url, "-o", "package.tar.gz")
	err := rizin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rizin_zip_url := "https://github.com/rizinorg/rizin/releases/download/v0.7.3/rizin-src-v0.7.3.tar.xz"
	rizin_cmd_zip := exec.Command("curl", "-L", rizin_zip_url, "-o", "package.zip")
	err = rizin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rizin_bin_url := "https://github.com/rizinorg/rizin/releases/download/v0.7.3/rizin-src-v0.7.3.tar.xz"
	rizin_cmd_bin := exec.Command("curl", "-L", rizin_bin_url, "-o", "binary.bin")
	err = rizin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rizin_src_url := "https://github.com/rizinorg/rizin/releases/download/v0.7.3/rizin-src-v0.7.3.tar.xz"
	rizin_cmd_src := exec.Command("curl", "-L", rizin_src_url, "-o", "source.tar.gz")
	err = rizin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rizin_cmd_direct := exec.Command("./binary")
	err = rizin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: capstone")
	exec.Command("latte", "install", "capstone").Run()
	fmt.Println("Instalando dependencia: libmagic")
	exec.Command("latte", "install", "libmagic").Run()
	fmt.Println("Instalando dependencia: libzip")
	exec.Command("latte", "install", "libzip").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: tree-sitter")
	exec.Command("latte", "install", "tree-sitter").Run()
	fmt.Println("Instalando dependencia: xxhash")
	exec.Command("latte", "install", "xxhash").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
