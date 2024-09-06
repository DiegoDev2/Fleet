package main

// Libstxxl - C++ implementation of STL for extra large data sets
// Homepage: https://stxxl.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibstxxl() {
	// Método 1: Descargar y extraer .tar.gz
	libstxxl_tar_url := "https://downloads.sourceforge.net/project/stxxl/stxxl/1.4.1/stxxl-1.4.1.tar.gz"
	libstxxl_cmd_tar := exec.Command("curl", "-L", libstxxl_tar_url, "-o", "package.tar.gz")
	err := libstxxl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libstxxl_zip_url := "https://downloads.sourceforge.net/project/stxxl/stxxl/1.4.1/stxxl-1.4.1.zip"
	libstxxl_cmd_zip := exec.Command("curl", "-L", libstxxl_zip_url, "-o", "package.zip")
	err = libstxxl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libstxxl_bin_url := "https://downloads.sourceforge.net/project/stxxl/stxxl/1.4.1/stxxl-1.4.1.bin"
	libstxxl_cmd_bin := exec.Command("curl", "-L", libstxxl_bin_url, "-o", "binary.bin")
	err = libstxxl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libstxxl_src_url := "https://downloads.sourceforge.net/project/stxxl/stxxl/1.4.1/stxxl-1.4.1.src.tar.gz"
	libstxxl_cmd_src := exec.Command("curl", "-L", libstxxl_src_url, "-o", "source.tar.gz")
	err = libstxxl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libstxxl_cmd_direct := exec.Command("./binary")
	err = libstxxl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
