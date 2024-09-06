package main

// Jnv - Interactive JSON filter using jq
// Homepage: https://github.com/ynqa/jnv

import (
	"fmt"
	
	"os/exec"
)

func installJnv() {
	// Método 1: Descargar y extraer .tar.gz
	jnv_tar_url := "https://github.com/ynqa/jnv/archive/refs/tags/v0.4.0.tar.gz"
	jnv_cmd_tar := exec.Command("curl", "-L", jnv_tar_url, "-o", "package.tar.gz")
	err := jnv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jnv_zip_url := "https://github.com/ynqa/jnv/archive/refs/tags/v0.4.0.zip"
	jnv_cmd_zip := exec.Command("curl", "-L", jnv_zip_url, "-o", "package.zip")
	err = jnv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jnv_bin_url := "https://github.com/ynqa/jnv/archive/refs/tags/v0.4.0.bin"
	jnv_cmd_bin := exec.Command("curl", "-L", jnv_bin_url, "-o", "binary.bin")
	err = jnv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jnv_src_url := "https://github.com/ynqa/jnv/archive/refs/tags/v0.4.0.src.tar.gz"
	jnv_cmd_src := exec.Command("curl", "-L", jnv_src_url, "-o", "source.tar.gz")
	err = jnv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jnv_cmd_direct := exec.Command("./binary")
	err = jnv_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
