package main

// Libresample - Audio resampling C library
// Homepage: https://ccrma.stanford.edu/~jos/resample/Available_Software.html

import (
	"fmt"
	
	"os/exec"
)

func installLibresample() {
	// Método 1: Descargar y extraer .tar.gz
	libresample_tar_url := "https://deb.debian.org/debian/pool/main/libr/libresample/libresample_0.1.3.orig.tar.gz"
	libresample_cmd_tar := exec.Command("curl", "-L", libresample_tar_url, "-o", "package.tar.gz")
	err := libresample_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libresample_zip_url := "https://deb.debian.org/debian/pool/main/libr/libresample/libresample_0.1.3.orig.zip"
	libresample_cmd_zip := exec.Command("curl", "-L", libresample_zip_url, "-o", "package.zip")
	err = libresample_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libresample_bin_url := "https://deb.debian.org/debian/pool/main/libr/libresample/libresample_0.1.3.orig.bin"
	libresample_cmd_bin := exec.Command("curl", "-L", libresample_bin_url, "-o", "binary.bin")
	err = libresample_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libresample_src_url := "https://deb.debian.org/debian/pool/main/libr/libresample/libresample_0.1.3.orig.src.tar.gz"
	libresample_cmd_src := exec.Command("curl", "-L", libresample_src_url, "-o", "source.tar.gz")
	err = libresample_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libresample_cmd_direct := exec.Command("./binary")
	err = libresample_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
