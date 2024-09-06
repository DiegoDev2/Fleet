package main

// Networkit - Performance toolkit for large-scale network analysis
// Homepage: https://networkit.github.io

import (
	"fmt"
	
	"os/exec"
)

func installNetworkit() {
	// Método 1: Descargar y extraer .tar.gz
	networkit_tar_url := "https://github.com/networkit/networkit/archive/refs/tags/11.0.tar.gz"
	networkit_cmd_tar := exec.Command("curl", "-L", networkit_tar_url, "-o", "package.tar.gz")
	err := networkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	networkit_zip_url := "https://github.com/networkit/networkit/archive/refs/tags/11.0.zip"
	networkit_cmd_zip := exec.Command("curl", "-L", networkit_zip_url, "-o", "package.zip")
	err = networkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	networkit_bin_url := "https://github.com/networkit/networkit/archive/refs/tags/11.0.bin"
	networkit_cmd_bin := exec.Command("curl", "-L", networkit_bin_url, "-o", "binary.bin")
	err = networkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	networkit_src_url := "https://github.com/networkit/networkit/archive/refs/tags/11.0.src.tar.gz"
	networkit_cmd_src := exec.Command("curl", "-L", networkit_src_url, "-o", "source.tar.gz")
	err = networkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	networkit_cmd_direct := exec.Command("./binary")
	err = networkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: cython")
exec.Command("latte", "install", "cython")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: tlx")
exec.Command("latte", "install", "tlx")
	fmt.Println("Instalando dependencia: libnetworkit")
exec.Command("latte", "install", "libnetworkit")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: scipy")
exec.Command("latte", "install", "scipy")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
}
