package main

// RubyAT32 - Powerful, clean, object-oriented scripting language
// Homepage: https://www.ruby-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installRubyAT32() {
	// Método 1: Descargar y extraer .tar.gz
	rubyat32_tar_url := "https://cache.ruby-lang.org/pub/ruby/3.2/ruby-3.2.5.tar.gz"
	rubyat32_cmd_tar := exec.Command("curl", "-L", rubyat32_tar_url, "-o", "package.tar.gz")
	err := rubyat32_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rubyat32_zip_url := "https://cache.ruby-lang.org/pub/ruby/3.2/ruby-3.2.5.zip"
	rubyat32_cmd_zip := exec.Command("curl", "-L", rubyat32_zip_url, "-o", "package.zip")
	err = rubyat32_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rubyat32_bin_url := "https://cache.ruby-lang.org/pub/ruby/3.2/ruby-3.2.5.bin"
	rubyat32_cmd_bin := exec.Command("curl", "-L", rubyat32_bin_url, "-o", "binary.bin")
	err = rubyat32_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rubyat32_src_url := "https://cache.ruby-lang.org/pub/ruby/3.2/ruby-3.2.5.src.tar.gz"
	rubyat32_cmd_src := exec.Command("curl", "-L", rubyat32_src_url, "-o", "source.tar.gz")
	err = rubyat32_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rubyat32_cmd_direct := exec.Command("./binary")
	err = rubyat32_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
