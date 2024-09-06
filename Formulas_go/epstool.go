package main

// Epstool - Edit preview images and fix bounding boxes in EPS files
// Homepage: http://www.ghostgum.com.au/software/epstool.htm

import (
	"fmt"
	
	"os/exec"
)

func installEpstool() {
	// Método 1: Descargar y extraer .tar.gz
	epstool_tar_url := "https://deb.debian.org/debian/pool/main/e/epstool/epstool_3.09.orig.tar.xz"
	epstool_cmd_tar := exec.Command("curl", "-L", epstool_tar_url, "-o", "package.tar.gz")
	err := epstool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	epstool_zip_url := "https://deb.debian.org/debian/pool/main/e/epstool/epstool_3.09.orig.tar.xz"
	epstool_cmd_zip := exec.Command("curl", "-L", epstool_zip_url, "-o", "package.zip")
	err = epstool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	epstool_bin_url := "https://deb.debian.org/debian/pool/main/e/epstool/epstool_3.09.orig.tar.xz"
	epstool_cmd_bin := exec.Command("curl", "-L", epstool_bin_url, "-o", "binary.bin")
	err = epstool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	epstool_src_url := "https://deb.debian.org/debian/pool/main/e/epstool/epstool_3.09.orig.tar.xz"
	epstool_cmd_src := exec.Command("curl", "-L", epstool_src_url, "-o", "source.tar.gz")
	err = epstool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	epstool_cmd_direct := exec.Command("./binary")
	err = epstool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ghostscript")
exec.Command("latte", "install", "ghostscript")
}
