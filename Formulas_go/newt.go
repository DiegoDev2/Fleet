package main

// Newt - Library for color text mode, widget based user interfaces
// Homepage: https://pagure.io/newt

import (
	"fmt"
	
	"os/exec"
)

func installNewt() {
	// Método 1: Descargar y extraer .tar.gz
	newt_tar_url := "https://releases.pagure.org/newt/newt-0.52.24.tar.gz"
	newt_cmd_tar := exec.Command("curl", "-L", newt_tar_url, "-o", "package.tar.gz")
	err := newt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	newt_zip_url := "https://releases.pagure.org/newt/newt-0.52.24.zip"
	newt_cmd_zip := exec.Command("curl", "-L", newt_zip_url, "-o", "package.zip")
	err = newt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	newt_bin_url := "https://releases.pagure.org/newt/newt-0.52.24.bin"
	newt_cmd_bin := exec.Command("curl", "-L", newt_bin_url, "-o", "binary.bin")
	err = newt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	newt_src_url := "https://releases.pagure.org/newt/newt-0.52.24.src.tar.gz"
	newt_cmd_src := exec.Command("curl", "-L", newt_src_url, "-o", "source.tar.gz")
	err = newt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	newt_cmd_direct := exec.Command("./binary")
	err = newt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: popt")
exec.Command("latte", "install", "popt")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: s-lang")
exec.Command("latte", "install", "s-lang")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
