package main

// Exiftool - Perl lib for reading and writing EXIF metadata
// Homepage: https://exiftool.org

import (
	"fmt"
	
	"os/exec"
)

func installExiftool() {
	// Método 1: Descargar y extraer .tar.gz
	exiftool_tar_url := "https://cpan.metacpan.org/authors/id/E/EX/EXIFTOOL/Image-ExifTool-12.76.tar.gz"
	exiftool_cmd_tar := exec.Command("curl", "-L", exiftool_tar_url, "-o", "package.tar.gz")
	err := exiftool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	exiftool_zip_url := "https://cpan.metacpan.org/authors/id/E/EX/EXIFTOOL/Image-ExifTool-12.76.zip"
	exiftool_cmd_zip := exec.Command("curl", "-L", exiftool_zip_url, "-o", "package.zip")
	err = exiftool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	exiftool_bin_url := "https://cpan.metacpan.org/authors/id/E/EX/EXIFTOOL/Image-ExifTool-12.76.bin"
	exiftool_cmd_bin := exec.Command("curl", "-L", exiftool_bin_url, "-o", "binary.bin")
	err = exiftool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	exiftool_src_url := "https://cpan.metacpan.org/authors/id/E/EX/EXIFTOOL/Image-ExifTool-12.76.src.tar.gz"
	exiftool_cmd_src := exec.Command("curl", "-L", exiftool_src_url, "-o", "source.tar.gz")
	err = exiftool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	exiftool_cmd_direct := exec.Command("./binary")
	err = exiftool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
