package main

// Kuto - Reverse JS bundler
// Homepage: https://github.com/samthor/kuto

import (
	"fmt"
	
	"os/exec"
)

func installKuto() {
	// Método 1: Descargar y extraer .tar.gz
	kuto_tar_url := "https://registry.npmjs.org/kuto/-/kuto-0.3.6.tgz"
	kuto_cmd_tar := exec.Command("curl", "-L", kuto_tar_url, "-o", "package.tar.gz")
	err := kuto_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kuto_zip_url := "https://registry.npmjs.org/kuto/-/kuto-0.3.6.tgz"
	kuto_cmd_zip := exec.Command("curl", "-L", kuto_zip_url, "-o", "package.zip")
	err = kuto_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kuto_bin_url := "https://registry.npmjs.org/kuto/-/kuto-0.3.6.tgz"
	kuto_cmd_bin := exec.Command("curl", "-L", kuto_bin_url, "-o", "binary.bin")
	err = kuto_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kuto_src_url := "https://registry.npmjs.org/kuto/-/kuto-0.3.6.tgz"
	kuto_cmd_src := exec.Command("curl", "-L", kuto_src_url, "-o", "source.tar.gz")
	err = kuto_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kuto_cmd_direct := exec.Command("./binary")
	err = kuto_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
