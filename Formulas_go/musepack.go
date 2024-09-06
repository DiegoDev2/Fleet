package main

// Musepack - Audio compression format and tools
// Homepage: https://www.musepack.net/

import (
	"fmt"
	
	"os/exec"
)

func installMusepack() {
	// Método 1: Descargar y extraer .tar.gz
	musepack_tar_url := "https://files.musepack.net/source/musepack_src_r475.tar.gz"
	musepack_cmd_tar := exec.Command("curl", "-L", musepack_tar_url, "-o", "package.tar.gz")
	err := musepack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	musepack_zip_url := "https://files.musepack.net/source/musepack_src_r475.zip"
	musepack_cmd_zip := exec.Command("curl", "-L", musepack_zip_url, "-o", "package.zip")
	err = musepack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	musepack_bin_url := "https://files.musepack.net/source/musepack_src_r475.bin"
	musepack_cmd_bin := exec.Command("curl", "-L", musepack_bin_url, "-o", "binary.bin")
	err = musepack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	musepack_src_url := "https://files.musepack.net/source/musepack_src_r475.src.tar.gz"
	musepack_cmd_src := exec.Command("curl", "-L", musepack_src_url, "-o", "source.tar.gz")
	err = musepack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	musepack_cmd_direct := exec.Command("./binary")
	err = musepack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libcuefile")
exec.Command("latte", "install", "libcuefile")
	fmt.Println("Instalando dependencia: libreplaygain")
exec.Command("latte", "install", "libreplaygain")
}
