package main

// Montage - Toolkit for assembling FITS images into custom mosaics
// Homepage: http://montage.ipac.caltech.edu

import (
	"fmt"
	
	"os/exec"
)

func installMontage() {
	// Método 1: Descargar y extraer .tar.gz
	montage_tar_url := "http://montage.ipac.caltech.edu/download/Montage_v6.0.tar.gz"
	montage_cmd_tar := exec.Command("curl", "-L", montage_tar_url, "-o", "package.tar.gz")
	err := montage_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	montage_zip_url := "http://montage.ipac.caltech.edu/download/Montage_v6.0.zip"
	montage_cmd_zip := exec.Command("curl", "-L", montage_zip_url, "-o", "package.zip")
	err = montage_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	montage_bin_url := "http://montage.ipac.caltech.edu/download/Montage_v6.0.bin"
	montage_cmd_bin := exec.Command("curl", "-L", montage_bin_url, "-o", "binary.bin")
	err = montage_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	montage_src_url := "http://montage.ipac.caltech.edu/download/Montage_v6.0.src.tar.gz"
	montage_cmd_src := exec.Command("curl", "-L", montage_src_url, "-o", "source.tar.gz")
	err = montage_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	montage_cmd_direct := exec.Command("./binary")
	err = montage_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cfitsio")
	exec.Command("latte", "install", "cfitsio").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
}
