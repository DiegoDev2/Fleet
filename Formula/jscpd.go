package main

// Jscpd - Copy/paste detector for programming source code
// Homepage: https://github.com/kucherenko/jscpd

import (
	"fmt"
	
	"os/exec"
)

func installJscpd() {
	// Método 1: Descargar y extraer .tar.gz
	jscpd_tar_url := "https://registry.npmjs.org/jscpd/-/jscpd-4.0.5.tgz"
	jscpd_cmd_tar := exec.Command("curl", "-L", jscpd_tar_url, "-o", "package.tar.gz")
	err := jscpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jscpd_zip_url := "https://registry.npmjs.org/jscpd/-/jscpd-4.0.5.tgz"
	jscpd_cmd_zip := exec.Command("curl", "-L", jscpd_zip_url, "-o", "package.zip")
	err = jscpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jscpd_bin_url := "https://registry.npmjs.org/jscpd/-/jscpd-4.0.5.tgz"
	jscpd_cmd_bin := exec.Command("curl", "-L", jscpd_bin_url, "-o", "binary.bin")
	err = jscpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jscpd_src_url := "https://registry.npmjs.org/jscpd/-/jscpd-4.0.5.tgz"
	jscpd_cmd_src := exec.Command("curl", "-L", jscpd_src_url, "-o", "source.tar.gz")
	err = jscpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jscpd_cmd_direct := exec.Command("./binary")
	err = jscpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
