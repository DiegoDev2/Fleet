package main

// Libxmu - X.Org: X miscellaneous utility routines library
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxmu() {
	// Método 1: Descargar y extraer .tar.gz
	libxmu_tar_url := "https://www.x.org/archive/individual/lib/libXmu-1.2.1.tar.xz"
	libxmu_cmd_tar := exec.Command("curl", "-L", libxmu_tar_url, "-o", "package.tar.gz")
	err := libxmu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxmu_zip_url := "https://www.x.org/archive/individual/lib/libXmu-1.2.1.tar.xz"
	libxmu_cmd_zip := exec.Command("curl", "-L", libxmu_zip_url, "-o", "package.zip")
	err = libxmu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxmu_bin_url := "https://www.x.org/archive/individual/lib/libXmu-1.2.1.tar.xz"
	libxmu_cmd_bin := exec.Command("curl", "-L", libxmu_bin_url, "-o", "binary.bin")
	err = libxmu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxmu_src_url := "https://www.x.org/archive/individual/lib/libXmu-1.2.1.tar.xz"
	libxmu_cmd_src := exec.Command("curl", "-L", libxmu_src_url, "-o", "source.tar.gz")
	err = libxmu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxmu_cmd_direct := exec.Command("./binary")
	err = libxmu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
}
