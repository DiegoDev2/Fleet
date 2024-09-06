package main

// Precomp - Command-line precompressor to achieve better compression
// Homepage: http://schnaader.info/precomp.php

import (
	"fmt"
	
	"os/exec"
)

func installPrecomp() {
	// Método 1: Descargar y extraer .tar.gz
	precomp_tar_url := "https://github.com/schnaader/precomp-cpp/archive/refs/tags/v0.4.7.tar.gz"
	precomp_cmd_tar := exec.Command("curl", "-L", precomp_tar_url, "-o", "package.tar.gz")
	err := precomp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	precomp_zip_url := "https://github.com/schnaader/precomp-cpp/archive/refs/tags/v0.4.7.zip"
	precomp_cmd_zip := exec.Command("curl", "-L", precomp_zip_url, "-o", "package.zip")
	err = precomp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	precomp_bin_url := "https://github.com/schnaader/precomp-cpp/archive/refs/tags/v0.4.7.bin"
	precomp_cmd_bin := exec.Command("curl", "-L", precomp_bin_url, "-o", "binary.bin")
	err = precomp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	precomp_src_url := "https://github.com/schnaader/precomp-cpp/archive/refs/tags/v0.4.7.src.tar.gz"
	precomp_cmd_src := exec.Command("curl", "-L", precomp_src_url, "-o", "source.tar.gz")
	err = precomp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	precomp_cmd_direct := exec.Command("./binary")
	err = precomp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
