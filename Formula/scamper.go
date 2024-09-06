package main

// Scamper - Advanced traceroute and network measurement utility
// Homepage: https://www.caida.org/catalog/software/scamper/

import (
	"fmt"
	
	"os/exec"
)

func installScamper() {
	// Método 1: Descargar y extraer .tar.gz
	scamper_tar_url := "https://www.caida.org/catalog/software/scamper/code/scamper-cvs-20240813.tar.gz"
	scamper_cmd_tar := exec.Command("curl", "-L", scamper_tar_url, "-o", "package.tar.gz")
	err := scamper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scamper_zip_url := "https://www.caida.org/catalog/software/scamper/code/scamper-cvs-20240813.zip"
	scamper_cmd_zip := exec.Command("curl", "-L", scamper_zip_url, "-o", "package.zip")
	err = scamper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scamper_bin_url := "https://www.caida.org/catalog/software/scamper/code/scamper-cvs-20240813.bin"
	scamper_cmd_bin := exec.Command("curl", "-L", scamper_bin_url, "-o", "binary.bin")
	err = scamper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scamper_src_url := "https://www.caida.org/catalog/software/scamper/code/scamper-cvs-20240813.src.tar.gz"
	scamper_cmd_src := exec.Command("curl", "-L", scamper_src_url, "-o", "source.tar.gz")
	err = scamper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scamper_cmd_direct := exec.Command("./binary")
	err = scamper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
