package main

// Ponysay - Cowsay but with ponies
// Homepage: https://github.com/erkin/ponysay/

import (
	"fmt"
	
	"os/exec"
)

func installPonysay() {
	// Método 1: Descargar y extraer .tar.gz
	ponysay_tar_url := "https://github.com/erkin/ponysay/archive/refs/tags/3.0.3.tar.gz"
	ponysay_cmd_tar := exec.Command("curl", "-L", ponysay_tar_url, "-o", "package.tar.gz")
	err := ponysay_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ponysay_zip_url := "https://github.com/erkin/ponysay/archive/refs/tags/3.0.3.zip"
	ponysay_cmd_zip := exec.Command("curl", "-L", ponysay_zip_url, "-o", "package.zip")
	err = ponysay_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ponysay_bin_url := "https://github.com/erkin/ponysay/archive/refs/tags/3.0.3.bin"
	ponysay_cmd_bin := exec.Command("curl", "-L", ponysay_bin_url, "-o", "binary.bin")
	err = ponysay_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ponysay_src_url := "https://github.com/erkin/ponysay/archive/refs/tags/3.0.3.src.tar.gz"
	ponysay_cmd_src := exec.Command("curl", "-L", ponysay_src_url, "-o", "source.tar.gz")
	err = ponysay_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ponysay_cmd_direct := exec.Command("./binary")
	err = ponysay_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gzip")
	exec.Command("latte", "install", "gzip").Run()
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
}
