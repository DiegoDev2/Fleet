package main

// Libsmi - Library to Access SMI MIB Information
// Homepage: https://www.ibr.cs.tu-bs.de/projects/libsmi/

import (
	"fmt"
	
	"os/exec"
)

func installLibsmi() {
	// Método 1: Descargar y extraer .tar.gz
	libsmi_tar_url := "https://www.ibr.cs.tu-bs.de/projects/libsmi/download/libsmi-0.5.0.tar.gz"
	libsmi_cmd_tar := exec.Command("curl", "-L", libsmi_tar_url, "-o", "package.tar.gz")
	err := libsmi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsmi_zip_url := "https://www.ibr.cs.tu-bs.de/projects/libsmi/download/libsmi-0.5.0.zip"
	libsmi_cmd_zip := exec.Command("curl", "-L", libsmi_zip_url, "-o", "package.zip")
	err = libsmi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsmi_bin_url := "https://www.ibr.cs.tu-bs.de/projects/libsmi/download/libsmi-0.5.0.bin"
	libsmi_cmd_bin := exec.Command("curl", "-L", libsmi_bin_url, "-o", "binary.bin")
	err = libsmi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsmi_src_url := "https://www.ibr.cs.tu-bs.de/projects/libsmi/download/libsmi-0.5.0.src.tar.gz"
	libsmi_cmd_src := exec.Command("curl", "-L", libsmi_src_url, "-o", "source.tar.gz")
	err = libsmi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsmi_cmd_direct := exec.Command("./binary")
	err = libsmi_cmd_direct.Run()
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
}
