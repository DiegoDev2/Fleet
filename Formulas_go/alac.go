package main

// Alac - Basic decoder for Apple Lossless Audio Codec files (ALAC)
// Homepage: https://web.archive.org/web/20150319040222/craz.net/programs/itunes/alac.html

import (
	"fmt"
	
	"os/exec"
)

func installAlac() {
	// Método 1: Descargar y extraer .tar.gz
	alac_tar_url := "https://web.archive.org/web/20150510210401/craz.net/programs/itunes/files/alac_decoder-0.2.0.tgz"
	alac_cmd_tar := exec.Command("curl", "-L", alac_tar_url, "-o", "package.tar.gz")
	err := alac_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	alac_zip_url := "https://web.archive.org/web/20150510210401/craz.net/programs/itunes/files/alac_decoder-0.2.0.tgz"
	alac_cmd_zip := exec.Command("curl", "-L", alac_zip_url, "-o", "package.zip")
	err = alac_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	alac_bin_url := "https://web.archive.org/web/20150510210401/craz.net/programs/itunes/files/alac_decoder-0.2.0.tgz"
	alac_cmd_bin := exec.Command("curl", "-L", alac_bin_url, "-o", "binary.bin")
	err = alac_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	alac_src_url := "https://web.archive.org/web/20150510210401/craz.net/programs/itunes/files/alac_decoder-0.2.0.tgz"
	alac_cmd_src := exec.Command("curl", "-L", alac_src_url, "-o", "source.tar.gz")
	err = alac_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	alac_cmd_direct := exec.Command("./binary")
	err = alac_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
