package main

// Libcds - C++ library of Concurrent Data Structures
// Homepage: https://libcds.sourceforge.net/doc/cds-api/index.html

import (
	"fmt"
	
	"os/exec"
)

func installLibcds() {
	// Método 1: Descargar y extraer .tar.gz
	libcds_tar_url := "https://github.com/khizmax/libcds/archive/refs/tags/v2.3.3.tar.gz"
	libcds_cmd_tar := exec.Command("curl", "-L", libcds_tar_url, "-o", "package.tar.gz")
	err := libcds_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcds_zip_url := "https://github.com/khizmax/libcds/archive/refs/tags/v2.3.3.zip"
	libcds_cmd_zip := exec.Command("curl", "-L", libcds_zip_url, "-o", "package.zip")
	err = libcds_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcds_bin_url := "https://github.com/khizmax/libcds/archive/refs/tags/v2.3.3.bin"
	libcds_cmd_bin := exec.Command("curl", "-L", libcds_bin_url, "-o", "binary.bin")
	err = libcds_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcds_src_url := "https://github.com/khizmax/libcds/archive/refs/tags/v2.3.3.src.tar.gz"
	libcds_cmd_src := exec.Command("curl", "-L", libcds_src_url, "-o", "source.tar.gz")
	err = libcds_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcds_cmd_direct := exec.Command("./binary")
	err = libcds_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
}
