package main

// Bsdiff - Generate and apply patches to binary files
// Homepage: https://www.daemonology.net/bsdiff/

import (
	"fmt"
	
	"os/exec"
)

func installBsdiff() {
	// Método 1: Descargar y extraer .tar.gz
	bsdiff_tar_url := "https://deb.debian.org/debian/pool/main/b/bsdiff/bsdiff_4.3.orig.tar.gz"
	bsdiff_cmd_tar := exec.Command("curl", "-L", bsdiff_tar_url, "-o", "package.tar.gz")
	err := bsdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bsdiff_zip_url := "https://deb.debian.org/debian/pool/main/b/bsdiff/bsdiff_4.3.orig.zip"
	bsdiff_cmd_zip := exec.Command("curl", "-L", bsdiff_zip_url, "-o", "package.zip")
	err = bsdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bsdiff_bin_url := "https://deb.debian.org/debian/pool/main/b/bsdiff/bsdiff_4.3.orig.bin"
	bsdiff_cmd_bin := exec.Command("curl", "-L", bsdiff_bin_url, "-o", "binary.bin")
	err = bsdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bsdiff_src_url := "https://deb.debian.org/debian/pool/main/b/bsdiff/bsdiff_4.3.orig.src.tar.gz"
	bsdiff_cmd_src := exec.Command("curl", "-L", bsdiff_src_url, "-o", "source.tar.gz")
	err = bsdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bsdiff_cmd_direct := exec.Command("./binary")
	err = bsdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bsdmake")
	exec.Command("latte", "install", "bsdmake").Run()
}
