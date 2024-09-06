package main

// Ginac - Not a Computer algebra system
// Homepage: https://www.ginac.de/

import (
	"fmt"
	
	"os/exec"
)

func installGinac() {
	// Método 1: Descargar y extraer .tar.gz
	ginac_tar_url := "https://www.ginac.de/ginac-1.8.7.tar.bz2"
	ginac_cmd_tar := exec.Command("curl", "-L", ginac_tar_url, "-o", "package.tar.gz")
	err := ginac_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ginac_zip_url := "https://www.ginac.de/ginac-1.8.7.tar.bz2"
	ginac_cmd_zip := exec.Command("curl", "-L", ginac_zip_url, "-o", "package.zip")
	err = ginac_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ginac_bin_url := "https://www.ginac.de/ginac-1.8.7.tar.bz2"
	ginac_cmd_bin := exec.Command("curl", "-L", ginac_bin_url, "-o", "binary.bin")
	err = ginac_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ginac_src_url := "https://www.ginac.de/ginac-1.8.7.tar.bz2"
	ginac_cmd_src := exec.Command("curl", "-L", ginac_src_url, "-o", "source.tar.gz")
	err = ginac_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ginac_cmd_direct := exec.Command("./binary")
	err = ginac_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cln")
	exec.Command("latte", "install", "cln").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
