package main

// Jq - Lightweight and flexible command-line JSON processor
// Homepage: https://jqlang.github.io/jq/

import (
	"fmt"
	
	"os/exec"
)

func installJq() {
	// Método 1: Descargar y extraer .tar.gz
	jq_tar_url := "https://github.com/jqlang/jq/releases/download/jq-1.7.1/jq-1.7.1.tar.gz"
	jq_cmd_tar := exec.Command("curl", "-L", jq_tar_url, "-o", "package.tar.gz")
	err := jq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jq_zip_url := "https://github.com/jqlang/jq/releases/download/jq-1.7.1/jq-1.7.1.zip"
	jq_cmd_zip := exec.Command("curl", "-L", jq_zip_url, "-o", "package.zip")
	err = jq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jq_bin_url := "https://github.com/jqlang/jq/releases/download/jq-1.7.1/jq-1.7.1.bin"
	jq_cmd_bin := exec.Command("curl", "-L", jq_bin_url, "-o", "binary.bin")
	err = jq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jq_src_url := "https://github.com/jqlang/jq/releases/download/jq-1.7.1/jq-1.7.1.src.tar.gz"
	jq_cmd_src := exec.Command("curl", "-L", jq_src_url, "-o", "source.tar.gz")
	err = jq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jq_cmd_direct := exec.Command("./binary")
	err = jq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: oniguruma")
exec.Command("latte", "install", "oniguruma")
}
