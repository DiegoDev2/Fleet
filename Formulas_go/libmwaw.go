package main

// Libmwaw - Library for converting legacy Mac document formats
// Homepage: https://sourceforge.net/p/libmwaw/wiki/Home/

import (
	"fmt"
	
	"os/exec"
)

func installLibmwaw() {
	// Método 1: Descargar y extraer .tar.gz
	libmwaw_tar_url := "https://downloads.sourceforge.net/project/libmwaw/libmwaw/libmwaw-0.3.22/libmwaw-0.3.22.tar.xz"
	libmwaw_cmd_tar := exec.Command("curl", "-L", libmwaw_tar_url, "-o", "package.tar.gz")
	err := libmwaw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmwaw_zip_url := "https://downloads.sourceforge.net/project/libmwaw/libmwaw/libmwaw-0.3.22/libmwaw-0.3.22.tar.xz"
	libmwaw_cmd_zip := exec.Command("curl", "-L", libmwaw_zip_url, "-o", "package.zip")
	err = libmwaw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmwaw_bin_url := "https://downloads.sourceforge.net/project/libmwaw/libmwaw/libmwaw-0.3.22/libmwaw-0.3.22.tar.xz"
	libmwaw_cmd_bin := exec.Command("curl", "-L", libmwaw_bin_url, "-o", "binary.bin")
	err = libmwaw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmwaw_src_url := "https://downloads.sourceforge.net/project/libmwaw/libmwaw/libmwaw-0.3.22/libmwaw-0.3.22.tar.xz"
	libmwaw_cmd_src := exec.Command("curl", "-L", libmwaw_src_url, "-o", "source.tar.gz")
	err = libmwaw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmwaw_cmd_direct := exec.Command("./binary")
	err = libmwaw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: librevenge")
exec.Command("latte", "install", "librevenge")
}
