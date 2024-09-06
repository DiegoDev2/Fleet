package main

// Mako - Production-grade web bundler based on Rust
// Homepage: https://makojs.dev

import (
	"fmt"
	
	"os/exec"
)

func installMako() {
	// Método 1: Descargar y extraer .tar.gz
	mako_tar_url := "https://registry.npmjs.org/@umijs/mako/-/mako-0.8.8.tgz"
	mako_cmd_tar := exec.Command("curl", "-L", mako_tar_url, "-o", "package.tar.gz")
	err := mako_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mako_zip_url := "https://registry.npmjs.org/@umijs/mako/-/mako-0.8.8.tgz"
	mako_cmd_zip := exec.Command("curl", "-L", mako_zip_url, "-o", "package.zip")
	err = mako_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mako_bin_url := "https://registry.npmjs.org/@umijs/mako/-/mako-0.8.8.tgz"
	mako_cmd_bin := exec.Command("curl", "-L", mako_bin_url, "-o", "binary.bin")
	err = mako_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mako_src_url := "https://registry.npmjs.org/@umijs/mako/-/mako-0.8.8.tgz"
	mako_cmd_src := exec.Command("curl", "-L", mako_src_url, "-o", "source.tar.gz")
	err = mako_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mako_cmd_direct := exec.Command("./binary")
	err = mako_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
