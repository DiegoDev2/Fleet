package main

// RbenvBundler - Makes shims aware of bundle install paths
// Homepage: https://github.com/carsomyr/rbenv-bundler

import (
	"fmt"
	
	"os/exec"
)

func installRbenvBundler() {
	// Método 1: Descargar y extraer .tar.gz
	rbenvbundler_tar_url := "https://github.com/carsomyr/rbenv-bundler/archive/refs/tags/1.0.1.tar.gz"
	rbenvbundler_cmd_tar := exec.Command("curl", "-L", rbenvbundler_tar_url, "-o", "package.tar.gz")
	err := rbenvbundler_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rbenvbundler_zip_url := "https://github.com/carsomyr/rbenv-bundler/archive/refs/tags/1.0.1.zip"
	rbenvbundler_cmd_zip := exec.Command("curl", "-L", rbenvbundler_zip_url, "-o", "package.zip")
	err = rbenvbundler_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rbenvbundler_bin_url := "https://github.com/carsomyr/rbenv-bundler/archive/refs/tags/1.0.1.bin"
	rbenvbundler_cmd_bin := exec.Command("curl", "-L", rbenvbundler_bin_url, "-o", "binary.bin")
	err = rbenvbundler_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rbenvbundler_src_url := "https://github.com/carsomyr/rbenv-bundler/archive/refs/tags/1.0.1.src.tar.gz"
	rbenvbundler_cmd_src := exec.Command("curl", "-L", rbenvbundler_src_url, "-o", "source.tar.gz")
	err = rbenvbundler_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rbenvbundler_cmd_direct := exec.Command("./binary")
	err = rbenvbundler_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rbenv")
	exec.Command("latte", "install", "rbenv").Run()
}
