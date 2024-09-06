package main

// Jimtcl - Small footprint implementation of Tcl
// Homepage: https://jim.tcl.tk/index.html

import (
	"fmt"
	
	"os/exec"
)

func installJimtcl() {
	// Método 1: Descargar y extraer .tar.gz
	jimtcl_tar_url := "https://github.com/msteveb/jimtcl/archive/refs/tags/0.83.tar.gz"
	jimtcl_cmd_tar := exec.Command("curl", "-L", jimtcl_tar_url, "-o", "package.tar.gz")
	err := jimtcl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jimtcl_zip_url := "https://github.com/msteveb/jimtcl/archive/refs/tags/0.83.zip"
	jimtcl_cmd_zip := exec.Command("curl", "-L", jimtcl_zip_url, "-o", "package.zip")
	err = jimtcl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jimtcl_bin_url := "https://github.com/msteveb/jimtcl/archive/refs/tags/0.83.bin"
	jimtcl_cmd_bin := exec.Command("curl", "-L", jimtcl_bin_url, "-o", "binary.bin")
	err = jimtcl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jimtcl_src_url := "https://github.com/msteveb/jimtcl/archive/refs/tags/0.83.src.tar.gz"
	jimtcl_cmd_src := exec.Command("curl", "-L", jimtcl_src_url, "-o", "source.tar.gz")
	err = jimtcl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jimtcl_cmd_direct := exec.Command("./binary")
	err = jimtcl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
