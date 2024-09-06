package main

// Exif - Read, write, modify, and display EXIF data on the command-line
// Homepage: https://libexif.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installExif() {
	// Método 1: Descargar y extraer .tar.gz
	exif_tar_url := "https://github.com/libexif/exif/releases/download/exif-0_6_22-release/exif-0.6.22.tar.xz"
	exif_cmd_tar := exec.Command("curl", "-L", exif_tar_url, "-o", "package.tar.gz")
	err := exif_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	exif_zip_url := "https://github.com/libexif/exif/releases/download/exif-0_6_22-release/exif-0.6.22.tar.xz"
	exif_cmd_zip := exec.Command("curl", "-L", exif_zip_url, "-o", "package.zip")
	err = exif_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	exif_bin_url := "https://github.com/libexif/exif/releases/download/exif-0_6_22-release/exif-0.6.22.tar.xz"
	exif_cmd_bin := exec.Command("curl", "-L", exif_bin_url, "-o", "binary.bin")
	err = exif_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	exif_src_url := "https://github.com/libexif/exif/releases/download/exif-0_6_22-release/exif-0.6.22.tar.xz"
	exif_cmd_src := exec.Command("curl", "-L", exif_src_url, "-o", "source.tar.gz")
	err = exif_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	exif_cmd_direct := exec.Command("./binary")
	err = exif_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libexif")
exec.Command("latte", "install", "libexif")
	fmt.Println("Instalando dependencia: popt")
exec.Command("latte", "install", "popt")
}
