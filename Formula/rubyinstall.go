package main

// RubyInstall - Install Ruby, JRuby, Rubinius, TruffleRuby, or mruby
// Homepage: https://github.com/postmodern/ruby-install

import (
	"fmt"
	
	"os/exec"
)

func installRubyInstall() {
	// Método 1: Descargar y extraer .tar.gz
	rubyinstall_tar_url := "https://github.com/postmodern/ruby-install/releases/download/v0.9.3/ruby-install-0.9.3.tar.gz"
	rubyinstall_cmd_tar := exec.Command("curl", "-L", rubyinstall_tar_url, "-o", "package.tar.gz")
	err := rubyinstall_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rubyinstall_zip_url := "https://github.com/postmodern/ruby-install/releases/download/v0.9.3/ruby-install-0.9.3.zip"
	rubyinstall_cmd_zip := exec.Command("curl", "-L", rubyinstall_zip_url, "-o", "package.zip")
	err = rubyinstall_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rubyinstall_bin_url := "https://github.com/postmodern/ruby-install/releases/download/v0.9.3/ruby-install-0.9.3.bin"
	rubyinstall_cmd_bin := exec.Command("curl", "-L", rubyinstall_bin_url, "-o", "binary.bin")
	err = rubyinstall_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rubyinstall_src_url := "https://github.com/postmodern/ruby-install/releases/download/v0.9.3/ruby-install-0.9.3.src.tar.gz"
	rubyinstall_cmd_src := exec.Command("curl", "-L", rubyinstall_src_url, "-o", "source.tar.gz")
	err = rubyinstall_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rubyinstall_cmd_direct := exec.Command("./binary")
	err = rubyinstall_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
