package main

// Mkcue - Generate a CUE sheet from a CD
// Homepage: https://packages.debian.org/sid/mkcue

import (
	"fmt"
	
	"os/exec"
)

func installMkcue() {
	// Método 1: Descargar y extraer .tar.gz
	mkcue_tar_url := "https://deb.debian.org/debian/pool/main/m/mkcue/mkcue_1.orig.tar.gz"
	mkcue_cmd_tar := exec.Command("curl", "-L", mkcue_tar_url, "-o", "package.tar.gz")
	err := mkcue_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mkcue_zip_url := "https://deb.debian.org/debian/pool/main/m/mkcue/mkcue_1.orig.zip"
	mkcue_cmd_zip := exec.Command("curl", "-L", mkcue_zip_url, "-o", "package.zip")
	err = mkcue_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mkcue_bin_url := "https://deb.debian.org/debian/pool/main/m/mkcue/mkcue_1.orig.bin"
	mkcue_cmd_bin := exec.Command("curl", "-L", mkcue_bin_url, "-o", "binary.bin")
	err = mkcue_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mkcue_src_url := "https://deb.debian.org/debian/pool/main/m/mkcue/mkcue_1.orig.src.tar.gz"
	mkcue_cmd_src := exec.Command("curl", "-L", mkcue_src_url, "-o", "source.tar.gz")
	err = mkcue_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mkcue_cmd_direct := exec.Command("./binary")
	err = mkcue_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
