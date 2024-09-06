package main

// DvdVr - Utility to identify and extract recordings from DVD-VR files
// Homepage: https://www.pixelbeat.org/programs/dvd-vr/

import (
	"fmt"
	
	"os/exec"
)

func installDvdVr() {
	// Método 1: Descargar y extraer .tar.gz
	dvdvr_tar_url := "https://www.pixelbeat.org/programs/dvd-vr/dvd-vr-0.9.7.tar.gz"
	dvdvr_cmd_tar := exec.Command("curl", "-L", dvdvr_tar_url, "-o", "package.tar.gz")
	err := dvdvr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dvdvr_zip_url := "https://www.pixelbeat.org/programs/dvd-vr/dvd-vr-0.9.7.zip"
	dvdvr_cmd_zip := exec.Command("curl", "-L", dvdvr_zip_url, "-o", "package.zip")
	err = dvdvr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dvdvr_bin_url := "https://www.pixelbeat.org/programs/dvd-vr/dvd-vr-0.9.7.bin"
	dvdvr_cmd_bin := exec.Command("curl", "-L", dvdvr_bin_url, "-o", "binary.bin")
	err = dvdvr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dvdvr_src_url := "https://www.pixelbeat.org/programs/dvd-vr/dvd-vr-0.9.7.src.tar.gz"
	dvdvr_cmd_src := exec.Command("curl", "-L", dvdvr_src_url, "-o", "source.tar.gz")
	err = dvdvr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dvdvr_cmd_direct := exec.Command("./binary")
	err = dvdvr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
