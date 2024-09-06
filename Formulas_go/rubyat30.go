package main

// RubyAT30 - Powerful, clean, object-oriented scripting language
// Homepage: https://www.ruby-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installRubyAT30() {
	// Método 1: Descargar y extraer .tar.gz
	rubyat30_tar_url := "https://cache.ruby-lang.org/pub/ruby/3.0/ruby-3.0.7.tar.xz"
	rubyat30_cmd_tar := exec.Command("curl", "-L", rubyat30_tar_url, "-o", "package.tar.gz")
	err := rubyat30_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rubyat30_zip_url := "https://cache.ruby-lang.org/pub/ruby/3.0/ruby-3.0.7.tar.xz"
	rubyat30_cmd_zip := exec.Command("curl", "-L", rubyat30_zip_url, "-o", "package.zip")
	err = rubyat30_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rubyat30_bin_url := "https://cache.ruby-lang.org/pub/ruby/3.0/ruby-3.0.7.tar.xz"
	rubyat30_cmd_bin := exec.Command("curl", "-L", rubyat30_bin_url, "-o", "binary.bin")
	err = rubyat30_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rubyat30_src_url := "https://cache.ruby-lang.org/pub/ruby/3.0/ruby-3.0.7.tar.xz"
	rubyat30_cmd_src := exec.Command("curl", "-L", rubyat30_src_url, "-o", "source.tar.gz")
	err = rubyat30_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rubyat30_cmd_direct := exec.Command("./binary")
	err = rubyat30_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
