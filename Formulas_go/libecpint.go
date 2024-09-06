package main

// Libecpint - Library for the efficient evaluation of integrals over effective core potentials
// Homepage: https://github.com/robashaw/libecpint

import (
	"fmt"
	
	"os/exec"
)

func installLibecpint() {
	// Método 1: Descargar y extraer .tar.gz
	libecpint_tar_url := "https://github.com/robashaw/libecpint/archive/refs/tags/v1.0.7.tar.gz"
	libecpint_cmd_tar := exec.Command("curl", "-L", libecpint_tar_url, "-o", "package.tar.gz")
	err := libecpint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libecpint_zip_url := "https://github.com/robashaw/libecpint/archive/refs/tags/v1.0.7.zip"
	libecpint_cmd_zip := exec.Command("curl", "-L", libecpint_zip_url, "-o", "package.zip")
	err = libecpint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libecpint_bin_url := "https://github.com/robashaw/libecpint/archive/refs/tags/v1.0.7.bin"
	libecpint_cmd_bin := exec.Command("curl", "-L", libecpint_bin_url, "-o", "binary.bin")
	err = libecpint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libecpint_src_url := "https://github.com/robashaw/libecpint/archive/refs/tags/v1.0.7.src.tar.gz"
	libecpint_cmd_src := exec.Command("curl", "-L", libecpint_src_url, "-o", "source.tar.gz")
	err = libecpint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libecpint_cmd_direct := exec.Command("./binary")
	err = libecpint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libcerf")
exec.Command("latte", "install", "libcerf")
	fmt.Println("Instalando dependencia: pugixml")
exec.Command("latte", "install", "pugixml")
}
