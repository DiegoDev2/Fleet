package main

// CucumberRuby - Cucumber for Ruby
// Homepage: https://cucumber.io

import (
	"fmt"
	
	"os/exec"
)

func installCucumberRuby() {
	// Método 1: Descargar y extraer .tar.gz
	cucumberruby_tar_url := "https://github.com/cucumber/cucumber-ruby/archive/refs/tags/v9.2.0.tar.gz"
	cucumberruby_cmd_tar := exec.Command("curl", "-L", cucumberruby_tar_url, "-o", "package.tar.gz")
	err := cucumberruby_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cucumberruby_zip_url := "https://github.com/cucumber/cucumber-ruby/archive/refs/tags/v9.2.0.zip"
	cucumberruby_cmd_zip := exec.Command("curl", "-L", cucumberruby_zip_url, "-o", "package.zip")
	err = cucumberruby_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cucumberruby_bin_url := "https://github.com/cucumber/cucumber-ruby/archive/refs/tags/v9.2.0.bin"
	cucumberruby_cmd_bin := exec.Command("curl", "-L", cucumberruby_bin_url, "-o", "binary.bin")
	err = cucumberruby_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cucumberruby_src_url := "https://github.com/cucumber/cucumber-ruby/archive/refs/tags/v9.2.0.src.tar.gz"
	cucumberruby_cmd_src := exec.Command("curl", "-L", cucumberruby_src_url, "-o", "source.tar.gz")
	err = cucumberruby_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cucumberruby_cmd_direct := exec.Command("./binary")
	err = cucumberruby_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ruby")
exec.Command("latte", "install", "ruby")
}
