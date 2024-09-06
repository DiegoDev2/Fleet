package main

// RbenvBundlerRubyVersion - Pick a ruby version from bundler's Gemfile
// Homepage: https://github.com/aripollak/rbenv-bundler-ruby-version

import (
	"fmt"
	
	"os/exec"
)

func installRbenvBundlerRubyVersion() {
	// Método 1: Descargar y extraer .tar.gz
	rbenvbundlerrubyversion_tar_url := "https://github.com/aripollak/rbenv-bundler-ruby-version/archive/refs/tags/v1.0.0.tar.gz"
	rbenvbundlerrubyversion_cmd_tar := exec.Command("curl", "-L", rbenvbundlerrubyversion_tar_url, "-o", "package.tar.gz")
	err := rbenvbundlerrubyversion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rbenvbundlerrubyversion_zip_url := "https://github.com/aripollak/rbenv-bundler-ruby-version/archive/refs/tags/v1.0.0.zip"
	rbenvbundlerrubyversion_cmd_zip := exec.Command("curl", "-L", rbenvbundlerrubyversion_zip_url, "-o", "package.zip")
	err = rbenvbundlerrubyversion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rbenvbundlerrubyversion_bin_url := "https://github.com/aripollak/rbenv-bundler-ruby-version/archive/refs/tags/v1.0.0.bin"
	rbenvbundlerrubyversion_cmd_bin := exec.Command("curl", "-L", rbenvbundlerrubyversion_bin_url, "-o", "binary.bin")
	err = rbenvbundlerrubyversion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rbenvbundlerrubyversion_src_url := "https://github.com/aripollak/rbenv-bundler-ruby-version/archive/refs/tags/v1.0.0.src.tar.gz"
	rbenvbundlerrubyversion_cmd_src := exec.Command("curl", "-L", rbenvbundlerrubyversion_src_url, "-o", "source.tar.gz")
	err = rbenvbundlerrubyversion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rbenvbundlerrubyversion_cmd_direct := exec.Command("./binary")
	err = rbenvbundlerrubyversion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rbenv")
	exec.Command("latte", "install", "rbenv").Run()
}
