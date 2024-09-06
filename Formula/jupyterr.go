package main

// JupyterR - R support for Jupyter
// Homepage: https://irkernel.github.io

import (
	"fmt"
	
	"os/exec"
)

func installJupyterR() {
	// Método 1: Descargar y extraer .tar.gz
	jupyterr_tar_url := "https://github.com/IRkernel/IRkernel/archive/refs/tags/1.3.2.tar.gz"
	jupyterr_cmd_tar := exec.Command("curl", "-L", jupyterr_tar_url, "-o", "package.tar.gz")
	err := jupyterr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jupyterr_zip_url := "https://github.com/IRkernel/IRkernel/archive/refs/tags/1.3.2.zip"
	jupyterr_cmd_zip := exec.Command("curl", "-L", jupyterr_zip_url, "-o", "package.zip")
	err = jupyterr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jupyterr_bin_url := "https://github.com/IRkernel/IRkernel/archive/refs/tags/1.3.2.bin"
	jupyterr_cmd_bin := exec.Command("curl", "-L", jupyterr_bin_url, "-o", "binary.bin")
	err = jupyterr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jupyterr_src_url := "https://github.com/IRkernel/IRkernel/archive/refs/tags/1.3.2.src.tar.gz"
	jupyterr_cmd_src := exec.Command("curl", "-L", jupyterr_src_url, "-o", "source.tar.gz")
	err = jupyterr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jupyterr_cmd_direct := exec.Command("./binary")
	err = jupyterr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: jupyterlab")
	exec.Command("latte", "install", "jupyterlab").Run()
	fmt.Println("Instalando dependencia: r")
	exec.Command("latte", "install", "r").Run()
	fmt.Println("Instalando dependencia: zeromq")
	exec.Command("latte", "install", "zeromq").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
