package main

// Intercal - Esoteric, parody programming language
// Homepage: http://catb.org/~esr/intercal/

import (
	"fmt"
	
	"os/exec"
)

func installIntercal() {
	// Método 1: Descargar y extraer .tar.gz
	intercal_tar_url := "http://catb.org/~esr/intercal/intercal-0.31.tar.gz"
	intercal_cmd_tar := exec.Command("curl", "-L", intercal_tar_url, "-o", "package.tar.gz")
	err := intercal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	intercal_zip_url := "http://catb.org/~esr/intercal/intercal-0.31.zip"
	intercal_cmd_zip := exec.Command("curl", "-L", intercal_zip_url, "-o", "package.zip")
	err = intercal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	intercal_bin_url := "http://catb.org/~esr/intercal/intercal-0.31.bin"
	intercal_cmd_bin := exec.Command("curl", "-L", intercal_bin_url, "-o", "binary.bin")
	err = intercal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	intercal_src_url := "http://catb.org/~esr/intercal/intercal-0.31.src.tar.gz"
	intercal_cmd_src := exec.Command("curl", "-L", intercal_src_url, "-o", "source.tar.gz")
	err = intercal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	intercal_cmd_direct := exec.Command("./binary")
	err = intercal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
