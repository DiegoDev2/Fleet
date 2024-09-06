package main

// Shivavg - OpenGL based ANSI C implementation of the OpenVG standard
// Homepage: https://sourceforge.net/projects/shivavg/

import (
	"fmt"
	
	"os/exec"
)

func installShivavg() {
	// Método 1: Descargar y extraer .tar.gz
	shivavg_tar_url := "https://downloads.sourceforge.net/project/shivavg/ShivaVG/0.2.1/ShivaVG-0.2.1.zip"
	shivavg_cmd_tar := exec.Command("curl", "-L", shivavg_tar_url, "-o", "package.tar.gz")
	err := shivavg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shivavg_zip_url := "https://downloads.sourceforge.net/project/shivavg/ShivaVG/0.2.1/ShivaVG-0.2.1.zip"
	shivavg_cmd_zip := exec.Command("curl", "-L", shivavg_zip_url, "-o", "package.zip")
	err = shivavg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shivavg_bin_url := "https://downloads.sourceforge.net/project/shivavg/ShivaVG/0.2.1/ShivaVG-0.2.1.zip"
	shivavg_cmd_bin := exec.Command("curl", "-L", shivavg_bin_url, "-o", "binary.bin")
	err = shivavg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shivavg_src_url := "https://downloads.sourceforge.net/project/shivavg/ShivaVG/0.2.1/ShivaVG-0.2.1.zip"
	shivavg_cmd_src := exec.Command("curl", "-L", shivavg_src_url, "-o", "source.tar.gz")
	err = shivavg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shivavg_cmd_direct := exec.Command("./binary")
	err = shivavg_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: freeglut")
	exec.Command("latte", "install", "freeglut").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
