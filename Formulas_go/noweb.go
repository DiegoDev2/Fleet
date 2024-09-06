package main

// Noweb - WEB-like literate-programming tool
// Homepage: https://www.cs.tufts.edu/~nr/noweb/

import (
	"fmt"
	
	"os/exec"
)

func installNoweb() {
	// Método 1: Descargar y extraer .tar.gz
	noweb_tar_url := "https://github.com/nrnrnr/noweb/archive/refs/tags/v2_13.tar.gz"
	noweb_cmd_tar := exec.Command("curl", "-L", noweb_tar_url, "-o", "package.tar.gz")
	err := noweb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	noweb_zip_url := "https://github.com/nrnrnr/noweb/archive/refs/tags/v2_13.zip"
	noweb_cmd_zip := exec.Command("curl", "-L", noweb_zip_url, "-o", "package.zip")
	err = noweb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	noweb_bin_url := "https://github.com/nrnrnr/noweb/archive/refs/tags/v2_13.bin"
	noweb_cmd_bin := exec.Command("curl", "-L", noweb_bin_url, "-o", "binary.bin")
	err = noweb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	noweb_src_url := "https://github.com/nrnrnr/noweb/archive/refs/tags/v2_13.src.tar.gz"
	noweb_cmd_src := exec.Command("curl", "-L", noweb_src_url, "-o", "source.tar.gz")
	err = noweb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	noweb_cmd_direct := exec.Command("./binary")
	err = noweb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
	fmt.Println("Instalando dependencia: icon")
exec.Command("latte", "install", "icon")
}
