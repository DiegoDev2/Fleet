package main

// Clens - Library to help port code from OpenBSD to other operating systems
// Homepage: https://github.com/conformal/clens

import (
	"fmt"
	
	"os/exec"
)

func installClens() {
	// Método 1: Descargar y extraer .tar.gz
	clens_tar_url := "https://github.com/conformal/clens/archive/refs/tags/CLENS_0_7_0.tar.gz"
	clens_cmd_tar := exec.Command("curl", "-L", clens_tar_url, "-o", "package.tar.gz")
	err := clens_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clens_zip_url := "https://github.com/conformal/clens/archive/refs/tags/CLENS_0_7_0.zip"
	clens_cmd_zip := exec.Command("curl", "-L", clens_zip_url, "-o", "package.zip")
	err = clens_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clens_bin_url := "https://github.com/conformal/clens/archive/refs/tags/CLENS_0_7_0.bin"
	clens_cmd_bin := exec.Command("curl", "-L", clens_bin_url, "-o", "binary.bin")
	err = clens_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clens_src_url := "https://github.com/conformal/clens/archive/refs/tags/CLENS_0_7_0.src.tar.gz"
	clens_cmd_src := exec.Command("curl", "-L", clens_src_url, "-o", "source.tar.gz")
	err = clens_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clens_cmd_direct := exec.Command("./binary")
	err = clens_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libbsd")
exec.Command("latte", "install", "libbsd")
}
