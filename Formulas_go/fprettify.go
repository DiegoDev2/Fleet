package main

// Fprettify - Auto-formatter for modern fortran source code
// Homepage: https://github.com/fortran-lang/fprettify/

import (
	"fmt"
	
	"os/exec"
)

func installFprettify() {
	// Método 1: Descargar y extraer .tar.gz
	fprettify_tar_url := "https://github.com/fortran-lang/fprettify/archive/refs/tags/v0.3.7.tar.gz"
	fprettify_cmd_tar := exec.Command("curl", "-L", fprettify_tar_url, "-o", "package.tar.gz")
	err := fprettify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fprettify_zip_url := "https://github.com/fortran-lang/fprettify/archive/refs/tags/v0.3.7.zip"
	fprettify_cmd_zip := exec.Command("curl", "-L", fprettify_zip_url, "-o", "package.zip")
	err = fprettify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fprettify_bin_url := "https://github.com/fortran-lang/fprettify/archive/refs/tags/v0.3.7.bin"
	fprettify_cmd_bin := exec.Command("curl", "-L", fprettify_bin_url, "-o", "binary.bin")
	err = fprettify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fprettify_src_url := "https://github.com/fortran-lang/fprettify/archive/refs/tags/v0.3.7.src.tar.gz"
	fprettify_cmd_src := exec.Command("curl", "-L", fprettify_src_url, "-o", "source.tar.gz")
	err = fprettify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fprettify_cmd_direct := exec.Command("./binary")
	err = fprettify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
