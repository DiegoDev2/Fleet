package main

// Libngspice - Spice circuit simulator as shared library
// Homepage: https://ngspice.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installLibngspice() {
	// Método 1: Descargar y extraer .tar.gz
	libngspice_tar_url := "https://downloads.sourceforge.net/project/ngspice/ng-spice-rework/43/ngspice-43.tar.gz"
	libngspice_cmd_tar := exec.Command("curl", "-L", libngspice_tar_url, "-o", "package.tar.gz")
	err := libngspice_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libngspice_zip_url := "https://downloads.sourceforge.net/project/ngspice/ng-spice-rework/43/ngspice-43.zip"
	libngspice_cmd_zip := exec.Command("curl", "-L", libngspice_zip_url, "-o", "package.zip")
	err = libngspice_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libngspice_bin_url := "https://downloads.sourceforge.net/project/ngspice/ng-spice-rework/43/ngspice-43.bin"
	libngspice_cmd_bin := exec.Command("curl", "-L", libngspice_bin_url, "-o", "binary.bin")
	err = libngspice_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libngspice_src_url := "https://downloads.sourceforge.net/project/ngspice/ng-spice-rework/43/ngspice-43.src.tar.gz"
	libngspice_cmd_src := exec.Command("curl", "-L", libngspice_src_url, "-o", "source.tar.gz")
	err = libngspice_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libngspice_cmd_direct := exec.Command("./binary")
	err = libngspice_cmd_direct.Run()
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
