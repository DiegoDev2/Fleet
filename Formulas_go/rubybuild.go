package main

// RubyBuild - Install various Ruby versions and implementations
// Homepage: https://github.com/rbenv/ruby-build

import (
	"fmt"
	
	"os/exec"
)

func installRubyBuild() {
	// Método 1: Descargar y extraer .tar.gz
	rubybuild_tar_url := "https://github.com/rbenv/ruby-build/archive/refs/tags/v20240903.tar.gz"
	rubybuild_cmd_tar := exec.Command("curl", "-L", rubybuild_tar_url, "-o", "package.tar.gz")
	err := rubybuild_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rubybuild_zip_url := "https://github.com/rbenv/ruby-build/archive/refs/tags/v20240903.zip"
	rubybuild_cmd_zip := exec.Command("curl", "-L", rubybuild_zip_url, "-o", "package.zip")
	err = rubybuild_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rubybuild_bin_url := "https://github.com/rbenv/ruby-build/archive/refs/tags/v20240903.bin"
	rubybuild_cmd_bin := exec.Command("curl", "-L", rubybuild_bin_url, "-o", "binary.bin")
	err = rubybuild_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rubybuild_src_url := "https://github.com/rbenv/ruby-build/archive/refs/tags/v20240903.src.tar.gz"
	rubybuild_cmd_src := exec.Command("curl", "-L", rubybuild_src_url, "-o", "source.tar.gz")
	err = rubybuild_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rubybuild_cmd_direct := exec.Command("./binary")
	err = rubybuild_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
