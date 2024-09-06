package main

// Schroedinger - High-speed implementation of the Dirac codec
// Homepage: https://launchpad.net/schroedinger

import (
	"fmt"
	
	"os/exec"
)

func installSchroedinger() {
	// Método 1: Descargar y extraer .tar.gz
	schroedinger_tar_url := "https://launchpad.net/schroedinger/trunk/1.0.11/+download/schroedinger-1.0.11.tar.gz"
	schroedinger_cmd_tar := exec.Command("curl", "-L", schroedinger_tar_url, "-o", "package.tar.gz")
	err := schroedinger_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	schroedinger_zip_url := "https://launchpad.net/schroedinger/trunk/1.0.11/+download/schroedinger-1.0.11.zip"
	schroedinger_cmd_zip := exec.Command("curl", "-L", schroedinger_zip_url, "-o", "package.zip")
	err = schroedinger_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	schroedinger_bin_url := "https://launchpad.net/schroedinger/trunk/1.0.11/+download/schroedinger-1.0.11.bin"
	schroedinger_cmd_bin := exec.Command("curl", "-L", schroedinger_bin_url, "-o", "binary.bin")
	err = schroedinger_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	schroedinger_src_url := "https://launchpad.net/schroedinger/trunk/1.0.11/+download/schroedinger-1.0.11.src.tar.gz"
	schroedinger_cmd_src := exec.Command("curl", "-L", schroedinger_src_url, "-o", "source.tar.gz")
	err = schroedinger_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	schroedinger_cmd_direct := exec.Command("./binary")
	err = schroedinger_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: orc")
	exec.Command("latte", "install", "orc").Run()
}
