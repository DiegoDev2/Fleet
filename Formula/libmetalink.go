package main

// Libmetalink - C library to parse Metalink XML files
// Homepage: https://launchpad.net/libmetalink/

import (
	"fmt"
	
	"os/exec"
)

func installLibmetalink() {
	// Método 1: Descargar y extraer .tar.gz
	libmetalink_tar_url := "https://launchpad.net/libmetalink/trunk/libmetalink-0.1.3/+download/libmetalink-0.1.3.tar.xz"
	libmetalink_cmd_tar := exec.Command("curl", "-L", libmetalink_tar_url, "-o", "package.tar.gz")
	err := libmetalink_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmetalink_zip_url := "https://launchpad.net/libmetalink/trunk/libmetalink-0.1.3/+download/libmetalink-0.1.3.tar.xz"
	libmetalink_cmd_zip := exec.Command("curl", "-L", libmetalink_zip_url, "-o", "package.zip")
	err = libmetalink_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmetalink_bin_url := "https://launchpad.net/libmetalink/trunk/libmetalink-0.1.3/+download/libmetalink-0.1.3.tar.xz"
	libmetalink_cmd_bin := exec.Command("curl", "-L", libmetalink_bin_url, "-o", "binary.bin")
	err = libmetalink_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmetalink_src_url := "https://launchpad.net/libmetalink/trunk/libmetalink-0.1.3/+download/libmetalink-0.1.3.tar.xz"
	libmetalink_cmd_src := exec.Command("curl", "-L", libmetalink_src_url, "-o", "source.tar.gz")
	err = libmetalink_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmetalink_cmd_direct := exec.Command("./binary")
	err = libmetalink_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
