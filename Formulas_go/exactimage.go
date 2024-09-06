package main

// ExactImage - Image processing library
// Homepage: https://exactcode.com/opensource/exactimage/

import (
	"fmt"
	
	"os/exec"
)

func installExactImage() {
	// Método 1: Descargar y extraer .tar.gz
	exactimage_tar_url := "https://dl.exactcode.de/oss/exact-image/exact-image-1.0.2.tar.bz2"
	exactimage_cmd_tar := exec.Command("curl", "-L", exactimage_tar_url, "-o", "package.tar.gz")
	err := exactimage_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	exactimage_zip_url := "https://dl.exactcode.de/oss/exact-image/exact-image-1.0.2.tar.bz2"
	exactimage_cmd_zip := exec.Command("curl", "-L", exactimage_zip_url, "-o", "package.zip")
	err = exactimage_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	exactimage_bin_url := "https://dl.exactcode.de/oss/exact-image/exact-image-1.0.2.tar.bz2"
	exactimage_cmd_bin := exec.Command("curl", "-L", exactimage_bin_url, "-o", "binary.bin")
	err = exactimage_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	exactimage_src_url := "https://dl.exactcode.de/oss/exact-image/exact-image-1.0.2.tar.bz2"
	exactimage_cmd_src := exec.Command("curl", "-L", exactimage_src_url, "-o", "source.tar.gz")
	err = exactimage_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	exactimage_cmd_direct := exec.Command("./binary")
	err = exactimage_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libagg")
exec.Command("latte", "install", "libagg")
}
