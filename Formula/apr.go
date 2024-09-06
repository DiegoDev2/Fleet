package main

// Apr - Apache Portable Runtime library
// Homepage: https://apr.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installApr() {
	// Método 1: Descargar y extraer .tar.gz
	apr_tar_url := "https://www.apache.org/dyn/closer.lua?path=apr/apr-1.7.5.tar.bz2"
	apr_cmd_tar := exec.Command("curl", "-L", apr_tar_url, "-o", "package.tar.gz")
	err := apr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apr_zip_url := "https://www.apache.org/dyn/closer.lua?path=apr/apr-1.7.5.tar.bz2"
	apr_cmd_zip := exec.Command("curl", "-L", apr_zip_url, "-o", "package.zip")
	err = apr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apr_bin_url := "https://www.apache.org/dyn/closer.lua?path=apr/apr-1.7.5.tar.bz2"
	apr_cmd_bin := exec.Command("curl", "-L", apr_bin_url, "-o", "binary.bin")
	err = apr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apr_src_url := "https://www.apache.org/dyn/closer.lua?path=apr/apr-1.7.5.tar.bz2"
	apr_cmd_src := exec.Command("curl", "-L", apr_src_url, "-o", "source.tar.gz")
	err = apr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apr_cmd_direct := exec.Command("./binary")
	err = apr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
