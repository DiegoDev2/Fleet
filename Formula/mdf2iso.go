package main

// Mdf2iso - Tool to convert MDF (Alcohol 120% images) images to ISO images
// Homepage: https://packages.debian.org/sid/mdf2iso

import (
	"fmt"
	
	"os/exec"
)

func installMdf2iso() {
	// Método 1: Descargar y extraer .tar.gz
	mdf2iso_tar_url := "https://deb.debian.org/debian/pool/main/m/mdf2iso/mdf2iso_0.3.1.orig.tar.gz"
	mdf2iso_cmd_tar := exec.Command("curl", "-L", mdf2iso_tar_url, "-o", "package.tar.gz")
	err := mdf2iso_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mdf2iso_zip_url := "https://deb.debian.org/debian/pool/main/m/mdf2iso/mdf2iso_0.3.1.orig.zip"
	mdf2iso_cmd_zip := exec.Command("curl", "-L", mdf2iso_zip_url, "-o", "package.zip")
	err = mdf2iso_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mdf2iso_bin_url := "https://deb.debian.org/debian/pool/main/m/mdf2iso/mdf2iso_0.3.1.orig.bin"
	mdf2iso_cmd_bin := exec.Command("curl", "-L", mdf2iso_bin_url, "-o", "binary.bin")
	err = mdf2iso_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mdf2iso_src_url := "https://deb.debian.org/debian/pool/main/m/mdf2iso/mdf2iso_0.3.1.orig.src.tar.gz"
	mdf2iso_cmd_src := exec.Command("curl", "-L", mdf2iso_src_url, "-o", "source.tar.gz")
	err = mdf2iso_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mdf2iso_cmd_direct := exec.Command("./binary")
	err = mdf2iso_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
