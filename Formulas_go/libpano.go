package main

// Libpano - Build panoramic images from a set of overlapping images
// Homepage: https://panotools.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibpano() {
	// Método 1: Descargar y extraer .tar.gz
	libpano_tar_url := "https://downloads.sourceforge.net/project/panotools/libpano13/libpano13-2.9.22/libpano13-2.9.22.tar.gz"
	libpano_cmd_tar := exec.Command("curl", "-L", libpano_tar_url, "-o", "package.tar.gz")
	err := libpano_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpano_zip_url := "https://downloads.sourceforge.net/project/panotools/libpano13/libpano13-2.9.22/libpano13-2.9.22.zip"
	libpano_cmd_zip := exec.Command("curl", "-L", libpano_zip_url, "-o", "package.zip")
	err = libpano_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpano_bin_url := "https://downloads.sourceforge.net/project/panotools/libpano13/libpano13-2.9.22/libpano13-2.9.22.bin"
	libpano_cmd_bin := exec.Command("curl", "-L", libpano_bin_url, "-o", "binary.bin")
	err = libpano_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpano_src_url := "https://downloads.sourceforge.net/project/panotools/libpano13/libpano13-2.9.22/libpano13-2.9.22.src.tar.gz"
	libpano_cmd_src := exec.Command("curl", "-L", libpano_src_url, "-o", "source.tar.gz")
	err = libpano_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpano_cmd_direct := exec.Command("./binary")
	err = libpano_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
}
