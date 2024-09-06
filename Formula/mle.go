package main

// Mle - Flexible terminal-based text editor
// Homepage: https://github.com/adsr/mle

import (
	"fmt"
	
	"os/exec"
)

func installMle() {
	// Método 1: Descargar y extraer .tar.gz
	mle_tar_url := "https://github.com/adsr/mle/archive/refs/tags/v1.7.2.tar.gz"
	mle_cmd_tar := exec.Command("curl", "-L", mle_tar_url, "-o", "package.tar.gz")
	err := mle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mle_zip_url := "https://github.com/adsr/mle/archive/refs/tags/v1.7.2.zip"
	mle_cmd_zip := exec.Command("curl", "-L", mle_zip_url, "-o", "package.zip")
	err = mle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mle_bin_url := "https://github.com/adsr/mle/archive/refs/tags/v1.7.2.bin"
	mle_cmd_bin := exec.Command("curl", "-L", mle_bin_url, "-o", "binary.bin")
	err = mle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mle_src_url := "https://github.com/adsr/mle/archive/refs/tags/v1.7.2.src.tar.gz"
	mle_cmd_src := exec.Command("curl", "-L", mle_src_url, "-o", "source.tar.gz")
	err = mle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mle_cmd_direct := exec.Command("./binary")
	err = mle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: uthash")
	exec.Command("latte", "install", "uthash").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
