package main

// Libdsk - Library for accessing discs and disc image files
// Homepage: https://www.seasip.info/Unix/LibDsk/

import (
	"fmt"
	
	"os/exec"
)

func installLibdsk() {
	// Método 1: Descargar y extraer .tar.gz
	libdsk_tar_url := "https://www.seasip.info/Unix/LibDsk/libdsk-1.4.2.tar.gz"
	libdsk_cmd_tar := exec.Command("curl", "-L", libdsk_tar_url, "-o", "package.tar.gz")
	err := libdsk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdsk_zip_url := "https://www.seasip.info/Unix/LibDsk/libdsk-1.4.2.zip"
	libdsk_cmd_zip := exec.Command("curl", "-L", libdsk_zip_url, "-o", "package.zip")
	err = libdsk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdsk_bin_url := "https://www.seasip.info/Unix/LibDsk/libdsk-1.4.2.bin"
	libdsk_cmd_bin := exec.Command("curl", "-L", libdsk_bin_url, "-o", "binary.bin")
	err = libdsk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdsk_src_url := "https://www.seasip.info/Unix/LibDsk/libdsk-1.4.2.src.tar.gz"
	libdsk_cmd_src := exec.Command("curl", "-L", libdsk_src_url, "-o", "source.tar.gz")
	err = libdsk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdsk_cmd_direct := exec.Command("./binary")
	err = libdsk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
