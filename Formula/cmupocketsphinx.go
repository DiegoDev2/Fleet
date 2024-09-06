package main

// CmuPocketsphinx - Lightweight speech recognition engine for mobile devices
// Homepage: https://cmusphinx.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installCmuPocketsphinx() {
	// Método 1: Descargar y extraer .tar.gz
	cmupocketsphinx_tar_url := "https://github.com/cmusphinx/pocketsphinx/archive/refs/tags/v5.0.3.tar.gz"
	cmupocketsphinx_cmd_tar := exec.Command("curl", "-L", cmupocketsphinx_tar_url, "-o", "package.tar.gz")
	err := cmupocketsphinx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmupocketsphinx_zip_url := "https://github.com/cmusphinx/pocketsphinx/archive/refs/tags/v5.0.3.zip"
	cmupocketsphinx_cmd_zip := exec.Command("curl", "-L", cmupocketsphinx_zip_url, "-o", "package.zip")
	err = cmupocketsphinx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmupocketsphinx_bin_url := "https://github.com/cmusphinx/pocketsphinx/archive/refs/tags/v5.0.3.bin"
	cmupocketsphinx_cmd_bin := exec.Command("curl", "-L", cmupocketsphinx_bin_url, "-o", "binary.bin")
	err = cmupocketsphinx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmupocketsphinx_src_url := "https://github.com/cmusphinx/pocketsphinx/archive/refs/tags/v5.0.3.src.tar.gz"
	cmupocketsphinx_cmd_src := exec.Command("curl", "-L", cmupocketsphinx_src_url, "-o", "source.tar.gz")
	err = cmupocketsphinx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmupocketsphinx_cmd_direct := exec.Command("./binary")
	err = cmupocketsphinx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
