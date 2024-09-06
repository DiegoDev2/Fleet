package main

// Ocicl - OCI-based ASDF system distribution and management tool for Common Lisp
// Homepage: https://github.com/ocicl/ocicl

import (
	"fmt"
	
	"os/exec"
)

func installOcicl() {
	// Método 1: Descargar y extraer .tar.gz
	ocicl_tar_url := "https://github.com/ocicl/ocicl/archive/refs/tags/v2.4.3.tar.gz"
	ocicl_cmd_tar := exec.Command("curl", "-L", ocicl_tar_url, "-o", "package.tar.gz")
	err := ocicl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ocicl_zip_url := "https://github.com/ocicl/ocicl/archive/refs/tags/v2.4.3.zip"
	ocicl_cmd_zip := exec.Command("curl", "-L", ocicl_zip_url, "-o", "package.zip")
	err = ocicl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ocicl_bin_url := "https://github.com/ocicl/ocicl/archive/refs/tags/v2.4.3.bin"
	ocicl_cmd_bin := exec.Command("curl", "-L", ocicl_bin_url, "-o", "binary.bin")
	err = ocicl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ocicl_src_url := "https://github.com/ocicl/ocicl/archive/refs/tags/v2.4.3.src.tar.gz"
	ocicl_cmd_src := exec.Command("curl", "-L", ocicl_src_url, "-o", "source.tar.gz")
	err = ocicl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ocicl_cmd_direct := exec.Command("./binary")
	err = ocicl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: oras")
	exec.Command("latte", "install", "oras").Run()
	fmt.Println("Instalando dependencia: sbcl")
	exec.Command("latte", "install", "sbcl").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
