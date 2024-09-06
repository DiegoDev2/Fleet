package main

// Jpeginfo - Prints information and tests integrity of JPEG/JFIF files
// Homepage: https://www.kokkonen.net/tjko/projects.html

import (
	"fmt"
	
	"os/exec"
)

func installJpeginfo() {
	// Método 1: Descargar y extraer .tar.gz
	jpeginfo_tar_url := "https://www.kokkonen.net/tjko/src/jpeginfo-1.7.1.tar.gz"
	jpeginfo_cmd_tar := exec.Command("curl", "-L", jpeginfo_tar_url, "-o", "package.tar.gz")
	err := jpeginfo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jpeginfo_zip_url := "https://www.kokkonen.net/tjko/src/jpeginfo-1.7.1.zip"
	jpeginfo_cmd_zip := exec.Command("curl", "-L", jpeginfo_zip_url, "-o", "package.zip")
	err = jpeginfo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jpeginfo_bin_url := "https://www.kokkonen.net/tjko/src/jpeginfo-1.7.1.bin"
	jpeginfo_cmd_bin := exec.Command("curl", "-L", jpeginfo_bin_url, "-o", "binary.bin")
	err = jpeginfo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jpeginfo_src_url := "https://www.kokkonen.net/tjko/src/jpeginfo-1.7.1.src.tar.gz"
	jpeginfo_cmd_src := exec.Command("curl", "-L", jpeginfo_src_url, "-o", "source.tar.gz")
	err = jpeginfo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jpeginfo_cmd_direct := exec.Command("./binary")
	err = jpeginfo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
}
