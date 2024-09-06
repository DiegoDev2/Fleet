package main

// RubyAT27 - Powerful, clean, object-oriented scripting language
// Homepage: https://www.ruby-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installRubyAT27() {
	// Método 1: Descargar y extraer .tar.gz
	rubyat27_tar_url := "https://cache.ruby-lang.org/pub/ruby/2.7/ruby-2.7.8.tar.xz"
	rubyat27_cmd_tar := exec.Command("curl", "-L", rubyat27_tar_url, "-o", "package.tar.gz")
	err := rubyat27_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rubyat27_zip_url := "https://cache.ruby-lang.org/pub/ruby/2.7/ruby-2.7.8.tar.xz"
	rubyat27_cmd_zip := exec.Command("curl", "-L", rubyat27_zip_url, "-o", "package.zip")
	err = rubyat27_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rubyat27_bin_url := "https://cache.ruby-lang.org/pub/ruby/2.7/ruby-2.7.8.tar.xz"
	rubyat27_cmd_bin := exec.Command("curl", "-L", rubyat27_bin_url, "-o", "binary.bin")
	err = rubyat27_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rubyat27_src_url := "https://cache.ruby-lang.org/pub/ruby/2.7/ruby-2.7.8.tar.xz"
	rubyat27_cmd_src := exec.Command("curl", "-L", rubyat27_src_url, "-o", "source.tar.gz")
	err = rubyat27_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rubyat27_cmd_direct := exec.Command("./binary")
	err = rubyat27_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
