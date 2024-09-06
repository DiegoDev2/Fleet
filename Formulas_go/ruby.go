package main

// Ruby - Powerful, clean, object-oriented scripting language
// Homepage: https://www.ruby-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installRuby() {
	// Método 1: Descargar y extraer .tar.gz
	ruby_tar_url := "https://cache.ruby-lang.org/pub/ruby/3.3/ruby-3.3.5.tar.gz"
	ruby_cmd_tar := exec.Command("curl", "-L", ruby_tar_url, "-o", "package.tar.gz")
	err := ruby_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ruby_zip_url := "https://cache.ruby-lang.org/pub/ruby/3.3/ruby-3.3.5.zip"
	ruby_cmd_zip := exec.Command("curl", "-L", ruby_zip_url, "-o", "package.zip")
	err = ruby_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ruby_bin_url := "https://cache.ruby-lang.org/pub/ruby/3.3/ruby-3.3.5.bin"
	ruby_cmd_bin := exec.Command("curl", "-L", ruby_bin_url, "-o", "binary.bin")
	err = ruby_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ruby_src_url := "https://cache.ruby-lang.org/pub/ruby/3.3/ruby-3.3.5.src.tar.gz"
	ruby_cmd_src := exec.Command("curl", "-L", ruby_src_url, "-o", "source.tar.gz")
	err = ruby_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ruby_cmd_direct := exec.Command("./binary")
	err = ruby_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
