package main

// Libnetworkit - NetworKit is an OS-toolkit for large-scale network analysis
// Homepage: https://networkit.github.io

import (
	"fmt"
	
	"os/exec"
)

func installLibnetworkit() {
	// Método 1: Descargar y extraer .tar.gz
	libnetworkit_tar_url := "https://github.com/networkit/networkit/archive/refs/tags/11.0.tar.gz"
	libnetworkit_cmd_tar := exec.Command("curl", "-L", libnetworkit_tar_url, "-o", "package.tar.gz")
	err := libnetworkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnetworkit_zip_url := "https://github.com/networkit/networkit/archive/refs/tags/11.0.zip"
	libnetworkit_cmd_zip := exec.Command("curl", "-L", libnetworkit_zip_url, "-o", "package.zip")
	err = libnetworkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnetworkit_bin_url := "https://github.com/networkit/networkit/archive/refs/tags/11.0.bin"
	libnetworkit_cmd_bin := exec.Command("curl", "-L", libnetworkit_bin_url, "-o", "binary.bin")
	err = libnetworkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnetworkit_src_url := "https://github.com/networkit/networkit/archive/refs/tags/11.0.src.tar.gz"
	libnetworkit_cmd_src := exec.Command("curl", "-L", libnetworkit_src_url, "-o", "source.tar.gz")
	err = libnetworkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnetworkit_cmd_direct := exec.Command("./binary")
	err = libnetworkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: tlx")
exec.Command("latte", "install", "tlx")
	fmt.Println("Instalando dependencia: ttmath")
exec.Command("latte", "install", "ttmath")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
}
