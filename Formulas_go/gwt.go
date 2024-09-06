package main

// Gwt - Google web toolkit
// Homepage: https://www.gwtproject.org/

import (
	"fmt"
	
	"os/exec"
)

func installGwt() {
	// Método 1: Descargar y extraer .tar.gz
	gwt_tar_url := "https://github.com/gwtproject/gwt/releases/download/2.11.0/gwt-2.11.0.zip"
	gwt_cmd_tar := exec.Command("curl", "-L", gwt_tar_url, "-o", "package.tar.gz")
	err := gwt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gwt_zip_url := "https://github.com/gwtproject/gwt/releases/download/2.11.0/gwt-2.11.0.zip"
	gwt_cmd_zip := exec.Command("curl", "-L", gwt_zip_url, "-o", "package.zip")
	err = gwt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gwt_bin_url := "https://github.com/gwtproject/gwt/releases/download/2.11.0/gwt-2.11.0.zip"
	gwt_cmd_bin := exec.Command("curl", "-L", gwt_bin_url, "-o", "binary.bin")
	err = gwt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gwt_src_url := "https://github.com/gwtproject/gwt/releases/download/2.11.0/gwt-2.11.0.zip"
	gwt_cmd_src := exec.Command("curl", "-L", gwt_src_url, "-o", "source.tar.gz")
	err = gwt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gwt_cmd_direct := exec.Command("./binary")
	err = gwt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
