package main

// Libgfshare - Library for sharing secrets
// Homepage: https://www.digital-scurf.org/software/libgfshare

import (
	"fmt"
	
	"os/exec"
)

func installLibgfshare() {
	// Método 1: Descargar y extraer .tar.gz
	libgfshare_tar_url := "https://www.digital-scurf.org/files/libgfshare/libgfshare-2.0.0.tar.bz2"
	libgfshare_cmd_tar := exec.Command("curl", "-L", libgfshare_tar_url, "-o", "package.tar.gz")
	err := libgfshare_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgfshare_zip_url := "https://www.digital-scurf.org/files/libgfshare/libgfshare-2.0.0.tar.bz2"
	libgfshare_cmd_zip := exec.Command("curl", "-L", libgfshare_zip_url, "-o", "package.zip")
	err = libgfshare_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgfshare_bin_url := "https://www.digital-scurf.org/files/libgfshare/libgfshare-2.0.0.tar.bz2"
	libgfshare_cmd_bin := exec.Command("curl", "-L", libgfshare_bin_url, "-o", "binary.bin")
	err = libgfshare_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgfshare_src_url := "https://www.digital-scurf.org/files/libgfshare/libgfshare-2.0.0.tar.bz2"
	libgfshare_cmd_src := exec.Command("curl", "-L", libgfshare_src_url, "-o", "source.tar.gz")
	err = libgfshare_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgfshare_cmd_direct := exec.Command("./binary")
	err = libgfshare_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
