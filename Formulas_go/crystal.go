package main

// Crystal - Fast and statically typed, compiled language with Ruby-like syntax
// Homepage: https://crystal-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installCrystal() {
	// Método 1: Descargar y extraer .tar.gz
	crystal_tar_url := "https://github.com/crystal-lang/crystal/archive/refs/tags/1.13.2.tar.gz"
	crystal_cmd_tar := exec.Command("curl", "-L", crystal_tar_url, "-o", "package.tar.gz")
	err := crystal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crystal_zip_url := "https://github.com/crystal-lang/crystal/archive/refs/tags/1.13.2.zip"
	crystal_cmd_zip := exec.Command("curl", "-L", crystal_zip_url, "-o", "package.zip")
	err = crystal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crystal_bin_url := "https://github.com/crystal-lang/crystal/archive/refs/tags/1.13.2.bin"
	crystal_cmd_bin := exec.Command("curl", "-L", crystal_bin_url, "-o", "binary.bin")
	err = crystal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crystal_src_url := "https://github.com/crystal-lang/crystal/archive/refs/tags/1.13.2.src.tar.gz"
	crystal_cmd_src := exec.Command("curl", "-L", crystal_src_url, "-o", "source.tar.gz")
	err = crystal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crystal_cmd_direct := exec.Command("./binary")
	err = crystal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bdw-gc")
exec.Command("latte", "install", "bdw-gc")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
