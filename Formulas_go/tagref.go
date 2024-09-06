package main

// Tagref - Refer to other locations in your codebase
// Homepage: https://github.com/stepchowfun/tagref

import (
	"fmt"
	
	"os/exec"
)

func installTagref() {
	// Método 1: Descargar y extraer .tar.gz
	tagref_tar_url := "https://github.com/stepchowfun/tagref/archive/refs/tags/v1.10.0.tar.gz"
	tagref_cmd_tar := exec.Command("curl", "-L", tagref_tar_url, "-o", "package.tar.gz")
	err := tagref_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tagref_zip_url := "https://github.com/stepchowfun/tagref/archive/refs/tags/v1.10.0.zip"
	tagref_cmd_zip := exec.Command("curl", "-L", tagref_zip_url, "-o", "package.zip")
	err = tagref_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tagref_bin_url := "https://github.com/stepchowfun/tagref/archive/refs/tags/v1.10.0.bin"
	tagref_cmd_bin := exec.Command("curl", "-L", tagref_bin_url, "-o", "binary.bin")
	err = tagref_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tagref_src_url := "https://github.com/stepchowfun/tagref/archive/refs/tags/v1.10.0.src.tar.gz"
	tagref_cmd_src := exec.Command("curl", "-L", tagref_src_url, "-o", "source.tar.gz")
	err = tagref_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tagref_cmd_direct := exec.Command("./binary")
	err = tagref_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
