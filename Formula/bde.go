package main

// Bde - Basic Development Environment: foundational C++ libraries used at Bloomberg
// Homepage: https://github.com/bloomberg/bde

import (
	"fmt"
	
	"os/exec"
)

func installBde() {
	// Método 1: Descargar y extraer .tar.gz
	bde_tar_url := "https://github.com/bloomberg/bde/archive/refs/tags/4.8.0.0.tar.gz"
	bde_cmd_tar := exec.Command("curl", "-L", bde_tar_url, "-o", "package.tar.gz")
	err := bde_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bde_zip_url := "https://github.com/bloomberg/bde/archive/refs/tags/4.8.0.0.zip"
	bde_cmd_zip := exec.Command("curl", "-L", bde_zip_url, "-o", "package.zip")
	err = bde_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bde_bin_url := "https://github.com/bloomberg/bde/archive/refs/tags/4.8.0.0.bin"
	bde_cmd_bin := exec.Command("curl", "-L", bde_bin_url, "-o", "binary.bin")
	err = bde_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bde_src_url := "https://github.com/bloomberg/bde/archive/refs/tags/4.8.0.0.src.tar.gz"
	bde_cmd_src := exec.Command("curl", "-L", bde_src_url, "-o", "source.tar.gz")
	err = bde_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bde_cmd_direct := exec.Command("./binary")
	err = bde_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
